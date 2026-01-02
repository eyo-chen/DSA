# Problem Explanation

## Use Temporary Array
The idea is that we use a temporary array to store the compressed characters.<br>
We iterate through the original array,<br>
For each character, we count the number of consecutive characters.<br>
If the count is greater than 1, we add the character and the count to the temporary array.<br>
Otherwise, we only add the character to the temporary array.<br>
Finally, we copy the temporary array to the original array and return the length of the temporary array.<br>

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(n)

## Use Pointer
The idea is almost the same as the previous one, but we use a pointer to store the compressed characters.<br>
Instead of using a temporary array, we use a pointer<br>
This pointer is used to let use know where to start mutating the original array.<br>
For example,<br>
["a","a","b","b","c","c","c"] -> after compressed -> ["a",2,"b",2,"c",3]<br>
The pointer starts from 0<br>
After counting the number of consecutive characters for "a", i = 2 and pointer = 0<br>
We mutate the array[ptr] = "a" and ptr++, so now the pointer is 1<br>
Then, we mutate the array[ptr] = "2" and ptr++, so now the pointer is 2<br>
We continue the same process for "b" and "c"<br>

Let's clarify the purpose of each variable:
- `ptr`: it's used to know where to start mutating the original array
  - No matter is to mutate the value to current character or count
  - After mutating, we always update the pointer by 1
  - It's value also represents the length of the compressed array which is the return value of the function
- `i`: it just helps us to iterate through the array
  - Note that when we count the number of consecutive characters, we also update the value of `i`
  - For example, ["a","a","b","b","c","c","c"] -> ["a",2,"b",2,"c",3]
    - When counting the number of consecutive characters for "a", we update the value of `i` from 0 to 2
    - When counting the number of consecutive characters for "b", we update the value of `i` from 2 to 4
    - When counting the number of consecutive characters for "c", we update the value of `i` from 4 to 6
- `count`: it's used to count the number of consecutive characters

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)