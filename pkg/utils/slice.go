package utils

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func RemoveElem[T comparable](s []T, e T) []T {
	for i := 0; i < len(s); i++ {
		if s[i] == e {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}
