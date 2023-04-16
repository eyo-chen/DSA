//////////////////////////////////////////////////////
// *** Subsets II ***
//////////////////////////////////////////////////////
/*
Given an integer array nums that may contain duplicates, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Example 1:
Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

Example 2:
Input: nums = [0]
Output: [[],[0]]
 
Constraints:
1 <= nums.length <= 10
-10 <= nums[i] <= 10
*/
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
/*
First time writing this problem, fail
Second time writing this problem, succeed
I guess is because I write some similar duplicate problems before

The solution of this problem is very similar to subSet 1
But now we have more constraint, we don't wanna choose duplicate element
But the solution of solving duplicate issue is very similar to previous problem, like combinationSum 2 or permutation 2
In short, these problems all have similar pattern to follow

Like subSet 1, for each call stack frame, we can
Choose the element
Don't choose the element

If we use same solution to this problem, then we can find the duplicate problem
For example, input = [1,2,2], use [1,2a,2b] to represent, the recursive tree would be like
                                       [1,2a,2b]                                   -> 0
                     []                                  [1]                               -> 1st  choose or not choose 1
             []                [2a]                  [1]                 [1,2a]            -> 2nd  choose or not choose 2a
      []         [2b]*     [2a]       [2a,2b]     [1]   [1,2b]*       [1,2a]   [1,2a,2b]   -> 3rd  choose or not choose 2b
             
[2b] is equal to [2a]
[1,2b] is equal to [1,2a]

How can we solve this issue?
Recall to combinationSum2 and permutateUnique problem,
the choosing order DOES matter here
In short, we can NOT choose 2b before 2a
As we could see, all the duplicate problem happen after we skip 2a, and choose 2b
See the * part, those two parts are exact duplicate parts
And both of them are skiping 2a, and choosing 2b
And that's the main problem
So our algorithm have to make sure
We can ONLY choose second duplicate element if previous duplicate element has been chosen
For example,
We can ONLY choose 2b if 2a has been chosen

We can accomplish this validation by usedArr

if ( arr[index] !== arr[index - 1] ||
    (arr[index] === arr[index - 1] && usedArr[index - 1])
   ) 
1. arr[index] !== arr[index - 1]
=> it means if current element is not as same as previous one
=> because we have sorted the array first
=> if this is true, it means current element is NOT the duplicate element
=> we can choose it without further validation

2. (arr[index] === arr[index - 1] && usedArr[index - 1])
=> arr[index] === arr[index - 1], it means current element is as same as previous one
=> In other words, current element is duplicate element
=> then we wanna further validate
=> usedArr[index - 1], it means Has previous element been chosen?
=> If yes, now we can choose
=> If not, remember that We can ONLY choose second duplicate element if previous duplicate element has been chosen
=> so now we can NOT choose

************************************************************
(It's all as same as subSet1, because the worst case is all element is unique)
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
var subsetsWithDup = function (arr) {
  const res = [];
  // sort the array first
  const sortedArr = arr.sort((a, b) => a - b);
  const usedArr = new Array(arr.length).fill(false);

  recursiveHelper(sortedArr, 0, usedArr, [], res);

  return res;
};

function recursiveHelper(arr, index, usedArr, tmp, res) {
  // base case
  if (index === arr.length) {
    res.push([...tmp]);
    return;
  }

  // Validation
  if (
    arr[index] !== arr[index - 1] ||
    (arr[index] === arr[index - 1] && usedArr[index - 1])
  ) {
    // choose
    tmp.push(arr[index]);
    usedArr[index] = true;

    // explore
    recursiveHelper(arr, index + 1, usedArr, tmp, res);

    // Note here, we only unchoose if we have chosen before
    // If we never choose, we don't need to unchoose
    usedArr[index] = false;
    tmp.pop();
  }

  // explore
  recursiveHelper(arr, index + 1, usedArr, tmp, res);

  return;
}

/*
Solution 2
For example, the input is [1,2,2]
                                       [1,2a,2b]                                   -> 0
                     []                                  [1]                               -> 1st  choose or not choose 1
             []                [2a]                  [1]                 [1,2a]            -> 2nd  choose or not choose 2a
      []         [2b]*     [2a]       [2a,2b]     [1]   [1,2b]*       [1,2a]   [1,2a,2b]   -> 3rd  choose or not choose 2b
The main idea is that
Once we decide NOT to choose 2, then we have to skie all the 2 afterwards
https://www.youtube.com/watch?v=Vn2v6ajA7U0
See this in detailed explanation

************************************************************
n = the length of candiate t = target
Time complexity: O(2 ^ n)
=> same as above

Space complexity: O(n)
=> same as above
*/
var subsetsWithDup1 = function (arr) {
  const res = [];
  const sortedArr = arr.sort((a, b) => a - b);

  recursiveHelper1(sortedArr, 0, [], res);

  return res;
};

