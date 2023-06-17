package main

import (
	"fmt"
	"go.uber.org/dig"
	"testing"
)

func Test_test2(t *testing.T) {
	container := dig.New()
	container.Provide(getCarA)
	container.Provide(getCarB)

	err := container.Invoke(test2)

	if err != nil {
		fmt.Println(err)
	}
}
