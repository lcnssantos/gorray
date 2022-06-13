package goarray

type Element interface {
	interface{}
}

type Array[T Element] struct {
	data []T
}

func New[T Element](data *[]T) Array[T] {
	if data != nil {
		return Array[T]{
			data: *data,
		}
	}

	return Array[T]{
		data: []T{},
	}
}
