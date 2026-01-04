# Problem Statement
## 2274. Maximum Consecutive Floors Without Special Floors

Alice manages a company and has rented some floors of a building as office space. Alice has decided some of these floors should be special floors, used for relaxation only.

You are given two integers bottom and top, which denote that Alice has rented all the floors from bottom to top (inclusive). You are also given the integer array special, where special[i] denotes a special floor that Alice has designated for relaxation.

Return the maximum number of consecutive floors without a special floor.

Example 1:<br>
Input: bottom = 2, top = 9, special = [4,6]<br>
Output: 3<br>
Explanation: The following are the ranges (inclusive) of consecutive floors without a special floor:<br>
- (2, 3) with a total amount of 2 floors.<br>
- (5, 5) with a total amount of 1 floor.<br>
- (7, 9) with a total amount of 3 floors.<br>
Therefore, we return the maximum number which is 3 floors.<br>


Example 2:<br>
Input: bottom = 6, top = 8, special = [7,6,8]<br>
Output: 0<br>
Explanation: Every floor rented is a special floor, so we return 0.<br>
 

Constraints:<br>
1 <= special.length <= 10^5<br>
1 <= bottom <= special[i] <= top <= 10^9<br>
All the values of special are unique.<br>

# Solution Explanation
Both solutions are easy to understand; the tricky part is the time complexity.<br>
The first solution will encounter a *Time Limit Exceeded* error, while the second won't.<br>
Let's analyze the time complexity to understand why.<br>
Let's say bottom = b, top = t, length of special = n<br>

## Loop from bottom to top
- The first for-loop to build the hash table takes O(n)<br>
- The second for-loop to find the answer takes O(t-b)<br>
- The overall time complexity is O(n) + O(t-b)<br>

## Sort the special array
- The sorting part takes O(n log n)<br>
- The for-loop to find the answer takes O(n)<br>
- The overall time complexity is O(n log n) + O(n)<br>

## Analysis
Taking a closer look at the constraints, we can see:<br>
- 1 <= special.length <= 10^5
- 1 <= bottom <= special[i] <= top <= 10^9

This tells us that n ≤ 10^5, but (t-b) can be up to 10^9.<br>
The first solution has O(t-b), which can be extremely large and cause TLE.<br>
In contrast, the second solution only depends on the length of the special array (n).<br>
That's why the second solution is much more efficient for this problem.<br>

## Example
Consider: bottom = 1, top = 1,000,000,000, special = [500,000,000]
- Solution 1: ~1 billion iterations → TLE ❌
- Solution 2: ~17 operations (sorting + checking gaps) → Passes ✅