package main

func GetConcatenation(nums []int) []int {
	ans := make([]int, 0, len(nums)*2)

	for i := 0; i < 2; i++ {
		for k := 0; k < len(nums); k++ {
			ans = append(ans, nums[k])
		}
	}

	return ans
}

func GetConcatenation1(nums []int) []int {
	l := len(nums)
	ans := make([]int, l*2)

	for i := 0; i < l; i++ {
		ans[i] = nums[i]
		ans[i+l] = nums[i]
	}

	return ans
}

func GetConcatenation2(nums []int) []int {
	length := len(nums)
	arr := make([]int, length*2)

	for i := 0; i < length*2; i++ {
		arr[i] = nums[i%length]
	}

	return arr
}
