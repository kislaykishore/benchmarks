package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math"
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
	bestTimeTaken := time.Duration(math.MaxInt64)
	for i := *goMaxProcs; i > 0; i = i - 10 {
		runtime.GOMAXPROCS(i)
		for j := *numThreads; j > 0; j = j - 10 {
			clearPageCache()
			bw, timeTaken, cancelled := computeBandwidth(j, *mountPoint, bestTimeTaken)
			if cancelled {
				fmt.Printf("Run with GOMAXPROCS: %d, Goroutines: %d was cancelled as it exceeded best time (%.2fs).\n", i, j, bestTimeTaken.Seconds())
				continue
			}
			fmt.Printf("Running with GOMAXPROCS: %d, Goroutines: %d -> BW: %.2f GiB/s, Time: %.2fs\n", i, j, bw, timeTaken.Seconds())
			if bw > maxBW {
				bestGOMaxProcs = i
				bestNumThreads = j
				maxBW = bw
				bestTimeTaken = timeTaken
			}
		}
	}

	fmt.Printf("\n--- Results ---\nMax bandwidth: %.2f GiB/s\n", maxBW)
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

func computeBandwidth(numThreads int, mountPoint string, bestTimeTaken time.Duration) (float64, time.Duration, bool) {
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
		return 0, 0, false
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan string, len(files))
	results := make(chan int64, len(files))

	var wg sync.WaitGroup
	buffer := make([]byte, numThreads*1024*1024)
	for w := range numThreads {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			for {
				select {
				case path, ok := <-jobs:
					if !ok {
						return
					}
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
				case <-ctx.Done():
					return
				}
			}
		}(w)
	}

	startTime := time.Now()

	// Goroutine to cancel if time exceeds bestTimeTaken
	go func() {
		if bestTimeTaken >= time.Duration(math.MaxInt64) {
			return
		}
		timer := time.NewTimer(bestTimeTaken)
		defer timer.Stop()
		select {
		case <-timer.C:
			cancel()
		case <-ctx.Done():
		}
	}()

jobLoop:
	for _, file := range files {
		select {
		case jobs <- file:
		case <-ctx.Done():
			break jobLoop
		}
	}
	close(jobs)

	wg.Wait()
	close(results)

	duration := time.Since(startTime)
	if ctx.Err() != nil {
		return 0, duration, true
	}

	var totalBytesRead int64
	for bytesRead := range results {
		totalBytesRead += bytesRead
	}

	if duration.Seconds() == 0 {
		return 0, duration, false
	}

	return float64(totalBytesRead) / (1024 * 1024 * 1024) / duration.Seconds(), duration, false
}
