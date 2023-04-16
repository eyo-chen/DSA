//////////////////////////////////////////////////////
// *** Permutations ***
//////////////////////////////////////////////////////
/*
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:
Input: nums = [1]
Output: [[1]]

Constraints:
1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.
*/
/*
I came up the solution by myself(second time writing the problem)

In order to find all the permutation, have to really think about "decision space"
For example, nums = [1,2,3]
We know that we're gonna have three slots to fill up 
__ __ __
=> At this point, our decision space is [1,2,3], which means
=> We have three choices at this state of call stack
=> Choose 1, 2 or 3

Use 1 to fill up the first slot
1 __ __
=> Once choosing the 1, now the decision space is [2,3], which means
=> now we can only have two choices, we're not allowed to choose 1 anymore

Use 2 to fill up the second slot
1 2 __
=> Now the decision space is just [3]

1 2 3
=> Base case

Use 2 to fill up the first slot
2 __ __
=> Once choosing the 2, now the decision space is [1,3]

so on and so forth

The main difficulty is how to keep the decision space
Which means for any single state of call stack,
we have to know what's our decision space,
what element we can choose right now, and what element we can't choose

In this solution, i use exploredArr to keep track
For example, exploredArr = [false, false, false], it means decision space is [1,2,3]
__ __ __
=> At this point, exploredArr = [false, false, false], which means
=> decision space is [1,2,3]

Use 1 to fill up the first slot
1 __ __
=> Once choosing the 1, now the exploredArr = [true, false, false], which means
=> decision space is [2,3]
=> now we can only have two choices, we're not allowed to choose 1 anymore

Use 2 to fill up the second slot
1 2 __
=> the exploredArr = [true, true, false]
=> Now the decision space is just [3]

1 2 3
=> Base case
=> the exploredArr = [true, true, true]

Use 2 to fill up the first slot
2 __ __
=> the exploredArr = [false, true, false]
=> Once choosing the 2, now the decision space is [1,3]

so on and so forth
Hope now it's clear

Choice
=> Choose element in the decision space

Contraint
=> Have to keep decision space 

Goal
=> the length of working permutation is equal to the length of input

Think the process of recursion in prermutation like a tree
                                         [a, b, c]
                [a, _, _]                [b, _, _]                 [c, _, _]
        [a, b, _]      [a, c, _]   [b, a, _]     [b, c, _]    [c, a, _]   [c, b, _]
        [a, b, c]      [a, c, b]   [b, a, c]     [b, c, a]    [c, a, b]   [c, b, a]

Note the decision space(DS)(only left hand side)
                                         [a, b, c]
                [a, _, _]         DS = [T,F,F]       
        [a, b, _]   DS = [T,T,F]    [a, c, _]    DS = [T,F,T]
        [a, b, c]   DS = [T,T,T]    [a, c, b]    DS = [T,T,T]
=> As we could see, when [a, b, c] backtrack to [a, b, _], decision space also need to set back to false
=> Same thing for when [a, b, _] backtrack to [a, _, _]

************************************************************
n = the legnth of input
Time complexity: O(n * n!)
=> Look at the recursive tree
=> For first call stack, we have three choices
=> For second call stack, we have two choices
=> For third call stack, we have one choice
=> So that the branching factor is n, n - 1, n - 2, .... 1
=> that's the core idea of O(n!)
=> res.push([...tmp]); takes O(n) 

Space complexity: O(n)
=> the deepest height of recursive tree is n
*/
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permute = function (nums) {
  if (nums.length === 1) return [nums];

  const res = [];
  const exploredArr = new Array(nums.length).fill(false); // decision space

  recursiveHelper(nums, exploredArr, [], res);

  return res;

  function recursiveHelper(nums, exploredArr, tmp, res) {
    // the length of working permutation is eqaul to the length of input
    if (tmp.length === nums.length) {
      res.push([...tmp]);
      return;
    }

    for (let i = 0; i < nums.length; i++) {
      // if it's false, it means this element is in the decision space, we can choose
      if (exploredArr[i] === false) {
        tmp.push(nums[i]);
        exploredArr[i] = true;

        recursiveHelper(nums, exploredArr, tmp, res);

        tmp.pop();
        exploredArr[i] = false; // set it back to false when backtracking
      }
    }

    return;
  }
};

/*
https://www.geeksforgeeks.org/write-a-c-program-to-print-all-permutations-of-a-given-string/
This is the second solution using swap
Can't full understand right now, maybe comeback later
*/

/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
// Firs time writing this problem
function permute1(strArr) {
  const res = [];

  helperFunction(strArr, [], res);

  return res;

  function helperFunction(str, tmp, res) {
    // base case
    if (tmp.length === str.length) {
      res.push([...tmp]);
      return;
    }

    // recursive case (loop through from start)
    for (let i = 0; i < str.length; i++) {
      const choose = str[i];

      // we don't add the one we've already chosen
      if (tmp.includes(choose)) continue;

      tmp.push(choose);

      helperFunction(str, tmp, res);

      // remove (back to previos level)
      tmp.pop();
    }
  }
}

// SECOND ONE
/*
  The one is quciker in leetcode
  */
function permute2(arr) {
  const res = [];

  // if input array is less than two, we could just return the simple output
  if (arr.length === 2) {
    res.push([...arr]);
    arr.reverse();
    res.push([...arr]);
  } else if (arr.length === 1) return [arr];
  // greater than or equal to three, do the recursive work
  else helperFunction(arr, [], true, res);

  return res;

  function helperFunction(arr, tmp, first = true, res) {
    // base case
    /*
      If the different length between original array and tmp array, we could do the base case work
      For example, original array [1,2,3,4,5]
      If tmp array now is [1,2,3],
      We immediately know that the remaining work is add [4,5] and [5,4]
      Se we could
      1. add normal order array
      2. reverse array
      3. and add new one
      */
    /*
      helpToFilter() is the helper function to help us to extract the array from original array
      For example helpToFilter([1,2,3], [2]), will return us [1,3]
      The main purpose of that is we want the remaining array of last two element, so we could add and reverse
      If now is -> original array [1,2,3,4,5], tmp array [1,2,3]
      We want to extract [4,5], and do the work
      */
    if (arr.length - tmp.length === 2) {
      const remainingArr = helpToFilter(arr, tmp);
      res.push([...tmp, ...remainingArr]);
      remainingArr.reverse();
      res.push([...tmp, ...remainingArr]);
      return;
    }

    // recursive case
    // We split two part of recursive case
    /*
      First part,
      The first level of tree(see above)
      Because in this level, we do NOT need to care any potential duplicate case
      All we need to do is select the current order's element, and add into array
      After recursive work, we remove it to next part
      For example,
      i = 0, select [a], and go do remaining work, [a,b], [a,c]......
      i = 1, select [b], and go do remaining work, [b,a], [b,c]......
      */
    if (first) {
      for (let i = 0; i < arr.length; i++) {
        tmp.push(arr[i]);
        helperFunction(arr, tmp, false, res);
        tmp.pop();
      }
    } else {
      /*
      Second part,
      Here, we have to care do NOT select duplicate element(same concept above in solution1)
      */
      for (let i = 0; i < arr.length; i++) {
        // use nested array to make sure the element we're gonna add does NOT include in tmp array
        let noDuplicate = true;
        for (let j = 0; j < tmp.length; j++) {
          if (arr[i] === tmp[j]) noDuplicate = false;
        }
        // after that, we could add the element
        if (noDuplicate) {
          tmp.push(arr[i]);
          helperFunction(arr, tmp, false, res);
          tmp.pop();
        }
      }
    }
  }
}
