// main.go Copyright (c) 2016 Grant Brady
// Licensed under the MIT License. See LICENSE.TXT
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bragr/gn2/gn2"
	"io/ioutil"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dataFile := flag.String("data", "training.json", "Training data")
	numTests := flag.Int("tests", 1000, "Number of tests")
	population := flag.Int64("pop", 20, "Population")
	layers := flag.Int64("layers", 3, "Number of layers")
	width := flag.Int64("width", 10, "Width of layers")
	survivors := flag.Int64("survivors", 3, "Survivors per generation")
	mutants := flag.Int64("mutants", 3, "Mutants per survivor")
	children := flag.Int64("children", 2, "Width of layers")
	outputRate := flag.Int("output-rate", 10, "How many generations between updates")
	outputFitness := flag.Bool("output-fitness", false, "Output the current fitness")
	outputBest := flag.Bool("output-best", false, "a bool")
	flag.Parse()

	jsonData, err := ioutil.ReadFile(*dataFile)
	check(err)
	var data gn2.TrainingData
	err = json.Unmarshal(jsonData, &data)
	check(err)

	// Make a new "species" to train on
	inputSize := int64(len(data.Input[0]))
	outputSize := int64(len(data.Output[0]))
	species := gn2.NewSpecies(*population, inputSize, outputSize, *layers, *width)

	// Train the species
	for i := 0; i < *numTests; i++ {
		species.Compete(data.Input, data.Output)
		if i%*outputRate == 0 {
			if *outputFitness {
				fmt.Printf("Fitness: %f\n", species[0].Fitness)
			} else {
				fmt.Printf(".")
			}
		}
		species.Breed(*survivors, *mutants, *children)
	}
	sort.Sort(species)
	if *outputBest {
		// Print the results for the winner
		fmt.Println()
		species[0].Net.PrintNet()
	}
	for i, input := range data.Input {
		output := species[0].Net.Update(input)
		difference := make([]float64, len(output), len(output))
		for j, _ := range output {
			difference[j] = output[j] - data.Output[i][j]
		}

		fmt.Printf("[%d] Output:\t%.4v\tDifference: %.4v\n", i, output, difference)
		fmt.Printf("[%d] Answer:\t%.4v\tInput:\t\t%v\n\n", i, data.Output[i], input)
	}

}
