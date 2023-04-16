//////////////////////////////////////////////////////
// *** Combination Sum ***
//////////////////////////////////////////////////////
/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

Example 1:
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

Example 2:
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:
Input: candidates = [2], target = 1
Output: []
*/
/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
/*
This is the most-intuitive solution

Idea is just doing the DFS, search any single possibilities
In order to avoid duplicate result, we 
1) sort the array 
2) then compare it to the element in hashTable
when hitting the base case

The key point of this solution is complete searching
so that for each function calls, the for-loop start at 0

candidates = [2,3,6,7], target = 7
                                  7
                        [2]   [3]   [6]    [7]
           [2,2] [2,3] [2,6] [2,7]
                  .....

************************************************************
c = the length of candidatest array, t = target

Time compelxity: O(c ^ t)
because the branching factor is c, and the deepeset height of recursive tree is t
maybe also need to mutiply O(n * log(n)) since we sort the array when hitting the base case

Space comelxity: O(t)
the deepest height of recursive tree
*/
function combinationSum(candidates, target) {
  const res = [];

  // help to stop the for-loop earily
  // because when the element is too large, then there's no reason to keep looping through later element
  const sorterdCandidates = candidates.sort((a, b) => a - b);

  function combinationSumRecursiveHelper(
    candidates,
    target,
    res,
    tmpRes,
    hash
  ) {
    // base case
    if (target === 0) {
      // first sort the array, so that we can compare if it's duplicate
      const sortedTmpRes = [...tmpRes].sort((a, b) => a - b);

      // avoid storing duplicate array
      if (!hash[sortedTmpRes]) {
        res.push(sortedTmpRes);
        hash[sortedTmpRes] = true;
      }

      return;
    }

    // i start at 0 for every recusive callStack
    for (let i = 0; i < candidates.length; i++) {
      const curVal = candidates[i];

      const remainingVal = target - curVal;

      if (remainingVal >= 0) {
        tmpRes.push(curVal);

        combinationSumRecursiveHelper(
          candidates,
          remainingVal,
          res,
          tmpRes,
          hash
        );

        tmpRes.pop();
      }
      // because array is sorted, we can stop earily, there's no reason to keep exploring further value once the target is negative
      else break;
    }
  }

  combinationSumRecursiveHelper(sorterdCandidates, target, res, [], {});

  return res;
}

/*
  As we could see, the main difficulties are how can we avoid duplicate result??
  
  The previous solution is just sort the array, and use hashTable
  
  In this solution, we change our decision for each recursive callStack
  so does change the recursive tree, also time complexity
  
  For each recursive callStack, we only have two decisions
  1. choose the element
  2. do not choose the element ever again
  (use index pointer to help us skip the element later)
  
                                                        7 
                            [2]                                                  []
                  [2,2]                 [2]
        [2,2,2]           [2,2]     [2,3]   [2]
  [2,2,2,2]  [2,2,2]  [2,2,3]  [2,2]
                          [2,2,6]  [2,2]
                                [2,2,7] [2,2]
  
  The key of this recursive tree is that we only have two decision for every callStack
  1. Choose the element, do NOT update the index, which means keep choosing the element in the next callStack
  2. Don't choose the element, keep the tmp array same, and update the index, so that we won't choose the element ever again along with the process
  
  The letfest part of tree will be all the first element, like [2,2,2,2,2]
  Because we keep choosing the same element over and over again
  
  The righest part of tree will be all the empty array, like []
  Because we keep DON'T choosing the element
  
  ************************************************************
  c = the length of candidatest array, t = target
  
  Time compelxity: O(2 ^ t)
  because the branching factor is always 2, and the deepeset height of recursive tree is t
  
  Space comelxity: O(t)
  the deepest height of recursive tree
  */
