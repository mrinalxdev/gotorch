package core

import "sync"

type AutogradEngine struct {
	graph *ComputationalGraph
	tapeEnabled bool
}

type ComputationalGraph struct {
	Nodes map[int64] *Tensor
	mutex sync.RWMutex
	deviceMnger *DeviceManger
}

type DeviceManger struct {
	currentDevice string
	gpuAvailable bool
	gpuMemory map[int64] interface{}
}
func NewAutogradEngine() *AutogradEngine {
	return &AutogradEngine{
		graph : &ComputationalGraph{
			Nodes : make(map[int64]*Tensor),
			deviceMnger: &DeviceManger{
				currentDevice: "cpu",
				gpuMemory: make(map[int64]interface{}),
			},
		},
		tapeEnabled: true,
	}
}

func (e *AutogradEngine) Backward(tensor *Tensor){
	if !tensor.RequiredGrad {
		return
	}
	visited := make(map[int]bool)
	queue := []*Tensor{tensor}

	for i := range tensor.Grad{
		tensor.Grad[i] = 1.0
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[int(current.ID)] {
			continue
		}
		visited[int(current.ID)] = true

		if current.Op != nil {
			grads := current.Op.Backward(current, current.Op.Inputs)
			for i, input := range current.Op.Inputs{
				for j := range input.Grad {
					input.Grad[j] += grads[i].Data[j]
				}
				queue = append(queue, input)
			}
		}
	}
}