package cartesian

// nextIndex retrieves the next index for the next set
func nextIndex(ix []int, lens func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}

func Product[T any](sets ...[]T) [][]T {
	lengths := func(i int) int { return len(sets[i]) }
	var product [][]T
	for ix := make([]int, len(sets)); ix[0] < lengths(0); nextIndex(ix, lengths) {
		var r []T
		for j, k := range ix {
			r = append(r, sets[j][k])
		}
		product = append(product, r)
	}
	return product
}
