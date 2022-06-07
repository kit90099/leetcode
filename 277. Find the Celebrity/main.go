package main

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */

func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		cand := 0
		for other := 1; other < n; other++ {
			if !knows(other, cand) || knows(cand, other) {
				cand = other
			}
		}

		for other := 0; other < n; other++ {
			if cand == other {
				continue
			}

			if !knows(other, cand) || knows(cand, other) {
				return -1
			}
		}

		return cand
	}
}
