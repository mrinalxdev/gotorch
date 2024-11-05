package ops

import (
	"math"

	"github.com/mrinalxdev/gotorch/core"
)

type Activation struct {
	engine *core.AutogradEngine
}

func NewActivation() *Activation {
	return &Activation {
		engine : core.NewAutogradEngine(),
	}
}


func (a *Activation) ReLU(x *core.Tensor) *core.Tensor {
	result := make([]float64, len(x.Data))
	for i, v := range x.Data {
		result[i] = math.Max(0, v)
	}

	op := &core.Operation{
		Forward: func(inputs []*core.Tensor) *core.Tensor {
			return core.NewTensor(result, x.Shape, true)
		},
		Backward: func(grad *core.Tensor, inputs []*core.Tensor) []*core.Tensor {
			inputGrad := make([]float64, len(inputs[0].Data))
			for i := range inputGrad {
				if inputs[0].Data[i] > 0 {
					inputGrad[i] = grad.Data[i]
				}
			}
			return []*core.Tensor{core.NewTensor(inputGrad, inputs[0].Shape, true)}
		},
		Inputs: []*core.Tensor{x},
		Name:   "ReLU",
	}

	output := op.Forward([]*core.Tensor{x})
	output.Op = op
	return output
}