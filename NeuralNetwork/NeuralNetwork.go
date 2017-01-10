package NeuralNetwork

import (
	"math"
)

// NeuralNetwork provides a Neural Network type.
// This neural network is intended to be evolved, so it does not have a learning
// rate or threshold as would be expected from a typical backpropogated neural
// network.
type NeuralNetwork struct {
	num_inputs  int
	num_hiddens int
	num_outputs int

	activations_hidden []chan float64
	activations_output []chan float64

	hidden_outputs []float64

	layer_hidden []neuron
	layer_output []neuron
}

// Fire() feeds a given input array into the network and activates the neurons.
// It returns the output of the network as an array of floats.
// The input array must match the input size of the network.
// The output array will match the output size of the network.
func (net NeuralNetwork) Fire(inputs []float64) []float64 {
	// Launch a goroutine for each hidden neuron to calculate the output.
	for i := 0; i < net.num_hiddens; i++ {
		go func() {
			net.activations_hidden[i] <- net.layer_hidden[i].Activate(inputs)
		}()
	}

	// Make a new slice to store the final outputs.
	// This slice can't be reused like the hidden outputs because we need
	// to pass by value.
	final_outputs := make([]float64, net.num_outputs)

	// Receive values from earlier goroutines.
	for j := 0; j < net.num_hiddens; j++ {
		net.hidden_outputs[j] = <-net.activations_hidden[j]
	}

	// Launch a goroutine for each output neuron to calculate the final output.
	for i := 0; i < net.num_outputs; i++ {
		go func() {
			net.activations_output[i] <- net.layer_output[i].Activate(net.hidden_outputs)
		}()
	}

	// Receive values from output goroutines.
	for j := 0; j < net.num_outputs; j++ {
		final_outputs[j] = <-net.activations_output[j]
	}

	return final_outputs

}

// Given a number of inputs, an array of inputs, and an array of weight values,
// calculate the output of a single neuron. This function works regardless of
// whether the neuron is in a hidden layer or the output layer.
// This version of activate_neuron has been replaced by the method in Neuron.go
// TODO Consider for deletion.
func activate_neuron(num_inputs int, neuron_inputs []float64, neuron_weights []float64) float64 {
	weighted_input := 0.0
	// Add each input value multiplied by the weight associated to calculate
	// the sum weighted input for this neuron.
	for i := 0; i < num_inputs; i++ {
		weighted_input += neuron_inputs[i] * neuron_weights[i]
	}
	// Using the inverse tangent as a sigmoid function
	return math.Tanh(weighted_input)
}
