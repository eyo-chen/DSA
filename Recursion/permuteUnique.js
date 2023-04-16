//////////////////////////////////////////////////////
// *** Permutations II ***
//////////////////////////////////////////////////////
/*
Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.

Example 1:
Input: nums = [1,1,2]
Output:
[[1,1,2],
 [1,2,1],
 [2,1,1]]

Example 2:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 
Constraints:
1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/
/*
This problem is very similar to permutation 1
but the main difference is how do we to handle duplicate element in the nums array
Let's try to use the solution of permutation 1 in this problem, and see the outcome
For example, nums = [1,1,2], for the sake of differentiation, we use [1a, 1b, 2] to represent
                                               [1a, 1b, 2]
                        1a                          1b                    2
            1b                  2         1a                2      1a          1b
            2                   1b        2                1a      1b          1a
The result would be like this
[1a, 1b, 2], [1a, 2, 1b], [1b, 1a, 2], [1b, 2, 1a], [2, 1a, 1b], [2, 1b, 1a]
[1,  1,  2], [1,  2,  1], [1,  1,  2], [1,  2,  1], [2,  1,  1], [2,  1,  1]

Can you see the problem?
We can have one order
What order?
Order 1 -> Choose 1a first, then choose 1b  [1a, 1b, 2], [2, 1a, 1b]
Order 2 -> Choose 1b first, then choose 1a  [1b, 1a, 2], [2, 1b, 1a]
=> those two order are duplicate, so we have to decide to only choose one order
It's obvious to choose order 1, because it's easier to implement in the for-loop

So now we know the core logic
1. use the same logic as permutaion 1
2. BUT, have to first sort the array, so that we can found the duplicate easily
3. BUT, we have more constraint
   => if nums[i] === nums[i - 1] -> found the duplicate (ex: 1b)
   => then need to check the order
   => !usedArr[i - 1] -> make sure we're not choosing 1b before 1a (order!!!)
   => i - 1 means previous same element
   => usedArr[i - 1] means hey, have we used the previus element?
   => we only wanna choose the current duplicate element AFTER choosing previous duplicate element
   => choose 1a, then choose 1b
4. if (usedArr[i] || (nums[i] === nums[i - 1] && !usedArr[i - 1])) continue;
   => usedArr[i] -> if this element has been used, skip it
   => (nums[i] === nums[i - 1] && !usedArr[i - 1])
       -> if current element is duplicate, and we haven't used previous duplicate element, skip it

************************************************************
n = the legnth of input (both are same as permutation 1 because of the worst case is all element is unique)
Time complexity: O(n * n!)

Space complexity: O(n)
*/
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permuteUnique = function (nums) {
  const res = [];
  const usedArr = new Array(nums.length).fill(false);

  // sort first
  const sortedNums = nums.sort((a, b) => a - b);

  recursiveHelper(sortedNums, usedArr, [], res);

  return res;

  function recursiveHelper(nums, usedArr, tmp, res) {
    // base case
    if (tmp.length === nums.length) {
      res.push([...tmp]);
      return;
    }

    for (let i = 0; i < nums.length; i++) {
      // constraint
      if (usedArr[i] || (nums[i] === nums[i - 1] && !usedArr[i - 1])) continue;

      tmp.push(nums[i]);
      usedArr[i] = true;
      recursiveHelper(nums, usedArr, tmp, res);

      tmp.pop();
      usedArr[i] = false;
    }

    return;
  }
};
////////////////////////////////////////////////////////////////////////
/*
Solution 2
Instead of reusing the logic of permutation 1
We first convert nums array to hashTable
key -> each unique element in the nums
value -> frequence for each unique element in the nums
For example, nums = [1,1,2]
Hahstable
{
 1: 2
 2: 1
}

We just simply keeping using the element in order until the frequence is 0
                                         {1: 2, 2: 1}
                            1                         2
                        {1:1, 2:1}                {1:2,2:0}
            1                      2                  1
        {1:0,2:1}             {1:1,2:0}           {1:1,2:0}
           2                      1                   1
        {1:0,2:0}             {1:0,2:0}            {1:0,2:0}

This solution works for both duplcate and non-duplicate problem

I think this has same time and space complexity as solution 1
*/
var permuteUnique1 = function (nums) {
  const res = [];
  const hashTable = {};
  // build hashTable
  for (let i = 0; i < nums.length; i++) {
    if (hashTable[nums[i]] === undefined) hashTable[nums[i]] = 1;
    else hashTable[nums[i]]++;
  }

  // for iteration inside recursive function
  const newNums = Array.from(new Set(nums));

  recursiveHelper1(nums, newNums, hashTable, [], res);

  return res;
};

function recursiveHelper1(nums, newNums, hashTable, tmp, res) {
  if (tmp.length === nums.length) {
    res.push([...tmp]);
    return;
  }

  for (let i = 0; i < newNums.length; i++) {
    // frequency is 0, skip it
    if (hashTable[newNums[i]] === 0) continue;

    // choose
    tmp.push(newNums[i]);
    hashTable[newNums[i]]--;

    // explore
    recursiveHelper1(nums, newNums, hashTable, tmp, res);

    // unchoose
    tmp.pop();
    hashTable[newNums[i]]++;
  }

  return;
}

console.log(permuteUnique([1, 1, 2, 3, 3]));
console.log(permuteUnique1([1, 1, 2, 3, 3]));
