#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<string>> solveNQueens(int n) {
    vector<vector<string>> ans;
    vector<string> temp(n, string(n, '.'));

    helper(ans, temp, 0, n);

    return ans;
  }

  void helper(vector<vector<string>>& ans, vector<string>& temp, int row,
              int n) {
    if (row == n) {
      ans.push_back(temp);
      return;
    }

    for (int c = 0; c < n; c++) {
      if (!isValid(temp, row, c, n)) continue;

      temp[row][c] = 'Q';
      helper(ans, temp, row + 1, n);
      temp[row][c] = '.';
    }
  }

  bool isValid(vector<string>& temp, int row, int col, int n) {
    for (int r = row - 1; r >= 0; r--) {
      string prevRow = temp[r];
      int rowDiff = row - r;

      // check if the queen is in the same column
      if (prevRow[col] == 'Q') return false;

      // check if the queen is in the same negative diagonal
      if (col - rowDiff >= 0 && prevRow[col - rowDiff] == 'Q') return false;

      // check if the queen is in the same positive diagonal
      if (col + rowDiff < n && prevRow[col + rowDiff] == 'Q') return false;
    }

    return true;
  }
};
