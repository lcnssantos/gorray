package goarray

func (a *Array[T]) Map(cb func(element T) interface{}) Array[interface{}] {
	output := make([]interface{}, 0)

	for _, item := range a.data {
		output = append(output, cb(item))
	}

	return New(&output)
}