function combinationSum1(candidates, target) {
  const res = [];

  recursiveHelper(candidates, target, res, [], 0);

  return res;

  function recursiveHelper(candidates, target, res, tmp, index) {
    // base case1: get the answer
    if (target === 0) {
      res.push([...tmp]);
      return;
    }

    // base case2: index out of the bound or the target is negative
    if (index >= candidates.length || target < 0) return;

    const curVal = candidates[index];

    // choose the current element, do NOT update the index, which means keep choosing the same element in next callStack
    tmp.push(curVal);
    recursiveHelper(candidates, target - curVal, res, tmp, index);

    // do NOT choose the current element, update the index, which means we won't even choose this element ever again along with the process
    tmp.pop();
    recursiveHelper(candidates, target, res, tmp, index + 1);
  }
}

/*
  This yet another solution to avoid duplicate result
  It's quite similar to the solution 1
  
  However, instead of staring at 0 of for-loop for every recursive calls
  use index pointer to not to start at 0 every time
  
  The idea is once we choose an element, we won't choose any element before this particular element
  For example, if we're given candidates [2, 3, 6, 7]
  For each recursive calls, once we choose the 3 which is index 1
  Then the next and later recusive process, we start the for-loop from index 1, which means we won't choose 2 ever again
  
  So, if we're at index 0, we can have any decomposition of [2,3,6,7]
  but, if we're at index 1, we can only have decomposition of [3,6,7]
  so we won't have the result is [2,2,3] and [3,2,2]
  becuase once we choose 3, there's no way to choose 2 again
  
  The recusive tree would look like this
                                                             7
                          [2]                [3]                        [6]              [7]
              [2,2]             [2,3]       [2,6]  [2,7]
        [2,2,2]  [2,2,3]   [2,3,3] [2,3,6]  [2,6,7]
        [2,2,6]  [2,2,7]   [2,3,7]
  
  
  ************************************************************
  c = the length of candidatest array, t = target
  
  Time compelxity: O(c ^ t)
  because the branching factor is c, and the deepeset height of recursive tree is t
  maybe also need to mutiply O(n * log(n)) since we sort the array when hitting the base case
  
  Space comelxity: O(t)
  the deepest height of recursive tree
  
  Note that the time complexity seems is as same as first soultion
  It is. But it acutally won't always have c branching factor
  So it basically costs less time
  
  */
function combinationSum2(candidates, target) {
  const res = [];

  const sortedCandidates = candidates.sort((a, b) => a - b);

  return recursiveHelper(sortedCandidates, target, res, [], 0);

  function recursiveHelper(candidates, target, res, tmp, index) {
    if (target === 0) {
      res.push([...tmp]);
      return res;
    }

    for (let i = index; i < candidates.length; i++) {
      const curVal = candidates[i];
      const remainingTarget = target - curVal;

      if (remainingTarget >= 0) {
        // with
        tmp.push(curVal);
        recursiveHelper(candidates, remainingTarget, res, tmp, i); // have to pass i

        // without
        tmp.pop();
      }
      // because the array is sorted, once the element is too large, we don't need to do any further exploration
      else break;
    }

    return res;
  }
}

// [35,]

// console.log(combinationSum([2, 3, 6, 7], 7));
// console.log(combinationSum1([2, 3, 6, 7], 7));
// console.log(combinationSum2([7, 3, 6, 2], 7));
// console.log(combinationSum3([7, 3, 6, 2], 7));

/*
  function combinationSum3(candidates, target) {
    const res = [];
  
    const sortedCandidates = candidates.sort((a, b) => a - b);
  
    return recursiveHelper(sortedCandidates, target, res, [], 0);
  
    function recursiveHelper(candidates, target, res, tmp, index) {
      if (target === 0) {
        res.push([...tmp]);
        return res;
      }
  
      for (let i = index; i < candidates.length; i++) {
        const curVal = candidates[i];
        const remainingTarget = target - curVal;
  
        if (remainingTarget >= 0) {
          tmp.push(curVal);
          recursiveHelper(candidates, remainingTarget, res, tmp, index); // pass i 
  
          tmp.pop();
          recursiveHelper(candidates, target, res, tmp, index + 1); // don't need this line
        }
  
        else break;
      }
  
      return res;
    }
  }
  
  This solution is completely incorrect
  Look that we have two recusive calls inside for-loop
  which means we'll have eight branching factor for each callStacks
  */

