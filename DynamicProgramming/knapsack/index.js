// *** The 0-1 Knapsack Problem ***
//////////////////////////////////////////////////////
/*
We are given an array of values values and an array of weights weights:
values[i] corresponds to the value of the i'th item
weights[i] corresponds to the weight of the i'th item

Given these two lists and an integer maxWeight, find a subset of the items in this "knapsack" that has maximal overall value, yet stays <= maxWeight in total weight.

Example 1:
Input:
values = [60, 50, 70, 30]
weights = [5, 3, 4, 2]
maxWeight = 8

Output: 120
Explanation: We take items 1 and 2 (zero-indexed) for a total value of 120 and a total weight of 7.

Example 2:
Input:
values = [60, 100, 120, 80, 30]
weights = [10, 20, 30, 40, 50]
maxWeight = 400

Output: 390
Explanation: We take all items for a total value of 390 and a total weight of 150, still below 400.
*/
/**
 * @param {Array<number>} values
 * @param {Array<number>} weights
 * @param {number} maxWeightConstraint
 * @return {number}
 */
/*
This is classic dynamic programming problem

For most of the dynamic programming problem
We have to think about 1) subproblems 2) relation between subproblems

Subproblem
=> For each item, can either choose or not choose

Relation
=> Find the maximum between choose or not choose 

For example,
values = [60, 50, 70, 30]
weights = [5, 3, 4, 2]
maxWeight = 8

         0   1   2   3   4   5   6   7   8
[]         
[60, 5]
[50, 3]
[70, 4]
[30, 2]

The zero amount and empty knapsack is our base case
Another reason for empty array is because the logic would need to -1 for each row
So the empty array(base case) helps us to not have array[-1], and cause potential error

         0   1   2   3   4   5   6   7   8
[]         
[60, 5]
[50, 3]                  .
[70, 4]
[30, 2]
subproblem -> What's the maximum value when we're only allowed to use two item [60,5] and [50,3] when the target weight is 4?

         0   1   2   3   4   5   6   7   8
[]         
[60, 5]
[50, 3]                  
[70, 4]         .
[30, 2]
subproblem -> What's the maximum value when we're only allowed to use two item [60,5], [70,4] and [50,3] when the target weight is 2?


Initialize the cell 0
         0   1   2   3   4   5   6   7   8
[]       0   0   0   0   0   0   0   0   0
[60, 5]  0   0   0   0   0   0   0   0   0
[50, 3]  0   0   0   0   0   0   0   0   0
[70, 4]  0   0   0   0   0   0   0   0   0
[30, 2]  0   0   0   0   0   0   0   0   0

We can just start at table[1][1] because the answer of 0 weight and empty array are just 0

Item 1
The weight of first item is 5, so we can only choose this item until the target weight is 5
         0   1   2   3   4   5   6   7   8
[]       0   0   0   0   0   0   0   0   0
[60, 5]  0   0   0   0   0  60  60  60  60
[50, 3]  0   0   0   0   0   0   0   0   0
[70, 4]  0   0   0   0   0   0   0   0   0
[30, 2]  0   0   0   0   0   0   0   0   0
If the target weight is 8 and I'm only allowed to use item1, the maximum value is just  60

Item 2
Target weight 1 -> can't use the item, so go up to see the subanswer, which is also 0
Now the weight is 3, so we can put 50 at target weight 3 and 4
Target weight 5 -> Have to choices
                      1. Choose -> After using the item, the remaining weight is 2, so go to see table[2 - 1][5 - 3], the value is 0, and 0 + 50 = 50
                      2. Not choose -> Go up to see the subanswer = 60
                      maximum = Max(50, 60)
The main logic
1. Choose
=> Have to go back to the (target weight - weight), and go up one level

2. Not chosoe
=> Only go up to see the subproblem

See the example above
         0   1   2   3   4   5   6   7   8
[]       0   0   0   0   0   0   0   0   0
[60, 5]  0   0   0   0   0  60  60  60  60
[50, 3]  0   0   0  50  50   .   0   0   0
[70, 4]  0   0   0   0   0   0   0   0   0
[30, 2]  0   0   0   0   0   0   0   0   0
Subproblem: What's the maximum value when we're only allowed to use two item [60,5] and [50,3] when the target weight is 5?

At this point, we know we now have two choices
1. Choose
=> The subproblem becomes What's the maximum value when we're only allowed to use two item [60,5] when the target weight is 2?
=> After getting the answer, we just add the value of item
=> Go right means subtract(using) the weight 
=> Go up means after using the item

2. Not choose
=> The subproblem becomes What's the maximum value when we're only allowed to use two item [60,5] when the target weight is 5?
=> Only go up, don't need to subtract the weight because we're not using the item

so on and so forth

************************************************************
n = the legnth of item(values and weights), w = maxWeightConstraint
Time compelxity: O(n * w)
=> Nested for loop

Space comelxity: O(n * w)
=> DP table
*/
function knapsack(values, weights, maxWeightConstraint) {
  const table = [];

  for (let i = 0; i < values.length + 1; i++) {
    const row = new Array(maxWeightConstraint + 1).fill(0);

    table.push(row);
  }

  for (let item = 1; item < values.length + 1; item++) {
    /*
      Be careful here,
      We add the empty array as our base case, and it's in the first row
      Need to subtract one to get the correct item
      */
    const itemWeight = weights[item - 1];
    const itemValue = values[item - 1];

    for (let maxWeight = 1; maxWeight < maxWeightConstraint + 1; maxWeight++) {
      // If item weight is greater than targer weight, we can't choose the item
      const choose =
        itemWeight > maxWeight
          ? table[item - 1][maxWeight]
          : table[item - 1][maxWeight - itemWeight] + itemValue;

      const notChoose = table[item - 1][maxWeight];

      table[item][maxWeight] = Math.max(choose, notChoose);
    }
  }

  return table[values.length][maxWeightConstraint];
}

