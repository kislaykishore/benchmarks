package main_test

import (
	"crypto/rand"
	"maps"
	"math/big"
	"slices"
	"testing"
)

type key struct {
	part1 string
	part2 string
}

func createRandomStringOfLength(length int) string {
	str := rand.Text()
	for len(str) < length {
		str += rand.Text()
	}
	return str[:length]

}

func BenchmarkSequentialStructMapAccessPerf(b *testing.B) {
	// Initialize the map
	part1Length := 30
	part2Length := 60

	entries := 500
	// Prepare data
	mp := make(map[key]int)

	for range entries {
		k := key{
			part1: createRandomStringOfLength(part1Length),
			part2: createRandomStringOfLength(part2Length),
		}

		v, err := rand.Int(rand.Reader, big.NewInt(12345))
		if err != nil {
			panic(err)
		}
		mp[k] = int(v.Int64())
	}

	keys := slices.Collect(maps.Keys(mp))
	idx := 0
	for b.Loop() {
		idx = idx % len(keys)
		k := keys[idx]
		idx++
		_ = mp[k]
	}
}
