package neuralnetwork

import (
	"encoding/json"
	"math/rand"
	"os"

	"github.com/CRRDerek/Neuroevolution/games"
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

	layer_hidden []*neuron
	layer_output []*neuron
}

// Predict() feeds a given input array into the network and activates the neurons.
// It returns the output of the network as an array of floats.
// The input array must match the input size of the network.
// The output array will match the output size of the network.
func (net neuralNetwork) Predict(inputs []float64) []float64 {
	// Launch a goroutine for each hidden neuron to calculate the output.
	for i := 0; i < net.num_hiddens; i++ { // TODO Why is this minus one!?
		index := i
		go func() {
			//			fmt.Println("Preparing to send hidden neuron ", index)
			net.activations_hidden[index] <- net.layer_hidden[index].Activate(inputs)
			//			fmt.Println("Preparing to send hidden neuron ", index)
		}()
	}

	// Make a new slice to store the final outputs.
	// This slice can't be reused like the hidden outputs because we need
	// to pass by value.
	final_outputs := make([]float64, net.num_outputs)

	// Receive values from earlier goroutines.
	for j := 0; j < net.num_hiddens; j++ {
		//		fmt.Println("Preparing to receive hidden neuron ", j)
		net.hidden_outputs[j] = <-net.activations_hidden[j]
		//		fmt.Println("Received hidden neuron ", j)
	}

	// Launch a goroutine for each output neuron to calculate the final output.
	for i := 0; i < net.num_outputs; i++ { // TODO Why is this minus one!?
		index := i
		go func() {
			net.activations_output[index] <- net.layer_output[index].Activate(net.hidden_outputs)
		}()
	}

	// Receive values from output goroutines.
	for j := 0; j < net.num_outputs; j++ { // TODO Why is this minus one!?
		final_outputs[j] = <-net.activations_output[j]
	}

	return final_outputs

}

// Generate a single hidden layer neural network with randomly assigned weights.
// This will be used at the beginning of an evolutionary algorithm to randomly
// seed the population.
func RandomNetwork(num_inputs, num_hiddens, num_outputs int) games.Agent {
	n := new(neuralNetwork)
	n.num_inputs = num_inputs
	n.num_hiddens = num_hiddens
	n.num_outputs = num_outputs

	n.activations_hidden = make([]chan float64, num_hiddens)
	n.activations_output = make([]chan float64, num_outputs)

	n.hidden_outputs = make([]float64, num_hiddens)

	n.layer_hidden = make([]*neuron, num_hiddens)
	n.layer_output = make([]*neuron, num_outputs)

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

	return n
}

// Implements the Agent interface. Used by evolutionary algorithms to mate two
// networks together.
//
// Other must also be this type of neural network for this to work!
func (n neuralNetwork) Mate(other games.Agent) games.Agent {
	o := other.(*neuralNetwork)
	return mate(&n, o)
}

// When two neural networks love each other very much...
func mate(p1 *neuralNetwork, p2 *neuralNetwork) *neuralNetwork {
	n := new(neuralNetwork)
	n.num_inputs = p1.num_inputs
	n.num_hiddens = p1.num_hiddens
	n.num_outputs = p1.num_outputs

	n.activations_hidden = make([]chan float64, p1.num_hiddens)
	n.activations_output = make([]chan float64, p1.num_outputs)

	n.hidden_outputs = make([]float64, p1.num_hiddens)

	n.layer_hidden = make([]*neuron, p1.num_hiddens)
	n.layer_output = make([]*neuron, p1.num_outputs)

	// Initialize hidden activation channels and hidden neurons
	for i := 0; i < p1.num_hiddens; i++ {
		n.activations_hidden[i] = make(chan float64)
		n.hidden_outputs[i] = 0.0

		// Get a random float to determine the parent of this gene
		randomFloat := rand.Float64()
		// If the random value is less than 0.5, choose the mother
		if randomFloat < 0.5 {
			n.layer_hidden[i] = CopyNeuron(p1.layer_hidden[i])
			// If the random value is greater than 0.5 but less than 0.995, choose
			// the father.
		} else {
			n.layer_hidden[i] = CopyNeuron(p2.layer_hidden[i])
		}
	}

	// Initialize output activation channels and output neurons
	for i := 0; i < p1.num_outputs; i++ {
		n.activations_output[i] = make(chan float64)

		// Get a random float to determine the parent of this gene
		randomFloat := rand.Float64()
		// If the random value is less than 0.5, choose the mother
		if randomFloat < 0.5 {
			n.layer_output[i] = CopyNeuron(p1.layer_output[i])
			// If the random value is greater than 0.5 but less than 0.995, choose
			// the father.
		} else {
			n.layer_output[i] = CopyNeuron(p2.layer_output[i])
		}
	}

	// Check for mutation
	randomFloat := rand.Float64()
	if randomFloat > 0.8 {
		gene := rand.Intn(n.num_hiddens + n.num_outputs)
		if gene >= n.num_hiddens {
			gene -= n.num_hiddens
			n.layer_output[gene] = RandomNeuron(n.num_hiddens)
		} else {
			n.layer_hidden[gene] = RandomNeuron(n.num_inputs)
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
	return &n, err
}
