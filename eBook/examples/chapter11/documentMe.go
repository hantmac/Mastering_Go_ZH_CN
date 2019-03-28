// This package is for showcasing the documentation capabilities of Go
// It is a naive package!
package documentMe

// Pie is a global variable
// This is a silly comment!
const Pie = 3.1415912

// The S1() function finds the length of a string
// It iterates over the string using range
func S1(s string) int {
	if s == "" {
		return 0
	}
	n := 0
	for range s {
		n++
	}
	return n
}

// The F1() function returns the double value of its input integer
// A better function name would have been Double()!
func F1(n int) int {
	return 2 * n
}
