package main

import (
	"encoding/binary"
	"fmt"
    "math"
	"math/rand"
	"os"
)

type (
    neuron []float64
    nLayer []neuron
    neuralNet []nLayer
)

func randWeight() float64 {
	return rand.Float64() - rand.Float64()
}

func NewNeuron(numInputs int64) neuron {
	var n neuron
	for i := int64(0); i < numInputs; i++ {
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
			for _, weight := range neuron {
				fmt.Println(weight)
			}
		}
	}
}

func sigmoid(input float64) float64 {
    p := 1.0
    return 1 / ( 1 + math.Exp(-1*input/p))
}

func (n neuron) updateNeuron(inputs []float64) float64 {
    var output float64 = 0.0
    for i, weight := range n {
        output += weight * inputs[i]
    }
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

func main() {
	seedRand()

	net := NewNeuralNet(1, 1, 3, 10)
	net.printNet()

	w := net.getWeights()

	derp := make([]float64, 0)
	for i := 0; i < len(w); i++ {
		derp = append(derp, float64(i))
	}
	net.setWeights(derp)
	net.printNet()
	net.setWeights(w)
	net.printNet()
}
