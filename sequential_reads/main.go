package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Define flags
	goMaxProcs := flag.Int("gomaxprocs", runtime.GOMAXPROCS(0), "GoMaxProcs value")
	numThreads := flag.Int("max-num-goroutines", runtime.NumCPU(), "Number of GoRoutines")
	mountPoint := flag.String("mount-point", "", "Mount point")

	// Parse the command-line flags from os.Args[1:]. Must be called
	// after all flags are defined and before flags are accessed by the program.
	flag.Parse()

	if *mountPoint == "" {
		log.Fatal("-mount-point is a required flag")
	}

	maxBW := 0.0
	bestGOMaxProcs := 0
	bestNumThreads := 0
	for i := 1; i < *goMaxProcs; i = i + 10 {
		runtime.GOMAXPROCS(*goMaxProcs)
		for i := 1; i < *numThreads; i = i + 10 {
			clearPageCache()
			bw := computeBandwidth(*numThreads, *mountPoint)
			if bw > maxBW {
				bestGOMaxProcs = *goMaxProcs
				bestNumThreads = *numThreads
				maxBW = bw
			}
		}
	}

	fmt.Printf("Best GOMAXPROCS: %d, Goroutines: %d\n", bestGOMaxProcs, bestNumThreads)
}

func clearPageCache() {
	// This requires root privileges and is for Linux systems.
	// In a real application, you'd handle permissions and OS differences.
	cmd := exec.Command("sudo", "sh", "-c", "echo 3 > /proc/sys/vm/drop_caches")
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("Failed to clear page cache: %v", err))
	}
}
func computeBandwidth(numThreads int, mountPoint string) float64 {
	files := []string{}
	err := filepath.WalkDir(mountPoint, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Type().IsRegular() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to walk directory %s: %v", mountPoint, err)
	}

	if len(files) == 0 {
		log.Printf("No files found in %s", mountPoint)
		return 0
	}

	jobs := make(chan string, len(files))
	results := make(chan int64, len(files))

	var wg sync.WaitGroup
	buffer := make([]byte, numThreads*1024*1024)
	for w := range numThreads {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			for path := range jobs {
				f, err := os.Open(path)
				if err != nil {
					log.Printf("Failed to open file %s: %v", path, err)
					continue
				}

				n, err := io.CopyBuffer(io.Discard, f, buffer[k*1024*1024:(k+1)*1024*1024])
				f.Close()
				if err != nil {
					log.Printf("Failed to read file %s: %v", path, err)
					continue
				}
				results <- n
			}
		}(w)
	}

	startTime := time.Now()

	for _, file := range files {
		jobs <- file
	}
	close(jobs)

	wg.Wait()
	close(results)

	duration := time.Since(startTime)

	var totalBytesRead int64
	for bytesRead := range results {
		totalBytesRead += bytesRead
	}

	if duration.Seconds() == 0 {
		return 0
	}

	return totalBytesRead / (1024 * 1024 * 1024 * duration.Seconds())
}
