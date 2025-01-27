# Problem Explanation

The idea is to first sort the intervals by the start time.<br>
Once we sort the intervals, we can iterate through the intervals and merge them if they overlap.<br>
For example, if the intervals are<br>
[[1,3],[2,6],[8,10],[15,18]]<br>
We can put the intervals into diagram, like this:<br> 
```
1   3
  2      6
            8    10
                       15  18
```
How can we know if two intervals overlap?<br>
***If the end time of the first interval is greater than or equal to the start time of the second interval, then the two intervals overlap.***<br>
For example, [1,3] and [2,6] overlap because 3 >= 2.

How can we merge the intervals?<br>
***If the two intervals overlap, we can merge them by taking the maximum of the end times of the two intervals.***<br>
For example, [1,3] and [2,6] overlap, so we can merge them into [1,6].<br>
Note that we don't need to consider the start time of the merged interval because we already sorted the intervals by the start time.<br>

Let's walk through the process of merging the intervals:<br>
Because we need to compare two intervals at a time, we can first initialize the first interval as the previous interval.<br>
prev = [1,3]<br>
Then, we start the iteration from the second interval.<br>
cur = [2,6]<br>
Check if two intervals(prev & cur) overlap, if they do, use the logic to merge them.<br>
They are overlapping, so we merge them into [1,6].<br>
prev = [1,6]<br>
Then, we continue the iteration.<br>
cur = [8,10]<br>
Check if two intervals(prev & cur) overlap, if they do, use the logic to merge them.<br>
They are not overlapping, so we add the previous interval to the result and update the previous interval to the current interval.<br>
prev = [8,10]<br>
Then, we continue the iteration.<br>
cur = [15,18]<br>
Check if two intervals(prev & cur) overlap, if they do, use the logic to merge them.<br>
They are not overlapping, so we add the previous interval to the result and update the previous interval to the current interval.<br>
prev = [15,18]<br>
Finally, we add the last interval to the result.<br>
ans = [[1,6], [8,10], [15,18]]<br>

Why do we need to sort the intervals by the start time?<br>
Let's see the case we don't sort the intervals by the start time.<br>
[[3,6], [1,4], [8,10], [2,5]]<br>
```
     3       6
1       4
                8      10
   2       5
```
We first check [3,6] and [1,4], they overlap, so we merge them into [1,6].<br>
Then, we check [1,6] and [8,10], they don't overlap, so we add [1,6] to the result and update the previous interval to [8,10].<br>
Finally, we check [8,10] and [2,5], they overlap, so we merge them into [2,10].<br>
The result is [[1,6], [2,10]]<br>

Let's see the case we sort the intervals by the start time.<br>
[[1,4], [2,5], [3,6], [8,10]]<br>
```
1       4
  2       5
    3       6
               8      10
```
It's pretty obvious that the final result is [[1,6], [8,10]]<br>
Which is different from the result we get when we don't sort the intervals by the start time.<br>
