# Explanation

## Solution 1
Let's see one example to walk through the entire recursion process.<br>
```
                                       [1,2a,2b]                                           -> 0
                     []                                  [1]                               -> 1st  choose or not choose 1
             []                [2a]                  [1]                 [1,2a]            -> 2nd  choose or not choose 2a
      []         [2b]*     [2a]       [2a,2b]     [1]   [1,2b]*       [1,2a]   [1,2a,2b]   -> 3rd  choose or not choose 2b
```
Look at the position with asterisk `*`, `[2b]` and `[1,2b]`<br>
- For `[2b]`, it's identical to `[2a]`
- For `[1,2b]`, it's identical to `[1,2a]`

So, the next question is how to avoid this duplication?<br>
From the diagram, we can see that ***the duplication happens when we choose `2b` without choosing `2a` first.***<br>
That means, ***if we have two identical numbers, we can only choose the second one if we have chosen the first one.***<br>
In order to implement this logic, we need to sort the input array first.<br>

```go
	if idx-1 >= 0 && nums[idx-1] == nums[idx] && !hashTable[idx-1] {
		return
	}
```
This is the code snippet to implement the above logic.<br>
We can't choose the current value if:
- The current value is the same as the previous value
- The previous value is not chosen


## Solution 2
Solution 2 is very similar to Solution 1.<br>
But we choose the value first, let's see the diagram again.<br
```
                                            [1,2a,2b]                                           -> 0
                         [1]                                    []                              -> 1st  choose or not choose 1
             [1,2a]                 [1]              [2a]                 []                    -> 2nd  choose or not choose 2a
    [1,2a,2b]      [1,2a]     [1,2b]*    [1]     [2a,2b]   [2a]       [2b]*   []                  -> 3rd  choose or not choose 2b
```
Again, look at the position with asterisk `*`, `[1,2b]` and `[2b]`<br>

From the diagram, we can see that ***the duplication happens when we choose `2b` without choosing `2a` first.***<br>
The logic here is that once we choose a value, we can't choose the identical value in the next step.<br>

```go
	helper1(nums, idx+1, ans, append(cur, nums[idx]))
	for i := idx + 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			idx++
		} else {
			break
		}
	}
	helper1(nums, idx+1, ans, cur)
```
- We first choose the current value
- Then we skip all the identical values
