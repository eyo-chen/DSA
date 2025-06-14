# Problem Explanation

## Brute Force (FourSumCount)
The brute force approach iterates over all possible combinations of indices (i, j, k, l) from the four arrays using four nested loops. For each combination, it checks if nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0. If true, it increments a counter.

### How It Works
Use four nested loops to iterate through all indices i, j, k, and l from 0 to n-1. <br>
For each tuple (i, j, k, l), compute the sum nums1[i] + nums2[j] + nums3[k] + nums4[l].<br>
If the sum equals 0, increment the result counter.<br>

### Complexity Analysis
#### Time Complexity: O(n^4)
- there are n choices for each of the four indices, leading to n × n × n × n = n^4 iterations.
#### Space Complexity: O(1)
- only a single counter variable is used.


## Hash Map (FourSumCount1)
The optimized solution uses a hash map to reduce the time complexity by grouping the arrays into two pairs: (nums1, nums2) and (nums3, nums4). <br>
It counts all sums of pairs from nums1 and nums2, then finds complementary sums from nums3 and nums4 that add to zero. <br>

### How It Works
Step 1: Build Hash Map for nums1 and nums2:<br>
Iterate over all pairs (i, j) from nums1 and nums2.<br>
Compute the sum nums1[i] + nums2[j] and store it in a hash map (sumCount), where the key is the sum and the value is the count of how many times that sum occurs.<br>

Step 2: Check nums3 and nums4:<br>
Iterate over all pairs (k, l) from nums3 and nums4.<br>
Compute the sum nums3[k] + nums4[l].<br>
Look for the complementary sum -(nums3[k] + nums4[l]) in the hash map. Add the count of this sum to the result, as it represents the number of tuples (i, j, k, l) where nums1[i] + nums2[j] + nums3[k] + nums4[l] = 0.<br>



### Why It Works
The solution ensures all valid tuples are counted because:<br>

It considers all pairs (i, j) from nums1 and nums2, covering every possible sum.<br>
For each pair (k, l) from nums3 and nums4, it checks if there exists a sum from nums1 and nums2 that makes the total zero.<br>
The hash map efficiently matches complementary sums, ensuring no valid tuple is missed, including special cases like i == j == k == l or where nums1[i] + nums3[k] == 0 and nums2[j] + nums4[l] == 0.<br>


### Complexity Analysis

#### Time Complexity: O(n^2)
- Computing sums for nums1 and nums2 takes n × n = n^2 iterations
- Checking sums for nums3 and nums4 also takes n^2 iterations, with O(1) hash map lookups.

#### Space Complexity: O(n^2)
- the hash map may store up to n^2 unique sums from nums1 and nums2.

