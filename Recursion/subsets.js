//////////////////////////////////////////////////////
// *** Subsets ***
//////////////////////////////////////////////////////
/*
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
Input: nums = [0]
Output: [[],[0]]
 
Constraints:
1 <= nums.length <= 10
-10 <= nums[i] <= 10
All the numbers of nums are unique.
*/
/*
I came up the solution by myself (second time doing this problem)

The solution is quite simple
Again, try to figure out the choice, constraint and goal
For each recursive stack or the pointer we're at, we can either choose the character or unchoose the character
For example, if input is [1,2,3]
one of the ouput is [], it means we unchoose three times
one of the output is [1], it means choose the one, and unchoose two times

Now it's clear
Choice => Choose or unchoose the character
Constraint => no constraint in this question, we just wanna generate all possibilities
Goal => our index pointer is out of the bound, aka is equal to the length of input

If the input is [1,2,3], the recursive tree would be like
                                       [1,2,3]                                   -> 0
                     []                                  [1]                     -> 1st
             []                [2]               [1]              [1,2]          -> 2nd
      []         [3]     [2]       [2,3]     [1]   [1,3]       [1,2]   [1,2,3]   -> 3rd

1st => choose "1" or unchoose
2nd => choose "2" or unchoose
3rd => choose "3" or unchoose

************************************************************
n = the legnth of nums
Time complexity: O(2 ^ n) or O((2 ^ n) * n)
=> For each call stack, we always only explore 2 choices, choose or unchoose
=> so the branching factor is 2
=> the deepest height is n
=> res.push([...tmp]); -> this is O(n) works
=> the time complexity would be O((2 ^ n) * n) if we consider this work

Space complexity: O(n)
=> the deepest height of recursive tree is n
*/
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var subsets = function (nums) {
  const res = [];

  recursiveHelper(nums, 0, [], res);

  return res;

  function recursiveHelper(nums, index, tmp, res) {
    // base case
    if (index === nums.length) {
      // have to copy the whole new array to get the new reference
      res.push([...tmp]);
      return;
    }

    // choose
    tmp.push(nums[index]);
    recursiveHelper(nums, index + 1, tmp, res);

    // unchoose
    tmp.pop();
    recursiveHelper(nums, index + 1, tmp, res);

    return;
  }
};

/*
This is yet another completely different solution
The idea is kinda like DP

For example, if nums = [1,2,3]
We first have default empty array in the res array
res = [[]]

First for loop,
1. take out the element in the res, and encapsulate into new array
2. preRes = [[]]
3. Add 1 in every single element in that new array
4. preRes = [[1]]
5. Add it to the res
6. res = [[], [1]]
=> [] is old element, [1] is new added element

Second for loop,
1. take out the element in the res, and encapsulate into new array
2. preRes = [[], [1]]
3. Add 2 in every single element in that new array
4. preRes = [[2], [1,2]]
5. Add it to the res
6. res = [[], [1], [2], [1,2]]
=> [], [1] is old element, [2], [1,2] is new added element

Third for loop,
1. take out the element in the res, and encapsulate into new array
2. preRes = [[], [1], [2], [1,2]]
3. Add 3 in every single element in that new array
4. preRes = [[3], [1,3], [2,3], [1,2,3]]
5. Add it to the res
6. res = [[], [1], [2], [1,2], [3], [1,3], [2,3], [1,2,3]]
=> [], [1], [1,2] is old element, [3], [1,3], [2,3], [1,2,3] is new added element

Hope this is clear

************************************************************
Time complexity: O((2 ^ n) * n
=> The outer loop is O(n)
=> The inner loop will scale as process keep goin
=> O(2 ^ n) scale

Space complexity: O(2 ^ (n - 1))
=> preRes will have (2 ^ (n - 1)) at most
*/
var subsets1 = function (nums) {
  const res = [[]];

  for (let i = 0; i < nums.length; i++) {
    const preRes = [...res];

    for (let j = 0; j < preRes.length; j++) {
      preRes[j] = [...preRes[j], nums[i]];
    }
    console.log(preRes);
    res.push(...preRes);
  }

  return res;
};
console.log(subsets1([1, 2, 3]));
console.log(subsets([1, 2, 3]));
