#include <queue>
#include <set>
#include <vector>

using namespace std;

// Using set to store seen points
class Solution {
 public:
  int maxAreaOfIsland(vector<vector<int>>& grid) {
    set<pair<int, int>> seen;
    int maxCount = 0;

    for (int r = 0; r < grid.size(); r++) {
      for (int c = 0; c < grid[r].size(); c++) {
        if (grid[r][c] == 0) continue;

        pair<int, int> point = {r, c};
        if (seen.find(point) != seen.end()) continue;

        int count = helper(grid, r, c, seen);
        maxCount = max(maxCount, count);
      }
    }

    return maxCount;
  }

  // BFS
  int helper(vector<vector<int>>& grid, int row, int col,
             set<pair<int, int>>& seen) {
    int count = 0;
    queue<pair<int, int>> q;
    q.push({row, col});

    while (!q.empty()) {
      pair<int, int> point = q.front();
      q.pop();

      int r = point.first;
      int c = point.second;

      if (r < 0 || c < 0 || r >= grid.size() || c >= grid[0].size()) continue;
      if (grid[r][c] == 0) continue;
      if (seen.find(point) != seen.end()) continue;

      q.push({r + 1, c});
      q.push({r - 1, c});
      q.push({r, c + 1});
      q.push({r, c - 1});

      count++;
      seen.insert(point);
    }

    return count;
  }
};

// Using vector to store seen points
class Solution {
 public:
  int maxAreaOfIsland(vector<vector<int>>& grid) {
    int maxCount = 0;
    vector<vector<bool>> seen(grid.size(), vector<bool>(grid[0].size(), false));

    for (int r = 0; r < grid.size(); r++) {
      for (int c = 0; c < grid[r].size(); c++) {
        if (grid[r][c] == 0 || seen[r][c]) continue;

        int count = helper(grid, r, c, seen);
        maxCount = max(maxCount, count);
      }
    }

    return maxCount;
  }

  int helper(vector<vector<int>>& grid, int row, int col,
             vector<vector<bool>>& seen) {
    int count = 0;
    queue<pair<int, int>> q;
    q.push({row, col});
    seen[row][col] = true;

    while (!q.empty()) {
      pair<int, int> point = q.front();
      q.pop();

      int r = point.first;
      int c = point.second;

      if (r < 0 || c < 0 || r >= grid.size() || c >= grid[0].size()) continue;
      if (grid[r][c] == 0) continue;

      static const vector<pair<int, int>> directions = {
          {1, 0}, {-1, 0}, {0, 1}, {0, -1}};
      for (const auto& dir : directions) {
        int newRow = r + dir.first;
        int newCol = c + dir.second;

        if (newRow >= 0 && newCol >= 0 && newRow < grid.size() &&
            newCol < grid[0].size() && grid[newRow][newCol] == 1 &&
            !seen[newRow][newCol]) {
          q.push({newRow, newCol});
          seen[newRow][newCol] = true;
        }
      }

      count++;
    }

    return count;
  }
};
