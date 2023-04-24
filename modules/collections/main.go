package collections

func Except[T comparable](everything, except []T) []T {
	var result []T

	for _, iv := range everything {
		found := false

		for _, jv := range except {
			if jv == iv {
				found = true
				break
			}
		}

		if !found {
			result = append(result, iv)
		}
	}

	return result
}

func Intersect[T comparable](x, y []T) []T {
	var result []T

	for _, iv := range x {
		for _, jv := range y {
			if jv == iv {
				result = append(result, iv)
				break
			}
		}
	}

	return result
}
