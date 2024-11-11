# Problem Explanation

## Use Temporary Array
The idea is to create a new array to store the sorted elements.<br>
After that, we can copy the sorted elements back to the original array(nums1).<br>

For example, nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3<br>
We can use two pointers to loop through nums1 and nums2, and compare the elements between two pointers.<br>
When the elements in nums1 is less than the elements in nums2, we can put the elements in nums1 to the sorted array.<br>
Otherwise, we can put the elements in nums2 to the sorted array.<br>

In short, we just put all the elements in nums1 and nums2 to the sorted array in order.<br>
Then, we can copy the sorted elements back to the original array(nums1).<br>

### Complexity Analysis
#### Time Complexity O(m + n)
- We need to loop through nums1 and nums2 to put all the elements to the sorted array.

#### Space Complexity O(m + n)
- We need to create a new array to store the sorted elements.

## Use Two Pointers
The idea is to use two pointers to loop through nums1 and nums2 from ***the end to the beginning.***<br>
If the elements in nums1 is greater than the elements in nums2, we can put the elements in nums1 to the end of the array.<br>
Otherwise, we can put the elements in nums2 to the end of the array.<br>

***The core logic is that we fill up the sorted elements from the end of the array(nums1).***<br>

### Complexity Analysis
#### Time Complexity O(m + n)
- We need to loop through nums1 and nums2 to put all the elements to the end of the array.

#### Space Complexity O(1)