# Problem Explanation

## Marking Visited Numbers
The idea is to mark the visited numbers by negating the value at the index of the number.<br>
For example, if we see the number 4, we will negate(multiply by -1) the value at index 3 (4-1).<br>
If we see the number 4 again, we will check if the value at index 3 is negative. If it is, that means 4 is a duplicate number.<br>

Let's see the following example:<br>
nums = [4,3,2,7,8,2,3,1]
- index 0, nums[0] = 4
  - idx = 4 - 1 = 3
  - check if nums[3] is negative, it's not
  - mark nums[3] as negative, it represents that we have seen 4
  - nums = [4,3,2,-7,8,2,3,1]
- index 1, nums[1] = 3
  - idx = 3 - 1 = 2
  - check if nums[2] is negative, it's not
  - mark nums[2] as negative, it represents that we have seen 3
  - nums = [4,3,-2,-7,8,2,3,1]
- index 2, nums[2] = -2
  - idx = 2 - 1 = 1
  - check if nums[1] is negative, it's not
  - mark nums[1] as negative, it represents that we have seen 2
  - nums = [4,-3,-2,-7,8,2,3,1]
- index 3, nums[3] = -7
  - idx = 7 - 1 = 6
  - check if nums[6] is negative, it's not
  - mark nums[6] as negative, it represents that we have seen 7
  - nums = [4,-3,-2,-7,8,2,-3,1]
- index 4, nums[4] = 8
  - idx = 8 - 1 = 7
  - check if nums[7] is negative, it's not
  - mark nums[7] as negative, it represents that we have seen 8
  - nums = [4,-3,-2,-7,8,2,-3,-1]
- index 5, nums[5] = 2
  - idx = 2 - 1 = 1
  - check if nums[1] is negative, it is
  - add 2 to the result array
  - result = [2]
  - nums = [4,-3,-2,-7,8,2,-3,-1]
- index 6, nums[6] = -3
  - idx = 3 - 1 = 2
  - check if nums[2] is negative, it is
  - add 3 to the result array
  - result = [2, 3]
  - nums = [4,-3,-2,-7,8,2,-3,-1]
- index 7, nums[7] = -1
  - idx = 1 - 1 = 0
  - check if nums[0] is negative, it's not
  - mark nums[0] as negative, it represents that we have seen 1
  - nums = [-4,-3,-2,-7,8,2,-3,-1]

- return result = [2, 3]

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)

## Sorting With Swapping
Because we're guaranteed that each number is within the range of 1 to n, we can use the value of the number as the index to swap the number to its correct position.<br>
We keep swapping the number to its correct position until the number is at its correct position or the number is already in its correct position.<br>

Let's find out how to sort the array with swapping<br>
First, we have to realize the ultimate goal is to sort the array in ascending order.<br>
In other words, the value of the number should be the same as the index of the number + 1.<br>
For example, 2 should be at index 1, 3 should be at index 2, 4 should be at index 3, etc.<br>

Suppose, nums = [4,3,2,7,8,2,3,1]<br>
- index 0, nums[0] = 4
  - index 0 should be 1, let's move 4 out of this index, but where to put it?
  - we can put it at the correct position, which is index 3 (4 - 1)
  - swap nums[0] with nums[3], nums = [7,3,2,4,8,2,3,1]
- index 0, nums[0] = 7
  - index 0 should be 1, let's move 7 out of this index, but where to put it?
  - we can put it at the correct position, which is index 6 (7 - 1)
  - swap nums[0] with nums[6], nums = [3,3,2,4,8,2,7,1]
- index 0, nums[0] = 3
  - index 0 should be 1, let's move 3 out of this index, but where to put it?
  - we can put it at the correct position, which is index 2 (3 - 1)
  - swap nums[0] with nums[2], nums = [2,3,3,4,8,2,7,1]
- index 0, nums[0] = 2
  - index 0 should be 1, let's move 2 out of this index, but where to put it?
  - we can put it at the correct position, which is index 1 (2 - 1)
  - swap nums[0] with nums[1], nums = [3,2,3,4,8,2,7,1]
- index 0, nums[0] = 3
  - index 0 should be 1, let's move 3 out of this index, but where to put it?
  - we can put it at the correct position, which is index 2 (3 - 1)
  - !!! However, we see that nums[2] is already 3, which means index 2 already have the correct number. There's no point to swap it again.!!!
    - So, that's one of our stopping conditions
    - ***We want to stop swapping if the index is gonna swap already has the correct number***
    - the condition is ***nums[i] == nums[nums[i] - 1]***

Now, we can't swap index 0 anymore because the current value of index 0 (value = 3) is already the same as the value at index 2 (value = 3). So, we move on to the next index.
- index 1, nums[1] = 2
  - index 1 should be 2, which is already in its correct position
  - !!!That's our another stopping condition!!!
    - ***We want to stop swapping if the index is already in its correct position***
    - the condition is ***nums[i] == i + 1***


Now, we know the swapping logic<br>
We just want to keep swapping until<br>
- the current index is in its correct position
- the current value is the same as the value at the correct position
One one of them is true, we stop swapping.

#### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)
