// species.go Copyright (c) 2016 Grant Brady
// Licensed under the MIT License
// See LICENSE.TXT
package gn2

import ()

// An individual neurala and its fitness
type genome struct {
	Fitness float64
	Net     neuralNet
}

// A species consists of a set of genomes
type Species []genome

// Let(), Swap(), and Less() implement the sort prototype to sort based on
// fitness
func (s Species) Len() int {
	return len(s)
}
func (s Species) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Species) Less(i, j int) bool {
	return s[i].Fitness < s[j].Fitness
}
