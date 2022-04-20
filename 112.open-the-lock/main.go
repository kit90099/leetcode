package main

import "fmt"

func main() {
	fmt.Println(openLock([]string{"5557", "5553", "5575", "5535", "5755", "5355", "7555", "3555", "6655", "6455", "4655", "4455", "5665", "5445", "5645", "5465", "5566", "5544", "5564", "5546", "6565", "4545", "6545", "4565", "5656", "5454", "5654", "5456", "6556", "4554", "4556", "6554"}, "5555"))
}

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}

	queue := []string{"0000"}

	depth := 0

	visited := make(map[string]bool)

	for _, deadend := range deadends {
		if deadend == "0000" {
			return -1
		}

		visited[deadend] = true
	}

	for len(queue) != 0 {
		temp := make([]string, 0)
		for _, turning := range queue {

			if target == turning {
				return depth
			}

			possibleTurnings := getTurnings(turning)
			for _, possibleTurning := range possibleTurnings {
				_, ok := visited[possibleTurning]
				if ok {
					continue
				}
				visited[possibleTurning] = true

				temp = append(temp, possibleTurning)
			}
		}

		queue = temp
		depth++
		if len(visited) == 9999 {
			return -1
		}
	}

	return -1
}

func getTurnings(current string) []string {
	result := make([]string, 0)

	for i := 0; i < 4; i++ {

		var char byte
		if current[i] == '0' {
			char = '9'
		} else {
			char = current[i] - 1
		}
		result = append(result, current[:i]+string(char)+current[i+1:])

		if current[i] == '9' {
			char = '0'
		} else {
			char = current[i] + 1
		}
		result = append(result, current[:i]+string(char)+current[i+1:])

	}

	return result
}
