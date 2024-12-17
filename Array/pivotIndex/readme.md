# Problem Explanation

## Build Two Arrays
The idea is to build two arrays, `arr1` and `arr2`.<br>
`arr1` is the prefix sum array of `nums`, which means to calculate the sum from the first element to the last element.<br>
`arr2` is the suffix sum array of `nums`, which means to calculate the sum from the last element to the first element.<br>

After building these two arrays, we can iterate through the array to see if there is element in `arr1` is equal to the element in `arr2`.<br>
If there is, it means we have found the pivot index.

For example, nums = [1, 7, 3, 6, 5, 6]
```
arr1 = [1,   8, 11, 17, 22, 28]
arr2 = [28, 27, 20, 17, 11, 6]
```
i = 3, arr1[3] = 17, arr2[3] = 17, so the pivot index is 3.

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(n)

## Optimized Approach
The idea is to calculate the sum of the array first<br>
Then iterate through the array to calculate the sum of the numbers to the left and right of the current index at the same time.<br>
For sum of the numbers to the left, we can simply add the current element to the sum of the numbers to the left.<br>
For sum of the numbers to the right, we can subtract the current element and the sum of the numbers to the left from the total sum.<br>
If the sum of the numbers to the left is equal to the sum of the numbers to the right, then we have found the pivot index.

For example, nums = [1, 7, 3, 6, 5, 6]
```
sum = 28
leftSum = 0
```
i = 0, leftSum = 0, rightSum = 27, not equal, continue<br>
i = 1, leftSum = 1, rightSum = 20, not equal, continue<br>
i = 2, leftSum = 8, rightSum = 17, not equal, continue<br>
i = 3, leftSum = 11, rightSum = 11, equal, so the pivot index is 3.<br>


### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(1)