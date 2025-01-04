package main

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	// Note that it has to be <=, because when left == right, we still need to check if the target is at the mid position.
	for left <= right {
		mid := (right-left)/2 + left

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

/*
# Understanding Binary Search's O(log N) Time Complexity

## Visual Representation
We can model the process of binary search as a tree structure:
```
                [-1,0,3,5,9,12]
                         [5,9,12]
                              [12]
```

## Understanding the Pattern
Looking at our example, we can observe:
- It takes two steps to find our target
- The input array length is 6
- At each level, the array size is reduced:
  - Level 1: from 6 to 3 elements
  - Level 2: from 3 to 1 element

## Mathematical Model
We can express this pattern mathematically:
- Level 1: 3 = 6 / 2¹
- Level 2: 1 ≈ 6 / 2²

## Finding the Time Complexity
For an input array of size n:
1. We know we'll find our answer when we've reduced the array to a single element(n = 1)
2. This can be expressed as: 1 = n / 2ᵏ
   - Where k represents the number of steps needed
	 - We want to know what does k represent?
	 - If we know K, we know how many steps we need to take to find the target.
3. Solving for k:
   - 2ᵏ = n
   - k = log₂(n)

## Conclusion
Therefore, for an input of size n, it takes roughly log(n) steps to find our target in the worst case, making binary search's time complexity O(log N).
*/
