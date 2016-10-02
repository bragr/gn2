// util.go Copyright (c) 2016 Grant Brady
// Licensed under the MIT License
// See LICENSE.TXT
package gn2

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"os"
)

var needSeed bool = true

// Return a random float64 -1.0 <= n <= 1 but only if rand has been seeded
func randWeight() float64 {
	if needSeed {
		SeedRand()
	}
	return rand.Float64() - rand.Float64()
}

// Normalize the output of the nodes
func sigmoid(input float64) float64 {
	p := 1.0
	return 1.0 / (1.0 + math.Exp(-1*input/p))
}

// Seed random from system entropy. Unix systems only
func SeedRand() {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		fmt.Printf("Failed to open random source: %s\n", err)
		os.Exit(1)
	}
	seed := make([]byte, 8)
	i, err := f.Read(seed)
	if err != nil {
		fmt.Printf("Failed to read random source: %s\n", err)
		os.Exit(1)
	}
	if i != 8 {
		fmt.Printf("Was expecting to read 8 bytes of seeds, read %d bytes instead\n", i)
		os.Exit(1)
	}
	seed64 := int64(binary.LittleEndian.Uint64(seed))
	rand.Seed(seed64)
}
