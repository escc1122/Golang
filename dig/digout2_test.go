package main

import (
	"fmt"
	"go.uber.org/dig"
	"testing"
)

func Test_test4(t *testing.T) {
	container := dig.New()
	container.Provide(getCar4)

	err := container.Invoke(test4)

	if err != nil {
		fmt.Println(err)
	}
}
