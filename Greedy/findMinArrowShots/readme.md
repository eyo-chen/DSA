# Problem Explanation

The problem is basically asking to find the minimum number of merging overlapping intervals.<br>
Suppose we have following inputs:<br>
[[10,16],[2,8],[1,6],[7,12]]<br>
```
0   1   2   3   4   5   6   7   8   9   10  11  12  13  14  15  16  17  18  19  20
                                        -------------------------
        -------------------------
    ---------------------
                             --------------------
```
The arrow can go through at 4 and 10, so the minimum number of arrows is 2.<br>

The idea to solve this problem is following:<br>
1. Sort the input list by the start of the interval.
2. Initialize the count of arrows to n, where n is the length of the input list.
   - Every time we find an overlapping interval, we can merge them and update the end of the interval.
2. Loop through the list and compare two intervals at a time
3. If two intervals are overlapping, we can merge them to be the new compared interval, and decrement the count of arrows by 1.
4. If two intervals are not overlapping, we can update the compared interval to be the new interval and continue the loop.

How to merge two intervals:<br>
[a1,a2] and [b1,b2] are two intervals.<br>
The new merged interval is [max(a1,b1),min(a2,b2)].<br>

Let's walk through the example:<br>
[[10,16],[2,8],[1,6],[7,12]]<br>
1. Sort the input list by the start of the interval.<br>
```
0   1   2   3   4   5   6   7   8   9   10  11  12  13  14  15  16  17  18  19  20
    ---------------------
        -------------------------
                             --------------------
                                        -------------------------
```
2. Initialize the count of arrows to 4.<br>
3. Loop through the list and compare two intervals at a time<br>

- First iteration:<br>
    - [1,6] and [2,8] are overlapping, so we can merge them to be [2,6].<br>
    - The new compared interval is [2,6].<br>
    - Decrement the count of arrows by 1.<br>

- Second iteration:<br>
    - [2,6] and [7,12] are not overlapping.<br>
    - The new compared interval is [7,12].<br>

- Third iteration:<br>
    - [7,12] and [10,16] are overlapping, so we can merge them to be [10,12].<br>
    - The new compared interval is [10,12].<br>
    - Decrement the count of arrows by 1.<br>

- Done.<br>

# Complexity Analysis
## Time Complexity O(NlogN)
- We need to sort the input list by the start of the interval, which takes O(NlogN) time.

## Space Complexity O(1)
- We are not using any extra space, so the space complexity is O(1).