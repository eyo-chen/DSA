package main

var mapParantheses = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func IsValid(s string) bool {
	sta := []byte{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		// if the character is not in the map, it is an open parentheses
		if _, ok := mapParantheses[c]; !ok {
			sta = append(sta, c)
			continue
		}

		// it is a close parentheses

		// if the stack is empty or the top of the stack is not the corresponding open parentheses
		// Note: we should check if the stack is empty first to avoid index out of range
		// e.g. input: "]"
		// sta[len(sta)-1] will be out of range
		// also, if it's a close parentheses, and stack is empty, it means the input is invalid
		// we should return false directly
		if len(sta) == 0 || sta[len(sta)-1] != mapParantheses[c] {
			return false
		}
		sta = sta[:len(sta)-1]
	}

	// we should check if the stack is empty
	// e.g. input: "("
	// because we only have one open parentheses, it will be added to the stack and exit the loop
	// but we need to check if the stack is empty to make sure the input is valid
	return len(sta) == 0
}
