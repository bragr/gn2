package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
)

type (
	neuron    []float64
	nLayer    []neuron
	neuralNet []nLayer
)

func randWeight() float64 {
	return rand.Float64() - rand.Float64()
}

func NewNeuron(numInputs int64) neuron {
	var n neuron
	for i := int64(0); i < numInputs+1; i++ {
		n = append(n, randWeight())
	}
	return n
}

func NewNeuronLayer(numNeurons int64, numInputs int64) nLayer {
	var layer nLayer
	for i := int64(0); i < numNeurons; i++ {
		layer = append(layer, NewNeuron(numInputs))
	}
	return layer
}

func NewNeuralNet(numInputs int64, numOutputs int64, numHiddenLayers int64, numNeuronsPerLayer int64) neuralNet {
	var net neuralNet
	if numHiddenLayers > 0 {
		net = append(net, NewNeuronLayer(numNeuronsPerLayer, numInputs))
		for i := int64(0); i < numHiddenLayers-1; i++ {
			net = append(net, NewNeuronLayer(numNeuronsPerLayer, numNeuronsPerLayer))
		}
		net = append(net, NewNeuronLayer(numOutputs, numNeuronsPerLayer))
	} else {
		net = append(net, NewNeuronLayer(numOutputs, numInputs))
	}
	return net
}

func (net neuralNet) getWeights() []float64 {
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

func (net neuralNet) setWeights(newWeights []float64) {
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

func (net neuralNet) printNet() {
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

func sigmoid(input float64) float64 {
	p := 1.0
	return 1.0 / (1.0 + math.Exp(-1*input/p))
}

func (n neuron) bias() float64 {
	return n[len(n)-1]
}

func (n neuron) updateNeuron(inputs []float64) float64 {
	var output float64 = 0.0
	for i := 0; i < len(n)-1; i++ {
		output += n[i] * inputs[i]
	}
	output -= n.bias()
	return sigmoid(output)
}

func (layer nLayer) updateLayer(inputs []float64) []float64 {
	var output []float64
	for _, n := range layer {
		output = append(output, n.updateNeuron(inputs))
	}
	return output
}

func (net neuralNet) update(inputs []float64) []float64 {
	var outputs []float64

	for _, layer := range net {
		outputs = layer.updateLayer(inputs)
		inputs = outputs
	}

	return outputs
}

func (net neuralNet) mutate(mutationRate, maxPerturbation float64, numInputs, numOutputs, numHiddenLayers, numNeuronsPerLayer int64) neuralNet {
	mutatedNet := NewNeuralNet(numInputs, numOutputs, numHiddenLayers, numNeuronsPerLayer)

	genome := net.getWeights()

	for i, _ := range genome {
		if rand.Float64() < mutationRate {
			genome[i] += randWeight() * maxPerturbation
		}
	}
	mutatedNet.setWeights(genome)

	return mutatedNet
}

func seedRand() {
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

type chromosome struct {
	Fitness float64
	Net     neuralNet
}
type Species []chromosome

func (s Species) Len() int {
	return len(s)
}
func (s Species) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Species) Less(i, j int) bool {
	return s[i].Fitness < s[j].Fitness
}

func main() {
	seedRand()

	var inputs [11]float64
	var answers [11]float64
	for i := 0; i <= 10; i++ {
		x := float64(i) / 10.0
		inputs[i] = x
		answers[i] = 1 - x
	}

	species := make(Species, 20)
	for i, _ := range species {
		species[i].Fitness = 0.0
		species[i].Net = NewNeuralNet(1, 1, 3, 10)
	}

	for i := 0; i < 100000; i++ {
		fmt.Printf(".")
		for c, _ := range species {
			for y, input := range inputs {
				species[c].Fitness += math.Abs(answers[y] - species[c].Net.update([]float64{input})[0])
			}
		}

		// Get best 5 nets, make 3 mutated children for each, and start again
		sort.Sort(species)
		for chromo := 0; chromo < 5; chromo++ {
			species[chromo].Fitness = 0.0
			species[3*chromo+5].Fitness = 0.0
			species[3*chromo+5].Net = species[chromo].Net.mutate(0.025, 0.3, 1, 1, 3, 10)
			species[3*chromo+6].Fitness = 0.0
			species[3*chromo+6].Net = species[chromo].Net.mutate(0.05, 0.3, 1, 1, 3, 10)
			species[3*chromo+7].Fitness = 0.0
			species[3*chromo+7].Net = species[chromo].Net.mutate(0.1, 0.3, 1, 1, 3, 10)
		}
	}
	fmt.Println()

	species[0].Net.printNet()
	for i, input := range inputs {
		fmt.Printf("Input: %f, Output: %f, Answer: %f\n", input, species[0].Net.update([]float64{input})[0], answers[i])
	}

}
