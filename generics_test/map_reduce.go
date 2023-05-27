package main

type MapReduce[T comparable] struct {
	s []T
}

func (m *MapReduce[T]) Map(f func(T) T) *MapReduce[T] {
	r := make([]T, len(m.s))
	for i, v := range m.s {
		r[i] = f(v)
	}
	m.s = r
	return m
}

func (m *MapReduce[T]) Reduce(f func(T, T) T, initializer T) T {
	r := initializer
	for _, v := range m.s {
		r = f(r, v)
	}
	return r
}

func (m *MapReduce[T]) Filter(f func(T) bool) *MapReduce[T] {
	var r []T
	for _, v := range m.s {
		if f(v) {
			r = append(r, v)
		}
	}
	m.s = r
	return m
}

func (m *MapReduce[T]) Get() []T {
	return m.s
}

// Map turns a []T1 to a []T2 using a mapping function.
// This function has two type parameters, T1 and T2.
// There are no constraints on the type parameters,
// so this works with slices of any type.
func Map[T1, T2 comparable](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce reduces a []T1 to a single value using a reduction function.
func Reduce[T1, T2 comparable](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter filters values from a slice using a filter function.
// It returns a new slice with only the elements of s
// for which f returned true.
func Filter[T comparable](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}
