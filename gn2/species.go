// species.go Copyright (c) 2016 Grant Brady
// Licensed under the MIT License
// See LICENSE.TXT
package gn2

import (
	"math"
	"sort"
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

// Choose the survivors of this population. Return the species to full population by creating new
// genomes that are mutant clones of survivors, children (interlaced copies of survivors), and, if
// more genomes are still needed, new random genomes.
func (s Species) Breed(survivors, mutantCopyRate, childRate int64) {
	sort.Sort(s)

	// Survivors
	for g := int64(0); g < survivors && g < int64(len(s)); g++ {
		s[g].Fitness = 0.0
	}
	newPop := survivors

	// Mutant clones
	for g := newPop; g < int64(len(s)) && g < (mutantCopyRate*survivors)+newPop; g++ {
		parent := (g - newPop) / mutantCopyRate
		s[g].Fitness = 0.0
		s[g].Net = s[parent].Net.Mutate(0.05, 0.3, s[g].netInputs, s[g].netOutputs, s[g].netHiddenLayers, s[g].netNeuronsPerLayer)
	}
	newPop += (mutantCopyRate * survivors)

	// Children
	childrenPerParent := make(map[int64]int64)
	for ; newPop < int64(len(s)); newPop++ {
		s[newPop].Fitness = 0.0

		parent1 := int64(-1)
		parent2 := int64(-1)

		// Get the first parent that needs a child
		for p := int64(0); p < survivors; p++ {
			if childrenPerParent[p] < childRate {
				parent1 = p
				childrenPerParent[p] = childrenPerParent[p] + 1
				break
			}
		}
		if parent1 == -1 {
			break
		}

		// Get the second parent that needs a child
		for p := parent1; p < survivors; p++ {
			if childrenPerParent[p] < childRate {
				parent2 = p
				childrenPerParent[p] = childrenPerParent[p] + 1
				break
			}
		}
		if parent2 == -1 {
			break
		}

		net1 := s[parent1].Net.GetWeights()
		net2 := s[parent2].Net.GetWeights()

		for i := int64(1); i < int64(len(net1)); i += 2 {
			net1[i] = net2[i]
		}
		s[newPop].Net.SetWeights(net1)
	}

	// New random nets (filler)
	for ; newPop < int64(len(s)); newPop++ {
		s[newPop].Fitness = 0.0
		s[newPop].Net = NewNeuralNet(s[newPop].netInputs, s[newPop].netOutputs, s[newPop].netHiddenLayers, s[newPop].netNeuronsPerLayer)
	}
}
