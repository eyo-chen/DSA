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
