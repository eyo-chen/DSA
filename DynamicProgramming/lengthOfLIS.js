//////////////////////////////////////////////////////
// *** Longest Increasing Subsequence ***
//////////////////////////////////////////////////////
/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].

Example 1:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:
Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1
*/
/**
 * @param {number[]} nums
 * @return {number}
 */
/*
I came up this by myself

The subprblem is
what's the Longest Increasing Subsequence in any sepcific index?

How to know the longest increasing subsequence
For each cell,
1. compare to all the cell before this index
2. Check if I can extend the length, can only extend if the value is greater

Initialize all the default value to 1
Because the longest increasing subsequence for each cell is gurantee 1

For example,
[10,9,2,5,3,7,101,18]

10   9   2   5   3   7   101   18
1    1   1   1   1   1    1     1

i = 1, search from 0 ~ 0 (k)
Can 9 extend the longest increasing subsequence of [10] -> No

i = 2, search from 0 ~ 1 (k)
Can 2 extend the longest increasing subsequence of [10] -> No
Can 2 extend the longest increasing subsequence of [9] -> No

i = 3, search from 0 ~ 2 (k)
Can 5 extend the longest increasing subsequence of [10] -> No
Can 5 extend the longest increasing subsequence of [9] -> No
Can 5 extend the longest increasing subsequence of [2] -> Yes, so the answer = Max(table[3], table[2] + 1) = 2

i = 4, search from 0 ~ 3 (k)
Can 3 extend the longest increasing subsequence of [10] -> No
Can 3 extend the longest increasing subsequence of [9] -> No
Can 3 extend the longest increasing subsequence of [2] -> Yes, so the answer = Max(table[3], table[2] + 1) = 2
Can 3 extend the longest increasing subsequence of [5] -> No

So on and so forth


************************************************************
n = the legnth of array
Time compelxity: O(n ^ 2)
=> Nested for loop

Space comelxity: O(n)
*/
function lengthOfLIS(nums) {
  const table = new Array(nums.length).fill(1);

  // the variable to keep tracking the final answer
  let res = 1;

  for (let i = 1; i < nums.length; i++) {
    const curValue = nums[i];

    for (let k = 0; k < i; k++) {
      // if the value is greater, then check which value is maximum
      // + 1 means extend the sequence
      if (curValue > nums[k]) table[i] = Math.max(table[i], table[k] + 1);
    }

    res = Math.max(res, table[i]);
  }

  return res;
}

/*
This is another approach
Bruth force solution
  
Find all the subsequence
And when finding any subsequence
1. Check if it's increasing subsequence
2. If yes, then check if this lengt of subsequence is greater than the previous answer
  
************************************************************
n = the legnth of array
Time compelxity: O(2 ^ n * n)
=> The basic create all the subsequence is 2 ^ n
=> We also do the n work for validateIncreasing
  
Space comelxity: O(n)
*/
function lengthOfLIS1(nums) {
  if (nums.length === 1) return 1;

  return recursiveHelper(nums, 0, []);

  function recursiveHelper(nums, index, tmp) {
    if (index === nums.length) {
      if (validateIncreasing(tmp)) {
        return tmp.length;
      }

      return -1;
    }

    tmp.push(nums[index]);
    const choose = recursiveHelper(nums, index + 1, tmp);

    tmp.pop();
    const notChoose = recursiveHelper(nums, index + 1, tmp);

    return Math.max(choose, notChoose);
  }

  function validateIncreasing(arr) {
    if (arr.length === 1) return true;

    for (let i = 1; i < arr.length; i++) {
      if (arr[i - 1] >= arr[i]) return false;
    }

    return true;
  }
}

// console.log(
//   lengthOfLIS([1, 2, 12, 2, 123, 3, 12, 33, 164, 1, 4, 12, 1, 3, 4, 5, 6])
// );
// console.log(
//   lengthOfLIS1([1, 2, 12, 2, 123, 3, 12, 33, 164, 1, 4, 12, 1, 3, 4, 5, 6])
// );
////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
/**
 * @param {number[]} array
 * @return {number[]}
 */
