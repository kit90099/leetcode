package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	prev1 := 0
	prev2 := 1

	for i := 2; i <= n; i++ {
		newValue := prev1 + prev2

		prev1 = prev2
		prev2 = newValue
	}

	return prev2
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input n: ")
	text, _, _ := reader.ReadLine()
	n, _ := strconv.Atoi(string(text))

	fmt.Printf("Fib(%s)=%d", text, Fib(n))
}
