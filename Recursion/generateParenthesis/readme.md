# Problem Explanation

The key idea to solve this problem is to think the following points
- How to validate the parenthesis?
  - ***WE CAN ONLY ADD ")" WHEN THERE IS REMAINING "(" FOR US TO MATCH***
- When can add "(" ?
  - We can keep adding "(" until we reach the limit(n)
- When can add ")" ?
  - We can only add ")" when there is remaining "(" for us to match


Based on the above idea, we use two index to keep track of the number of "(" and ")" we have added
- We can add "(" when the leftIndex < limit(n)
- We can add ")" when the rightIndex < leftIndex

## Choices and Constraints

- **Choice:** We can either add "(" or ")" to the string
- **Constraint:** ***WE CAN ONLY ADD ")" WHEN THERE IS REMAINING "(" FOR US TO MATCH***
- **Goal:** When the string is equal to the limit(n * 2)

## Recursive Tree Visualization
When n = 2
<pre>
                                     ""
</pre>
At first level,<br/>
we can add "(", because leftIndex < limit(n)<br/>
we can't add ")", because rightIndex = leftIndex<br/>
<pre>
                                     ""
                                    "("
</pre>
At second level,<br/>
we can add "(", because leftIndex < limit(n)<br/>
we can add ")", because rightIndex < leftIndex<br/>
<pre>
                                     ""
                                     "("            
                  "(("                                "()"        
</pre>
At third level,<br/>
<pre>
                                     ""
                                     "("            
                  "(("                                "()"        
                      "(()"                 "()("            "())" x
</pre>
At fourth level,<br/>
<pre>
                                     ""
                                     "("            
                  "(("                                "()"        
                      "(()"                 "()("            "())"x
                       "(())"             "()()"     
</pre>


# Complexity Analysis

n = the legnth of input digits

## Time Complexity: O(2^(n*2))
- Branching Factor = 2
  - We only have two choices, "(" or ")"
- Depth = n * 2
  - If the length of input is 4, then we'll have 8 height of the tree
- Each call stack = O(1)
    - Because we only do constant work

## Space Complexity: O(n)
- The deepest call stack is n * 2