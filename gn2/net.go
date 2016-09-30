// net.go Copyright (c) 2016 Grant Brady
// Licensed under the MIT License
// See LICENSE.TXT
package gn2

import (
	"fmt"
	"sort"
)

//
// Base neuron type
// The slice consists of n weights. The n+1 element is the neuron bias
type neuron []float64

// Basic neuron factory. Creates a neuron with numInputs input with randomized
// input weights and neuron bias.
func newNeuron(numInputs int64) neuron {
	var n neuron
	// One extra for the bias
	for i := int64(0); i < numInputs+1; i++ {
		n = append(n, randWeight())
	}
	return n
}

// Get the bias of the neuron
func (n neuron) bias() float64 {
	return n[len(n)-1]
}

// Calculate the output of the neuron for a given set of inputs
func (n neuron) updateNeuron(inputs []float64) float64 {
	var output float64 = 0.0
	for i := 0; i < len(n)-1; i++ {
		output += n[i] * inputs[i]
	}
	output -= n.bias()
	return sigmoid(output)
}

//
// layer of neurons
// A simple list of neurons
type nLayer []neuron

// Basic neuron layer factory. Create numNeurons neurons each taking numInputs
// inputs
func newNeuronLayer(numNeurons, numInputs int64) nLayer {
	var layer nLayer
	for i := int64(0); i < numNeurons; i++ {
		layer = append(layer, newNeuron(numInputs))
	}
	return layer
}

// Calculate the output of each neuron in the layer for the given set of inputs
func (layer nLayer) updateLayer(inputs []float64) []float64 {
	var output []float64
	for _, n := range layer {
		output = append(output, n.updateNeuron(inputs))
	}
	return output
}

//
// a full neural net
// A simple list of neuron layers
type NeuralNet []nLayer

// Neural net factory. Creates a net that takes numInputs inputs, numOutputs
// outputs, with numHiddenLayers hidden layers, and numNeuronPerLayer neurons
// in the hidden layers
func NewNeuralNet(numInputs, numOutputs, numHiddenLayers, numNeuronsPerLayer) NeuralNet {
	var net NeuralNet
	if numHiddenLayers > 0 {
		net = append(net, newNeuronLayer(numNeuronsPerLayer, numInputs))
		for i := int64(0); i < numHiddenLayers-1; i++ {
			net = append(net, newNeuronLayer(numNeuronsPerLayer, numNeuronsPerLayer))
		}
		net = append(net, newNeuronLayer(numOutputs, numNeuronsPerLayer))
	} else {
		net = append(net, newNeuronLayer(numOutputs, numInputs))
	}
	return net
}

// Return a list containing all the edge weights and neuron biases based on a
// simple serializations of the nested lists. This does the opposite of
// SetWeights()
func (net NeuralNet) GetWeights() []float64 {
	weights := make([]float64, 0)

	for _, layer := range net {
		for _, neuron := range layer {
			for _, weight := range neuron {
				weights = append(weights, weight)
			}
		}
	}
	return weights
}

// Set the edge weights and neuron biases by a simple deserialization of a list
// weights to the nested list. This does the opposite of GetWeights()
func (net NeuralNet) SetWeights(newWeights []float64) {
	index := 0
	for l, _ := range net {
		for n, _ := range net[l] {
			for e, _ := range net[l][n] {
				net[l][n][e] = newWeights[index]
				index++
			}
		}
	}
}

// Print all the neuron layers, nerons, edge weights, and neuron biases
func (net NeuralNet) PrintNet() {
	for c, layer := range net {
		fmt.Printf("Layer: %d (%d neurons)\n", c, len(layer))
		for i, neuron := range layer {
			fmt.Printf("Neuron: %d\n", i)
			for i, weight := range neuron {
				if i == len(neuron)-1 {
					fmt.Printf("Bias: %f\n", weight)
				} else {
					fmt.Println(weight)
				}
			}
		}
	}
}

// Return the output of a network for a given input
func (net NeuralNet) Update(inputs []float64) []float64 {
	var outputs []float64

	for _, layer := range net {
		outputs = layer.updateLayer(inputs)
		inputs = outputs
	}

	return outputs
}

func (net NeuralNet) Mutate(mutationRate, maxPerturbation float64, numInputs, numOutputs, numHiddenLayers, numNeuronsPerLayer int64) neuralNet {
	mutatedNet := NewNeuralNet(numInputs, numOutputs, numHiddenLayers, numNeuronsPerLayer)

	genome := net.GetWeights()

	for i, _ := range genome {
		if rand.Float64() < mutationRate {
			genome[i] += randWeight() * maxPerturbation
		}
	}
	mutatedNet.SetWeights(genome)

	return mutatedNet
}
