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
