# Problem Explanation
The problem should be straightforward.

We just need to loop through the numbers from the end to the beginning and add them together.<br>
We also have to take care of the carry.<br>

After the loop, if there is a carry left, we need to append it to the answer.<br>
Finally, we need to reverse the answer because we have been appending the digits from the end to the beginning.

There are four things to note:
1. We can use byte slice to store the answer
   - It's easier to reverse and convert to string at the end
2. We can subtract '0' from the character to convert it to integer
   - e.g. '6' = 54, '0' = 48, so '6' - '0' = 54 - 48 = 6
3. % gives us the last digit of the number
   - e.g. 18 % 10 = 8
4. / gives us the carry
   - e.g. 18 / 10 = 1

# Complexity Analysis
## Time Complexity O(max(N, M))
- N is the length of num1
- M is the length of num2
- We loop through both numbers once, so the time complexity is O(max(N, M))

## Space Complexity O(1)
- We don't use any additional space except for the answer, which is a byte slice
- So the space complexity is O(1)