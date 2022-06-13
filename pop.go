package goarray

func (a *Array[T]) Pop() T {
	toReturn := a.data[len(a.data)-1]
	output := make([]T, len(a.data)-1)

	for i := 0; i < len(a.data)-1; i++ {
		output[i] = a.data[i]
	}

	a.data = output

	return toReturn
}
