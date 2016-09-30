# Go Neural Network (GN2)
Simple neural net and genetic algorithm implemented in Go. This is not a serious project
but it may be helpful to you if you are just getting started with neural networks (as I am).

## Todo (in no particular order)
* Move neural network code into its own package.
* Comment all the things (hopefully the source will be able to act a mini neural network tutorial
* Add support for backpropagation in addition to genetic algorithm
* Add support for executing multiple networks simultaneously using goroutines
* Create an interface for training sets/problems to allow new ones to be created more simply
* Hardware acceleration? (Possible, but unlikely in the near term. There does seem to be some CUDA support for Go)
* Move params and magic numbers into one place and make them easily configurable by command lines args or config files 
