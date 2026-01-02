# Problem Explanation
The core idea to solve this problem is to think "How can I validate a substring is a valid string?"<br>
For example, <br>
- Is substring "AABAB" a valid string when k = 2?<br>
  - Yes, because we can replace two "B" with "A".<br>
- Is substring "AABAB" a valid string when k = 1?<br>
  - No, because we only have one replacement and we need to replace two "B" to make it a valid string.<br>

So, what's the formula to check if a substring is valid?<br>
We can think "Can I make k times of replacement to make the strings all same?"<br>
Suppose the substring is "AABAB", k = 2<br>
Should we use "A" to replace "B" or use "B" to replace "A"?<br>
It's pretty obvious that we should use "B" to replace "A" because "B" is less frequent than "A".<br>
So, the idea is that we want to replace the character that is NOT the most frequent character in the current window.<br>
If the number of character we want to replace is less than or equal to k, then the substring is valid.<br>
In other words, the formula is:<br>
(length of substring) - (largest frequency of character) <= k

For example, substring "AABAB" has following frequency map {"A": 3, "B": 2}<br>
So, the formula is 5 - 3 <= 2, which is true<br>
If the substring is "ACABAB", the frequency map is {"A": 3, "B": 2, "C": 1}<br>
So, the formula is 6 - 3 <= 2, which is false<br>

The idea is something like this:<br>
"Hey, because there's no way I replace the longest character with other character to make the string valid, so why not just try to replace other characters to make the string valid?"

Now, we know how to validate a substring is valid string<br>
Next, how we update the window?<br>
The idea is that we keep updating right pointer to the right in each iteration<br>
If the current window is valid substring, we simply update the answer<br>
If the current window is not valid substring,<br>
we need to move the left pointer to the right until the current window is valid substring<br>
or the left pointer is equal to right pointer<br>

# Complexity Analysis
## Time Complexity O(n)
- where n is the length of string s
- we have to iterate the string s in each iteration

## Space Complexity O(1)
- we use a frequency array to store the frequency of each character in the current window
- the frequency array will only contain 26 character at most



