package goarray

func (a *Array[T]) Reverse() Array[T] {
	output := make([]T, 0)

	for i := len(a.data) - 1; i >= 0; i-- {
		output = append(output, a.data[i])
	}

	return New(&output)
}
