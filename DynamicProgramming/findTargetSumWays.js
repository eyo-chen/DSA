/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var findTargetSumWays = function (nums, target) {
  return recursiveHelper(nums, target, 0, 0);
};

function recursiveHelper(nums, target, index, workingTarget) {
  if (target === workingTarget && index === nums.length) return 1;
  if (index > nums.length) return 0;

  const curVal = nums[index];
  let res = 0;

  res += recursiveHelper(nums, target, index + 1, workingTarget + curVal);

  res += recursiveHelper(nums, target, index + 1, workingTarget - curVal);

  return res;
}
console.log(findTargetSumWays([1, 1, 1, 1, 1], 3));
