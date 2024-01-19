#include <vector>

using namespace std;

class Solution {
 public:
  int totalNQueens(int n) {
    vector<int> colPosition;

    return helper(colPosition, 0, n);
  }

  int helper(vector<int>& colPosition, int row, int n) {
    if (row == n) {
      return 1;
    }

    int count = 0;

    for (int c = 0; c < n; c++) {
      colPosition.push_back(c);
      if (isValid(row, c, colPosition)) {
        count += helper(colPosition, row + 1, n);
      }
      colPosition.pop_back();
    }

    return count;
  }

  bool isValid(int row, int col, vector<int>& colPosition) {
    for (int r = 0; r < colPosition.size() - 1; r++) {
      int curColPosition = colPosition[r];
      int absRowDistance = abs(row - r);
      int absColDistance = abs(curColPosition - col);

      if (curColPosition == col) return false;
      if (absRowDistance == absColDistance) return false;
    }

    return true;
  }
};