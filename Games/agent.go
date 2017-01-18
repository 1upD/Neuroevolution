package games

// Agent is an interface for classifiers that can be used with the neuroevolution
// algorithm. Primarily, it is used for neural networks. It could be implemented
// by a single hidden layer network, a perceptron, or a deep network.
//
// Other classifiers may implement this interface as well as long as they can be
// mated to produce new agents.
type Agent interface {
	Predict(inputs []float64) []float64 // All agents must be capable of producing
	// predictions as a list of probabilities of each output, given an input seequence

	Mate(other Agent) Agent // All agents must be able to mate with other agents
	// of the same type for the evolutionary algorithm
}
