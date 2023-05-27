package main

type WowStruct[T int | float32, S []T] struct {
	Data     S
	MaxValue T
	MinValue T
}

type valueType interface {
	int32 | float32
}

// https://alankrantas.medium.com/%E7%B0%A1%E5%96%AE%E7%8E%A9-go-1-18-%E6%B3%9B%E5%9E%8B-1d09da07b70
type Data[T valueType] struct {
	data []T
}

func (d *Data[T]) addData(newValues ...T) {
	for _, item := range newValues {
		d.data = append(d.data, item)
	}
}
func (d *Data[T]) sum() T {
	var s T
	for _, item := range d.data {
		s += item
	}
	return s
}
