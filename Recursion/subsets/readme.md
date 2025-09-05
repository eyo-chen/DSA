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


## Go Slice Reference Bug in Backtracking Algorithms

### Problem Description

A subtle but critical bug commonly occurs in Go backtracking algorithms when generating combinations, subsets, or permutations. The issue stems from storing slice references instead of creating independent copies, leading to corrupted results due to shared underlying memory.

### Example: Buggy Subset Generation Code

```go
func subsets(nums []int) [][]int {
    ans := [][]int{}
    helper(nums, 0, &ans, []int{})
    return ans
}

func helper(nums []int, index int, ans *[][]int, tmp []int) {
    if index == len(nums) {
        *ans = append(*ans, tmp)  // ❌ BUG: Storing reference, not copy
        return
    }
    helper(nums, index + 1, ans, tmp)
    helper(nums, index + 1, ans, append(tmp, nums[index]))
}
```

### Root Cause: Slice Memory Sharing

#### The Core Issue

When you execute `*ans = append(*ans, tmp)`, you're not creating a copy of the `tmp` slice. Instead, you're storing a reference to the same underlying array that `tmp` points to. This becomes problematic because:

1. **Shared Memory**: Multiple "different" subsets in your result actually point to the same memory location
2. **Later Modifications**: As recursion continues, modifications to `tmp` affect previously "saved" subsets
3. **Unpredictable Results**: The final output depends on the last state of the shared memory

### Why This Bug Is Tricky to Catch

#### 1. Appears to Work for Small Inputs

```go
// Input: [1,2,3] might produce correct output:
// [[],[3],[2],[2,3],[1],[1,3],[1,2],[1,2,3]]
```

This happens because:
- Go's slice capacity management may allocate separate memory regions for small inputs
- The sequence of recursive calls might not trigger the specific conditions needed for corruption
- **This is essentially a "lucky coincidence" - the bug still exists**

#### 2. Becomes Obvious with Larger Inputs

Try the same code with `[1,2,3,4,5,6,7,8]` and you'll likely see:
- Duplicate subsets in the output
- Missing expected combinations
- Subsets containing unexpected values

### Debugging Tips

Add logging to expose the issue:

```go
func helper(nums []int, index int, ans *[][]int, tmp []int) {
    if index == len(nums) {
        fmt.Printf("Adding subset: %v (address: %p)\n", tmp, tmp)
        *ans = append(*ans, tmp)
        return
    }
    // ... rest of function
}
```

If you see the same memory address for different subsets, you've found the bug.
