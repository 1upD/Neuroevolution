package neuralnetwork

import (
	"math"
	"math/rand"
)

// NeuralNetwork provides a Neural Network type.
// This neural network is intended to be evolved, so it does not have a learning
// rate or threshold as would be expected from a typical backpropogated neural
// network.
type neuralNetwork struct {
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
func (net neuralNetwork) Fire(inputs []float64) []float64 {
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

func RandomNetwork(num_inputs, num_hiddens, num_outputs) *neuralNetwork {
	n := new(neuralNetwork)
	n.num_inputs = num_inputs
	n.num_hiddens = num_hiddens
	n.num_outputs = num_outputs

	n.activations_hidden = make([]chan float64, num_hiddens)
	n.activations_ouput = make([]chan float64, num_outputs)

	n.hidden_outputs = make([]float64, num_hiddens)

	n.layer_hidden = make([]neuron, num_hiddens)
	n.layer_output = make([]neuron, num_outputs)

	// Initialize hidden activation channels and hidden neurons
	for i := 0; i < num_hiddens; i++ {
		n.activations_hidden[i] = make(chan float64)
		n.layer_hidden[i] = RandomNeuron(num_inputs)
		n.hidden_outputs[i] = 0.0
	}

	// Initialize output activation channels and output neurons
	for i := 0; i < num_outputs; i++ {
		n.activations_output[i] = make(chan float64)
		n.layer_output[i] = RandomNeuron(num_hiddens)
	}

}

// Generate a single hidden layer neural network with randomly assigned weights.
// This will be used at the beginning of an evolutionary algorithm to randomly
// seed the population.
func RandomNetwork(num_inputs, num_hiddens, num_outputs int) *neuralNetwork {
	n := new(neuralNetwork)
	n.num_inputs = num_inputs
	n.num_hiddens = num_hiddens
	n.num_outputs = num_outputs

	n.activations_hidden = make([]chan float64, num_hiddens)
	n.activations_ouput = make([]chan float64, num_outputs)

	n.hidden_outputs = make([]float64, num_hiddens)

	n.layer_hidden = make([]neuron, num_hiddens)
	n.layer_output = make([]neuron, num_outputs)

	// Initialize hidden activation channels and hidden neurons
	for i := 0; i < num_hiddens; i++ {
		n.activations_hidden[i] = make(chan float64)
		n.layer_hidden[i] = RandomNeuron(num_inputs)
		n.hidden_outputs[i] = 0.0
	}

	// Initialize output activation channels and output neurons
	for i := 0; i < num_outputs; i++ {
		n.activations_output[i] = make(chan float64)
		n.layer_output[i] = RandomNeuron(num_hiddens)
	}

}

// When two neural networks love each other very much...
func Mate(mommy *neuralNetwork, daddy *neuralNetwork) *neuralNetwork {
	n := new(neuralNetwork)
	n.num_inputs = mommy.num_inputs
	n.num_hiddens = mommy.num_hiddens
	n.num_outputs = mommy.num_outputs

	n.activations_hidden = make([]chan float64, mommy.num_hiddens)
	n.activations_ouput = make([]chan float64, mommy.num_outputs)

	n.hidden_outputs = make([]float64, mommy.num_hiddens)

	n.layer_hidden = make([]neuron, mommy.num_hiddens)
	n.layer_output = make([]neuron, mommy.num_outputs)

	// Initialize hidden activation channels and hidden neurons
	for i := 0; i < mommy.num_hiddens; i++ {
		n.activations_hidden[i] = make(chan float64)
		n.hidden_outputs[i] = 0.0

		// Get a random float to determine the parent of this gene
		randomFloat = rand.Float64()
		// If the random value is less than 0.5, choose the mother
		if randomFloat < 0.5(
			n.layer_hidden[i] = CopyNeuron(mommy.layer_hidden[i])
		)
		// If the random value is greater than 0.5 but less than 0.995, choose
		// the father.
		else if randomFloat < 0.995(
			n.layer_hidden[i] = CopyNeuron(daddy.layer_hidden[i])
		)
		// Mutation!
		else randomFloat < 0.99(
			n.layer_hidden[i] = RandomNeuron(mommy.num_inputs)
		)
	}

	// TODO Finish mating function!
	// Initialize output activation channels and output neurons
//	for i := 0; i < num_outputs; i++ {
//		n.activations_output[i] = make(chan float64)
//		n.layer_output[i] = RandomNeuron(num_hiddens)
//	}

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
