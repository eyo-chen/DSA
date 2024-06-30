#include <queue>
#include <vector>

using namespace std;

// put all the locations into the queue, then check validation
class Solution {
 public:
  int nearestExit(vector<vector<char>>& maze, const vector<int>& entrance) {
    int steps = 0;
    queue<pair<int, int>> q;
    q.push({entrance[0], entrance[1]});

    while (!q.empty()) {
      int qSize = q.size();

      for (int i = 0; i < qSize; ++i) {
        auto [r, c] = q.front();
        q.pop();

        // check if the location is valid
        if (r < 0 || c < 0 || r >= maze.size() || c >= maze[0].size() ||
            maze[r][c] != '.') {
          continue;
        }

        // check if the location is the exit
        if ((r != entrance[0] || c != entrance[1]) &&
            (r == 0 || c == 0 || r == maze.size() - 1 ||
             c == maze[0].size() - 1)) {
          return steps;
        }

        maze[r][c] = '-';

        q.push({r + 1, c});
        q.push({r - 1, c});
        q.push({r, c + 1});
        q.push({r, c - 1});
      }

      ++steps;
    }

    return -1;
  }
};

// check validation first, then put all the validate locations into the queue
class Solution {
 public:
  int nearestExit(vector<vector<char>>& maze, const vector<int>& entrance) {
    queue<pair<int, int>> q;
    q.push({entrance[0], entrance[1]});

    const vector<pair<int, int>> offsets = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};

    // set the initial steps to 1
    int steps = 1;

    while (!q.empty()) {
      int qSize = q.size();

      for (int i = 0; i < qSize; ++i) {
        pair<int, int> location = q.front();
        q.pop();

        int row = location.first;
        int col = location.second;

        for (const pair<int, int>& offset : offsets) {
          int r = row + offset.first;
          int c = col + offset.second;

          // check if the location is valid
          if (r < 0 || c < 0 || r >= maze.size() || c >= maze[0].size() ||
              maze[r][c] != '.') {
            continue;
          }

          // check if the location is the exit
          if ((r != entrance[0] || c != entrance[1]) &&
              (r == 0 || c == 0 || r == maze.size() - 1 ||
               c == maze[0].size() - 1)) {
            return steps;
          }

          q.push({r, c});
          maze[r][c] = '-';
        }
      }
      ++steps;
    }

    return -1;
  }
};
