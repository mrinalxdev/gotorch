package core

import (
	"sync"

)

type Tensor struct {
	Data []float64
	Shape []int
	Grad []float64
	RequiredGrad bool
	ID int64
	Op *Operation
}

type Operation struct {
	Forward func([]*Tensor) *Tensor
	Backward func(*Tensor, []*Tensor) []*Tensor
	Inputs []*Tensor
	Name string
}

func NewTensor(data []float64, shape []int, requiredGrad bool) *Tensor {
	if len(data) != product(shape) {
		panic("Data size does not match shape")
	}

	return &Tensor {
		Data : data,
		Shape : shape,
		Grad : make([]float64, len(data)),
		RequiredGrad: requiredGrad,
		ID: generateID(),
	}
}

var idCounter int64
var idMutex sync.Mutex

func generateID() int64 {
	idMutex.Lock()
	defer idMutex.Unlock()
	idCounter ++
	return idCounter
}

func product(shape []int) int {
	result := 1
	for _, dim := range shape {
		result *= dim
	}
	return result
}