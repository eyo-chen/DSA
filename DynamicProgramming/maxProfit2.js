//////////////////////////////////////////////////////
// ***  Best Time to Buy and Sell Stock II ***
//////////////////////////////////////////////////////
/**
 * @param {number[]} prices
 * @return {number}
 */
/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.

Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.

Example 2:
Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.

Example 3:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
 
Constraints:
1 <= prices.length <= 3 * 104
0 <= prices[i] <= 104
*/
/*
First, try to think the simple example, [1,3,6,2], k = 1
Try to first think the base case, or easy case
Day 1
We know that the maximum profit of first day is always gonna be 0, because we can't find any before day to buy the stock
So we only have one scenario
1. Buy at day 1, sell at day 1 => 1 - 1 = 0
=> maximum profit is 0 at day 1

Day 2
Now there are 1 + 2 scenarios
1. Buy at day 1, sell at day 1 => 1 - 1 = 0
2. Buy at day 1, sell at day 2 => 3 - 1 = 2
3. Buy at day 2, sell at day 2 => 3 - 3 = 0
=> maximum profit is 2 at day 2  

Day 3
Now there are 3 + 3 scenarios
1. Buy at day 1, sell at day 1 => 1 - 1 = 0
2. Buy at day 1, sell at day 2 => 3 - 1 = 2
3. Buy at day 1, sell at day 3 => 6 - 1 = 5
4. Buy at day 2, sell at day 2 => 3 - 3 = 0
5. Buy at day 2, sell at day 3 => 6 - 3 = 3
6. Buy at day 3, sell at day 3 => 6 - 6 = 0
=> maximum profit is 5 at day 3

Day 4
Now there are 3 + 8 scenarios
1. Buy at day 1, sell at day 1 => 1 - 1 = 0
2. Buy at day 1, sell at day 2 => 3 - 1 = 2
3. Buy at day 1, sell at day 3 => 6 - 1 = 5
4. Buy at day 1, sell at day 4 => 2 - 1 = 1
5. Buy at day 2, sell at day 2 => 3 - 3 = 0
6. Buy at day 2, sell at day 3 => 6 - 3 = 3
7. Buy at day 2, sell at day 4 => 2 - 3 = -1
8. Buy at day 2, sell at day 2 => 3 - 3 = 0
9. Buy at day 3, sell at day 3 => 6 - 6 = 0
10. Buy at day 3, sell at day 4 => 2 - 6 = -4
11. Buy at day 4, sell at day 4 => 2 - 2 = 0
=> maximum profit is 5 at day 4

This is bruth force mindset
How can we approve?

Day 1
=> maximum profit is 0 at day 1
=> What does this means?
=> It means we buy and sell the stock at the same day

Day 2
=> maximum profit is 2 at day 2 
=> It means we buy at day 1 and sell at day 2

Day 3
=> maximum profit is 5 at day 3
=> It means we buy at day 1 and sell at day 3

Day 4
=> maximum profit is 5 at day 4
=> It means we buy at day 1 and sell at day 3
=> In other words, we don't have any transaction at day 4

See the subproblem??
For any given day, we always have two choies
1. Sell at the given day
2. Don't do the transaction

Let's deeply look both subproblem
1. Sell at the given day
=> If I sell at the day[i], how can i know when should i buy to have the maximum profit
=> We're just supposed we buy at day[i - 1] and sell at day[i]
=> But that's just the part of maximum profit, not our total maximum profit
=> it's okay, we just ask the subproblem to day[i - 1]
=> Hey, the day before me (day[i - 1]), what's your total maximum profit?
=> I just add this, and that's my total maximum profit?
=> total maximum profit 
    = (day[i] - day[i - 1]) + total maximum profit[i - 1]

2. Don't do the transaction
=> Easy, ask, Hey, the day before me (day[i - 1]), what's your total maximum profit?
=> total maximum profit 
    = total maximum profit[i - 1]


Day 1, i = 0
1  3  6  2 <- table
0

Day 2, i = 1
answer = Max(table[1 - 1], price[1] - price[1 - 1] + table[1 - 1])
       = Max(0, 3 - 1 + 0)
       = Max(0, 2)
       = 2
1  3  6  2 <- table
0  2

Day 3, i = 2
answer = Max(table[2 - 1], price[2] - price[2 - 1] + table[2 - 1])
       = Max(2, 6 - 3 + 2)
       = Max(2, 5)
       = 5
1  3  6  2 <- table
0  2  5

Day 3, i = 3
answer = Max(table[3 - 1], price[3] - price[3 - 1] + table[3 - 1])
       = Max(5, 2 - 6 + 5)
       = Max(5, 1)
       = 5
1  3  6  2 <- table
0  2  5  5
*/
/**
 * @param {number[]} prices
 * @return {number}
 */
/*
Build DP table

************************************************************
n = the legnth of array
Time compelxity: O(n)

Space comelxity: O(n)
*/
function maxProfit(prices) {
  const table = new Array(prices.length).fill(0);

  for (let i = 1; i < table.length; i++) {
    const transaction = prices[i] - prices[i - 1] + table[i - 1];
    const notTransaction = table[i - 1];

    table[i] = Math.max(transaction, notTransaction);
  }

  return table[prices.length - 1];
}

/*
Only use pointer

************************************************************
n = the legnth of array
Time compelxity: O(n)

Space comelxity: O(1)
*/
function maxProfit2(prices) {
  let maxProfit = 0;

  for (let i = 1; i < prices.length; i++) {
    maxProfit = Math.max(maxProfit, prices[i] - prices[i - 1] + maxProfit);
  }

  return maxProfit;
}
console.log(maxProfit([7, 1, 5, 3, 6, 4, 11, 213, 54, 3, 12, 32, 66, 45, 12]));
console.log(maxProfit1([7, 1, 5, 3, 6, 4, 11, 213, 54, 3, 12, 32, 66, 45, 12]));
console.log(maxProfit2([7, 1, 5, 3, 6, 4, 11, 213, 54, 3, 12, 32, 66, 45, 12]));
