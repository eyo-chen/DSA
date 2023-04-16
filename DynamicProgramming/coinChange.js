//////////////////////////////////////////////////////
// *** Coin Change ***
//////////////////////////////////////////////////////
/*
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

Example 1:
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:
Input: coins = [2], amount = 3
Output: -1

Example 3:
Input: coins = [1], amount = 0
Output: 0
*/
/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */
/*
This is top-down approach with memoization
For example, coins = [2,3,5], amount = 8

For top root node 8, "what's the total amount for 8?"
I don't know, but we can ask subproblem after using each one coin
Sub Problem1: "what's the total amount for 6?" (after using 2)
Sub Problem2: "what's the total amount for 5?" (after using 3)
Sub Problem1: "what's the total amount for 1?" (after using 1)
We're waiting those three sub problems finishing, and we just choose the shortest answer

The main idea of this approach is
1. Try to find all possibilities
2. Remember the sub answer we've solved
3. Find the global optimize sub answer

Note this pattern of recursion
Instead of finding all possible combination(array), we only want to find the minimum amount(number)
So we won't pass any reference array along with the recusive process
We put this line of code in every call stack or recusive state
let shorteseCombination = -1;

Because for each recusive state, we all want to answer the same problem
"What's the minimum amount I can get back in this particular recusive state?"
From the above example, 
For top root node 8, "what's the total amount for 8?"
We don't know yet, so just put let shorteseCombination = -1; in this state
And we will have three different sub answers

Same thing for Sub Problem1: "what's the total amount for 6?" (after using 2)
We still want to find the minimum amount(number) in this state
so again put this let shorteseCombination = -1; in this state
And we will have three different sub answers

Again, for each recursive state, we first put "let shorteseCombination = -1;"
And once all the sub problem get back, we will or won't change the value of this varible
After all the process in this state, we return this value of this variable
Give it back(up) to the previos call stack

So on and so forth

************************************************************
c =  the length of coins, a = amount
Time: O(c * a)
After memoization, we'll only have a sub problem
For example, if amount = 5, we won't have the sub problem asking what's the amount of 6
For each sub problem, we'll try c amount

Space: O(a)
The highest length of recusive tree
*/
function coinChange(coins, amount, memo = {}) {
  // We already have an answer cached. Return it.
  if (memo[amount]) return memo[amount];

  /*
    The minimum coins needed to make change for 0 is always 0 coins no matter what coins we have.
  */
  if (amount === 0) return 0;

  /*
    Minimum coins to make change for a negative amount is -1.
    This is just a base case we arbitrarily define.
  */
  if (amount < 0) return -1;

  let shorteseCombination = -1;

  for (const coin of coins) {
    const subProblem = amount - coin;

    const subAnswer = coinChange(coins, subProblem, memo);

    if (subAnswer !== -1) {
      // use one more coin
      const tmpCombination = subAnswer + 1;

      /*
        shorteseCombination === -1 means we haven't found any possible answer
        So if we got whatever answer, just temporarily use that as tmp combination
  
        tmpCombination < shorteseCombination
        if we've found the answer before, now we want to make sure the tmp combination i use is the shortest
        */
      if (shorteseCombination === -1 || tmpCombination < shorteseCombination)
        shorteseCombination = tmpCombination;
    }
  }

  memo[amount] = shorteseCombination;
  return shorteseCombination;
}

/*
  Bottom-Up - Tabulation
  This is reference from DP course from freeCodeCamp
  
  The more classic idea of DP is below, this has different mindset to solve the problem
  
  Main Idea
  1. Create DP Table
     => Note that we have to create (amount + 1) array because we want the array from 0 ~ amount
     For example, amount = 3, we want the array be like
     [0, 1, 2, 3] -> this is an array with four length
  
  2. Initilaize each element with null which means no value or invalid value
     => In second tabulation approach, we initialize value. It's basically same overall idea, but has small difference
  
  3. For each cell of DP table, we want to loop through each coin
     => Before doing this, have to make sure the cell is not equal to null. Because it means there's no value can generate this value in this element
     For example, coins = [3], amount = 7
     the second cell of DP table would be null because there's no way using coin 2
  
  4. For each coin, we want to 
     (1) add one to the value of current element of DP table, which means we use one more coin
     (2) Check (using coin + current amount) is less then amount
         => Make sure it's still within the valid range
         => This is actually unnecessary since this is only make the DP table longer(not a big deal)
     (3) If the cell of table[coin + i] is null, then it means we haven't investigated this element, so just add the value
     (4) If it's not null, it means we have investigated this element, so we have to compare it's value and the current value we hold(tmpRes)
  
  ************************************************************
  c = length of coins, a = length of amount
  
  Time compelxity: O(c * a)
  => Because we have nested for-loop. For each element of DP table, we basically loop through each coin at worst case
  => Beside them, all we do is constant work
  
  Space comelxity: O(a)
  => The most expensive space is the DP table array which is used (amount + 1) space
  */
