// main.go Copyright (c) 2016 Grant Brady
// Licensed under the MIT License. See LICENSE.TXT
package main

import (
	"fmt"
	"github.com/bragr/gn2/gn2"
	"math/rand"
)

// Make the training data set. In this case a simple inversion
func genTraingingData(dataSize int64) (inputs, outputs [][]float64) {
	for i := int64(0); i < dataSize; i++ {
		inputs = append(inputs, []float64{rand.Float64()})
	}
	for i := int64(0); i < dataSize; i++ {
		outputs = append(outputs, []float64{1.0 - inputs[i][0]})
	}
	inputs = append(inputs, []float64{1.0})
	outputs = append(outputs, []float64{0.0})
	inputs = append(inputs, []float64{0.0})
	outputs = append(outputs, []float64{1.0})
	return inputs, outputs
}

func main() {
	// Seed rand from system entropy source
	gn2.SeedRand()

	// Make a new "species" to train on
	species := gn2.NewSpecies(20, 1, 1, 6, 20)
	inputs, outputs := genTraingingData(2000)

	// Train the species
	for i := 0; i < 1000; i++ {
		if i%10 == 0 {
			fmt.Printf(".")
		}
		species.Compete(inputs, outputs)
		species.Breed(3, 3, 1)
	}
	fmt.Println()

	// Print the results for the winner
	species[0].Net.PrintNet()
	inputs, outputs = genTraingingData(1000)
	for i, input := range inputs {
		output := species[0].Net.Update(input)[0]
		fmt.Printf("Input: %f, Output: %f, Answer: %f Difference: %f%%\n", input[0], output, outputs[i][0], output/outputs[i][0])
	}

}
