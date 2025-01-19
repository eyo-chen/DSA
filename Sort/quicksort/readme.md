# Problem Explanation

Quick Sort is a divide-and-conquer algorithm that efficiently sorts elements by partitioning the array into smaller sub-arrays. It selects a pivot element and rearranges the array so that all elements less than the pivot come before it, and all elements greater come after. This process is recursively applied to the sub-arrays until the entire array is sorted.

Core concept of quick sort:<br>
1. Choose a pivot element from the array.
   - A pivot is essentially an element that is used to divide the array into two parts.
   - The purpose of the pivot is to make sure that all elements less than the pivot come before it, and all elements greater come after.
   - The selection of the pivot is crucial for the performance of the algorithm.
2. Partition the array:
   - The idea is to rearrange the elements in the array so that all elements less than the pivot come before it, and all elements greater come after.
   - This is done by swapping elements that are greater than the pivot with the last smaller element than pivot.
   - After the loop, the pivot element will be in its correct position.
3. Recursively apply the same process to the sub-arrays:


Let's focus on the partition logic.<br>
Suppose we have the following array: [10,7,8,9,1,5]<br>
We choose the last element as pivot, which is 5.<br>
Again, our goal is to rearrange the array so that all elements less than the pivot come before it, and all elements greater come after.<br>
We will use two pointers:
- i: to iterate through the array
- idxSmallerThanPivot: to keep track where's the last smaller element than pivot

`i`'s job is to scan through the array, and check if the current element is smaller than pivot.<br>
If it is, we swap the current element with the last smaller element than pivot(`idxSmallerThanPivot`), and move the `idxSmallerThanPivot` to the right.<br>

`idxSmallerThanPivot`'s job is to keep track where's the last smaller element than pivot.<br>
After the loop, we swap the pivot element with the last smaller element than pivot(`idxSmallerThanPivot`), and return the `idxSmallerThanPivot` as the pivot index.

Let's see the process in detail:
[10,7,3,9,1,5], pivot = 5

i = 0, arr[i] = 10, 10 > 5, continue<br>
i = 1, arr[i] = 7, 7 > 5, continue<br>
i = 2, arr[i] = 3, 3 < 5<br>
Now,<br>
`i` says: "Hey, I found an element smaller than pivot"<br>
`idxSmallerThanPivot` says: "Great, Give it to me, let's swap it"<br>
After the swap, `idxSmallerThanPivot` moves to the right, and `i` moves to the right.<br>
arr = [3,7,10,9,1,5], idxSmallerThanPivot = 1, i = 3<br>

i = 3, arr[i] = 9, 9 > 5, continue<br>
i = 4, arr[i] = 1, 1 < 5<br>
Now,<br>
`i` says: "Hey, I found an element smaller than pivot"<br>
`idxSmallerThanPivot` says: "Great, Give it to me, let's swap it"<br>
After the swap, `idxSmallerThanPivot` moves to the right, and `i` moves to the right.<br>
arr = [3,1,10,9,7,5], idxSmallerThanPivot = 2, i = 5<br>

After for loop, the array looks like this:<br>
```
     i        -> idxSmallerThanPivot
[3,1,10,9,7,5]
```
The position of `idxSmallerThanPivot` is the last position of smaller elements than pivot.<br>
Aka, it's exactly the position where the pivot should be.<br>
So, we swap the pivot element with the last smaller element than pivot(`idxSmallerThanPivot`), and return the `idxSmallerThanPivot` as the pivot index.<br>
arr = [3,1,5,9,7,10]<br>
And we do the same thing for the left and right sub-array.

# Complexity Analysis
## Time Complexity O(nlogn) or O(n^2)
- Worst case:
  - If the pivot is always the smallest or largest element, the partition will be unbalanced, and the algorithm will degrade to O(n^2).
  - Why is that?
  - Suppose we have the following array: [5,3,9,6,1] n = 5
  - If we choose the last element as pivot, the result of partition will be: [1,3,5,6,9]
  - The left sub-array is empty(n = 0), and the right sub-array is [3,5,6,9](n = 4)
  - Can you see the patternr?
  - We do not successfully reduce the problem size, the array is only (n - 1), not (n/2)
  - So, the worst case is when the pivot is always the smallest or largest element.
  - For each level, we keep having (n - 1) elements to sort, and the problem size is not reduced.
  - So, the time complexity is O(n^2).
- Best case:
  - If the pivot is always the middle element, the partition will be balanced, and the algorithm will degrade to O(nlogn).
  - Why is that?
  - Suppose we have the following array: [5,3,9,6,1] n = 5
  - If we choose the middle element as pivot, the result of partition will be: [1,3,5,6,9]
  - The left sub-array is [1,3](n = 2), and the right sub-array is [6,9](n = 2)
  - We successfully reduce the problem size, the array is only (n/2), not (n - 1)
  - So, the best case is when the pivot is always the middle element.

## Space Complexity O(logn)
- The space complexity is O(logn) because we are using recursion, and the stack space is O(logn).
