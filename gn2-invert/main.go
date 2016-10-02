// main.go Copyright (c) 2016 Grant Brady
// Licensed under the MIT License. See LICENSE.TXT
package main

import (
	"fmt"
	"github.com/bragr/gn2/gn2"
	"math/rand"
	"sort"
)

func main() {
	// Seed rand from system entropy source
	gn2.SeedRand()

	// Make the training data set. In this case a simple inversion
	traingSetSize := 1000
	inputs := make([]float64, traingSetSize)
	for i := 0; i < traingSetSize; i++ {
		inputs[i] = rand.Float64()
	}
	sort.Sort(sort.Float64Slice(inputs))
	answers := make([]float64, traingSetSize)
	for i := 0; i < traingSetSize; i++ {
		answers[i] = 1 - inputs[i]
	}

	// Make a new "species" to train on
	species := gn2.NewSpecies(20, 1, 1, 4, 15)

	// Train the species
	for i := 0; i < 1000; i++ {
		fmt.Printf(".")
		species.Compete(inputs, answers)

		// Get best 5 nets, make 3 mutated children for each, and start again
		sort.Sort(species)
		for chromo := 0; chromo < 5; chromo++ {
			species[chromo].Fitness = 0.0
			species[3*chromo+5].Fitness = 0.0
			species[3*chromo+5].Net = species[chromo].Net.Mutate(0.025, 0.3, 1, 1, 3, 10)
			species[3*chromo+6].Fitness = 0.0
			species[3*chromo+6].Net = species[chromo].Net.Mutate(0.05, 0.3, 1, 1, 3, 10)
			species[3*chromo+7].Fitness = 0.0
			species[3*chromo+7].Net = species[chromo].Net.Mutate(0.1, 0.3, 1, 1, 3, 10)
		}
	}
	fmt.Println()

	// Print the results for the winner
	species[0].Net.PrintNet()
	for i, input := range inputs {
		output := species[0].Net.Update([]float64{input})[0]
		fmt.Printf("Input: %f, Output: %f, Answer: %f Difference: %f%%\n", input, output, answers[i], output/answers[i])
	}

}
