# Problem Explanation

The core idea is to find all the partition of the input string

Along with the process, we only keep exploring any further partition if the current snippet is palindrome

## Choices and Constraints

- **Choice:** Decompose from index ~ length of input
  - For example, "aab"
    - If we're at index 0
      - We can decompose from index 0, 1, 2 -> "aab"
      - We can decompose from index 0, 1    -> "aa"
      - We can decompose from index 0       -> "a"
    - If we're at index 1
      - We can decompose from index 1, 2    -> "ab"
      - We can decompose from index 1       -> "a"
    - If we're at index 2
      - We can decompose from index 2       -> "b"
- **Constraint:** Each of snippet(substring) have to be palindrome
- **Goal:** Index is out of the bound of input string

## Recursive Tree Visualization
<pre>
                                                "aab"
                   "a"                          "aa"                       "aab"
        "a"           "ab"                       "b"
  "b"   
</pre>
For any working index, we can decompose from the working index to the end of input string

For example, input string is "aab"<br/>
At index 0,
- s.substr(0, 1) = "a"
  - select "a", and can keep decomposing from index 1("ab")
- s.substr(0, 2) = "aa"
  - select "aa", and can keep decomposing from index 2("b")
- s.substr(0, 3) = "aab"
  - select "aab", we can't keep decomposing

Continue to index 1(s.substr(0, 1) = "a")<br/>
The remaining string is "ab"
- s.substr(1, 1) = "a"
  - select "a", and can keep decomposing from index 2("b")
- s.substr(1, 2) = "ab"
  - select "ab", we can't keep decomposing

Continue to index 2(s.substr(1, 1) = "a")
- s.substr(2, 1) = "b"
  - select "b", we can't keep decomposing


Here, **each level doesn't perfectly represent the workind index**<br/>
For example,<br/>
At second level,<br/>
"a" -> working index up to 1<br/>
"ab" -> working index up to 2<br/>
"aab" -> working index up to 3<br/>

At third level,<br/>
"a" -> working index up to 2<br/>
"ab" -> working index up to 3<br/>
"b" -> working index up to 2<br/>

As we can see, the index point is not perfectly aligned with the level<br/>
So it's not a perfect tree

## Caveats
### Caveat 1: The logic of substr method
```c++
s.substr(index, i);
```
`index` is the starting position of the substring, and `i` is the length of the substring.
- "123456789".substr(0, 1) -> "1"
- "123456789".substr(0, 2) -> "12"
- "123456789".substr(3, 1) -> "4"
- "123456789".substr(3, 2) -> "45"
- "123456789".substr(3, 3) -> "456"

```c++
string subs = s.substr(index, i - index + 1);
```
- `index` is the starting position of the substring
- `i` is the workind pointer
- `i - index + 1` is the length of the substring


# Complexity Analysis

n = the legnth of input string

## Time Complexity: O(n^n * n)
- Branching Factor = n
  - At worst, we can decompose from index 0 to n - 1
- Depth = n
  - At worst, we can decompose n times
- Each call stack = O(n)
    - Because of the palindrome check

## Space Complexity: O(n)