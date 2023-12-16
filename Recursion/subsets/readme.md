# Problem Explanation

For each recursive call stack, we can either choose the element or unchoose the element

For example, if input is [1,2,3]<br/> 
One of the ouput is [ ], it means we unchoose three times<br/> 
One of the output is [1], it means choose the one, and unchoose two times



## Choices and Constraints

- **Choice:** Choose or unchoose the character
- **Constraint:** None
- **Goal:** Our index pointer is out of the bound, aka is equal to the length of input

## Recursive Tree Visualization
If the input is [1,2,3], the recursive tree would be like
<pre>
                                       [1,2,3]                                   -> 0
                     []                                  [1]                     -> 1st
             []                [2]               [1]              [1,2]          -> 2nd
      []         [3]     [2]       [2,3]     [1]   [1,3]       [1,2]   [1,2,3]   -> 3rd
</pre>
1st level => choose "1" or unchoose<br/> 
2nd level => choose "2" or unchoose<br/> 
3rd level => choose "3" or unchoose<br/> 


# Complexity Analysis

n = the legnth of nums

## Time Complexity: O(2^n)
- Branching Factor = 2
   - For each call stack, we always only explore 2 choices, choose or unchoose
- Depth = n
    - If the length of input is 4, then we'll have 4 height of the tree
- Each call stack = O(1)
    - Because we only do constant work
    - Note that in this part of code, we may do `O(n)` work
      ```c++
      if (index == nums.size()) {
        ans.push_back(tmp);
        return;
      }
      ```
      - Note that `std::vector<T>::push_back()` creates a **COPY** of the argument and stores it in the vector in c++
      - So the time complexity could be `O(n * 2^n)` if we count this part of code

## Space Complexity: O(n)
- The deepest height of recursive tree is n