package main

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}

	left, right := 0, len(s)-1

	for right > left {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp

		left++
		right--
	}
}
