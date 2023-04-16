//////////////////////////////////////////////////////
// *** Best Time to Buy and Sell Stock IV ***
//////////////////////////////////////////////////////
/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day, and an integer k.

Find the maximum profit you can achieve. You may complete at most k transactions.

Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

Example 1:
Input: k = 2, prices = [2,4,1]
Output: 2
Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.

Example 2:
Input: k = 2, prices = [3,2,6,5,0,3]
Output: 7
Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4. Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.

Constraints:
0 <= k <= 100
0 <= prices.length <= 1000
0 <= prices[i] <= 1000
*/
/**
 * @param {number} k
 * @param {number[]} prices
 * @return {number}
 */
/*
This is hard question, so it's hard to come up the solution 
Before doing this problem, should go back to see the Best Time to Buy and Sell Stock which is the foundation of solving this problem

Base on our foundation, we know that for each day, we have two choices
1. Sell at the given day
2. Don't do the transaction

The second choice is easy, go back to see previous day
In the previous problem, we say
total maximum profit 
  = (day[i] - day[i - 1]) + total maximum profit[i - 1]
But now we can't do as many transactions as we want, we have limited transactions

For example, prices = [3,2,6,5,0,3], k = 2, i is row, j is col
Build the 2D array, col is the days of prices (array), row is the k transaction
  3   2   6   5   0   3
0
1             *
2
This point means what's the maximum profit from day 0 to day 3 in 1 transaction

  3   2   6   5   0   3
0
1             
2                 *
This point means what's the maximum profit from day 0 to day 4 in 2 transaction

1. Sell at the given j day
=> What's does this means?
=> Try to think how can we find the maximum profit after making the transaction
   [3,2,6,5,0,3], for example, we're at
     3   2   6   5   0   3
0
1             
2            *
    1. If we sell at day 3, we can guarantee get 6 dollars, no matter what day we bought
    2. But if we have to sell at day 3, that means we have to buy the stock BEFORE day 3, which means we should buy the stock at day 1 or day 2
    3. day 2
       => 6 - 2 = 4
       => That's not all because now we only make one transaction, and we're allowed to have two tranasaction
       => After making this transaction, we have to ask what's the maximum profit i can get at day 2? -> subproblem
       => The maximum profit of buying at day 2 and selling at day 3
              = 6 - 2 + maximum profit at day 2
      
       day 1, same logic 
       => 6 - 3 = 3
       => That's not all because now we only make one transaction, and we're allowed to have two tranasaction
       => After making this transaction, we have to ask what's the maximum profit i can get at day 1? -> subproblem
       => The maximum profit of buying at day 1 and selling at day 3
              = 6 - 3 + maximum profit at day 1
=> This logic tells us
=> We have to try from 0 ~ j - 1 day to find the maximum profit at j day
=> total maximum profit 
    = (price[j] - price[m]) + (total maximum profit minus one transaction)[m]
=> Where m is from 0 ~ j - 1
=> Represent as table
=> total maximum profit 
    = (price[j] - price[m]) + table[i-1][m]
=> j is given day
=> m is from 0 ~ j
=> i - 1 means we use one transaction

2. Don't do the transaction
=> total maximum profit
    = table[i][j-1]
=> i means same transaction, not use the transaction
=> j - 1 means go back to previous day

************************************************************
n = the length of prices, k = k
Time: O(n * k * n)
=> At least work is n * k, to travere each cell of table
=> n is for each cell, we also try from 0 ~ given day(j)

Space: O(n * k)
=> table
*/
function maxProfit(k, prices) {
  if (k === 0 || prices.length <= 1) return 0;

  const table = [];

  // use k = 0 as base case, so need to k + 1
  for (let i = 0; i < k + 1; i++) {
    /*
    First row and first col is initialize with 0
    First row -> 0 transaction 
    First col -> buy and sell at same day
    */
    const row = new Array(prices.length).fill(0);
    table.push(row);
  }

  // start at table[1][1]
  // first row and first col has been initialized
  for (let i = 1; i < k + 1; i++) {
    for (let j = 1; j < prices.length; j++) {
      // two choices
      const transaction = findMaxAfterTransactions(prices, table, i, j);
      const NotTransaction = table[i][j - 1];

      table[i][j] = Math.max(transaction, NotTransaction);
    }
  }

  return table[k][prices.length - 1];
}

/*
=> total maximum profit 
    = (price[j] - price[m]) + table[i-1][m]
=> j is given day
=> m is from 0 ~ j
=> i - 1 means we use one transaction

In this case, 
maxAfterTransaction = total maximum profit
curPrice = price[j]
prices[index] = price[m]
table[i - 1][index]= table[i-1][m]
*/
function findMaxAfterTransactions(prices, table, i, j) {
  let maxAfterTransaction = -Infinity;
  const curPrice = prices[j];

  for (let index = 0; index < j; index++) {
    maxAfterTransaction = Math.max(
      maxAfterTransaction,
      curPrice - prices[index] + table[i - 1][index]
    );
  }

  return maxAfterTransaction;
}

