# Solution Explanation

## What is a Prefix Sum? 

**Intuition**: Imagine you’re tracking the total amount of money you’ve saved each day. Each day, you add some amount to your savings, and you want to know how much you’ve saved up to any given day or between two days. Instead of adding up all the amounts every time you want to check, you keep a running total. This running total is the prefix sum—it’s the sum of all elements (e.g., daily savings) from the start up to a certain point.

**Formal Definition**: For an array `nums`, the prefix sum at index `i` is the sum of all elements from index `0` to `i`. It’s like a cumulative total as you move through the array.

**Example**: Let’s say you have an array representing daily savings: `nums = [3, 5, 2, -1, 4]` (dollars saved each day, including a negative day where you spent money).

- The prefix sum array is calculated as follows:
  - `prefix[0] = nums[0] = 3`
  - `prefix[1] = nums[0] + nums[1] = 3 + 5 = 8`
  - `prefix[2] = nums[0] + nums[1] + nums[2] = 3 + 5 + 2 = 10`
  - `prefix[3] = nums[0] + nums[1] + nums[2] + nums[3] = 3 + 5 + 2 + (-1) = 9`
  - `prefix[4] = nums[0] + nums[1] + nums[2] + nums[3] + nums[4] = 3 + 5 + 2 + (-1) + 4 = 13`

- So, the prefix sum array is: `prefix = [3, 8, 10, 9, 13]`.

**Why is this useful?**:
- Suppose you want the sum of savings from day 2 to day 4 (elements `[2, -1, 4]`).
- Without prefix sums, you’d add: `2 + (-1) + 4 = 5`, which takes O(n) time for each query.
- With prefix sums, you calculate: `prefix[4] - prefix[1] = 13 - 8 = 5`.
  - Why? `prefix[4]` is the sum from day 1 to day 5, and `prefix[1]` is the sum from day 1 to day 2. Subtracting gives the sum from day 2 to day 4.
  - This is O(1) per query after building the prefix sum array in O(n).


## Prefix Sum with a Hash Map (Dynamic Example)

Sometimes, you don’t need to store the entire prefix sum array, especially for problems like finding subarrays with a specific sum. Instead, you compute the prefix sum on the fly and use a hash map to track sums you’ve seen.

**New Example**: Find all subarrays in `nums = [3, 5, 2, -1, 4]` where the sum equals `6`.

- Initialize a hash map: `prefixMap = {0: -1}` (sum `0` before index `0` to handle subarrays starting at index `0`).
- Track the running prefix sum (`currSum`) as you iterate:
  - Index `0`: `currSum = 3`, `prefixMap = {0: -1, 3: 0}`.
  - Index `1`: `currSum = 3 + 5 = 8`, `prefixMap = {0: -1, 3: 0, 8: 1}`.
  - Index `2`: `currSum = 8 + 2 = 10`, `prefixMap = {0: -1, 3: 0, 8: 1, 10: 2}`.
  - Index `3`: `currSum = 10 + (-1) = 9`, check `currSum - target = 9 - 6 = 3`. `prefixMap` has `3` at index `0`, so subarray `[5, 2, -1]` (indices `1` to `3`) sums to `6`.
  - Continue: `prefixMap = {0: -1, 3: 0, 8: 1, 10: 2, 9: 3}`.
  - Index `4`: `currSum = 9 + 4 = 13`, `currSum - target = 13 - 6 = 7`, not in `prefixMap`. No new subarray found.

**Result**: Found subarray `[5, 2, -1]` with sum `6`.

**Why Hash Map?**:
- The hash map stores each prefix sum and its earliest index.
- If `currSum - target` exists in the hash map, it means there’s a subarray ending at the current index with the desired sum.
- This is efficient (O(n) time, O(n) space) and handles negative numbers seamlessly.


## First Solution: MaxNonOverlapping

This solution uses prefix sums with a hash map to find the maximum number of non-overlapping subarrays in `nums` where each subarray sums to `target`.

