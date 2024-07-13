#include <queue>
#include <vector>

using namespace std;

class Solution {
 public:
  int orangesRotting(vector<vector<int>>& grid) {
    vector<pair<int, int>> offsets = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
    queue<pair<int, int>> q;
    int minute = 0;
    int total = 0;

    for (int r = 0; r < grid.size(); r++) {
      for (int c = 0; c < grid[0].size(); c++) {
        // add rotten oranges to queue
        if (grid[r][c] == 2) q.push({r, c});

        // count fresh oranges
        if (grid[r][c] == 1) total++;
      }
    }

    // BFS
    // continue until queue is empty or no fresh oranges left
    while (q.size() > 0 && total > 0) {
      int size = q.size();
      for (int i = 0; i < size; i++) {
        pair<int, int> point = q.front();
        q.pop();

        // check all 4 directions
        int row = point.first;
        int col = point.second;
        for (const pair<int, int> o : offsets) {
          int r = row + o.first;
          int c = col + o.second;

          // check if the orange is out of grid or not fresh
          if (r < 0 || c < 0 || r >= grid.size() || c >= grid[0].size() ||
              grid[r][c] != 1)
            continue;

          total--;
          grid[r][c] = 2;
          q.push({r, c});
        }
      }

      minute++;
    }

    return total != 0 ? -1 : minute;
  }
};