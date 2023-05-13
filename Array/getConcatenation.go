package main

import "fmt"

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(getConcatenation1(nums))
}

// original solution
// bad solution because of append, and append will create a new array
// so it will cost more time and space
// also not create fix length array
func getConcatenation(nums []int) []int {
	var arr []int

	index := len(nums) * 2

	for i := 0; i < index; i++ {
		if i < len(nums) {
			arr = append(arr, nums[i])
		} else {
			arr = append(arr, nums[i-len(nums)])
		}
	}

	return arr
}

func getConcatenation1(nums []int) []int {
	length := len(nums)
	arr := make([]int, length*2)

	for i := 0; i < length; i++ {
		arr[i] = nums[i]
		arr[i+length] = nums[i]
	}

	return arr
}

func getConcatenation2(nums []int) []int {
	length := len(nums)
	arr := make([]int, length*2)

	for i := 0; i < length*2; i++ {
		arr[i] = nums[i%length]
	}

	return arr
}

func getConcatenation3(nums []int) []int {
	return append(nums, nums...)
}
