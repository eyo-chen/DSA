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

// Updated at 2025/01/10
func GenerateParenthesis2(n int) []string {
	ans := []string{}
	helper2(n, &ans, []byte{}, 0, 0)
	return ans
}

func helper2(n int, ans *[]string, curString []byte, open, close int) {
	if open == n && close == n {
		*ans = append(*ans, string(curString))
		return
	}

	if open > n || close > n {
		return
	}

	helper2(n, ans, append(curString, '('), open+1, close)
	if open > close {
		helper2(n, ans, append(curString, ')'), open, close+1)
	}
}