function coinChange1(coins, amount) {
  const table = new Array(amount + 1).fill(null);

  // base case
  // We can always choose not use the coin if the amount is 0
  table[0] = 0;

  for (let i = 0; i < amount; i++) {
    // current total counts of the amount
    const subProblem = table[i];

    if (subProblem !== null) {
      for (const coin of coins) {
        // use one more coin
        const tmpRes = subProblem + 1;

        if (
          (coin + i <= amount && table[coin + i] === null) ||
          // we only want to update the value if tmpRes is less than
          tmpRes < table[coin + i]
        )
          table[coin + i] = tmpRes;
      }
    }
  }

  return table[amount] === null ? -1 : table[amount];
}

/*
  This is another way to implement the Bottom-Up approach
  
  The main difference is
  1. Initialize (amount + 1) to each cell of DP table
     => The reason using (amount + 1) this is the minimum of invalid guarantee value
     For example, amount = 5, the minimum of invalid value would be 6
     Because there's no way using 6 coin to change 5 amount
     => Later will do the comparison, so this is the clean way to set the invalid value
  
  2. We loop through each element of DP table, and also loop through each coin
  
  3. Before comparing old value(table[i]) and new value(table[i - coin] + 1) -> note that +1 means using one more coin
     Have to make sure (i - coin >= 0) which means current amount is greater than coin
     In other words, i >= coin
     The reason is if current amount is 3, and coin is 5
     Again, it's no possible to use coin 5 to change amount 3. Also, there's no array[3 - 5]
  
  4. At the end of return statement,
     If the very ending value is default value, return -1
     => table[amount] === amount + 1
     If it's not, just return table[amount]
  
  
  It's very important to undestand what's the meaning of the table
  If coins = [2,3,5], amount = 8
  
  0   1   2   3   4   5   6   7   8
  9   9   9   9   9   9   9   9   9
  
  For each cell of DP table, it represents the sub problem
  If we're at 3,
              #
  0   1   2   3   4   5   6   7   8
  9   9   9   9   9   9   9   9   9
  It asks "If the amount is 3, and I have coins[2,3,5], what's the minmum amount coin I can change?"
  If we're at 6,
                          #
  0   1   2   3   4   5   6   7   8
  9   9   9   9   9   9   9   9   9
  It asks "If the amount is 6, and I have coins[2,3,5], what's the minmum amount coin I can change?"
  
  So we just start from 1, and all the way to the final problem
  
  Base case is 0, because it's no possible to use any coin if the amount if 0, no matter what coins we're given
  0   1   2   3   4   5   6   7   8
  0   9   9   9   9   9   9   9   9
  
  For amount 1, 
  1 - 2 < 0, 1 - 3 < 0, 1 - 5 < 0
  We can't use any coins to make change to the amount of 1, all coins are greater than amount of 1
  0   1   2   3   4   5   6   7   8
  9   9   9   9   9   9   9   9   9
  
  For amount 2,
  2 - 2 = 0, go back to sub problem of 0
  min(0 + 1, 9) = 1 (+1 means using one coin)
  0   1   2   3   4   5   6   7   8
  9   9   1   9   9   9   9   9   9
  
  For amount 3, 
  3 - 2 = 1,  go back to sub problem of 1
  min(9 + 1, 9) = 9
  3 - 3 = 0,  go back to sub problem of 0
  min(0 + 1, 9) = 1
  0   1   2   3   4   5   6   7   8
  9   9   1   1   9   9   9   9   9
  
  For amount 4, 
  4 - 2 = 2,  go back to sub problem of 2
  min(1 + 1, 9) = 2
  4 - 3 = 1,  go back to sub problem of 1
  min(9 + 1, 2) = 2
  0   1   2   3   4   5   6   7   8
  9   9   1   1   2   9   9   9   9
  
  For amount 5, 
  5 - 2 = 3,  go back to sub problem of 3
  min(1 + 1, 9) = 2
  5 - 3 = 1,  go back to sub problem of 2
  min(1 + 1, 2) = 2
  5 - 5 = 0,  go back to sub problem of 0
  min(0 + 1, 2) = 1
  0   1   2   3   4   5   6   7   8
  9   9   1   1   1   9   9   9   9
  
  so on and so forth
  ************************************************************
  c = length of coins, a = length of amount
  
  Time compelxity: O(c * a)
  => Because we have nested for-loop. For each element of DP table, we basically loop through each coin at worst case
  => Beside them, all we do is constant work
  
  Space comelxity: O(a)
  => The most expensive space is the DP table array which is used (amount + 1) space
  */
function coinChange2(coins, amount) {
  /*
  the reason setting amount + 1 is because it's the minimum invalid value
  there's no way using amount + 1 totoal amount to change amount
  for example, coins = [1,2,3], amount = 5,
  there's no way using 6 numbers of coin to change amount 5
  the maximum is using 5 in total (all use 1 coin)

  the is helpful when we use Math.min(table[i], others...)
  */
  const table = new Array(amount + 1).fill(amount + 1);

  /*
    The answer to making change with minimum coins for 0
    will always be 0 coins no matter what the coins we are given are
  */
  table[0] = 0;

  // each i represent sub problem or sub amount
  for (let i = 1; i <= amount; i++) {
    for (const coin of coins) {
      // if i(amount) is 1, coin is 3, then there's no way using this coin
      if (i >= coin) table[i] = Math.min(table[i], table[i - coin] + 1);
    }
  }

  return table[amount] === amount + 1 ? -1 : table[amount];
}

// console.log(coinChange2([2, 3, 4, 7], 7));
// console.log(coinChange1([2], 3));
