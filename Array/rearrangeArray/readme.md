# Problem Explanation

## Solution 1: Build Two Arrays
The idea is to separate positive and negative numbers into two arrays, then interleave them into the result array.

1. Create two slices: `positive` for positive numbers and `negative` for negative numbers.
2. Iterate through `nums` and append each number to the appropriate slice based on its sign.
3. Iterate through the `positive` array and append one positive and one negative number to the result slice per iteration.

For example, `nums = [3,1,-2,-5,2,-4]`:
- `positive = [3,1,2]`, `negative = [-2,-5,-4]`
- Result: `[3,-2,1,-5,2,-4]`

### Complexity Analysis
- **Time Complexity**: O(n), where `n` is the length of `nums`. One pass to separate numbers, one pass to interleave.
- **Space Complexity**: O(n) for the `positive`, `negative`, and result slices.


## Solution 2: Two-Pointer Approach with Skipping
The idea is to use two pointers to find the next positive and negative numbers, appending them alternately to the result.

1. Initialize `positivePtr` and `negativePtr` to track the next positive and negative numbers.
2. Skip negative numbers to find the first positive and positive numbers to find the first negative.
3. While both pointers are within bounds:
   - Append the current positive and negative numbers to the result.
   - Advance `positivePtr` to the next positive number, skipping negatives.
   - Advance `negativePtr` to the next negative number, skipping positives.

For example, `nums = [3,1,-2,-5,2,-4]`:
- Start: `positivePtr = 0` (3), `negativePtr = 2` (-2)
- Append `[3,-2]`, move to `positivePtr = 1` (1), `negativePtr = 3` (-5)
- Append `[1,-5]`, move to `positivePtr = 4` (2), `negativePtr = 5` (-4)
- Append `[2,-4]`

**Note**: This solution requires bounds checks to avoid out-of-bounds errors (e.g., if all numbers are negative).

### Complexity Analysis
- **Time Complexity**: O(n) average case
- **Space Complexity**: O(n) for the result slice.


## Solution 3: Optimized Single-Pass Approach
The idea is to place positive and negative numbers directly into their final positions in a single pass.

1. Create a result array of length `n`.
2. Use two indices: `positiveIdx` (starting at 0 for even indices) and `negativeIdx` (starting at 1 for odd indices).
3. Iterate through `nums`:
   - If the number is positive, place it at `positiveIdx` and increment by 2.
   - If the number is negative, place it at `negativeIdx` and increment by 2.

For example, `nums = [3,1,-2,-5,2,-4]`:
- Place 3 at index 0, `positiveIdx = 2`
- Place 1 at index 2, `positiveIdx = 4`
- Place -2 at index 1, `negativeIdx = 3`
- Place -5 at index 3, `negativeIdx = 5`
- Place 2 at index 4, `positiveIdx = 6`
- Place -4 at index 5, `negativeIdx = 7`
- Result: `[3,-2,1,-5,2,-4]`

### Complexity Analysis
- **Time Complexity**: O(n), single pass through the array.
- **Space Complexity**: O(n) for the result array (optimal since a new array is required).

