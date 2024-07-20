# Problem Explanation

The idea to solve this problem is straightforward.<br>

1. Iterate through the whole grid
2. If find 1, increment the count of islands and call the helper function to mark all the connected 1's to 0 or other number.

The helper function could be BFS or DFS. Here, I have used DFS.

# Complexity Analysis
## Time Complexity: O(M*N)
- Where `M` is the number of rows and `N` is the number of columns in the grid.
- We are visiting each cell at most once.

## Space Complexity: O(M*N)
- Where `M` is the number of rows and `N` is the number of columns in the grid.
- The space complexity is used by the recursive stack.
- In the worst case, the depth of the recursive stack could be `M*N`.