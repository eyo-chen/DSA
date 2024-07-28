#include <queue>
#include <vector>

using namespace std;

struct location {
  int row;
  int col;
  int steps;
  int k;
};

// Using BFS
class Solution {
 public:
  int shortestPath(vector<vector<int>>& grid, int k) {
    // initialize the queue with the starting location
    queue<location> q({{0, 0, 0, k}});

    // offsets for the four directions
    vector<vector<int>> offsets = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};

    // visited locations
    // for each location, we need to keep track of the state of the key
    // for example, if we have visited location (1, 2) with 1 key, we should not
    // visit the same location with the same key again so we need a 3D array to
    // keep track of the visited locations
    vector<vector<vector<bool>>> visited(
        grid.size(),
        vector<vector<bool>>(grid[0].size(), vector<bool>(k + 1, false)));
    visited[0][0][k] = true;

    while (q.size() > 0) {
      location l = q.front();
      q.pop();

      // check if we have reached the destination
      // note that the index is 0-based, so we need to subtract 1 from the
      // length
      if (l.row == grid.size() - 1 && l.col == grid[0].size() - 1)
        return l.steps;

      // loop through the four directions(adjacent locations)
      for (const vector<int>& o : offsets) {
        int r = l.row + o[0];
        int c = l.col + o[1];

        // check if the location is out of the boundary
        if (r < 0 || c < 0 || r >= grid.size() || c >= grid[0].size()) continue;

        // check if the location is obstacle and we don't have any key left
        if (grid[r][c] == 1 && l.k <= 0) continue;

        // update the key (if we have encountered an obstacle)
        int newK = grid[r][c] == 1 ? l.k - 1 : l.k;

        // check if the location has been visited
        // if the location has been visited with the same key, we should not
        // visit it again
        if (visited[r][c][newK]) continue;

        // add the location to the queue
        location neighbor = {r, c, l.steps + 1, newK};
        q.push(neighbor);

        // mark the location as visited
        visited[r][c][newK] = true;
      }
    }

    return -1;
  }
};

class Solution {
 public:
  int shortestPath(vector<vector<int>>& grid, int k) {
    int result = dfs(grid, 0, 0, 0, k);
    return result == INT_MAX ? -1 : result;
  }

  int dfs(vector<vector<int>>& grid, int r, int c, int steps, int k) {
    // check if the location is out of the boundary
    if (r < 0 || c < 0 || r >= grid.size() || c >= grid[0].size())
      return INT_MAX;

    // check if the location is visited
    if (grid[r][c] == 2) return INT_MAX;

    // check if the location is obstacle and we don't have any key left
    if (grid[r][c] == 1 && k == 0) return INT_MAX;

    // check if we have reached the destination
    if (r == grid.size() - 1 && c == grid[0].size() - 1) return steps;

    // update the key (if we have encountered an obstacle)
    int newK = grid[r][c] == 1 ? k - 1 : k;

    int state = grid[r][c];

    // mark the location as visited
    // we can directly update the state of the location to 2 as visited, which
    // is incorrect in the BFS solution because we explore each path in the DFS
    // solution, so we guarantee that there's only one state for each location
    grid[r][c] = 2;

    // explore the four directions
    int upSteps = dfs(grid, r, c + 1, steps + 1, newK);
    int downSteps = dfs(grid, r, c - 1, steps + 1, newK);
    int rightSteps = dfs(grid, r + 1, c, steps + 1, newK);
    int leftSteps = dfs(grid, r - 1, c, steps + 1, newK);

    // restore the state of the location
    grid[r][c] = state;

    return min({upSteps, downSteps, rightSteps, leftSteps});
  }
};