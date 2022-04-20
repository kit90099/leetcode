package main

/* func main() {
	s, _ := os.ReadFile("s.txt")
	t, _ := os.ReadFile("t.txt")

	fmt.Println(minWindow(string(s), string(t)))
} */

func minWindow(s string, t string) string {
	min := ""

	left := 0

	needs := make(map[byte]int)
	windows := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		str := t[i]
		needs[str] = needs[str] + 1
	}

	for right := 0; right < len(s); right++ {
		char := s[right]
		//update window
		windows[char] = windows[char] + 1

		// judge whether move left ptr
		valid := true
		for j := 0; j < len(t); j++ {
			str := t[j]
			if needs[str] > windows[str] {
				valid = false
				break
			}
		}
		//start moving left
		if valid {
			for l := left; l <= right; l++ {
				toBeRemoved := s[l]
				window := windows[toBeRemoved] - 1
				if window < needs[toBeRemoved] {
					if min == "" || (right-left+1) < len(min) {
						min = s[left : right+1]
						windows[s[left]] = windows[s[left]] - 1
						left++
					}
					break
				}
				left++
				windows[toBeRemoved] = windows[toBeRemoved] - 1
			}
		}
	}

	return min
}
