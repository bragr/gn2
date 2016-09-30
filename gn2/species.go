// species.go Copyright (c) 2016 Grant Brady
// Licensed under the MIT License
// See LICENSE.TXT
package gn2

import ()

type chromosome struct {
	Fitness float64
	Net     neuralNet
}
type Species []chromosome

func (s Species) Len() int {
	return len(s)
}
func (s Species) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Species) Less(i, j int) bool {
	return s[i].Fitness < s[j].Fitness
}