/*
  This is bottom-up approach (tabulation)
  
  For example, candidates = [2,3,5], target = 5
  
  1. Initialize an 2D array
      0   1   2   3   4   5
  []  
  [2] 
  [3] 
  [5] 
  
  The left column represent what can i use to get to the target
  [2] means how can i use 2 to accomplish the target, so on and so forth
  
  The top row represent the target number
  If colum is [3], and top row is 3,
  it means how can i use 3 to accomplish the target 3
  Note that from top to bottom is accumulated
  which means when it's [3], it means we can use [2,3]
  when it's [5], it means our choice is [2,3,5]
  
  Note that for each cell, it represents the total of use the candidate and not use the candidate
  For example, 
      0   1   2   3   4   5
  []  
  [2]         
  [3]                     *   
  [5]
  If we're at this cell, we can have two choices
  Use the candidate 3
  => which means the follow-up question will be what's the previous computed result of 2
  
  Not use the candidate 3
  => which means we can only use 2, go back previous row to see the computed result
  
      0   1   2   3   4   5
  []  
  [2]                   not use
  [3]        use          *   
  [5] 
  
  
  2. Initialize the first row to null, 
     which means there's no way to construct any target value from empty array
  
        0        1        2          3        4        5
  []   null     null    null       null      null     null
  [2]                 
  [3]        
  [5] 
  
  3. Initialize the first column to []
     which means we can always choose do NOT use any candidate to construct the target value 0
  
        0        1        2          3        4        5
  []   null     null    null       null      null     null
  [2]   []              
  [3]   []     
  [5]   []
  
  4. Loop through each candidate
        Loop through each target value
     => get the case when using and not using the candidate
     => if it's not null, we can keep adding the new candidate
  
  For candidate = 2, i = 1
      For target = 1,
      there's no way to construct 1 from using 2
  
      For target = 2, 
      there's one way to construct 2 from using 2
      2 - 2 = 0, so goes to see 0, which is [], so just add 2 to empty array
  
      For target = 3, 
      there's no there's no way to construct 1 from only using 2
      Use => 3 - 2 = 1, table[1][1], so goes to see 1, which is null, means it can't construct 1
      Not use => table[0][2] = null
      Both are null, so just put null
  
      For target = 4,
      Use => 4 - 2 = 2, table[1][2] = [2], so just add 2 to this compurted array
      Not use => table[0][4] = null
  
      For target = 5,
      Use => 5 - 2 = 3, table[1][3] = null
      Not use => table[1][5] = null
  
        0        1        2          3        4        5
  []   null     null    null       null      null     null
  [2]   []      null     [2]       null      [2,2]    null
  [3]   []     
  [5]   []
  
  For candidate = 3, i = 2
      For target = 1,
      there's no way to construct 1 from using 3
  
      For target = 2, 
      Use => 2 - 3 = -1, no way
      Not use => table[1][2] = [2]
      So it means, we can still construct 2 from using 2 when we choose not to use 3
      Note that the the left column is accumulated
  
      For target = 3, 
      Use => 3 - 3 = 0, table[2][0] = [], push candidate 3 to empty array
      Not use => table[1][3] = null
  
      For target = 4,
      Use => 4 - 3 = 1, table[2][1] = null, so no way
      Not use => table[1][4] = [2,2]
      Same logic as above
  
      For target = 5,
      Use => 5 - 3 = 2, table[1][3] = [2], so push candidate 3 to this computed array
      Not use => table[1][5] = null
  
        0        1        2          3        4        5
  []   null     null    null       null      null     null
  [2]   []      null     [2]       null      [2,2]    null
  [3]   []      null     [2]        [3]      [2,2]    [2,3]
  [5]   []
  
  
  For candidate = 5, i = 3
        0        1        2          3        4        5
  []   null     null    null       null      null     null
  [2]   []      null     [2]       null      [2,2]    null
  [3]   []      null     [2]        [3]      [2,2]    [2,3]
  [5]   []      null     [2]        [3]      [2,2]    [[2,3],[5]]
  
  ************************************************************
  c = the length of candidatest array, t = target
  
  Time compelxity: O(c * t * t)
  Two nested for-loops -> c * t
  Also, we have another two for-loops inside nested for-loops
  I guess the worst case of these two for-loops just t
  
  
  Space comelxity: O(c * t * t)
  Basically same logic as above
  */
