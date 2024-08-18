# Problem Explanation

The core idea to solve this problem is <br>
For any given day, we always have two choices<br>
1. Sell the stock
2. Do not sell the stock

If we choose to sell the stock, what's the maximum profit we can make? <br>
We just suppose we buy at day[i - 1] and sell at day[i] <br>
So, the profit we can make is prices[i] - prices[i - 1] <br>
However, that's not the whole profit we can make. <br>
We need to consider the profit we made before day[i - 1] <br>
So, the total profit we can make is (day[i] - day[i - 1]) + total maximum profit[i - 1]<br>

If we choose not to sell the stock, what's the maximum profit we can make? <br>
Just see what's the maximum profit we can make at day[i - 1] <br>

Let's walk through an example to see how it works <br>
prices = [7, 1, 5, 3, 6, 4] <br>

Day 1, i = 0 <br>
- Sell the stock: profit = 0
  - because we can only buy and sell on the same day
- Do not sell the stock: profit = 0
- total maximum profit = 0

Day 2, i = 1 <br>
- Sell the stock: profit = -6
  - (day[i] - day[i - 1]) + total maximum profit[i - 1]
  - (1 - 7) + 0 = -6
- Do not sell the stock: profit = 0
- total maximum profit = 0

Day 3, i = 2 <br>
- Sell the stock: profit = 4
  - (day[i] - day[i - 1]) + total maximum profit[i - 1]
  - (5 - 1) + 0 = 4
- Do not sell the stock: profit = 0
- total maximum profit = 4

Day 4, i = 3 <br>
- Sell the stock: profit = 2
  - (day[i] - day[i - 1]) + total maximum profit[i - 1]
  - (3 - 5) + 4 = 2
- Do not sell the stock: profit = 4
- total maximum profit = 4

Day 5, i = 4 <br>
- Sell the stock: profit = 7
  - (day[i] - day[i - 1]) + total maximum profit[i - 1]
  - (6 - 3) + 4 = 7
- Do not sell the stock: profit = 4
- total maximum profit = 7

Day 6, i = 5 <br>
- Sell the stock: profit = 5
  - (day[i] - day[i - 1]) + total maximum profit[i - 1]
  - (4 - 6) + 7 = 5
- Do not sell the stock: profit = 7
- total maximum profit = 7

The final answer is 7

# Complexity Analysis
## Time Complexity O(1)
- We only need to iterate through the prices once

## Space Complexity O(1)
- We only need to store the total maximum profit