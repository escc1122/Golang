package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := someFunction()
	if err != nil {
		fmt.Printf("錯誤: %+v\n", err)
	}
}

func someFunction() error {
	err := anotherFunction()
	if err != nil {
		return err
		//return errors.Wrap(err, "另一個函數出錯")
	}
	return nil
}

func anotherFunction() error {
	return errors.New("這是一個錯誤")
}
