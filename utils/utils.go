package utils

func FindIndex(slice []int, number int) *int {
	for i, v := range slice {
		if v == number {
			return &i
		}
	}

	return nil
}
