package main

import (
	"fmt"
	"go.uber.org/dig"
	"testing"
)

func Test_test3(t *testing.T) {
	container := dig.New()
	container.Provide(getCar)

	err := container.Invoke(test3)

	if err != nil {
		fmt.Println(err)
	}
}
