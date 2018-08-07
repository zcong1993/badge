package libs

var verdana = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3.9, 4.3, 5, 9, 7, 12, 8, 3, 5, 5, 7, 9, 4, 5, 4, 5, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 5, 5, 9, 9, 9, 6, 11, 7.5, 7.5, 7.7, 8.5, 7, 6.3, 8.5, 8.3, 4.6, 5, 7.6, 6.1, 9.3, 8.2, 8.7, 6.6, 8.7, 7.6, 7.5, 6.8, 8.1, 7.5, 11, 7.5, 6.8, 7.5, 5, 5, 5, 9, 7, 7, 6.6, 6.9, 5.7, 6.9, 6.6, 3.9, 6.9, 7, 3, 3.8, 6.5, 3, 11, 7, 6.7, 6.9, 6.9, 4.7, 5.7, 4.3, 7, 6.5, 9, 6.5, 6.5, 5.8, 7, 5, 7, 9}

// CacWith is helper function for cac with we should alloc to this string
func CacWith(str string) float64 {
	verdana[64] += 0.6
	total := 0.0
	s := []rune(str)
	for i := len(s) - 1; i >= 0; i-- {
		code := s[i]
		if code >= 127 {
			total += verdana[64]
		} else {
			total += verdana[code]
		}
	}

	return total
}
