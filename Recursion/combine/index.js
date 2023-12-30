var combine = function (n, k) {
  const res = [];
  recursiveHelper(n, k, 1, [], res);
  return res;
};

function recursiveHelper(n, k, index, tmp, res) {
  // Goal
  if (tmp.length === k) {
    res.push([...tmp]);
    return;
  }

  for (let i = index; i <= n; i++) {
    // Choose
    tmp.push(i);

    // Explore
    recursiveHelper(n, k, i + 1, tmp, res);

    // Unchoose
    tmp.pop();
  }

  return;
}

console.log(combine(3, 2));
