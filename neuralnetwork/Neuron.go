package neuralnetwork

import (
	"math"
	"math/rand"
)

type neuron struct {
	// Number of inputs into this particle neuron.
	// In a neural network, a neuron in the hidden layer has the same number of
	// inputs as the number of inputs to the network, whlie a neuron in the
	// output layer has the same number of inputs as the number of hidden neurons
	// per layer.
	Num_inputs int

	// This variable is used to store the sum of weighted inputs.
	Weighted_input float64

	// This slice contains the weights of this particular neuron.
	// The slice must be the same size as the number of inputs to the neuron.
	Weights []float64
}

// Generates a neuron with random weights between -1 and 1.
// Useful for generating random neural networks or mutants.
// TODO Should there be a max random weight?
func RandomNeuron(num_inputs int) *neuron {
	n := new(neuron)
	n.Num_inputs = num_inputs
	n.Weighted_input = 0
	n.Weights = make([]float64, num_inputs)
	for i := 0; i < num_inputs; i++ {
		n.Weights[i] = (rand.Float64() * 2.0) - 1.0
	}
	return n

}

// Makes an identical copy of a neuron.
// Useful for mating.
func CopyNeuron(n *neuron) *neuron {
	newNeuron := new(neuron)
	newNeuron.Num_inputs = n.Num_inputs
	newNeuron.Weighted_input = 0
	newNeuron.Weights = make([]float64, newNeuron.Num_inputs)
	//copy(n.weights, newNeuron.weights)
	copy(newNeuron.Weights, n.Weights)
	return newNeuron
}

// Given an array of inputs, calculate the output of a single neuron. This
// function works for any neuron.
func (n neuron) Activate(neuron_inputs []float64) float64 {
	n.Weighted_input = 0.0
	// Add each input value multiplied by the weight associated to calculate
	// the sum weighted input for this neuron.
	for i := 0; i < n.Num_inputs; i++ {
		n.Weighted_input += neuron_inputs[i] * n.Weights[i]
	}
	// Using the hyperbolic tangent as a sigmoid activation function
	return math.Tanh(n.Weighted_input)
}
