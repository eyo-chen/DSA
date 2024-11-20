# Problem Explanation

## Brute Force
Use nested for loop and a hash table to solve this problem.<br>
Nested for loop is used to find the pair of numbers that sum up to k.<br>
Hash table is used to avoid duplicates.

For example, nums = [1, 2, 3, 4], k = 5
Start from the first number, 1, and find if there is a number that, when added to 1, equals to 5.
We loop through the rest of the numbers and find that 4 + 1 = 5.
Then we add 1 and 4 to the hash table to indicate that we have used them.
Then, we move to the next number, 2, and find if there is a number that, when added to 2, equals to 5.
We loop through the rest of the numbers and find that 3 + 2 = 5.
Then we add 2 and 3 to the hash table to indicate that we have used them.
We continue this process until we have checked all the numbers.

### Complexity Analysis
#### Time Complexity: O(n^2)
#### Space Complexity: O(n)

## Sorting and Two Pointers
Use sorting and two pointers to solve this problem.<br>
Sort the array first.<br>
Use two pointers to find the pair of numbers that sum up to k.<br>
One pointer starts from the beginning of the array and the other starts from the end of the array.<br>
If the sum of the two numbers is less than k, move the left pointer to the right.<br>
If the sum of the two numbers is greater than k, move the right pointer to the left.<br>
If the sum of the two numbers is equal to k, update the answer, and move both pointers to the middle and increment the answer.<br>

### Complexity Analysis
#### Time Complexity: O(n log n)
#### Space Complexity: O(1)

## Hash Table
Use a hash table to solve this problem.<br>
We use hash table to store the numbers and their counts.<br>
For each number, we check if the remaining number (k - n) is in the hash table.<br>
If it is, we increment the answer and decrement the count of the remaining number.<br>
If it is not, we increment the count of the number.<br>

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(n)
