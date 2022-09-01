package main

func main() {
	multiply("0", "0")
}

func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	res := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			product := int((num1[i] - '0') * (num2[j] - '0'))
			p1, p2 := i+j, i+j+1

			sum := product + res[p2]
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}

	i := 0
	for i < len(res) && res[i] == 0 {
		i++
	}

	str := ""
	for ; i < len(res); i++ {
		str += string(res[i] + '0')
	}

	if str == "" {
		return "0"
	}
	return str
}
