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

func randWeight() float64 {
	return rand.Float64() - rand.Float64()
}

func sigmoid(input float64) float64 {
	p := 1.0
	return 1.0 / (1.0 + math.Exp(-1*input/p))
}

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
