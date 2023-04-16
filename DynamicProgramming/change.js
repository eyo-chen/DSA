//////////////////////////////////////////////////////
// *** Coin Change 2 ***
//////////////////////////////////////////////////////
/*
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the number of combinations that make up that amount. If that amount of money cannot be made up by any combination of the coins, return 0.

You may assume that you have an infinite number of each kind of coin.

The answer is guaranteed to fit into a signed 32-bit integer.

Example 1:
Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

Example 2:
Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.

Example 3:
Input: amount = 10, coins = [10]
Output: 1
 
Constraints:
1 <= coins.length <= 300
1 <= coins[i] <= 5000
All the values of coins are unique.
0 <= amount <= 5000
*/
/**
 * @param {number} amount
 * @param {number[]} coins
 * @return {number}
 */
/*
Recursive Approach

amount = 5, coins = [1,2,5]

Base case
1. If amount === 0, we know we find one way to change the amount, return 1
2. If amount is negative, there's no way to change the amount, so return 0
3. The subanswer has been cached

Recursive case
For amount 5, 
we can have three choiecs, [1, 2, 5]
So we throw three subproblems
How many totoal amounts can change the amount of 4?
How many totoal amounts can change the amount of 2?
How many totoal amounts can change the amount of 0? -> Hit base case, return 1

We just keep throwing the sub problems until hitting the base case

Two things to note
1. How to avoid duplicate combination?
   [1,2,2] is eqaul to [2,1,2]
   => We use pointer to avoid loop-through repeated work
   => For example,
      First recursive calls,
          How many totoal amounts can change the amount of 4?
          We keep loop-through coins [1,2,5]
      However, 
      for the recursive calls
      How many totoal amounts can change the amount of 2?
      We only loop through coins [2,5]

                                   5
              [1,2,5]             [2,5]           [5]
                4                  2               0
    [1,2,5]  [2,5]     [5] 
       3      2        -2
2   1   -2

    => I hope this pattern is clear
    => The main logic is whenever I choose the i index element, the later recursive calls, we won't go back to any (i-1), (i-2)... index
    => As we could see from the recursive tree, the leftest tree is always gonna choose the first element, but the rightese always can choose the last element
    => Try to write this on paper, it should be more clear

2. How to cache the subanswer?
  => From the recursive above, there's three 2 amount, but the leftest and bottom 2 and the above 2 is not equal
  => One is "How many totoal amounts can change the amount of 2 by using coins[1,2,5]?"
  => The other one is "How many totoal amounts can change the amount of 2 by using coins[2,5]?"
  => Though the amount is the same, but we can use differnt coins, so the answer getting back is different
  => So we have to both use amount and index(how much coin can i later use?) to cahce the answer

************************************************************
c = the length of coins, a = amount
Time: O(c * a)
After memoization, we'll only have a sub problem
For example, if amount = 5, we won't have the sub problem asking what's the amount of 6
For each sub problem, we'll try c length of coins

Space: O(a)
The highest length of recusive tree
*/
function change(amount, coins) {
  return recuersiveHelper(amount, coins, 0, {});

  function recuersiveHelper(amount, coins, index, memo) {
    // three base cases
    if (amount === 0) return 1;
    if (amount < 0) return 0;
    const key = `${amount}-${index}`;
    if (memo[key] !== undefined) return memo[key];

    let res = 0;
    for (let i = index; i < coins.length; i++) {
      const curCoin = coins[i];
      res += recuersiveHelper(amount - curCoin, coins, i, memo);
    }

    memo[key] = res;
    return res;
  }
}

