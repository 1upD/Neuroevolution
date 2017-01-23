package classifiers

import (
	"encoding/json"
	"math/rand"
	"os"
)

// NeuralNetwork provides a Neural Network type.
// This neural network is intended to be evolved, so it does not have a learning
// rate or threshold as would be expected from a typical backpropogated neural
// network.
type neuralNetwork struct {
	Num_inputs  int
	Num_hiddens int
	Num_outputs int

	activations_hidden []chan float64
	activations_output []chan float64

	hidden_outputs []float64

	Layer_hidden []*neuron
	Layer_output []*neuron
}

// Classify() feeds a given input array into the network and activates the neurons.
// It returns the output of the network as an array of floats.
// The input array must match the input size of the network.
// The output array will match the output size of the network.
func (net neuralNetwork) Classify(inputs []float64) []float64 {
	// Launch a goroutine for each hidden neuron to calculate the output.
	for i := 0; i < net.Num_hiddens; i++ { // TODO Why is this minus one!?
		index := i
		go func() {
			//			fmt.Println("Preparing to send hidden neuron ", index)
			net.activations_hidden[index] <- net.Layer_hidden[index].Activate(inputs)
			//			fmt.Println("Preparing to send hidden neuron ", index)
		}()
	}

	// Make a new slice to store the final outputs.
	// This slice can't be reused like the hidden outputs because we need
	// to pass by value.
	final_outputs := make([]float64, net.Num_outputs)

	// Receive values from earlier goroutines.
	for j := 0; j < net.Num_hiddens; j++ {
		//		fmt.Println("Preparing to receive hidden neuron ", j)
		net.hidden_outputs[j] = <-net.activations_hidden[j]
		//		fmt.Println("Received hidden neuron ", j)
	}

	// Launch a goroutine for each output neuron to calculate the final output.
	for i := 0; i < net.Num_outputs; i++ { // TODO Why is this minus one!?
		index := i
		go func() {
			net.activations_output[index] <- net.Layer_output[index].Activate(net.hidden_outputs)
		}()
	}

	// Receive values from output goroutines.
	for j := 0; j < net.Num_outputs; j++ { // TODO Why is this minus one!?
		final_outputs[j] = <-net.activations_output[j]
	}

	return final_outputs

}

// Generate a single hidden layer neural network with randomly assigned weights.
// This will be used at the beginning of an evolutionary algorithm to randomly
// seed the population.
func RandomNetwork(num_inputs, Num_hiddens, Num_outputs int) Classifier {
	n := new(neuralNetwork)
	n.Num_inputs = num_inputs
	n.Num_hiddens = Num_hiddens
	n.Num_outputs = Num_outputs

	n.activations_hidden = make([]chan float64, Num_hiddens)
	n.activations_output = make([]chan float64, Num_outputs)

	n.hidden_outputs = make([]float64, Num_hiddens)

	n.Layer_hidden = make([]*neuron, Num_hiddens)
	n.Layer_output = make([]*neuron, Num_outputs)

	// Initialize hidden activation channels and hidden neurons
	for i := 0; i < Num_hiddens; i++ {
		n.activations_hidden[i] = make(chan float64)
		n.Layer_hidden[i] = RandomNeuron(num_inputs)
		n.hidden_outputs[i] = 0.0
	}

	// Initialize output activation channels and output neurons
	for i := 0; i < Num_outputs; i++ {
		n.activations_output[i] = make(chan float64)
		n.Layer_output[i] = RandomNeuron(Num_hiddens)
	}

	return n
}

// Implements the Classifier interface. Used by evolutionary algorithms to mate two
// networks together.
//
// Other must also be this type of neural network for this to work!
func (n neuralNetwork) Mate(other Classifier) Classifier {
	o := other.(*neuralNetwork)
	return mate(&n, o)
}

// When two neural networks love each other very much...
func mate(p1 *neuralNetwork, p2 *neuralNetwork) *neuralNetwork {
	n := new(neuralNetwork)
	n.Num_inputs = p1.Num_inputs
	n.Num_hiddens = p1.Num_hiddens
	n.Num_outputs = p1.Num_outputs

	n.activations_hidden = make([]chan float64, p1.Num_hiddens)
	n.activations_output = make([]chan float64, p1.Num_outputs)

	n.hidden_outputs = make([]float64, p1.Num_hiddens)

	n.Layer_hidden = make([]*neuron, p1.Num_hiddens)
	n.Layer_output = make([]*neuron, p1.Num_outputs)

	// Initialize hidden activation channels and hidden neurons
	for i := 0; i < p1.Num_hiddens; i++ {
		n.activations_hidden[i] = make(chan float64)
		n.hidden_outputs[i] = 0.0

		// Get a random float to determine the parent of this gene
		randomFloat := rand.Float64()
		// If the random value is less than 0.5, choose the mother
		if randomFloat < 0.5 {
			n.Layer_hidden[i] = CopyNeuron(p1.Layer_hidden[i])
			// If the random value is greater than 0.5 but less than 0.995, choose
			// the father.
		} else {
			n.Layer_hidden[i] = CopyNeuron(p2.Layer_hidden[i])
		}
	}

	// Initialize output activation channels and output neurons
	for i := 0; i < p1.Num_outputs; i++ {
		n.activations_output[i] = make(chan float64)

		// Get a random float to determine the parent of this gene
		randomFloat := rand.Float64()
		// If the random value is less than 0.5, choose the mother
		if randomFloat < 0.5 {
			n.Layer_output[i] = CopyNeuron(p1.Layer_output[i])
			// If the random value is greater than 0.5 but less than 0.995, choose
			// the father.
		} else {
			n.Layer_output[i] = CopyNeuron(p2.Layer_output[i])
		}
	}

	// Check for mutation
	randomFloat := rand.Float64()
	if randomFloat > 0.8 {
		gene := rand.Intn(n.Num_hiddens + n.Num_outputs)
		if gene >= n.Num_hiddens {
			gene -= n.Num_hiddens
			n.Layer_output[gene] = RandomNeuron(n.Num_hiddens)
		} else {
			n.Layer_hidden[gene] = RandomNeuron(n.Num_inputs)
		}

	}

	return n

}

// Encodes this network as a JSON file.
func (n *neuralNetwork) SaveJSON(filepath string) error {
	fi, err := os.Open(filepath)
	if err != nil {
		fi, err = os.Create(filepath)
		if err != nil {
			return err
		}
	}
	enc := json.NewEncoder(fi)

	return enc.Encode(n)
}

// Decodes a new network from a JSON file
func LoadJSON(filepath string) (*neuralNetwork, error) {
	fi, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	enc := json.NewDecoder(fi)

	n := neuralNetwork{}

	err = enc.Decode(&n)

	n.activations_hidden = make([]chan float64, n.Num_hiddens)
	n.activations_output = make([]chan float64, n.Num_outputs)

	n.hidden_outputs = make([]float64, n.Num_hiddens)

	// Initialize hidden activation channels
	for i := 0; i < n.Num_hiddens; i++ {
		n.activations_hidden[i] = make(chan float64)
		n.hidden_outputs[i] = 0.0
	}

	// Initialize output activation channels and output neurons
	for i := 0; i < n.Num_outputs; i++ {
		n.activations_output[i] = make(chan float64)
	}

	return &n, err
}
