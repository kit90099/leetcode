package main

func main() {
	solve([][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}})
}

type UF struct {
	count  int
	parent []int
}

func (uf *UF) union(p int, q int) {
	rootP := uf.find(p)
	rootQ := uf.find(q)

	if rootP == rootQ {
		return
	}

	uf.parent[rootQ] = rootP
	uf.count--
}

func (uf *UF) find(p int) int {
	if uf.parent[p] != p {
		uf.parent[p] = uf.find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *UF) connected(p int, q int) bool {
	rootP := uf.find(p)
	rootQ := uf.find(q)

	return rootP == rootQ
}

func solve(board [][]byte) {
	//todo: check if size == 0
	m := len(board)
	n := len(board[0])

	parent := make([]int, m*n+1)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

	uf := UF{
		count:  m*n + 1,
		parent: parent,
	}

	dummy := m * n

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			uf.union(n*i, dummy)
		}
		if board[i][n-1] == 'O' {
			uf.union(n*i+n-1, dummy)
		}
	}

	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			uf.union(i, dummy)
		}
		if board[m-1][i] == 'O' {
			uf.union((m-1)*n+i, dummy)
		}
	}

	d := [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				for k := 0; k < 4; k++ {
					x := i + d[k][0]
					y := j + d[k][1]
					if board[x][y] == 'O' {
						//fmt.Printf("x=%d, y=%d, i=%d, j=%d\n", x, y, i, j)
						uf.union(x*n+y, i*n+j)
					}
				}
			}
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				if !uf.connected(dummy, i*n+j) {
					board[i][j] = 'X'
				}
			}
		}
	}
}
