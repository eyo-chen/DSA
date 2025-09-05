#include <queue>
#include <set>
#include <stack>
#include <vector>

using namespace std;

class Solution {
 public:
  void solve(vector<vector<char>>& board) {
    set<pair<int, int>> safeRegion;
    int lastRow = board.size() - 1;
    int lastCol = board[0].size() - 1;

    // check the first and last row(=)
    for (int r = 0; r < board.size(); r++) {
      if (board[r][0] == 'O') helper(board, r, 0, safeRegion);
      if (board[r][lastCol] == 'O') helper(board, r, lastCol, safeRegion);
    }

    // check the first and last column(||)
    for (int c = 0; c < board[0].size(); c++) {
      if (board[0][c] == 'O') helper(board, 0, c, safeRegion);
      if (board[lastRow][c] == 'O') helper(board, lastRow, c, safeRegion);
    }

    // iterate through the board and change the 'O' to 'X' if it is not in the
    // safeRegion
    for (int r = 0; r < board.size(); r++) {
      for (int c = 0; c < board[r].size(); c++) {
        pair<int, int> curPoint = {r, c};
        if (board[r][c] == 'O' &&
            safeRegion.find(curPoint) == safeRegion.end()) {
          board[r][c] = 'X';
        }
      }
    }
  }

  void helper(vector<vector<char>>& board, int r, int c,
              set<pair<int, int>>& safeRegion) {
    int rowLen = board.size();
    int colLen = board[0].size();

    queue<pair<int, int>> q;
    q.push({r, c});

    while (!q.empty()) {
      pair<int, int> curPoint = q.front();
      q.pop();

      int r = curPoint.first;
      int c = curPoint.second;

      // check if the location is out of bound
      if (r < 0 || c < 0 || r >= rowLen || c >= colLen) continue;

      // check if the location is already explored
      if (safeRegion.find(curPoint) != safeRegion.end()) continue;

      // check if the location is not 'O'
      // if it's not 'O', then there's no need to explore the adjacent locations
      if (board[r][c] == 'X') continue;

      // mark the location as safe
      safeRegion.insert(curPoint);

      // push the adjacent locations into the data structure
      q.push({r + 1, c});
      q.push({r - 1, c});
      q.push({r, c + 1});
      q.push({r, c - 1});
    }
  }
};