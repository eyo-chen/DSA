# Problem Explanation

This solution is intuitive. For each call stack, we only have three choices: 

1. Decompose one substrings
2. Decompose two substrings
3. Decompose three substrings

Because the valid IP address is always between 0 and 255. 

For example, for the input "123456789":
- First choice: decompose one substring -> "1" ("1.23456789")
- Second choice: decompose two substrings -> "12" ("12.3456789")
- Third choice: decompose three substrings -> "123" ("123.456789")

We won't decompose four substrings("1234") because that would result in an invalid ID address.

After decomposing one substring, we can keep doing the same thing:

- For "1.23456789":
  - First choice: decompose one substring -> "2" ("1.2.3456789")
  - Second choice: decompose two substrings -> "23" ("1.23.456789")
  - Third choice: decompose three substrings -> "234" ("1.234.56789")

... and so on.

Another constraint is that we can only have four segments in the ID address:

- "123.456.78.9" -> valid
- "123.4.5.67.89" -> invalid
  - Though each section is between 0 and 255, the entire ID address can only have four segments

So, we won't always need to decompose the whole string to the end. We only need to decompose **THREE** times to get four segments.

See the following example, we only do two decomposition, but now we have three segments
- "1.2.3456789"
- "1.23.456789"
- "1.234.56789"

All of them are three segements, and only do two decomposition

## Choices and Constraints

- **Choice:** Three decompositions, 1, 2, or 3.
- **Constraint:** Each segment should be between 0 and 255, and can't have leading zeros.
- **Goal:** Decompose three times in total to get the four segments (base case).

## Recursive Tree Visualization
<pre>
                                      "245123888"
                   "2. ....."          "24. ....."         "245. ....."
 "2.4. ..."  "2.45. ..."  "2.451. ..."
</pre>

## Caveats
### Caveat 1: Solidify the base case
```c++
if (ans.size() == 4 && index == s.length()) {
  // ...
}
```
When do we know we have a valid IP address?
- We have four segments
  - `ans.size() == 4`
  - We go down the recursion tree three times
- We have explored the entire string
  - `index == s.length()`
  - It's eaiser to forget this condition
  - Without this condition, it gets the following wrong output
    - `2.5.5.2 | 2.5.5.25 | 2.5.5.255`

### Caveat 2: The logic of substr method
```c++
s.substr(index, i);
```
`index` is the starting position of the substring, and `i` is the length of the substring.
- "123456789".substr(0, 1) -> "1"
- "123456789".substr(0, 2) -> "12"
- "123456789".substr(3, 1) -> "4"
- "123456789".substr(3, 2) -> "45"
- "123456789".substr(3, 3) -> "456"


### Caveat 3: Check the looping condition
```c++
for (int i = 1; i <= 3 && index + i <= s.length(); i++) {
  // ...
}
```
- `i <= 3`
  - We can only decompose one, two, or three substrings
- `index + i <= s.length()`
  - We can't go beyond the end of the string
  - It's eaiser to forget this condition
  - Why does this condition matter?
    - Suppose the input string is "123456789"
    - When index is 8, we're at "9" position
      - When i = 1, s.substr(8, 1) -> "9"
      - When i = 2, s.substr(8, 2) -> "9"
      - When i = 3, s.substr(8, 3) -> "9"
    - As we can see, when i = 2 or 3, we're going beyond the end of the string
    - So it will end up having duplicate output  

# Complexity Analysis

n = the legnth of input string

## Time Complexity: O(1)
- Branching Factor = 3
   - Because each segment can be between 0 and 255
   - So we can only decompose one, two, or three substrings
- Depth = 3
    - Because we can only decompose three times to get four segments
- Each call stack = O(1)
    - Because we only do constant work
- Time Complexity = O(3^3) = O(1)
  - It doesn't mean it's fast
  - It means the time complexity is not related to the input size
  - No matter how long the input string is, we only do three decompositions to get four segments

## Space Complexity: O(1)
- Each call stack = O(1)
- We at most goes down three levels
