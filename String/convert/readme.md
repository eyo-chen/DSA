# Problem Explanation
The idea is not that hard, but the implementation is a bit tricky.<br>
What we need is a slice of strings<br>
Suppose the input string is "ABCDEFG" and numRows is 3<br>
And we just need to build the table like this:
```
[
  "A,E"
  "B,D,F"
  "C,G"
]
```
The length of slice is the input numRows, and each element in the slice is a string which represents different column.<br>
Finally, we just need to concatenate each string in the slice to get the result.

So, the problem becomes how to build the table like this.
There are multiple ways to do this, I use two approaches to solve it.

Approach 1 (I came up with)<br>
The core idea is to let the row over the edge and go to the other side.<br>
For example, when input numRows is 3,<br>
When row is equal to numRows, we go up two rows and change the direction.<br>
When row is equal to -1, we go down two rows and change the direction.<br>

Approach 2<br>
The core idea is to change the direction when the row hits the edge.<br>
For example, when input numRows is 3,<br>
When row is equal to numRows - 1, we go up one row and change the direction.<br>
When row is equal to 0, we go down one row and change the direction.<br>

# Complexity Analysis
## Time Complexity O(n)
- where n is the length of the input string
## Space Complexity O(r)
- where r is the input numRows
- we need to store the result in a slice of strings


