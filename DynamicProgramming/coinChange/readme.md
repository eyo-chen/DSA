# Problem Explanation

## Brute Force
The brute force approach is to try every coin and see if we can make change for the amount.<br>
The core idea is to choose a coin, and make the problem smaller by subtracting the coin from the amount.<br>

For example, suppose the coins are [1,2,5] and the amount is 11.<br>
We can first choose a coin, say 5, and the problem becomes 11 - 5 = 6.<br>
Then we continue ask "What's the minimum number of coins to make change for 6?"<br>
We can choose coin 1, and the problem becomes 6 - 1 = 5.<br>
Then we continue ask "What's the minimum number of coins to make change for 5?"<br
So on and so forth, until the amount is 0.

For each sub problem, we can choose any coin.

### Complexity Analysis
#### Time Complexity O(k^n)
- k is the number of coins, n is the amount.
- For each sub problem, we try every coin, so there's k choices.
- The depth of the recursion tree is n (the amount), so there's n levels.
- Therefore, the time complexity is O(k^n).

#### Space Complexity O(n)
- The depth of the recursion tree is n (the amount), so the call stack can go as deep as n.
- Therefore, the space complexity is O(n).

## Top-Down Dynamic Programming
To reduce the time complexity, we can use a memoization technique to store the results of the sub problems.
We store the results of the sub problems in a hash map, so we don't need to recompute the results of the sub problems.

#### Time Complexity O(k * n)
- We still have k choices for each sub problem.
- But now we don't have to compute the sub problems again.
- So the time complexity is O(k * n).

#### Space Complexity O(n)
- We still need to store the memoization table, so the space complexity is O(n).


## Bottom-Up Dynamic Programming
Instead going from the top to the bottom(amount to 0), we can go from the bottom to the top(0 to amount).

First create a DP table with length amount + 1, and fill it with a special value, say amount + 1.<br>
- The reason we need amount + 1 length is because<br>
  - For each element of the DP table, it represents the minimum number of coins needed for that amount.<br>
  - If the given amount is 5, we need [0, 1, 2, 3, 4, 5] DP table.<br>
    - For 0, it means "What's the minimum number of coins needed for 0?"<br>
    - For 1, it means "What's the minimum number of coins needed for 1?"<br>
    - For 2, it means "What's the minimum number of coins needed for 2?"<br>
    - .....<br>
    - For 5, it means "What's the minimum number of coins needed for 5?"<br>
    - Which is our target amount.
- The reason we need to fill it with amount + 1 is because<br>
  - amount + 1 is the maximum number of coins we can use to make change for the amount.
  - For example, if the coins are [1,2,5] and the amount is 11,<br>
    - The maximum number of coins we can use to make change for 11 is 11 coins (all 1 coins).
    - So we fill the DP table with 11 + 1 = 12.

How to fill the DP table?
The idea is that for each amount(element of the DP table), we want to find out what's the minimum number of coins needed for that amount.
Let's find out together. Suppose the coins are [1,2,5] and the amount is 11.
The DP table will be [0, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12].
As we discussed, the initial DP table is length amount + 1, and fill it with amount + 1.
Also, because the minimum number of coins needed for 0 is 0, so the first element of the DP table is always 0.
- For 1, what's the minimum number of coins needed for 1?
  - We can only choose coin 1 because 2 and 5 are greater than amount 1.
  - So the minimum number of coins needed for 1 is 1.
- For 2, what's the minimum number of coins needed for 2?
  - We can choose coin 1 and 2.
  - If we choose coin 1, the sub problem becomes "What's the minimum number of coins needed for 1(2-1)?"
    - We know that the DP[1] = 1, so the answer is DP[1] + 1 = 2.
  - If we choose coin 2, the sub problem becomes "What's the minimum number of coins needed for 0(2-2)?"
    - We know that the DP[0] = 0, so the answer is DP[0] + 1 = 1.
  - So the minimum number of coins needed for 2 is min(DP[1] + 1, DP[0] + 1) = 1.
- For 3, what's the minimum number of coins needed for 3?
  - We can choose coin 1 and 2.
  - If we choose coin 1, the sub problem becomes "What's the minimum number of coins needed for 2(3-1)?"
    - We know that the DP[2] = 1, so the answer is DP[2] + 1 = 2.
  - If we choose coin 2, the sub problem becomes "What's the minimum number of coins needed for 1(3-2)?"
    - We know that the DP[1] = 1, so the answer is DP[1] + 1 = 2.
  - So the minimum number of coins needed for 3 is min(DP[2] + 1, DP[1] + 1) = 2.

From the above, we know that the logic to fill up the DP table is:
- For each amount i,
  - Loop through each coin,
    - If the coin is less than or equal to the amount i,
      - The minimum number of coins needed for i is min(DP[i], DP[i - coin] + 1).
      - DP[i - coin] means "what's the minimum number of coins needed if I use the current coin?"
        - For example, if the coins are [1,2,5] and the amount is 11,
          - For amount 11, if I use coin 1, the sub problem becomes "what's the minimum number of coins needed for 10?"
          - For amount 11, if I use coin 2, the sub problem becomes "what's the minimum number of coins needed for 9?"
          - For amount 11, if I use coin 5, the sub problem becomes "what's the minimum number of coins needed for 6?"
      - + 1 means "I'm using one more coin."
  
  

### Complexity Analysis
#### Time Complexity O(k * n)
- k is the number of coins, n is the amount.
- For each amount i, we loop through each coin, so there's k choices.
- Therefore, the time complexity is O(k * n).

#### Space Complexity O(n)
- We need to store the DP table, so the space complexity is O(n).

