package testMe

func f1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return f1(n-1) + f1(n-2)
}

func f2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 2
	}
	return f2(n-1) + f2(n-2)
}

func s1(s string) int {
	if s == "" {
		return 0
	}
	n := 1
	for range s {
		n++
	}
	return n
}

func s2(s string) int {
	return len(s)
}
