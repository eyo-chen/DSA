#include <string>
#include <unordered_map>
#include <vector>

using namespace std;

// Brute Force Approach
class Solution {
 public:
  int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
    return helper(obstacleGrid, 0, 0);
  }

  int helper(vector<vector<int>>& obstacleGrid, int r, int c) {
    if (r < 0 || c < 0 || r >= obstacleGrid.size() ||
        c >= obstacleGrid[0].size()) {
      return 0;
    }

    if (obstacleGrid[r][c] == 1) {
      return 0;
    }

    if (r == obstacleGrid.size() - 1 && c == obstacleGrid[0].size() - 1) {
      return 1;
    }

    return helper(obstacleGrid, r + 1, c) + helper(obstacleGrid, r, c + 1);
  }
};

// Top-Down Approach with Memoization
class Solution {
 public:
  int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
    unordered_map<string, int> memo;
    return helper(obstacleGrid, 0, 0, memo);
  }

  int helper(vector<vector<int>>& obstacleGrid, int r, int c,
             unordered_map<string, int>& memo) {
    if (r < 0 || c < 0 || r >= obstacleGrid.size() ||
        c >= obstacleGrid[0].size()) {
      return 0;
    }

    if (obstacleGrid[r][c] == 1) {
      return 0;
    }

    if (r == obstacleGrid.size() - 1 && c == obstacleGrid[0].size() - 1) {
      return 1;
    }

    string key = to_string(r) + "," + to_string(c);
    if (memo.find(key) != memo.end()) {
      return memo[key];
    }

    int right = helper(obstacleGrid, r + 1, c, memo);
    int down = helper(obstacleGrid, r, c + 1, memo);

    memo[key] = right + down;

    return memo[key];
    helper(obstacleGrid, r, c + 1, memo);
  }
};

// Dynamic Programming Approach (Bottom-Up) With 2D Array
class Solution {
 public:
  int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
    if (obstacleGrid[0][0] == 1) return 0;

    int row = obstacleGrid.size();
    int col = obstacleGrid[0].size();

    vector<vector<int>> table(row, vector<int>(col, 0));
    table[0][0] = 1;

    for (int r = 1; r < row; r++) {
      // if the current cell is an obstacle, break the loop
      // leave the rest of the cells in that row as 0
      if (obstacleGrid[r][0] == 1) {
        break;
      }

      // if the current cell is not an obstacle, set the current cell to be
      // equal to the cell above it
      table[r][0] = table[r - 1][0];
    }

    for (int c = 1; c < col; c++) {
      // if the current cell is an obstacle, break the loop
      // leave the rest of the cells in that column as 0
      if (obstacleGrid[0][c] == 1) {
        break;
      }

      // if the current cell is not an obstacle, set the current cell to be
      // equal to the cell to the left of it
      table[0][c] = table[0][c - 1];
    }

    for (int r = 1; r < row; r++) {
      for (int c = 1; c < col; c++) {
        // if the current cell is an obstacle, set the current cell to be equal
        // to 0
        if (obstacleGrid[r][c] == 1) {
          continue;
        }

        // if the current cell is not an obstacle, set the current cell to be
        // equal to the sum of the cell above it and the cell to the left of it
        table[r][c] = table[r - 1][c] + table[r][c - 1];
      }
    }

    return table[row - 1][col - 1];
  }
};

// Dynamic Programming Approach (Bottom-Up) With 1D Array
class Solution {
 public:
  int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
    if (obstacleGrid[0][0] == 1) return 0;

    int row = obstacleGrid.size();
    int col = obstacleGrid[0].size();

    vector<int> table(col, 0);
    table[0] = 1;

    for (int r = 0; r < row; r++) {
      for (int c = 0; c < col; c++) {
        if (obstacleGrid[r][c] == 1) {
          table[c] = 0;
          continue;
        }

        if (c > 0) table[c] += table[c - 1];
      }
    }

    return table[col - 1];
  }
};