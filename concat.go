package goarray

func (a *Array[T]) Concat(inputs ...Array[T]) Array[T] {
	newArray := New(&a.data)

	for _, item := range inputs {
		newArray.data = append(newArray.data, item.data...)
	}

	return newArray
}
