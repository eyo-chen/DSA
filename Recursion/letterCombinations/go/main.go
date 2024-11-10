package main

var digitToChar = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ans := []string{}
	dfs(digits, &ans, []byte{}, 0)

	return ans
}

// it's important to use pointer for ans
// because we need to modify the ans in the caller
func dfs(digits string, ans *[]string, curStr []byte, index int) {
	if index >= len(digits) {
		*ans = append(*ans, string(curStr))
		return
	}

	digit := digits[index]
	mapString := digitToChar[digit]
	for i := 0; i < len(mapString); i++ {
		c := mapString[i]
		curStr = append(curStr, c)
		dfs(digits, ans, curStr, index+1)
		curStr = curStr[:len(curStr)-1]
	}
}