**Problem Recap**: Find the maximum number of non-overlapping subarrays in `nums` with sum equal to `target`. Non-overlapping means no shared indices.

**Approach**:
- **Prefix Sum with Hash Map**: For each segment starting at `startIdx`, compute prefix sums and store them in a hash map (`prefixSums`). If `currentSum - target` exists, a subarray summing to `target` is found.
- **Greedy Selection**: Select the earliest valid subarray in each segment to maximize count, then move `startIdx` to the subarray’s end to ensure non-overlapping.
- **Segmented Scanning**: Reset the hash map to `{0: -1}` for each segment to avoid overlap with prior subarrays.

**How It Works**:
1. Iterate `startIdx` from `0` to `len(nums)-1`.
2. For each `startIdx`, initialize `prefixSums = {0: -1}`, `currentSum = 0`.
3. Scan `endIdx` from `startIdx` to `len(nums)-1`:
   - Update `currentSum += nums[endIdx]`.
   - If `currentSum - target` exists in `prefixSums` at `prevIdx`, subarray from `prevIdx+1` to `endIdx` sums to `target`. Increment `subarrayCount`, set `startIdx = endIdx`, and break.
4. If no subarray found, increment `startIdx`.

**Example**: `nums = [1, 1, 1, 1, 1]`, `target = 2`
- `startIdx = 0`: At `endIdx = 1`, `currentSum = 2`, `currentSum - 2 = 0` in `prefixSums`. Subarray `[1, 1]` (indices `0` to `1`). `subarrayCount = 1`, `startIdx = 1`.
- `startIdx = 2`: At `endIdx = 3`, find subarray `[1, 1]` (indices `2` to `3`). `subarrayCount = 2`, `startIdx = 3`.
- `startIdx = 4`: No subarray sums to `2`. End with `subarrayCount = 2`.
- Result: `2` subarrays.

**Why Prefix Sum?**:
- Enables `O(1)` checks for subarrays summing to `target`.
- Hash map handles negative numbers and large arrays.
- Resetting hash map ensures non-overlapping by limiting prefix sums to the current segment.

**Drawback**:
- Nested loops can approach `O(n²)` in worst cases (e.g., sparse valid subarrays), causing Time Limit Exceeded (TLE) for large inputs (`nums.length <= 10^5`).



## Second Solution: MaxNonOverlapping1

This solution improves the first to achieve `O(n)` time complexity, addressing TLE while maintaining correctness.

**Issues with First Solution**:
- **Nested Loops**: Scanning from each `startIdx` reprocesses elements, leading to near-quadratic time.
- **Hash Map Resets**: Discards useful prefix sum data, forcing redundant scans.
- **Inefficient Non-Overlapping**: Restarting scans after finding a subarray increases time.

**Improvements**:
- **Single Pass**: Iterate through the array once, processing each element exactly once.
- **Persistent Hash Map**: Use a single hash map (`prefixMap`) across the array, storing all prefix sums and their earliest indices.
- **Non-Overlapping with `lastEnd`**: Track the end index of the last selected subarray (`lastEnd`). Count a new subarray only if its start (post-`prevIndex`) is after `lastEnd`. Update `lastEnd` to the current index after counting.
- **Greedy Strategy**: Select earliest valid subarrays in one pass, maximizing count.

**How It Works**:
1. Initialize `prefixMap = {0: -1}`, `currSum = 0`, `count = 0`, `lastEnd = -1`.
2. For each index `i`:
   - Update `currSum += nums[i]`.
   - If `currSum - target` exists in `prefixMap` at `prevIndex` and `prevIndex >= lastEnd`, count subarray from `prevIndex+1` to `i`. Increment `count`, set `lastEnd = i`.
   - Store `currSum` in `prefixMap` with index `i`.
3. Return `count`.

