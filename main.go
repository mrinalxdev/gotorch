package main

import (
	"fmt"

	"github.com/mrinalxdev/gotorch/core"
	"github.com/mrinalxdev/gotorch/ops"
)


func main(){
	inputData := []float64{1.0, 2.0, 3.0, 4.0}
	weightsData := []float64{0.1, 0.2, 0.3, 0.4}

	input := core.NewTensor(inputData, []int{2, 2}, true)
	weights := core.NewTensor(weightsData, []int{2, 2}, true)

	matOps := ops.NewMatrixOps()
	activation := ops.NewActivation()

	hidden := matOps.MatMul(input, weights)
	output := activation.ReLU(hidden)
	optimizer := core.NewSGD(0.01, 0.9)

	engine := core.NewAutogradEngine()
	engine.Backward(output)

	optimizer.Step([]*core.Tensor{weights})
	optimizer.ZeroGrad([]*core.Tensor{weights})
	fmt.Println("Output tensor :", output.Data)

}