function combinationSum3(candidates, target) {
  const table = [];

  for (let i = 0; i < candidates.length + 1; i++) {
    table.push(new Array(target + 1).fill());
  }

  // Initialize the first row to null
  for (let i = 0; i < target + 1; i++) {
    table[0][i] = null;
  }

  // Initialize the first column to []
  for (let i = 1; i < candidates.length + 1; i++) {
    table[i][0] = [[]];
  }

  for (let i = 1; i < candidates.length + 1; i++) {
    const curCandidate = candidates[i - 1];

    for (let k = 1; k < target + 1; k++) {
      // use
      const use = table[i][k - curCandidate];

      // not use
      const notUse = table[i - 1][k];

      table[i][k] = [];

      if (!!use) {
        for (let j = 0; j < use.length; j++) {
          table[i][k].push([...use[j], curCandidate]);
        }
      }

      if (!!notUse) {
        for (let j = 0; j < notUse.length; j++) {
          table[i][k].push([...notUse[j]]);
        }
      }

      /*
        if it's still empty array after 2 for-loops
        it means there's no way to construct either using candidate or not using candidate
        */
      if (table[i][k].length === 0) table[i][k] = null;
    }
  }

  // have to return [] if it's not possible
  return table[candidates.length][target] === null
    ? []
    : table[candidates.length][target];
}
/*
  https://www.youtube.com/watch?v=AUIfTelAGVc
  This is also another tabulation
  but only use 1D array, so the code is much cleaner
  
  For candidate is 2, 
  target 2 is equal to current candidate, so just table[2].push(...table[2 - 2], 2)
  target 4 is greater than 2, so go checking table[4 - 2] = [[2]], table[4].push(...table[4 - 2], 2), [[2,2]]
  
  For candidate is 3,
  target 3 is equal to current candidate, so just table[3].push(...table[3 - 3], 3)
  target 5 is greater than 3, so go checking table[5 - 3] = [[2]], table[5].push(...table[5 - 3], 3), [[2,3]]
  
  so on and so forth
  
  
        0        1        2          3        4        5
        []      []       [[2]]    [[3]]    [[2,2]]  [[2,3],[5]]
  
  The table will look like this
  ************************************************************
  c = the length of candidatest array, t = target
  
  Time compelxity: O(c * t * t)
  Two nested for-loops -> c * t
  Also, we have another 2 for-loops inside nested for-loops
  I guess the worst case of these two for-loops just t
  
  
  Space comelxity: O(c * t * t)
  Basically same logic as above
  */
function combinationSum4(candidates, target) {
  const table = new Array(target + 1).fill(null);

  // initialize all to empty array
  /*
    Note that can not do sth like this
    const table = new Array(target + 1).fill([]);
    because that will cause each element of table hold the same empty array, all of them hold the same reference
  
    do sth like this instead
    so each empty array is different
    */
  for (let i = 0; i < target + 1; i++) {
    table[i] = [];
  }

  for (const curCandidate of candidates) {
    for (let k = curCandidate; k < target + 1; k++) {
      if (k === curCandidate) table[k].push([curCandidate]);
      // k is greater than current candidate
      else {
        // find the previous computed table result
        const preTable = table[k - curCandidate];

        // loop through the computed result, and add new candidate
        for (let j = 0; j < preTable.length; j++) {
          table[k].push([...preTable[j], curCandidate]);
        }
      }
    }
  }

  return table[target];
}

// console.log(combinationSum4([1, 2, 3], 10));
// console.log(combinationSum4([2, 3, 7, 12], 15));
// console.log(combinationSum([2, 3, 7, 12], 15));
// console.log(combinationSum4([2, 3, 5], 5));
