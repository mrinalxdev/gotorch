package ops

import (
	"sync"

	"github.com/mrinalxdev/gotorch/core"
)

type MatrixOps struct {
	engine *core.AutogradEngine
}

func NewMatrixOps() *MatrixOps {
	return &MatrixOps{
		engine: core.NewAutogradEngine(),
	}
}

func (m *MatrixOps) MatMul(a, b *core.Tensor) *core.Tensor {
	if len(a.Shape) != 2 || len(b.Shape) != 2 {
		panic("MatMul requires 2D tensors")
	}

	if a.Shape[1] != b.Shape[0] {
		panic("Invalid dimensions for matrix multiplication")
	}

	rows, cols := a.Shape[0], b.Shape[1]
	result := make([]float64, rows*cols)

	var wg sync.WaitGroup
	for i := 0; i < rows; i++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			for j := 0; j < cols; j++ {
				sum := 0.0
				for k := 0; k < a.Shape[1]; k++ {
					sum += a.Data[row*a.Shape[1]+k] * b.Data[k*cols+j]
				}
				result[row*cols+j] = sum
			}
		}(i)
	}
	wg.Wait()

	op := &core.Operation{
		Forward: func(inputs []*core.Tensor) *core.Tensor {
			return core.NewTensor(result, []int{rows, cols}, true)
		},

		Backward: func(grad *core.Tensor, inputs []*core.Tensor) []*core.Tensor {
			aGrad := make([]float64, len(inputs[0].Data))
			bGrad := make([]float64, len(inputs[1].Data))

			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					for k := 0; k < a.Shape[1]; k++ {
						aGrad[i*a.Shape[1]+k] += grad.Data[i*cols+j] * b.Data[k*cols+j]
						bGrad[k*cols+j] += grad.Data[i*cols+j] * a.Data[i*a.Shape[1]+k]
					}
				}
			}

			return []*core.Tensor{
				core.NewTensor(aGrad, a.Shape, true),
				core.NewTensor(bGrad, b.Shape, true),
			}
		},

		Inputs: []*core.Tensor{a, b},
		Name:   "MatMul",
	}

	output := op.Forward([]*core.Tensor{a, b})
	output.Op = op
	return output
}