/*
This is variant, instead of returning the interger, the value of the length of Longest Increasing Subsequence, returns the string of Longest Increasing Subsequence(find the path of LIS)

Basically the same idea as above
But storing the array in each cell of table

************************************************************
n = the legnth of array
Time compelxity: O(n * 2 * n)
=> Because of table[i].unshift(...table[LISindex]);

Space comelxity: O(n * 2)
=> store array in each cell of table
*/
function lengthOfLIS2(nums) {
  const table = [];

  // to keep track where is the index of Longest Increasing Subsequence, so that we can finnaly return the answer in O(1) time
  let LISindex = 0;

  // intialize all cell with it's own value with one length, it's like base case
  // for each cell, the Longest Increasing Subsequence is just itself
  for (let i = 0; i < nums.length; i++) {
    table.push([nums[i]]);
  }

  for (let i = 0; i < nums.length; i++) {
    const curVal = nums[i];

    /*
    This is crucial
    keep track what's the temporary LIS at index i?
    For example, [10,9,2,5,3,7,101,18]
    For i = 5, nums[5] = 7;
    We first encounter nums[2] = 2, we temporarily set this as LIS, because we're not sure at this point, there may be other longer LIS at further index
    Temporarily now the table[5] = [7,2]
    Then encounter nums[3] = 5, now we want to compare the (length of table[3]) + 1 with tmpLISLength which has set above, so now we won't add 5 to table[5] because table[3].length + 1 does not greater than tmpLISLength

    If we only use table[i].length, then we will add it because table[3].length does greater than table[5].length
    Try to use table[i].length in the comparison, instead tmpLISLength
    */
    let tmpLISLength = table[i].length;

    for (let j = 0; j < i; j++) {
      if (curVal > nums[j] && table[j].length + 1 > tmpLISLength) {
        // update new tmpLISLength
        tmpLISLength = table[j].length + 1;

        /*
        Another crucial part
        Now LISindex helps us to keep track what's index is the final index of LIS we're gonna push at table[i]
        For example, [10,9,2,5,3,7,101,9999]
        For i = 7, j = 6, table[6] = [2,5,7,101]
        LISindex = j = 6
        After for-loop, we know we want to push([2,5,7,101]) at table[7](nums[7] = 9999)
        */
        LISindex = j;
      }
    }

    // if tmpLISLength has been updated, it means we at least expand once for any possible LIS before i
    if (tmpLISLength !== 1) {
      table[i].unshift(...table[LISindex]);

      /*
      Another crucial part
      If input is [10,9,2,5,3,7,101,18], we won't get into this statement
      So the LISindex is j we set above, which is 6, so we can finally return table[6]
      But 
      If input is [10,9,2,5,3,7,101,9999],
      the LISindex is still j without this if-statement, which means we won't have chance to update LISindex, so it will return wrong aswer if we indeed add the last element in the LIS
      so now we get into this statement, then update LISindex = i;
      It means
      Hey, because tmpLISLength has been updated, so i guarantee expand the LIS, but we can't update this inside the j-for-loop, because j can all the way up to i - 1, so update LISindex here
      */
      LISindex = i;
    }
  }

  return table[LISindex];
}

/*
Bruth force solution

************************************************************
n = the legnth of array
Time compelxity: O(2 ^ n * n)
=> The basic create all the subsequence is 2 ^ n
=> We also do the n work for validateIncreasing
  
Space comelxity: O(n)
*/
function lengthOfLIS3(nums) {
  return recursiveHelper(nums, 0, []);

  function recursiveHelper(nums, index, tmp) {
    if (index === nums.length) {
      if (validateIncreasing(tmp)) {
        return [...tmp];
      }

      return [];
    }

    tmp.push(nums[index]);
    const choose = recursiveHelper(nums, index + 1, tmp);

    tmp.pop();
    const notChoose = recursiveHelper(nums, index + 1, tmp);

    if (choose.length > notChoose.length) return choose;
    else return notChoose;
  }

  function validateIncreasing(arr) {
    if (arr.length === 1) return true;

    for (let i = 1; i < arr.length; i++) {
      if (arr[i - 1] >= arr[i]) return false;
    }

    return true;
  }
}

// console.log(
//   lengthOfLIS2([10, 1, 2, 12, 333, 1, 2, 3, 4, 5, 121, 6, 4232, 121])
// );
// console.log(
//   lengthOfLIS3([10, 1, 2, 12, 333, 1, 2, 3, 4, 5, 121, 6, 4232, 121])
// );

////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
