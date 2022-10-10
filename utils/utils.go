package utils

func Min(a, b int, type_1, type_2 string) (int, string) {
	if a >= b {
		return a, type_1
	}
	return b, type_2
}
