package classifiers

// Classifier is an interface for classifiers that can be used with the neuroevolution
// algorithm. Primarily, it is used for neural networks. It could be implemented
// by a single hidden layer network, a perceptron, or a deep network.
//
// Other classifiers may implement this interface as well as long as they can be
// mated to produce new agents.
type Classifier interface {
	Classify(inputs []float64) []float64 // All classifiers must be capable of producing
	// predictions as a list of probabilities of each output, given an input seequence

	Mate(other Classifier) Classifier // All agents must be able to mate with other agents
	// of the same type for the evolutionary algorithm

	// Agents must be able to save some sort of learned data to JSON
	SaveJSON(filepath string) error
}
