# Problem Explanation

## Use HashTable to count the number of duplicates
The idea is to use a hash table to count the number of duplicates of each number.
If the number of duplicates is greater than 2, we skip it.
Otherwise, we write it to the array.

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(n)

## Use a write index to write the result to the array
The idea is to simply loop through the array, and write the number to the array at the write index.<br>
When the number is equal to the current number, we increment the duplicate counter.<br>
Else, we reset the duplicate counter, and update the current number.<br>
If the duplicate counter is less than or equal to 2, we write the number to the array at the write index, and increment the write index.<br>

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(1)