/*
Only build two arrays O(n) space
************************************************************
n = the length of prices, k = k
Time: O(n * k * n)
=> At least work is n * k, to travere each cell of table
=> n is for each cell, we also try from 0 ~ given day(j)

Space: O(n)
*/
function maxProfit1(k, prices) {
  if (k === 0 || prices.length <= 1) return 0;

  let firstRow = new Array(prices.length).fill(0);
  let secondRow = new Array(prices.length).fill(0);

  for (let i = 1; i < k + 1; i++) {
    for (let j = 1; j < prices.length; j++) {
      const transaction = findMaxAfterTransactions1(prices, firstRow, j);
      const NotTransaction = secondRow[j - 1];

      secondRow[j] = Math.max(transaction, NotTransaction);
    }

    [secondRow, firstRow] = [firstRow, secondRow];
  }

  return firstRow[prices.length - 1];
}

function findMaxAfterTransactions1(prices, firstRow, i) {
  let maxAfterTransaction = -Infinity;
  const curPrice = prices[i];

  for (let index = 0; index < i; index++) {
    maxAfterTransaction = Math.max(
      maxAfterTransaction,
      curPrice - prices[index] + firstRow[index]
    );
  }

  return maxAfterTransaction;
}
/*
If we at this spot, i = 2, j = 4
  3   2   6   5   0   3
0
1             
2                 *

total maximum profit 
  = (price[j] - price[m]) + table[i-1][m]

m = 0, = price[4] - price[0] + table[1][0]
m = 1, = price[4] - price[1] + table[1][1]
m = 2, = price[4] - price[2] + table[1][2]
m = 3, = price[4] - price[3] + table[1][3]
=> we want to find the maximum in all of them

price[4] is duplicate, so it means 
m = 0, = - price[0] + table[1][0]
m = 1, = - price[1] + table[1][1]
m = 2, = - price[2] + table[1][2]
m = 3, = - price[3] + table[1][3]
=> find the maximum in all of them

In other words, 
m = 0, = table[1][0] - price[0]
m = 1, = table[1][1] - price[1]
m = 2, = table[1][2] - price[2]
m = 3, = table[1][3] - price[3]
=> find the maximum in all of them

Imagine we're at i = 2, j = 5
  3   2   6   5   0   3
0
1             
2                    *
We have to do these works
m = 0, = table[1][0] - price[0]
m = 1, = table[1][1] - price[1]
m = 2, = table[1][2] - price[2]
m = 3, = table[1][3] - price[3]
m = 4, = table[1][4] - price[4]
=> find the maximum in all of them

Compare this to the works above when i = 2, j = 4
See
Once we know the maximum in m = 3, we just need to compare that maximum with new m = 4, = table[1][4] - price[4] when m = 4,
So, we just need a variable to keep track 
what's the maximum difference between table[i][m] and price[m]
where m is from 0 ~ j

************************************************************
n = the length of prices, k = k
Time: O(n * k)
=> Because of new variable, we don't need to do n works in each cell of table

Space: O(n)
*/
function maxProfit2(k, prices) {
  if (k === 0 || prices.length <= 1) return 0;

  let firstRow = new Array(prices.length).fill(0);
  let secondRow = new Array(prices.length).fill(0);

  for (let i = 1; i < k + 1; i++) {
    /*
    Note the initial value of maxDiff in every new row is 
    0 - prices[0]
    because m = 0, = table[1][0] - price[0]
    */
    let maxDiff = 0 - prices[0];

    for (let j = 1; j < prices.length; j++) {
      // i = 2, j = 5, table[1][4] - price[4]
      maxDiff = Math.max(firstRow[j - 1] - prices[j - 1], maxDiff);

      const transaction = maxDiff + prices[j];
      const NotTransaction = secondRow[j - 1];

      secondRow[j] = Math.max(transaction, NotTransaction);
    }

    [secondRow, firstRow] = [firstRow, secondRow];
  }

  return firstRow[prices.length - 1];
}

var maxProfit44 = function (prices) {
  const table = [];

  for (let i = 0; i < 3; i++) {
    const row = new Array(prices.length).fill(0);
    table.push(row);
  }

  for (let i = 1; i < 3; i++) {
    let maxDiff = 0 - prices[0];

    for (let j = 1; j < prices.length; j++) {
      maxDiff = Math.max(maxDiff, table[i - 1][j - 1] - prices[j - 1]);

      const transaction = maxDiff + prices[j];
      const NotTransaction = table[i][j - 1];

      table[i][j] = Math.max(transaction, NotTransaction);
    }
  }

  return table[2][prices.length - 1];
};
// console.log(maxProfit1(2, [5, 11, 3, 50, 60, 90]));
// console.log(maxProfit44([5, 11, 3, 50, 60, 90]));
