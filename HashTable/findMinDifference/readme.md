# Problem Explanation

The most hard part of this problem is how to handle the wraparound difference between the first and last time point.<br>
For example: ["23:59", "00:01"]<br>
The difference can be 23*60+58 or 2<br>
It's easier to see the first case, but how to calculate the second case?<br>

The key point is to recognize that time is NOT a linear structure, it's a circular structure.<br>
What does it mean of "00:01"?<br>
We can think of it as "24:01"<br>
And the difference between "00:01" and "24:01" is exactly 24 hours<br>
So, if we add 24*60 to "00:01", we can get "24:01", which we can use to calculate the difference with other time points<br>

The core idea to handle the wraparound difference is following:
1. Find the minimum and maximum time points
2. Calculate (24*60 + miVal) - maxVal

## Brute Force
The idea is pretty straightforward:
1. Iterate through all time points
2. For each time point, iterate through all other time points to find the minimum difference
3. Update the minimum difference
4. Handle the wraparound difference

For each time stamp, we convert it to minutes, and compare it with other time stamps<br>

### Complexity Analysis
#### Time Complexity O(n^2)
- We have two nested loops, each iterating through all time points

#### Space Complexity O(1)
- We only use a few variables to store the minimum and maximum time points and the minimum difference

## Sorting
The idea is also straightforward:
1. Convert all time points to minutes
2. Sort the time points
3. Calculate the difference between each time points
4. Handle the wraparound difference

After sorting, we can guarantee that the minimum difference is between two adjacent time points<br>
So, we just need to find the minimum difference between two adjacent time points, and the minimum difference between the first and last time points<br>

### Complexity Analysis
#### Time Complexity O(nlogn)
- We need to sort the time points

#### Space Complexity O(n)
- We need to store the time points in a list

## Hash Table
This is more hard to think of, but it's a good way to solve the problem<br>
The core idea is to regonize that the time points are limited, which is `24*60`<br>
No matter we're given how many time points, the range of time points is still `24*60`<br>

After marking the input time points in a hash table, like following:
```
   0     1    2    ...  23*60   24*60
[true, true, true ..., true, false]
```
How to find the minimum difference?<br>
Suppose we're at index 2, we do we care about?<br>
We only care about the previous time point, which is index 1<br>
We don't need to consider index 0 because we know the difference between index 1 and 2 is definitely smaller than the difference between index 2 and 0<br>

That's the key point of this solution<br>
Once we build the hash table,<br>
We can iterate through the hash table, and compare the current time point with the previous time point to find the minimum difference<br>


### Complexity Analysis
#### Time Complexity O(n)
- We need to iterate through all time points

#### Space Complexity O(24*60) / O(1)
- We need to store the time points in a hash table
