package main

func superPow(a int, b []int) int {
	n := len(b)

	if n == 0 {
		return 1
	}

	last := b[n-1]
	b = b[:n-1]

	p1 := mypow(a, last)
	p2 := mypow(superPow(a, b), 10)

	return (p1 * p2) % 1337
}

func mypow(a, k int) int {
	if k == 0 {
		return 1
	}

	a %= 1337

	if k%2 == 1 {
		return (a * mypow(a, k-1)) % 1337
	} else {
		sub := mypow(a, k/2)
		return (sub * sub) % 1337
	}
}
