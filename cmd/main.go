package main

import (
	"fmt"

	"github.com/aandrade1234/go-algorithm-arena/algorithms"
	challengespkg "github.com/aandrade1234/go-algorithm-arena/challenges"
)

func main() {
	challenges := []challengespkg.Challenge{
		&algorithms.BirthdayCandles{},
		&algorithms.MinMaxSum{},
		&algorithms.LongestWord{},
	}

	for _, challenge := range challenges {
		result := "Success"
		if err := challenge.Execute(); err != nil {
			result = err.Error()
		}

		fmt.Printf("%s:%v\n", challenge.GetName(), result)
	}

}
