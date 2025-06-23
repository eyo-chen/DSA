# Problem Explanation

Do not try to over complicate this problem<br>
We just need to go left or right based on the value of the node<br>
If we hit a null node, we insert the new node there<br>
That's it<br>

```
    4
   / \
  2   7
```
If we want to insert 5, we simply insert it to the left of 7<br>
```
    4
   / \
  2   7
     /
    5
```

If we want to insert 3, we simply insert it to the right of 2<br>
```
    4
   / \
  2   7
   \
    3
```

# Complexity Analysis
## Time Complexity O(n) (O(log n) for balanced tree)
- In most cases(when the tree is balanced), we will be able to insert the node in O(log n) time
- In the worst case, we will have to traverse the entire tree to insert the node, which will take O(n) time
```
      7
        9
         12
           16
```

## Space Complexity O(h)
- The space complexity is O(h), where h is the height of the tree