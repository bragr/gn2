# Go Neural Network (GN2)

Simple neural net and genetic algorithm implemented in Go. This is not a serious project
but it may be helpful to you if you are just getting started with neural networks (as I am).

## Todo (in no particular order)
* Add support for backpropagation in addition to genetic algorithm
* Create an interface for training sets/problems to allow new ones to be created more simply
* Hardware acceleration? (Possible, but unlikely in the near term. There does seem to be some CUDA support for Go)
* Move params and magic numbers into one place and make them easily configurable by command lines args or config files 
* Add support for loading and saving nets (and possible whole species) to files
 * Either in some sort of json format or maybe Python pickling depending on what works best
