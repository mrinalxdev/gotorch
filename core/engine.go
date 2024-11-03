package core

import "sync"

type AutogradEngine struct {
	graph       *ComputationalGraph
	tapeEnabled bool
}

type ComputationalGraph struct {
	Nodes       map[int64]*Tensor
	mutex       sync.RWMutex
	deviceMnger *DeviceManger
}

type DeviceManger struct {
	currentDevice string
	gpuAvailable  bool
	gpuMemory     map[int64]interface{}
}

func NewAutogradEngine() *AutogradEngine {
	return &AutogradEngine{
		graph: &ComputationalGraph{
			Nodes: make(map[int64]*Tensor),
			deviceMnger: &DeviceManger{
				currentDevice: "cpu",
				gpuMemory:     make(map[int64]interface{}),
			},
		},
		tapeEnabled: true,
	}
}

func (e *AutogradEngine) Backward(tensor *Tensor) {
	if !tensor.RequiredGrad {
		return
	}
	visited := make(map[int]bool)
	queue := []*Tensor{tensor}

	for i := range tensor.Grad {
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
			for i, input := range current.Op.Inputs {
				for j := range input.Grad {
					input.Grad[j] += grads[i].Data[j]
				}
				queue = append(queue, input)
			}
		}
	}
}

type Optimizer interface {
	Step([]*Tensor)
	ZeroGrad([]*Tensor)
}

type SGD struct {
	LearningRate float64
	Momentum     float64
	Velocities   map[int64][]float64
}

func NewSGD(lr, momentum float64) *SGD {
	return &SGD{
		LearningRate: lr,
		Momentum:     momentum,
		Velocities:   make(map[int64][]float64),
	}
}

func (sgd *SGD) Step(parameters []*Tensor) {
	for _, param := range parameters {
		if !param.RequiredGrad {
			continue
		}

		velocity, exists := sgd.Velocities[param.ID]
		if !exists {
			velocity = make([]float64, len(param.Data))
			sgd.Velocities[param.ID] = velocity
		}

		for i := range param.Data {
			velocity[i] = sgd.Momentum*velocity[i] - sgd.LearningRate*param.Grad[i]
			param.Data[i] += velocity[i]
		}
	}
}

func (sgd *SGD) ZeroGrad(parameters []*Tensor){
	for _, param := range parameters {
		for i := range param.Grad {
			param.Grad[i] = 0
		}
	}
}
