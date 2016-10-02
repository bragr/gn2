// species.go Copyright (c) 2016 Grant Brady
// Licensed under the MIT License
// See LICENSE.TXT
package gn2

import (
	"math"
)

// An individual neural net and its fitness
type genome struct {
	Fitness            float64
	Net                NeuralNet
	netInputs          int64
	netOutputs         int64
	netHiddenLayers    int64
	netNeuronsPerLayer int64
}

// Genome factory
func newGenome(inputs, outputs, layers, neuronsPerLayer int64) genome {
	return genome{
		Fitness:            0.0,
		Net:                NewNeuralNet(inputs, outputs, layers, neuronsPerLayer),
		netInputs:          inputs,
		netOutputs:         outputs,
		netHiddenLayers:    layers,
		netNeuronsPerLayer: neuronsPerLayer}
}

// A species consists of a set of genomes
type Species []genome

// Species factory
func NewSpecies(population, inputs, outputs, layers, neuronsPerLayer int64) Species {
	var species Species
	for i := int64(0); i < population; i++ {
		species = append(species, newGenome(inputs, outputs, layers, neuronsPerLayer))
	}
	return species
}

// Let(), Swap(), and Less() implement the sort prototype to sort based on
// fitness
func (s Species) Len() int {
	return len(s)
}
func (s Species) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Species) Less(i, j int) bool {
	return s[i].Fitness < s[j].Fitness
}

// Feed the given input to all genomes in the species and judge their fitness
// against the answers
func (s Species) Compete(inputs, answers [][]float64) {
	for g, _ := range s {
		for i, input := range inputs {
			outputs := s[g].Net.Update(input)
			for j, _ := range outputs {
				s[g].Fitness += math.Abs(answers[i][j] - outputs[j])
			}
		}
	}
}