**Example**: `nums = [1, 1, 1, 1, 1]`, `target = 2`
- `i = 1`: `currSum = 2`, `currSum - 2 = 0` at `prevIndex = -1`, `-1 >= -1`. Subarray `[1, 1]` (indices `0` to `1`). `count = 1`, `lastEnd = 1`.
- `i = 2`: `currSum = 3`, `currSum - 2 = 1` at `prevIndex = 0`, `0 >= 1` is false (prevents overlap).
- `i = 3`: `currSum = 4`, `currSum - 2 = 2` at `prevIndex = 1`, `1 >= 1`. Subarray `[1, 1]` (indices `2` to `3`). `count = 2`, `lastEnd = 3`.
- End with `count = 2`.

**Why Better?**:
- **Time Complexity**: `O(n)` with one pass and `O(1)` hash map operations, fixing TLE.
- **Space Complexity**: `O(n)` for hash map, no redundant allocations.
- **Non-Overlapping**: `lastEnd` enforces non-overlapping without resetting hash map, preserving prefix sums.
- **Simpler Logic**: Eliminates nested loops, though `lastEnd` adds slight conceptual complexity.

**Trade-off**:
- `lastEnd` introduces complexity but is essential for correctness in a single pass.
- First solution is more intuitive (resetting hash map feels like starting fresh) but less efficient.

## How Prefix Sum Relate To This Problem?

### Why `currentSum - target` Identifies a Subarray Summing to `target`?

**Intuition**: The prefix sum approach is like keeping a running total as you move through the array. When you’re at index `i` with a prefix sum `currentSum`, you’re looking for a previous point in the array where the prefix sum was just the right amount less than `currentSum` so that the difference (the sum of elements between those points) equals `target`. That “right amount less” is exactly `currentSum - target`.

**Formal Explanation**:
- A **prefix sum** at index `i`, denoted `currentSum`, is the sum of all elements from index `0` to `i`: `nums[0] + nums[1] + ... + nums[i]`.
- The sum of a subarray from index `start+1` to `end` (i.e., `nums[start+1] + ... + nums[end]`) can be computed as the difference between two prefix sums:
  - `prefix[end] - prefix[start]`.
- If this subarray sums to `target`, then:
  - `prefix[end] - prefix[start] = target`.
- Rearranging, we get:
  - `prefix[start] = prefix[end] - target`.
- In our code, when we’re at index `end` (e.g., `i` or `endIdx`), `prefix[end]` is `currentSum`. So, we need to find a previous index `start` where the prefix sum was `currentSum - target`. If such a `start` exists in the hash map, the subarray from `start+1` to `end` sums to `target`.

**Why the Hash Map?**:
- We store prefix sums and their indices in a hash map (e.g., `prefixSums` or `prefixMap`).
- When we compute `currentSum` at index `end`, we check if `currentSum - target` is in the hash map. If it is, it means there’s an earlier index `start` where the prefix sum was `currentSum - target`, and the subarray between `start+1` and `end` has the desired sum.


### Step-by-Step Breakdown

Let’s break it down with an example to see why `currentSum - target` works.

**Example Array**: `nums = [3, 5, 2, -1, 4]`, `target = 6`.

- We want to find subarrays where the sum equals `6`.
- We’ll compute prefix sums on the fly and use a hash map to track them.

**Step 1: Initialize**:
- Hash map: `prefixSums = {0: -1}` (sum `0` at index `-1` to handle subarrays starting at `0`).
- `currentSum = 0`.

**Step 2: Iterate**:
- **Index 0**: `currentSum = 3`, `currentSum - target = 3 - 6 = -3`. Not in `prefixSums`. Add `3: 0` to `prefixSums = {0: -1, 3: 0}`.
- **Index 1**: `currentSum = 3 + 5 = 8`, `currentSum - target = 8 - 6 = 2`. Not in `prefixSums`. Add `8: 1` to `prefixSums = {0: -1, 3: 0, 8: 1}`.
- **Index 2**: `currentSum = 8 + 2 = 10`, `currentSum - target = 10 - 6 = 4`. Not in `prefixSums`. Add `10: 2` to `prefixSums = {0: -1, 3: 0, 8: 1, 10: 2}`.
- **Index 3**: `currentSum = 10 + (-1) = 9`, `currentSum - target = 9 - 6 = 3`. **Found** in `prefixSums` at index `0`!
  - This means the prefix sum at index `0` was `3`, and at index `3` it’s `9`.
  - The subarray from index `0+1 = 1` to `3` is `[5, 2, -1]`.
  - Sum: `9 - 3 = 6` (or directly: `5 + 2 + (-1) = 6`), which matches `target`.
