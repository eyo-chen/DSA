# Problem Explanation

## Top-Down Approach (Memoization)
The idea is using a helper function to recursively check if we can jump to the end of the array from the current index. We use a memoization table to store the results of the subproblems to avoid redundant calculations.

Every recursive call represents a subproblem: "Can we jump to the end of the array from the current index?"

Base case:
- If the current index is the last index, return true.
- If the value of current index is 0, return false.
- If the current index is in the memoization table, return the result.

Recursive case:
- Iterate through the possible steps from the current index.
- If the helper function returns true for any of the steps, return true.
- If no steps lead to the end, return false.

For each index(position), we can jump from 1 to steps(nums[index]).<br>
We also need to check if the current index(position) has been cached, if so, return the cached result immediately without further calculation.

Also, note the following conditions:
`i <= steps && index+i < len(nums)`
We want to loop from 1 to steps, and we also need to make sure the next index(position) is within the bounds of the array.
For example, if we are at index 0, and the value of nums[0] is 3, we can jump from 1 to 3 steps.
But if the length of the array is 2, we cannot jump to index 3. Therefore, we only loop from 1 to 2.

### Complexity
#### Time Complexity O(n^2)
- In the worst case, we might need to explore every possible jump from each index.
- For each index, we can potentially make up to n-1 recursive calls (where n is the length of the input array).
- The length is n, and each element can at most visit n times.
- So the time complexity is O(n^2).

#### Space Complexity O(n)
- The recursion stack, which in the worst case can go up to n levels deep.
- The memoization map, which can store up to n entries (one for each index).


## Bottom-Up Approach (Tabulation)
This approach is similar to the top-down approach, but it uses a table to store the results of the subproblems.

First, we initialize a table with the same length as the input array, and set the first element to true.

Then, we iterate through the array, and for each element, we check if it is true. If it is, we iterate through the possible steps from the current index.<br>
If the next index is within the bounds of the array, we set the next index to true.<br>
Finally, we return the last element of the table.

Note the following conditions:
```Go
if !table[index] || steps <= 0 {
    continue
}
```
If the current index is false, it means this position is not reachable, so we skip it.
If the steps is 0, it means we cannot jump from the current position, so we skip it.

### Complexity
#### Time Complexity O(n^2)
- We iterate through the array once, and for each element, we iterate through the possible steps.
- This results in a time complexity of O(n^2).

#### Space Complexity O(n)
- We use a table to store the results of the subproblems, which can store up to n entries (one for each index).

## Optimized Approach
This is a clever approach. We start from the last index, and for each element, we check if it is reachable. If it is, we update the last index to the current index.
Finally, we check if the first element is reachable. If it is, we return true. Otherwise, we return false.

Let's walk through the example `[2,3,1,1,4]`:
- We start from the last index, which is 4.
- We check if index 3 is able to reach the last index 4. Since index 3 is 1, it can reach the last index 4, so we update the last index to 3.
- We check if index 2 is able to reach the last index 3. Since index 2 is 1, it can reach the last index 3, so we update the last index to 2.
- We check if index 1 is able to reach the last index 2. Since index 1 is 3, it can reach the last index 2, so we update the last index to 1.
- We check if index 0 is able to reach the last index 1. Since index 0 is 2, it can reach the last index 1, so we update the last index to 0.
- Finally, we check if the first element is reachable. Since the first element is 2, it can reach the last index 0, so we update the last index to 0.
- Since the last index is 0, we return true.

Let's walk through the example `[3,2,1,0,4]`:
- We start from the last index, which is 4.
- We check if index 3 is able to reach the last index 4. Since index 3 is 0, it cannot reach the last index 4, so last index remains 4.
- We check if index 2 is able to reach the last index 4. Since index 2 is 1, it cannot reach the last index 4, so last index remains 4.
- We check if index 1 is able to reach the last index 4. Since index 1 is 2, it cannot reach the last index 4, so last index remains 4.
- We check if index 0 is able to reach the last index 4. Since index 0 is 3, it cannot reach the last index 4, so last index remains 4.
- Last index is not 0, so we return false.

### Complexity
#### Time Complexity O(n)
- We iterate through the array once, and for each element, we check if it is reachable.
- This results in a time complexity of O(n).

#### Space Complexity O(1)
- We only use a few variables to store the last index, so the space complexity is O(1).