/*
  DP Table (2D array)
  
  1. Initialize all the value with 0
  2. Give amount 0 the value 1 (base case)
     => For any combination of coins, the amount to change 0 is always 1, which is not using any coins
  3. For any cell, the answer = use + notUse
     => If use the coin, go back to see table[coinIndex][amount - coin]
     => If not use the coin, go up to see table[coinIndex - 1][amount]
     => For example, amount = 5, coins = [1,2,5]
        The problem "how many total amounts can change the amount of 5 by using [1,2,5]"
        The can break to two subproblems
        1. Use coin 5
          -> how many total amounts can change the amount of 0 by using [1,2,5]
        2. Do not use coin 5
          -> how many total amounts can change the amount of 5 by using [1,2]
  
      0   1   2   3   4   5
  []  1   0   0   0   0   0
  [1] 1
  [2] 1           *
  [5] 1               
  the point* means how many total amounts can change the amount of 3 by using [1,2]
  
      0   1   2   3   4   5
  []  1   0   0   0   0   0
  [1] 1   
  [2] 1           
  [5] 1               *
  the point* means how many total amounts can change the amount of 4 by using [1,2,5]
  
  ************************************************************
  c =  the length of coins, a = amount
  Time: O(c * a)
  => We'll traverse every cell in the 2D array
  
  Space: O(c * a)
  */
function change1(amount, coins) {
  const table = [];

  for (let i = 0; i < coins.length + 1; i++) {
    const row = new Array(amount + 1).fill(0);
    row[0] = 1;

    table.push(row);
  }

  for (let i = 1; i < coins.length + 1; i++) {
    for (let k = 1; k < amount + 1; k++) {
      // Note that the first row of coins is empty coin
      // when need to get the coin, have to index - 1
      const remainingAmount = k - coins[i - 1];

      // if remainingAmount is negative, it means we can't use this coin
      const use = remainingAmount < 0 ? 0 : table[i][remainingAmount];

      const notUse = table[i - 1][k];

      table[i][k] = use + notUse;
    }
  }

  return table[coins.length][amount];
}

/*
  This is basically the same logic as previos solution
  But it's more optimize
  Only use two arrays along with the process
  
  ************************************************************
  c =  the length of coins, a = amount
  Time: O(c * a)
  Space: O(a)
*/
function change2(amount, coins) {
  let first = new Array(amount + 1).fill(0);
  let second = new Array(amount + 1).fill(0);
  first[0] = 1;
  second[0] = 1;

  for (let i = 1; i < coins.length + 1; i++) {
    for (let k = 1; k < amount + 1; k++) {
      const remainingAmount = k - coins[i - 1];

      const use = remainingAmount < 0 ? 0 : second[remainingAmount];

      const notUse = first[k];

      second[k] = use + notUse;
    }

    // swap two rows
    [second, first] = [first, second];
  }

  return first[amount];
}
/*
  This is even more optimize, only use one array along with the process
    
  if(amount >= coin), can use the coin, go back to see previois subanswer
  
  amount = 5, coins = [1,2,5]
  
  0   1   2   3   4   5
  1   0   0   0   0   0
  
  First iteration, use coin [1]
  0   1   2   3   4   5
  1   1   1   1   1   1
  
  Second iteration, use coin [2]
  0   1   2   3   4   5
  1   1   2   2   3   3
  amount 2 >= coin 2, table[2] += table[2-2]
  amount 5 >= coin 2, table[5] += table[5-2]
  
  the logic is kinda similar to the previos solution
  Use coin => same, go back to see previous sub answer
  Not use coin => do nothing, because the array is accumulated
  
  so on and so forth

  ************************************************************
  c =  the length of coins, a = amount
  Time: O(c * a)
  
  Space: O(a)
*/
function change3(amount, coins) {
  const table = new Array(amount + 1).fill(0);
  table[0] = 1;

  for (let i = 0; i < coins.length; i++) {
    const curCoin = coins[i];

    for (let k = 1; k < amount + 1; k++) {
      if (k >= curCoin) table[k] += table[k - curCoin];
    }
  }

  return table[amount];
}

// console.log(change(5, [1, 2, 5]));
// console.log(change1(5, [1, 2, 5]));
// console.log(change2(5, [1, 2, 5]));
// console.log(change3(5, [1, 2, 5]));