function recursiveHelper1(arr, index, tmp, res) {
  if (index === arr.length) {
    res.push([...tmp]);
    return;
  }

  const curVal = arr[index];

  // choose
  tmp.push(curVal);
  recursiveHelper1(arr, index + 1, tmp, res);

  // update the index all the way to skip later duplicate element
  for (let i = index + 1; i < arr.length; i++) {
    if (arr[i] === curVal) index++;
    else break;
  }

  tmp.pop();
  recursiveHelper1(arr, index + 1, tmp, res);

  return;
}

/* 
Solution 3
This is another completely different approach
But this very similar to the last solution of subSets 1
Plz go back to see that solution first

The main point to avoid duplicate is use duplicateCount
For example, arr = [2,3,3]

We first have default empty array in the res array
res = [[]]

First for loop
1. duplicateCount = 1, which means there's no duplicate 2
2. take out the element in the res, and encapsulate into new array
3. preRes = [[]]
4. Loop through preRes
5. tmp = []
6. Loop through duplicateCount, and going to add 2
7. tmp = [2]
8. add tmp into res
9. res = [[], [2]]

Second for loop
1. duplicateCount = 2, which means there's duplicate 3
2. take out the element in the res, and encapsulate into new array
3. preRes = [[], [2]]
4. Loop through preRes
5. tmp = []
6. Loop through duplicateCount, and going to add 3
=> Note that here we're going to add 3 twice respectively
7. tmp = [3]
8. add tmp into res
9. res = [[], [2], [3]]
10. tmp = [3,3]
11. add tmp into res
12. res = [[], [2], [3], [3,3]]
13. tmp = [2]
14. keep adding 3 twice
15. tmp = [2,3]
16. add tmp into res
17. res = [[], [2], [3], [3,3], [2,3]]
18. tmp = [2,3,3]
19. add tmp into res
20, res = [[], [2], [3], [3,3], [2,3], [2,3,3]]

This is the process

There's two things to keep in mind when implement the code
1. Have to sort the array first
=> In order to have to correct duplicateCount, we have to sort the array first

2. The reason we keep doing this [...arr]
=> is because we want to keep copying the array
=> For example, now the res = [[], [2]], and we're going to add 3
=> const preRes = [...res];
=> Now preRes = [[], [2]], but [] and [2] still have the same reference as in the res array
=> What does this means?
=> if we only do this const tmp = preRes[j]
=> Later when we push new value into tmp, it will also add the value into the array in the res array
=> Why?
=> Because the [] and [2] in the preRes is as same as in the res
=> That's why we need const tmp = [...preRes[j]];
=> it means hey, give me the first element of preRes
=> It's [], and i'm gonna copy it
=> so later i can push new value into tmp
=> same thing for res.push([...tmp]);
=> we have to copy another new array into res
=> so it won't be mutated later we keep pushing new value into the element

If this is still not clear, try to change the code
and debug in the chrome browser

I guess it has same time and space complexity
*/
var subsetsWithDup3 = function (arr) {
  const res = [[]];

  arr.sort((a, b) => a - b);

  for (let i = 0; i < arr.length; ) {
    let duplicateCount = 0;

    while (
      duplicateCount + i < arr.length &&
      arr[i + duplicateCount] === arr[i]
    )
      duplicateCount++;

    const preRes = [...res];

    for (let j = 0; j < preRes.length; j++) {
      const tmp = [...preRes[j]];

      for (let k = 0; k < duplicateCount; k++) {
        tmp.push(arr[i]);
        res.push([...tmp]);
      }
    }

    i += duplicateCount;
  }
  return res;
};

console.log(subsetsWithDup3([2, 3, 3]));
console.log(subsetsWithDup1([2, 3, 3]));