/*
  Same logic
  
  ************************************************************
  n = the legnth of item(values and weights), w = maxWeightConstraint
  Time compelxity: O(n * w)
  => Nested for loop
  
  Space comelxity: O(w)
  */
function knapsack1(values, weights, maxWeightConstraint) {
  let firstRow = new Array(maxWeightConstraint + 1).fill(0);
  let secondRow = new Array(maxWeightConstraint + 1).fill(0);

  for (let item = 1; item < values.length + 1; item++) {
    const itemWeight = weights[item - 1];
    const itemValue = values[item - 1];

    for (let maxWeight = 1; maxWeight < maxWeightConstraint + 1; maxWeight++) {
      const choose =
        itemWeight > maxWeight
          ? firstRow[maxWeight]
          : firstRow[maxWeight - itemWeight] + itemValue;

      const notChoose = firstRow[maxWeight];

      secondRow[maxWeight] = Math.max(choose, notChoose);
    }

    [secondRow, firstRow] = [firstRow, secondRow];
  }

  return firstRow[maxWeightConstraint];
}

// recursive solution
function knapsack2(values, weights, maxWeightConstraint) {
  return recursiveHelper(values, weights, maxWeightConstraint, 0);

  function recursiveHelper(values, weights, maxWeightConstraint, index) {
    if (maxWeightConstraint <= 0 || index >= values.length) return 0;

    if (weights[index] > maxWeightConstraint) {
      return recursiveHelper(values, weights, maxWeightConstraint, index + 1);
    }

    const choose =
      recursiveHelper(
        values,
        weights,
        maxWeightConstraint - weights[index],
        index + 1
      ) + values[index];

    const notChoose = recursiveHelper(
      values,
      weights,
      maxWeightConstraint,
      index + 1
    );

    return Math.max(choose, notChoose);
  }
}

