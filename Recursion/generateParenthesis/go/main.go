package main

func GenerateParenthesis(n int) []string {
	ans := []string{}
	helper(n, []byte{}, &ans, 0, 0)

	return ans
}

func helper(n int, str []byte, ans *[]string, countOpen int, countClose int) {
	if len(str) == n*2 {
		*ans = append(*ans, string(str))
		return
	}

	if countOpen < n {
		helper(n, append(str, '('), ans, countOpen+1, countClose)
	}

	if countClose < countOpen {
		helper(n, append(str, ')'), ans, countOpen, countClose+1)
	}
}
