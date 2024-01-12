var combinationSum3 = function (k, n) {
  const res = [];

  recursiveHelper(k, n, 1, [], res);

  return res;
};

function recursiveHelper(remainingCounts, target, index, tmp, res) {
  // base case
  if (remainingCounts === 0 && target === 0) {
    res.push([...tmp]);
    return;
  }

  // base case: out of the bound
  if (remainingCounts < 0 || target < 0) return;

  // if i is greater that the target, there's no need to keep for-loop
  for (let i = index; i <= 9 && i <= target; i++) {
    // choose
    tmp.push(i);

    // explore
    recursiveHelper(remainingCounts - 1, target - i, i + 1, tmp, res);

    // unchoose
    tmp.pop();
  }

  return;
}
