package mock

import "math/rand"

type Randomizer struct {
	Value       interface{}
	Unvalue     interface{}
	ProbOfValue float64
}

func (receiver *Randomizer) Return() interface{} {
	r := 0 + rand.Float64()*(1.0)

	if r > receiver.ProbOfValue {
		return receiver.Unvalue
	} else {
		return receiver.Value
	}
}
