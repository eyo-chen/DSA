# Problem Explanation

## Add new interval and sort by start time
The idea is very straightforward.<br>
We can simply add the new interval into the intervals array and sort it by the start time.<br>
Then we can iterate through the intervals array and merge the intervals if they overlap.<br>
If they don't overlap, we can add the current interval into the result array and move to the next interval.<br>
Finally, we can add the last interval into the result array.<br>

Let's focus on the first two steps.<br>
The idea is that if we can put the new interval into the correct position, we can merge the intervals easily.<br>
For example, if the input is following:
```
[[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
```
After adding the new interval, and sorting the intervals, the array will be:
```
1  2
     3  5
       4            8
               6  7
                    8  10
                           12  16
```
Now, it's much easier to merge the intervals.<br>


Let's focus on the merging part.<br>
Before thinking about merging, it's important to recognize that the intervals are already sorted by the start time.<br>
So, there's only two cases that we need to consider.<br>
1. The current interval overlaps with the next interval.<br>
``` 
1   4       current interval
  3   5     next interval
```
or
```
1       5   current interval
   2  3     next interval
```
In either case, we can merge them by this simple rule:
```
current[1] = max(current[1], next[1])
```
Why we only need to consider updating the end time of the current interval because we know the start time of the current interval is definitely less than the start time of the next interval.<br>
Why? Because the intervals are sorted by the start time.

2. The current interval doesn't overlap with the next interval.<br>
```
1   4            current
        6   8    next
```
In this case, we can simply add the current interval into the result array and move to the next interval.


Let's try to walk through the process.
```
1  2
     3  5
       4            8
               6  7
                    8  10
                           12  16
```
- i = 1
  - current = [1,2]
  - intervals[i] = [3,5]
  - do not overlap, add current to result, current = intervals[i] = [3,5]
  - ans = [[1,2]]
- i = 2
  - current = [3,5]
  - intervals[i] = [4,8]
  - overlap, current[1] = max(current[1], intervals[i][1]) = max(5, 8) = 8
- i = 3
  - current = [3,8]
  - intervals[i] = [6,7]
  - overlap, current[1] = max(current[1], intervals[i][1]) = max(8, 7) = 8
- i = 4
  - current = [3,8]
  - intervals[i] = [8,10]
  - overlap, current[1] = max(current[1], intervals[i][1]) = max(8, 10) = 10
- i = 5
  - current = [3,10]
  - intervals[i] = [12,16]
  - no overlap, add current to result, current = intervals[i] = [12,16]
  - ans = [[1,2],[3,10]]
- add current to result, ans = [[1,2],[3,10],[12,16]]

### Complexity Analysis
#### Time Complexity O(nlogn)
- Sorting the intervals by the start time takes O(nlogn) time.
- Iterating through the intervals takes O(n) time.
- So, the total time complexity is O(nlogn).

#### Space Complexity O(n)
- We need to create a new array to store the result.
- So, the space complexity is O(n).

## Iterate through the intervals and merge if overlap
This is a more efficient solution than the first one.<br>
The core idea is that we loop through the intervals, and try to merge the intervals if they overlap.<br>
In this approach, there are only three cases that we need to consider.<br>
1. The new interval is before the current interval.<br>
2. The new interval is after the current interval.<br>
3. The new interval is overlapping with the current interval.<br>

For the first case,
```
4    5
          6    7

[1,2] -> newInterval
```
In this case, we can simply add the new interval to the result array and append the rest of the intervals to the result array.<br>
Then, we return the result array.<br>
How do we know if the new interval is before the current interval?<br>
If `end time of new interval` is less than `start time of current interval`, then the new interval is before the current interval.<br>

For the second case,
```
1    3
          4    5

[6,7] -> newInterval
```
In this case, we can simply add the current interval to the result array and continue to the next interval.<br>
How do we know if the new interval is after the current interval?<br>
If `start time of new interval` is greater than `end time of current interval`, then the new interval is after the current interval.<br>

For the third case,
```
          4    5

[4,8] -> newInterval
```
In this case, we can merge the intervals by updating the start and end time of the new interval.<br>
We can simply update both the start time and end time of the new interval to the minimum start time and maximum end time of the new interval and the current interval.<br>
Note that we don't need to add interval into the result in this case because we are merging the intervals.<br>
Also, we not sure if this updated new interval will overlap with the next interval, so we need to continue to the next interval.<br>

### Complexity Analysis
#### Time Complexity O(n)
- We only need to iterate through the intervals once.
- So, the time complexity is O(n).

#### Space Complexity O(n)
- We need to create a new array to store the result.
- So, the space complexity is O(n).
