package itertools

// CartesianPower returns the Cartesian product of size with itself n times.
// Equivalent to itertools.product(range(size), repeat=n) in Python.
func CartesianPower(size uint64, n int) [][]byte {

	// count := size^n
	count := uint64(1)
	for i := 0; i < n; i++ {
		count *= size
	}
	product := make([][]byte, 0, count)
	tuple := make([]byte, n) // tuple has length n
	for i := uint64(0); i < count; i++ {

		idx := n - 1
		for num := i; num > 0; num /= size {
			tuple[idx] = byte(num % size)
			idx--
		}

		// copy tuple
		product = append(product, append([]byte{}, tuple...))
	}
	return product
}
