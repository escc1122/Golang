package binder

import (
	"dig_test/internal/utils"
	"go.uber.org/dig"
	"sync"
)

var (
	binder *dig.Container
	once   sync.Once
)

func GetDigInstance() *dig.Container {
	once.Do(func() {
		binder = dig.New()
		binder.Provide(utils.GetLog)

	})
	return binder
}
