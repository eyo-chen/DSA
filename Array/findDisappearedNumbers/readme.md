# Problem Explanation

## Brute Force
The idea is to iterate from 1 to `len(nums)` and check if the number is in `nums`. If it is not, add it to the result.<br>
Because the problem statement mentions that the numbers in `nums` are between 1 and `len(nums)`, we can use this property to our advantage.<br>
For example, nums = [4,3,2,7,8,2,3,1]<br>
Because the length of `nums` is 8, the numbers that should be in `nums` are 1 through 8.<br>
We iterate from 1 to 8 and check if the number is in `nums`. If it is not, add it to the result.<br>

### Complexity Analysis
#### Time Complexity O(n^2)
#### Space Complexity O(1)

## Hash Table
The idea is to use a hash table to represent if a number is in `nums`.<br>
We can create a slice of booleans to represent the hash table.<br>
For example, nums = [4,3,2,7,8,2,3,1], hashTable = [f,f,f,f,f,f,f,f]<br>
We iterate through `nums` and set the index of the number to true in the hash table.<br>
After iterating through `nums`, the hash table will be [t,t,t,t,f,f,t,t]<br>
Then, we just iterate through the hash table and add the index to the result if the value is false.<br>

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(n)

## Sorting With Swapping
The idea is to sort the array with swapping.<br>
The detail can go to check the problem of `findDuplicates`.

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)

## Marking With Negatives
The idea is to set the value to negative as the index to mark the number as seen.<br>
For example, nums = [4,3,2,7,8,2,3,1]<br>
We iterate through `nums` and mark the number as seen by setting the value at the index of the number to negative.<br>
When we hit `4`, we go to index 4(4 - 1), and set nums[4] to -7<br>
When we hit `3`, we go to index 3(3 - 1), and set nums[3] to -7<br>
So on and so forth...<br>
After iterating through `nums`, the array will be [-4,-3,-2,-7,8,2,-3,-1]<br>
Then, we just iterate through the array and add the index to the result if the value is positive.<br>

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)
