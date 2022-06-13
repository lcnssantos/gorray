package goarray

func Reduce[T Element, Y Element](data Array[T], cb func(output *Y, element T) Y, initialElement *Y) Y {
	out := initialElement

	for _, element := range data.data {
		val := cb(out, element)
		out = &val
	}

	return *out
}
