# Problem Explanation

## Using Two Arrays
The idea is using two arrays to store the length of going up and going down.<br>
up array: each element represent how many steps is going up til this index<br>
down array: each element represent how many steps is going down til this index<br>

For example, array `[2,1,4,7,3,2,5]`<br>
up array:          `[0,0,1,2,0,0,0]`<br>
down array:        `[0,0,0,2,1,0,0]`<br>
For up array, index 3(value 7) means the length of going up til index 2 is 2 (go from start to end), 1 -> 4 -> 7<br>
For down array, index 3(value 7) means the length of going down til index 3 is 2 (go from end to start), 2 -> 3 -> 7<br>

After building up and down array, we can iterate through the array to find the longest mountain.

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(n)

## One Pass (Finding Peak)
The idea is using one pass to find any peak in the array, and then check the left and right side of the peak to calculate the length of mountain.<br>
For example, array `[2,1,4,7,3,2,5]`<br>
There's only one peak is index 3(value 7), how do we know it's a peak?<br>
We can check the previous and next element of current index, if the previous element is less than current element and the next element is less than current element, then it's a peak.<br>
After finding a peak, we can check the left and right side of the peak to calculate the length of mountain.<br>
To find the left-most side and right-most side, we keep moving left and right pointer until the value is not decreasing.<br>

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)

## One Pass (Two Pointers)
This solution is not as straightfoward as the previous one, but it's still O(n) in time complexity and O(1) in space complexity.<br>
The idea is following:<br>
1. Take the current index as starting point
2. Move the starting point to the right until the value is not increasing
3. Check if the starting point is still the same as current index
   - If yes, it means the current value is equal to the next value, so it's not a mountain, we need to move the starting point to the right to find the next starting point
4. Set the current index as peak
   - Note we know that the pointer is not as same as starting point, so we can set it as peak
5. Move the current index to the right until the value is not decreasing
6. After moving, the current index is the end of the mountain, we can calculate the length of mountain and update the longest length

Note that both (Finding Peak) and (Two Pointers) have same time and space complexity, but (Two Pointers) is more complex than (Finding Peak).

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)
