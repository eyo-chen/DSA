
# DP Approach
The original question ask us "how many ways can I form the word 'rabbit' using the string 'rabbbit'", we can break this down into smaller questions:
- "How many ways can I form 'rabbit' using 'rabbbi'?" (skipping the last character)
- "How many ways can I form 'rabbi' using 'rabbbi'?" (if the last characters match, which they do - both are 't')

Our final answer is just the sum of answers to these two smaller questions.<br>

At every step, when we're looking at a character in our source string, we face exactly two choices:
1. Skip this character
- We're saying "I don't need this character to build my target word. Let me see how many ways I can build the target using everything that came before this character."
2. Use this character (if it matches what I need)
- We're saying "This character is exactly what I need for the next letter in my target word. If I use it, then I need to figure out how many ways I could build the shorter target word using the earlier part of the source string."


Let's set up the DP table:
```
      "" r  a  b  b  i  t
""    0  0  0  0  0  0  0
r     0  0  0  0  0  0  0
a     0  0  0  0  0  0  0
b     0  0  0  0  0  0  0
b     0  0  0  0  0  0  0
b     0  0  0  0  0  0  0
i     0  0  0  0  0  0  0
t     0  0  0  0  0  0  0
```
row is the target string, column is the source string<br>

We can first initialize the first column to 1<br>
The first column asks us "How many ways can I form the empty string "" using "r", "ra", "rab", "rabb", .... "rabbbit""<br>
For all of the source string, the answer is 1 because we can always not include the current character<br>
```
      "" r  a  b  b  i  t
""    1  0  0  0  0  0  0
r     1  0  0  0  0  0  0
a     1  0  0  0  0  0  0
b     1  0  0  0  0  0  0
b     1  0  0  0  0  0  0
b     1  0  0  0  0  0  0
i     1  0  0  0  0  0  0
t     1  0  0  0  0  0  0
```

We leave the first row(except for the first cell) to 0<br>
The first row asks us "How many ways can I form the "r", "ra", "rab", "rabb", .... "rabbbit" using empty string"""<br>
It's obvious that the answer is 0 because there's no way we can form any of the target string using empty string<br>

Now, we can start building up the table.<br>
Based on previous observation, we know the formula is:<br>
```
dp[r][c] = dp[r][c-1] + (dp[r-1][c-1] if s[c-1] == t[r-1] else 0)
```
- `dp[r][c-1]` represents we are not using the current character from the source string<br>
  - For example, if the original problem is "how many ways can I form the word 'rabbit' using the string 'rabbbit'", we can choose to ignore the last 't' in the source string and just focus on forming 'rabbit' using 'rabbbi'.
- `dp[r-1][c-1]` represents we are using the current character from the source string<br>
  - For example, if the original problem is "how many ways can I form the word 'rabbit' using the string 'rabbbit'", we can choose to use the last 't' in the source string. After that, we need to find the number of ways to form 'rabbi' using 'rabbbi'.
  - Note that we can only use the current character from the source string if it matches the current character from the target string.