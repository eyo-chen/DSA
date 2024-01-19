/*
This is basically the exact same idea as N-Queens 1

The main difference is we don't need to generate the board
we only need to return the number

************************************************************
n = the input n
Time complexity: O(!n * n)
=> checkValidPlacement is O(n) work

Space complexity: O(n)
=> the deepest height of recursive tree is n
*/
var totalNQueens = function (n) {
  return recursiveHelper(n, 0, []);
};

function recursiveHelper(n, row, pos) {
  if (n === row) {
    return 1;
  }

  let res = 0;

  for (let i = 0; i < n; i++) {
    if (checkValidPlacement(row, i, pos)) {
      pos.push(i);
      res += recursiveHelper(n, row + 1, pos, res);
      pos.pop();
    }
  }

  return res;
}

function checkValidPlacement(row, col, pos) {
  for (let i = 0; i < pos.length; i++) {
    if (
      row === i ||
      col === pos[i] ||
      row + col === i + pos[i] ||
      row - col === i - pos[i]
    )
      return false;
  }

  return true;
}
