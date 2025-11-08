# Problem Explanation

## Using Recursion

The idea to solve this problem is easy.<br>
We traverse each node, and check if the current node is the same as the subtree we are looking for.<br>

For example, if we have the following tree:<br>
```
    3
   / \
  4   5
 / \
1   2
```

And we are looking for the subtree:<br>
```
  4
 / \
1   2
```

We first look at node 3, and ask Is the current node the same as the subtree we are looking for?<br>
If not, we go to the left and right child of node 3, and ask the same question.<br>
If yes, we return True.<br>

## Complexity Analysis
### Time Complexity O(N)
- Where N is the number of nodes in the tree.<br>

### Space Complexity O(H)
- Where H is the height of the tree.<br>

## Using Serialization
The idea is to serialize both trees and check if the serialized string of the subtree is a substring of the serialized string of the main tree.<br>
For example, if we have the following tree:<br>
```
    3
   / \
  4   5
 / \
1   2
```
And we are looking for the subtree:<br>
```
  4
 / \
1   2
```
We serialize both trees:<br>
- Main tree: "3,4,1,nil,nil,2,nil,nil,5,nil,nil"
- Subtree: "4,1,nil,nil,2,nil,nil"<br>

One thing to note that is that ***we can't use inorder traversal*** to serialize the tree, because it will not be able to distinguish between different structures of the tree.<br><br>
For example, the following two trees will have the same inorder traversal:<br>
```
      10         10
     /          /
    5          4
   /            \
  4              5
```
They both have the same serialized string: "nil,4,nil,5,nil,10,nil"<br>

<br>

Also, another thing to keep in mind is where to put the delimiter.<br>
```go
sb.WriteString(strconv.Itoa(root.Val))
sb.WriteString(",")
```
The bug happens when node values create false substring matches. For example:
```
root = [12]
subRoot = [2] 
```

If you serialize without **proper delimiters at the start**:
- `serialize([12])` = `"12,null,null"`  
- `serialize([2])` = `"2,null,null"`

The string `"2,null,null"` is contained in `"12,null,null"` (the '2' in '12'), so it returns `true` incorrectly!

### The Fix

Add a **delimiter before each value** to prevent partial matches:

```go
sb.WriteString(",")  // Add comma BEFORE the value
sb.WriteString(strconv.Itoa(root.Val))
```

Now:
- `serialize([12])` = `",12,null,null"`
- `serialize([2])` = `",2,null,null"`

The substring `",2,null,null"` is NOT in `",12,null,null"` ✓

This ensures that value `2` is only matched as a complete node value, not as a digit within `12`.