package popcount

func popcount(x uint64) int {
	count := 0
	for ; x > 0; x >>= 1 {
		if x&1 == 1 {
			count += 1
		}
	}
	return count
}

func sparsePopcount(x uint64) int {
	count := 0
	for ; x > 0; x &= x - 1 {
		count += 1
	}
	return count
}
