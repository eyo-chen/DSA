# Problem Explanation

## Brute force
The brute force approach is to iterate through the array and check the absolute difference between every pair of numbers.

How do we avoid duplicate pairs?<br>
We use hash table to store unique pairs.<br>
Every time we find a pair, <br>
We first find the smaller and larger number in this pair.<br>
Then we store this pair in the hash table.<br>
Because store two same keys in a hash table is still considered as one pair.<br>
So it helps us to avoid duplicate pairs.<br>

### Complexity Analysis
#### Time Complexity O(n^2)
- where n is the length of the array.
- the outer loop runs n times, and the inner loop runs n times in the worst case.

#### Space Complexity O(m)
- where m is the number of unique pairs.
- In the worst case, all pairs are unique, so m is equal to n(n-1)/2.


## Hash table(2 passes)
The idea is to use a hash table to store the number of times a number appears in the array.<br>
Then we iterate through the hash table and check if the number plus k exists in the hash table.<br>
If it does, we add it to the unique pairs.<br>

How does this avoid duplicate pairs?<br>
The key point is that we first store the number of times a number appears in the hash table.<br>
[3,1,4,1,5] -> {3:1, 1:2, 4:1, 5:1}<br>
When we iterate through the hash table, we won't iterate through the same number twice.<br>

Another key point is that we only need to check if the [num + k] exists in the hash table.<br>
If we also consider [num - k], we might end up with duplicate pairs.<br>
For example, [3,1,4,1,5] -> {3:1, 1:2, 4:1, 5:1}<br>
First, when encountering 3, we check both 3+2 and 3-2 in the hash table.<br>
We add (3,1) and (3,5) to the unique pairs.<br>
Later, when encountering 5, we check both 5+2 and 5-2 in the hash table.<br>
We add (5,3) to the unique pairs.<br>
We end up with duplicate pairs (3,5) and (5,3).<br>
That's why we only need to check if the [num + k] exists in the hash table.<br>



### Complexity Analysis
#### Time Complexity O(n)
- where n is the length of the array.
- The first pass is to store the number of times a number appears in the array, which takes O(n) time.
- The second pass is to iterate through the hash table and check if the number plus k exists in the hash table, which takes O(n) time.

#### Space Complexity O(n)
- where n is the length of the array.
- In the worst case, all numbers are unique, so we need to store all numbers in the hash table.


## Hash table(1 pass)
We build the hash table while iterating through the array.<br>
- If the current number exists in the hash table, we skip it.<br>
  - We only need to check if k is 0 and the number appears more than once.
- If the current number haven't been checked before, we check its [num + k] and [num - k] pair.
  - If the pair exists in the hash table, we update the answer.
  - We store the current number in the hash table.

How does this avoid duplicate pairs?<br>
The main reason is that we first check if the number exists in the hash table.<br>
If the number exists in the hash table, it means that we've checked its pair already.<br>
So we don't need to check it again.<br>
For example, [3,1,4,1,5]<br>
When we first encounter 1, we check its pair 1+2 and 1-2 in the hash table.<br>
When we encounter 1 again, we only check the case when k is 0, then skip it.<br>
So we don't end up with duplicate pairs.

### Complexity Analysis
#### Time Complexity O(n)
- where n is the length of the array.
- The first pass is to store the number of times a number appears in the array, which takes O(n) time.

#### Space Complexity O(n)
- where n is the length of the array.
- In the worst case, all numbers are unique, so we need to store all numbers in the hash table.
