# Problem Explanation

Let's break down the problem one by one<br>

First, how can we tell if two intervals overlap? ([a, b], [c, d])<br>
Two intervals overlap if c < b (second_end < first_start)<br>
In this case, we can guarantee the two intervals overlap<br>

Second, how can we know which interval to remove if we know two intervals overlap?<br>
We can remove the interval with the larger end value<br>
Why is that?<br>
Because the problem asks us to remove the minimum number of intervals<br>
If we remove the interval with the larger end value, that means we reduce the chance of overlapping with other intervals<br>
Therefore, whenever we find two intervals overlap, we remove the one with the larger end value<br>

Finally, how can we find the minimum number of intervals to remove?<br>
This is hard to come up with a solution at first glance<br>
But we can simply sort the intervals by their first element<br>
Then we can iterate through the intervals and check if the current interval overlaps with the previous interval<br>

Let's see the example below:
Input: [[1,2],[2,3],[3,4],[1,3]]
```
0   1   2   3   4   
    -----
        -----
            -----
    ---------
```

We first sort the intervals by their first element<br>
[[1,2],[1,3],[2,3],[3,4]]<br>
```
0   1   2   3   4   
    -----
    ---------
        -----
            -----
```
Then we iterate through the intervals<br>
- First, we compare [1, 2] and [1, 3]<br>
  - they overlap because 1 < 2 (second_end < first_start)<br>
  - we remove [1, 3] because 3 is the larger end value<br>

- Second, we compare [1, 2] and [2, 3]<br>
  - they do not overlap because 2 is not less than 2<br>

- Third, we compare [1, 2] and [3, 4]<br>
  - they do not overlap because 3 is not less than 2<br>


Let's summarize the steps<br>
1. Sort the intervals by their first element
2. Iterate through the intervals
3. If the current interval overlaps with the previous interval, remove the one with the larger end value
4. Return the number of intervals we removed

# Complexity Analysis
## Time Complexity O(NlogN)
- We sort the intervals by their first element O(NlogN)
- Then we iterate through the intervals O(N)

## Space Complexity O(1)
- We do not use any extra space