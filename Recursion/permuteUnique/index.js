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
