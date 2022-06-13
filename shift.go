package goarray

func (a *Array[T]) Shift() T {
	toReturn := a.data[0]

	output := make([]T, len(a.data)-1)

	for i := 1; i < len(a.data); i++ {
		output[i-1] = a.data[i]
	}

	a.data = output

	return toReturn
}
