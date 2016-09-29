package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
)

type Neuron struct {
	NumInputs   int64
	EdgeWeights []float64
}

type NLayer struct {
	NumNeurons int64
	Neurons    []*Neuron
}

type NNet struct {
	NumInputs       int64
	NumOutputs      int64
	NumLayers       int64
	NumNeuronsLayer int64
	NeuronLayer     []*NLayer
}

func randWeight() float64 {
	return rand.Float64() - rand.Float64()
}

func NewNeuron(numInputs int64) *Neuron {
	nn := new(Neuron)
	nn.NumInputs = numInputs

	for i := int64(0); i < numInputs; i++ {
		nn.EdgeWeights = append(nn.EdgeWeights, randWeight())
	}
	return nn
}

func NewNeuronLayer(numNeurons int64, numInputs int64) *NLayer {
	layer := new(NLayer)
	for i := int64(0); i < numNeurons; i++ {
		layer.Neurons = append(layer.Neurons, NewNeuron(numInputs))
	}
	return layer
}

func NewNeuralNet(numInputs int64, numOutputs int64, numHiddenLayers int64, numNeuronsPerLayer int64) *NNet {
	net := &NNet{NumInputs: numInputs, NumOutputs: numOutputs, NumLayers: numHiddenLayers, NumNeuronsLayer: numNeuronsPerLayer}
	if numHiddenLayers > 0 {
		net.NeuronLayer = append(net.NeuronLayer, NewNeuronLayer(numNeuronsPerLayer, numInputs))
		for i := int64(0); i < numHiddenLayers-1; i++ {
			net.NeuronLayer = append(net.NeuronLayer, NewNeuronLayer(numNeuronsPerLayer, numNeuronsPerLayer))
		}
		net.NeuronLayer = append(net.NeuronLayer, NewNeuronLayer(numOutputs, numNeuronsPerLayer))
	} else {
		net.NeuronLayer = append(net.NeuronLayer, NewNeuronLayer(numOutputs, numInputs))
	}
	return net
}

func (n *NNet) getWeights() []float64 {
	weights := make([]float64, 0)

	for _, layer := range n.NeuronLayer {
		for _, neuron := range layer.Neurons {
			for _, weight := range neuron.EdgeWeights {
				weights = append(weights, weight)
			}
		}
	}
	return weights
}

func (n *NNet) setWeights(newWeights []float64) {
	index := 0
	for x, _ := range n.NeuronLayer {
		for y, _ := range n.NeuronLayer[x].Neurons {
			for z, _ := range n.NeuronLayer[x].Neurons[y].EdgeWeights {
				n.NeuronLayer[x].Neurons[y].EdgeWeights[z] = newWeights[index]
				index++
			}
		}
	}
}

func (n *NNet) printNet() {
	for c, layer := range n.NeuronLayer {
		fmt.Printf("Layer: %d (%d neurons)\n", c, len(layer.Neurons))
		for i, neuron := range layer.Neurons {
			fmt.Printf("Neuron: %d\n", i)
			for _, weight := range neuron.EdgeWeights {
				fmt.Println(weight)
			}
		}
	}
}

func (n *NNet) update(inputs []float64) []float64 {
	outputs := make([]float64, 1)
	if len(inputs) != n.NumInputs {
		return outputs
	}
	for i := int64(0); i < n.NumNeuronsLayer+1; i++ {
		if i > 0 {
			inputs = outputs
		}
		outputs := make([]float64, 1)
		weightIndex := 0
		for j := int64(0); j < len(n.NeuronLayer[i].Neurons); j++ {
		}
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
