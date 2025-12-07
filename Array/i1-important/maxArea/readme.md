# Problem Explanation
## Brute Force
The idea is to calculate the area for every possible pair of the array and return the maximum area.<br>
For example, [1,8,6,2,5,4,8,3,7]<br>
We start with i = 0, k = 1, the area is (k - i) * min(1, 8) = 1<br>
Then we move k to 2, the area is (k - i) * min(1, 6) = 6<br>
Keep moving k to the end of the array, and update the maximum area<br>
Then, we move i to 1, and do the same thing as before<br>

### Complexity Analysis
#### Time Complexity O(n^2)
#### Space Complexity O(1)

## Two Pointers
The idea is to use two pointers, one starting from the beginning and the other starting from the end of the array.<br>
First, we calculate the area with the two pointers<br>
Then, it's time to move one of the pointers inward<br>
But, which pointer should we move inward?<br>
We should move the pointer with the smaller height inward<br>
Why is that?<br>
Consider this example, [1,8,6,2,5,4,8,3,7]<br>
When left = 1, right = 7, the area is (right - left) * min(1, 7) = 8<br>
At this point, should we move left(1) inward or right(7) inward?<br>
We know that the current max area is limited by two factors<br>
1. The distance between left and right<br>
2. The minimum height between left and right<br>

We know that the first factor remains the same no matter which pointer we move inward<br>
For the second factor, we know that ***it's the minimum height LIMITING the maximum area***<br>
So, there's no point to move the pointer with the larger height inward because the height is limited by the other pointer<br>
Therefore, we should move the pointer with the smaller height inward to have a chance to get a larger height in the next step<br>

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)
