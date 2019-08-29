package work

// Purposefully using slower recursive function
func FibonacciNumberAtIndex(idx int) int {
	if idx == 0 {
		return 0
	} else if idx == 1 {
		return 1
	} else {
		return FibonacciNumberAtIndex(idx-2) + FibonacciNumberAtIndex(idx-1)
	}
}