/*
  This is variant of previous problem
  items = [
    [1, 2],
    [4, 3],
    [5, 6],
    [6, 7]
  ]
  
  capacity = 10,
  
  return [maximum value, [index array]]
  
  
  We have two differnt solution
  
  1. Use the same idea of previous one, after building the table, backtrack to the path to find the index array
  
  ************************************************************
  n = the legnth of item(values and weights), w = maxWeightConstraint
  Time compelxity: O(n * w)
  => Nested for loop
  => While-loop won't above over  O(n * w)
  
  Space comelxity: O(w)
  */
function knapsack3(items, capacity) {
  const table = [];

  for (let i = 0; i < items.length + 1; i++) {
    const row = new Array(capacity + 1).fill(0);

    table.push(row);
  }

  for (let item = 1; item < items.length + 1; item++) {
    const [itemValue, itemWeight] = items[item - 1];

    for (let maxWeight = 1; maxWeight < capacity + 1; maxWeight++) {
      const choose =
        itemWeight > maxWeight
          ? table[item - 1][maxWeight]
          : table[item - 1][maxWeight - itemWeight] + itemValue;

      const notChoose = table[item - 1][maxWeight];

      table[item][maxWeight] = Math.max(choose, notChoose);
    }
  }

  const res = table[items.length][capacity];
  const indexArr = [];

  let i = items.length,
    k = capacity;

  while (i >= 1 && k >= 1) {
    const curVal = table[i][k];

    if (curVal !== table[i - 1][k]) {
      /*
          Be careful here,
          We add the empty array as our base case, and it's in the first row
          Need to subtract one to get the correct item
        */
      indexArr.push(i - 1);

      k -= items[i - 1][1];
      i--;
    } else {
      i--;
    }
  }

  return [res, indexArr.sort((a, b) => a - b)];
}
/*
  Solution 2
  => Change what stores in each cell of table
  
  ************************************************************
  n = the legnth of item(values and weights), w = maxWeightConstraint
  Time compelxity: O(n * w * n)
  => Nested for loop
  => The additional n is for the spread out the array [...item]
  => The worst case is spread out all the items 
  
  Space comelxity: O(n * w * n)
  => For each cell, we will additionally store another items array at worst
  */
function knapsack4(items, capacity) {
  const table = [];

  for (let i = 0; i < items.length + 1; i++) {
    const row = [];

    for (let k = 0; k < capacity + 1; k++) {
      row.push([0, []]);
    }

    table.push(row);
  }

  for (let item = 1; item < items.length + 1; item++) {
    const [itemValue, itemWeight] = items[item - 1];

    for (let maxWeight = 1; maxWeight < capacity + 1; maxWeight++) {
      const choose =
        itemWeight > maxWeight
          ? table[item - 1][maxWeight][0]
          : table[item - 1][maxWeight - itemWeight][0] + itemValue;

      const notChoose = table[item - 1][maxWeight][0];

      if (choose > notChoose) {
        table[item][maxWeight][0] = choose;

        table[item][maxWeight][1].push(
          ...table[item - 1][maxWeight - itemWeight][1],
          item - 1 // be careful here
        );
      } else {
        table[item][maxWeight][0] = notChoose;

        table[item][maxWeight][1].push(...table[item - 1][maxWeight][1]);
      }
    }
  }

  return table[items.length][capacity];
}

// console.log(knapsack([60, 100, 120, 80, 30], [10, 20, 30, 40, 50], 400));
// console.log(knapsack2([60, 100, 120, 80, 30], [10, 20, 30, 40, 50], 400));

// console.log(
//   knapsack2(
//     [
//       [1, 2],
//       [4, 3],
//       [5, 6],
//       [6, 7],
//     ],
//     10
//   )
// );
// console.log(
//   knapsack3(
//     [
//       [1, 2],
//       [4, 3],
//       [5, 6],
//       [6, 7],
//     ],
//     10
//   )
// );
