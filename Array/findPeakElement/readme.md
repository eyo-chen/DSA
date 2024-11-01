# Problem Explanation

## Linear Search
The idea is to iterate through the array and check if the current element is greater than its neighbors(left and right). If it is, then it is a peak element.<br>
However, there are some edge cases to consider:

- What if the array has only one element?<br>
  - This case is simple, just return 0. 
- What if the array has only two elements?<br>
  - We can simply check if the first element is greater than the second element, if so, return 0, otherwise return 1.
- What if the array is in ascending order or descending order?
  - This one is a bit tricky because we might not consider this case<br>
  - Let's first consider the ascending order:
    - For example, if the array is [1,2,3,4,5], the for loop only loop from index 1(value 2) to 3(value 4)
    - However, for 2, 3 and 4, there is no peak element. Therfore, we might return nothing after for loop.
    - In order to handle this case, we can simply return the last element's index.
  - Now consider the descending order:
    - For example, if the array is [5,4,3,2,1], the for loop only loop from index 1(value 4) to 3(value 2)
    - However, for 4, 3 and 2, there is no peak element. Therfore, we might return nothing after for loop.
    - In order to handle this case, we can re-use the logic of checking the case of only two elements.
    - If the first element is greater than the second element, immediately return 0. Because the problem stated that nums[-1] = nums[n] = -∞, if the first element is greater than the second element, it is already a peak element.
    - For example, if the array is [5,4,3,10,1].
    - We know that 10 is a peak. However, 5 is also a peak because nums[-1] = -∞. So, we can immediately return 0.

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)

## Binary Search
The idea is to use the binary search logic to find the peak element.<br>
We know that the binary search start with two pointers, left and right.<br>
Then, we calculate the mid index.<br>
If the mid index is a peak element, we return the mid index.<br>
Now, we have to cut the search scope.<br>
We can either cut the left part or the right part.<br>
But which part should we cut?<br>
Let's consider two simple examples:<br>
array = [1,2,3,4,5]<br>
When mid = 2(value = 3), which part should we cut?<br>
It's pretty obvious that we should cut the left part<br>
Because we know that 2 < 3, there's no way there's a peak element on the left side of mid.<br>
array = [5,4,3,2,1]<br>
When mid = 2(value = 3), which part should we cut?<br>
It's pretty obvious that we should cut the right part<br>
Because we know that 3 > 2, there's no way there's a peak element on the right side of mid.<br>
array = [1,3,2,4,1]<br>
When mid = 2(value = 2), which part should we cut?<br>
This one is a bit tricky.<br>
In this case, we can either cut the left part or the right part because 2 < 3 and 2 < 4.<br>
At the left side, the peak element is 3.<br>
At the right side, the peak element is 4.<br>
Therefore, going either direction is correct.<br>

Now, it's pretty clear that the logic of which part to cut is:<br>
- If nums[mid + 1] > nums[mid], we should cut the left part.<br>
- Otherwise, we should cut the right part.<br>

But, how exactly should we move the left and right pointer?<br>
- If we cut the left part, we should move the left pointer to mid + 1.<br>
- Otherwise, we should move the right pointer to mid.<br>

Why is `left = mid + 1` and `right = mid`?<br>
Let's consider the following example:<br>
array = [1,2,3,4,5]<br>
In this case, when mid = 2(value = 3), we should cut the left part.<br>
Look closly at the condition:<br>
`nums[mid + 1] > nums[mid]`<br>
We're checking `mid + 1` and `mid`<br>
If we know that mid is smaller than mid + 1, there's no way mid is the peak element.<br>
Therefore, we should move the left pointer to mid + 1.<br>

Now, let's consider the following example:<br>
array = [5,4,3,2,1]<br>
In this case, when mid = 2(value = 3), we should cut the right part.<br>
Look closly at the condition:<br>
`nums[mid + 1] > nums[mid]`<br>
In this case, we only know that mid is greater than mid + 1.<br>
That means mid itself could be the peak element.<br>
Therefore, we should move the right pointer to mid.<br>

### Complexity Analysis
#### Time Complexity O(log n)
#### Space Complexity O(1)
