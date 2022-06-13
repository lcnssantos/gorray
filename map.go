package goarray

func Map[T Element, V Element](data Array[T], cb func(element T) V) Array[V] {
	output := make([]V, 0)

	for _, item := range data.data {
		output = append(output, cb(item))
	}

	return New(&output)
}
