package lib

// WindowSum takes a data array and a window width and returns the windowed sums
// as an array.
func WindowSum(data []int, width int) []int {
	// TODO: check width < len(data)
	sums := make([]int, 0)

	sum := 0
	for i := 0; i < width; i++ {
		sum += data[i]
	}
	sums = append(sums, sum)

	for i := width; i < len(data); i++ {
		sum = sum + data[i] - data[i-width]
		sums = append(sums, sum)
	}

	return sums
}

// Reduce performs a reduction on the data array using the provided accumulation
// function fn, and returns the reduced value. The fn must be associative.
func Reduce(data []int, fn func(x, y int) int) int {
	total := 0
	a := data[0]

	for _, b := range data[1:] {
		total += fn(a, b)
		a = b
	}

	return total
}

// Increase returns 1 if y > x, 0 otherwise.
func Increase(x, y int) int {
	if y > x {
		return 1
	}
	return 0
}
