// main.go Copyright (c) 2016 Grant Brady
// Licensed under the MIT License. See LICENSE.TXT
package main

import (
	"fmt"
	"github.com/bragr/gn2/gn2"
	"math/rand"
	"sort"
)

// Make the training data set. In this case a simple inversion
func genTraingingData(dataSize int64) (inputs, outputs [][]float64) {
	for i := int64(0); i < dataSize; i++ {
		inputs = append(inputs, []float64{rand.Float64()})
	}
	for i := int64(0); i < dataSize; i++ {
		outputs = append(outputs, []float64{1.0 - inputs[i][0]})
	}
	return inputs, outputs
}

func main() {
	// Seed rand from system entropy source
	gn2.SeedRand()

	// Make a new "species" to train on
	species := gn2.NewSpecies(20, 1, 1, 4, 15)

	// Train the species
	for i := 0; i < 1000; i++ {
		if i%10 == 0 {
			fmt.Printf(".")
		}
		species.Compete(genTraingingData(1000))

		// Get best 5 nets, make 3 mutated children for each, and start again
		sort.Sort(species)
		for chromo := 0; chromo < 5; chromo++ {
			species[chromo].Fitness = 0.0
			species[3*chromo+5].Fitness = 0.0
			species[3*chromo+5].Net = species[chromo].Net.Mutate(0.025, 0.3, 1, 1, 4, 15)
			species[3*chromo+6].Fitness = 0.0
			species[3*chromo+6].Net = species[chromo].Net.Mutate(0.05, 0.3, 1, 1, 4, 15)
			species[3*chromo+7].Fitness = 0.0
			species[3*chromo+7].Net = species[chromo].Net.Mutate(0.1, 0.3, 1, 1, 4, 15)
		}
	}
	fmt.Println()

	// Print the results for the winner
	species[0].Net.PrintNet()
	inputs, outputs := genTraingingData(1000)
	for i, input := range inputs {
		output := species[0].Net.Update(input)[0]
		fmt.Printf("Input: %f, Output: %f, Answer: %f Difference: %f%%\n", input[0], output, outputs[i][0], output/outputs[i][0])
	}

}
