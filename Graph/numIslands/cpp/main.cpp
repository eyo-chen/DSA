#include <vector>

using namespace std;

class Solution {
 public:
  int numIslands(vector<vector<char>>& grid) {
    int result = 0;

    for (int r = 0; r < grid.size(); r++) {
      for (int c = 0; c < grid[r].size(); c++) {
        if (grid[r][c] == '1') {
          helper(grid, r, c);
          result++;
        }
      }
    }

    return result;
  }

  void helper(vector<vector<char>>& grid, int r, int c) {
    if (r < 0 || c < 0 || r >= grid.size() || c >= grid[0].size() ||
        grid[r][c] != '1')
      return;

    grid[r][c] = '0';
    helper(grid, r + 1, c);
    helper(grid, r - 1, c);
    helper(grid, r, c + 1);
    helper(grid, r, c - 1);
  }
};
