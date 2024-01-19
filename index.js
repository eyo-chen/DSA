/**
 * @param {number} n
 * @return {string[][]}
 */
var solveNQueens = function(n) {
  var ans = [];
  var temp = Array.from({ length: n }, () => '.'.repeat(n));

  helper(ans, temp, 0, n);

  return ans;
};

/**
* @param {string[][]} ans
* @param {string[]} temp
* @param {number} row
* @param {number} n
*/
var helper = function(ans, temp, row, n) {
  if (row === n) {
      ans.push([...temp]);
      return;
  }

  for (var c = 0; c < n; c++) {
      var isValid = true;

      for (var r = row - 1; r >= 0; r--) {
          var prevRow = temp[r];
          var diff = row - r;

          if (
              (c - diff >= 0 && prevRow[c - diff] === 'Q') ||
              (prevRow[c] === 'Q') ||
              (c + diff < n && prevRow[c + diff] === 'Q')
          ) {
              isValid = false;
              break;
          }
      }

      if (!isValid) continue;

      temp[row] = temp[row].substring(0, c) + 'Q' + temp[row].substring(c + 1);
      helper(ans, temp, row + 1, n);
      temp[row] = temp[row].substring(0, c) + '.' + temp[row].substring(c + 1);
  }
};

// Example usage:
var result = solveNQueens(4);
console.log(result);
