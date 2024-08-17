# Problem Explanation

At first, we can easily come up with a brute force solution. We can try all possible pairs of buy and sell days and find the maximum profit. But this solution will take O(n^2) time complexity. We can do better than this.<br>

We can solve this problem using dynamic programming. We can keep track of the minimum price so far and the maximum profit so far. We can iterate through the prices and update the minimum price and maximum profit accordingly. We can return the maximum profit at the end.<br>

The concept is like this:<br>
For each day, we ask two questions:<br>
1. What is the minimum price so far?
2. What is the maximum profit so far?<br>
   - To calculate the maximum profit so far, we can subtract the minimum price so far from the current price.<br>

# Complexity Analysis
## Time Complexity O(n)
## Space Complexity O(1)