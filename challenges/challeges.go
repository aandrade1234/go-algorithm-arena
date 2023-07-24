package challenges

import (
	"errors"
	"fmt"
)

type Challenge interface {
	GetInputs() []any
	GetOutputs() []any
	Execute() error
	GetName() string
}

type BaseChallenge struct {
}

func (b *BaseChallenge) BaseExecute(solutionName string, inputs, expectedOutputs []any, solution func(any) any) error {
	var err error
	for i, input := range inputs {
		output := solution(input)
		if output != expectedOutputs[i] {
			err = errors.Join(err, fmt.Errorf("%s: wrong answer for input %v expected %v recived %v: %w",
				solutionName, input, expectedOutputs[i], output, ErrWrongAnswer))
		}
	}
	return err
}
