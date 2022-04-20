package main

func main() {
	s := "rabbbit"
	t := "rabbit"

	println(numDistinct(s, t))
}

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)

	a := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		a[i] = make([]int, n+1)
		a[i][0] = 1
	}

	for i := 1; i <= m; i++ { // s
		for j := 1; j <= n; j++ { // t
			if s[i-1] == t[j-1] {
				a[i][j] = a[j-1][i-1] + a[i-1][j]
			} else {
				a[i][j] = a[i-1][j]
			}
		}
	}

	return a[m][n]
}
