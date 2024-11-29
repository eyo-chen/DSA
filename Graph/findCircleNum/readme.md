# Problem Explanation
The core idea is similar to other graph problems, where we start from a node and traverse all connected nodes.<br>
However, in this problem, we are given a matrix, where the value at `matrix[i][j]` indicates if there is a direct connection between city `i` and city `j`.<br>
If there is a connection between city `i` and city `j`, then `matrix[i][j] = 1`, otherwise `matrix[i][j] = 0`.<br>
This is very important to understand, because it's very different from the typical graph problems where we are given a list of edges.<br>
Therefore, keep in mind that `matrix[i][j] = 1` means there is a direct connection between city `i` and city `j`.<br>

The idea to solve this problem is following:<br>
1. Use a `visited` array to keep track of the cities that have been visited.
   - The length of the `visited` array is the same as the length of the `matrix`, which is the number of cities.
2. Use a `provinces` counter to count the number of provinces.
3. Iterate through each city, and if it hasn't been visited, increment the `provinces` counter and use BFS to mark all the cities that are connected to it as visited.
4. Return the `provinces` counter.

The BFS function is similar to other BFS problems, where we use a queue to traverse the graph.<br>
We start from the current(input) city, add it to the queue and mark it as visited.<br>
Then, we iterate through all the cities, and if there is a connection between the current city and the city, and the city hasn't been visited, we add it to the queue and mark it as visited.<br>
This is the way to traverse all the cities that are connected to the current city.<br>

# Complexity Analysis

## Time Complexity: O(n²)
- The main `findCircleNum` function iterates through all n cities once
- For each unvisited city, we perform a BFS
- In the worst case, BFS visits each city and checks all its connections
- The inner loop in BFS checks all n possible neighbors for each city
- Even though we use `visited` to prevent revisiting cities, we still need to check all connections in the adjacency matrix
- Total: O(n) * O(n) = O(n²)

## Space Complexity: O(n)
- `visited` array: O(n) space
- BFS queue: O(n) space in worst case (when all cities are connected in a line)
- Total: O(n)

