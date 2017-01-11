package evolution

type Agent interface {
	Predict(inputs []float64) []float64
	RandomAgent(num_inputs, num_hiddens, num_outputs int) Agent
	Mate(other Agent) Agent
}
