# Problem Statement
**252. Meeting Rooms**

Given an array of meeting time intervals consisting of start and end times `[[s1,e1],[s2,e2],...]` (si < ei), determine if a person could attend all meetings.

**Example 1:**

```
Input:[[0,30],[5,10],[15,20]]
Output: false
```

**Example 2:**

```
Input: [[7,10],[2,4]]
Output: true
```

**NOTE:** input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

# Problem Explanation

Let's think about how to come up with the solution to solve this problem.<br>
Suppose the input is following:<br>
```
[[5 9] [18 20] [15 18]]
```
We can model this as a time line:
```
      5        9
                               18  20
                           15  18
|------------------------------------------------------------------------------------|
```
From this time line, we can see that there's no overlap between the meetings.<br>
But can we change it a little bit to look better?
```
      5        9
                          15  18
                               18  20
|------------------------------------------------------------------------------------|
```
Now, it's better to see that there's no overlap between the meetings.<br>
What we do is basically sorting the meetings by the start time.<br>
Then, we check if the start time of the current meeting is greater than the end time of the previous meeting.<br>
If it is, then there's NO overlap, and we continue.<br>
Otherwise, we return false.<br>

In this example, suppose we're at (15, 18)<br>
We can see that the start time of the current meeting(15) is greater than the end time of the previous meeting(9).<br>
So there's NO overlap, and we continue.<br>

# Complexity Analysis
## Time Complexity O(nlogn)
- The time complexity is O(nlogn) because we need to sort the meetings by the start time.<br>
## Space Complexity O(1)
- The space complexity is O(1) because we don't need to use any extra space.
