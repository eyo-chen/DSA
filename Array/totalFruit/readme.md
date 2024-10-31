# Problem Explanation

Note that the problem is essentially asking for the longest(continuous) subarray with at most 2 distinct integers.

## Brute Force
The brute force approach is to iterate through the array and for each element, find the longest subarray that starts with that element and has at most 2 distinct integers. This approach has a time complexity of O(n^2) because for each element, we iterate through the rest of the array to find the longest subarray.

For example, for the array [1,2,3,2,2]<br>
We start with 1, find the longest subarray starting with 1 and has at most 2 distinct integers, which is [1,2].<br>
Then we start with 2, find the longest subarray starting with 2 and has at most 2 distinct integers, which is [2,3,2,2].<br>
Then we start with 3, find the longest subarray starting with 3 and has at most 2 distinct integers, which is [3,2,2].<br>
Then we start with 2, the subarray is [2,2], so the length is 2.<br>
Finally we start with 2, the subarray is [2], so the length is 1.<br>
So the longest subarray with at most 2 distinct integers is [2,3,2,2] with the length of 4.

For each iteration to find the longest subarray, we need a hash table to store the distinct integers.<br>

### Complexity Analysis
#### Time Complexity O(n^2)
- For each iteration, we need to iterate through the rest of the array to find the longest subarray, so the time complexity is O(n^2).
#### Space Complexity O(n)
For each iteration, we need a hash table to store the distinct integers, so the space complexity is O(n).

## Two Pointers
The two pointers approach is to use two pointers to represent the start and end of the subarray. We move the right pointer to the right and add the current element to the hash table. If the hash table has more than 2 distinct integers, we move the left pointer to the right until the hash table has at most 2 distinct integers.

Note that we need to store the frequency of the integers in the hash table.<br>
Suppsoe the array is [1,3,1,3,2a,2b,4,4]<br>
When we're at index 4 (value 2a), the hash table is {1:2, 3:2, 2:1}.<br>
We know that we have to move the left pointer to the index 3 (value 3)<br>
But how should we move it?<br>
We can simply keep moving the left pointer to the right until the length of hash table is 2.<br>
Whenever we move the left pointer, we decrement the frequency of the integer at the left pointer.<br>
If the frequency of the integer at the left pointer is 0, we remove it from the hash table.<br>

### Complexity Analysis
#### Time Complexity O(n)
- For each element, we process it at most twice, so the time complexity is O(n).
#### Space Complexity O(1)
- Even though we use a hash table, but the size of the hash table is at most 2, so the space complexity is O(1).