- Add `9: 3` to `prefixSums = {0: -1, 3: 0, 8: 1, 10: 2, 9: 3}`.
- **Index 4**: `currentSum = 9 + 4 = 13`, `currentSum - target = 13 - 6 = 7`. Not in `prefixSums`. Add `13: 4`.

**Why `currentSum - target = 3` Worked**:
- At index `3`, `currentSum = 9` (sum from `0` to `3`: `3 + 5 + 2 + (-1)`).
- We needed a subarray ending at `3` that sums to `6`.
- If there’s an earlier index `start` where the prefix sum was `9 - 6 = 3`, then:
  - Subarray sum from `start+1` to `3` = `prefix[3] - prefix[start] = 9 - 3 = 6`.
- The hash map had `3` at index `0`, so the subarray from `1` to `3` sums to `target`.

### Connecting to the Problem and Your Code

In the context of the "Maximum Number of Non-Overlapping Subarrays With Sum Equals Target" problem, `currentSum - target` is used in both your solutions (`MaxNonOverlapping` and `MaxNonOverlapping1`) to find subarrays summing to `target`. Let’s see how it fits:

**First Solution (`MaxNonOverlapping`)**:
- For each segment starting at `startIdx`, you compute `currentSum` from `startIdx` to `endIdx`.
- When `currentSum - target` exists in `prefixSums`, you’ve found a subarray from `prevIdx+1` to `endIdx` that sums to `target`.
- Example: `nums = [1, 1, 1, 1, 1], target = 2`:
  - At `startIdx = 0`, `endIdx = 1`: `currentSum = 1 + 1 = 2`, `currentSum - target = 2 - 2 = 0`. `prefixSums` has `0: -1`, so subarray `[1, 1]` (indices `0` to `1`) sums to `2`.
- The hash map reset ensures non-overlapping by limiting prefix sums to the current segment.

**Second Solution (`MaxNonOverlapping1`)**:
- In a single pass, you compute `currSum` up to index `i`.
- If `currSum - target` exists at `prevIndex` and `prevIndex >= lastEnd`, the subarray from `prevIndex+1` to `i` sums to `target` and is non-overlapping.
- Example: At `i = 1`, `currSum = 2`, `currSum - target = 0` at `prevIndex = -1`, subarray `[1, 1]` is valid.
- `lastEnd` ensures no overlap without resetting the hash map.

**Why `currentSum - target` in Both**:
- It’s the mathematical consequence of wanting `prefix[end] - prefix[start] = target`.
- The hash map stores `prefix[start]` values, so checking `currentSum - target` finds the `start` that makes the subarray sum work.

---

### Common Points of Confusion and Clarifications

1. **Why Subtract `target`?**
   - You’re looking for a subarray where the sum is `target`. Since `currentSum` includes all elements up to `end`, you need to “remove” the sum of elements up to `start` to get just the subarray’s sum. Subtracting `target` from `currentSum` gives the prefix sum you need at `start`.

2. **What if `currentSum - target` Isn’t in the Hash Map?**
   - If it’s not there, no subarray ending at the current index sums to `target`. You continue to the next index, adding the current `currentSum` to the hash map.

3. **Why Initialize Hash Map with `{0: -1}`?**
   - This handles subarrays starting at index `0`. If `currentSum = target` at index `i`, then `currentSum - target = 0`, and the hash map’s `0: -1` entry indicates the subarray from `0` to `i`.

4. **Negative Numbers?**
   - Prefix sums work with negative numbers because the difference `currentSum - (currentSum - target)` still equals `target`, regardless of whether sums are positive or negative.

