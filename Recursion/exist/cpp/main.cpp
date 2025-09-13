#include <string>
#include <vector>

using namespace std;

/*
By mutating the board cell to '*' to indicate that it is visited
There are two base cases:
1. If the index is equal to the length of the word, then we have found the
answer
2. If the current cell is out of bound or the current cell is visited or the
current cell is not equal to the target character, then we return false
*/
class Solution {
 public:
  bool exist(vector<vector<char>>& board, string word) {
    for (int r = 0; r < board.size(); r++) {
      for (int c = 0; c < board[r].size(); c++) {
        if (helper(board, word, 0, r, c)) {
          return true;
        }
      }
    }

    return false;
  }

  bool helper(vector<vector<char>>& board, string word, int index, int r,
              int c) {
    if (index == word.length()) {
      return true;
    }

    if (r < 0 || r >= board.size() || c < 0 || c >= board[0].size() ||
        board[r][c] == '*' || board[r][c] != word[index]) {
      return false;
    }

    char cr = board[r][c];
    board[r][c] = '*';

    // It's okay to directly explore the four directions
    // Because the second base case will handle the invalid case
    bool res = helper(board, word, index + 1, r + 1, c) ||
               helper(board, word, index + 1, r - 1, c) ||
               helper(board, word, index + 1, r, c + 1) ||
               helper(board, word, index + 1, r, c - 1);

    board[r][c] = cr;

    return res;
  }
};

/*
This is my initial solution
Use a 2D vector to record the visited state(whether it is visited or not)
vector<vector<bool>> record

Note that remember to restore the prior state when we backtrack(unchoose)
*/
class Solution {
 public:
  bool exist(vector<vector<char>>& board, string word) {
    vector<vector<bool>> record(board.size(),
                                vector<bool>(board[0].size(), true));

    for (int r = 0; r < board.size(); r++) {
      for (int c = 0; c < board[r].size(); c++) {
        if (board[r][c] == word[0]) {
          record[r][c] = false;
          if (helper(board, record, word, 1, r, c)) {
            return true;
          }
          record[r][c] = true;
        }
      }
    }

    return false;
  }

  bool helper(vector<vector<char>>& board, vector<vector<bool>>& record,
              string word, int index, int r, int c) {
    if (index == word.length()) {
      return true;
    }

    char curTarget = word[index];

    // right
    if (c < board[0].size() - 1 && record[r][c + 1] &&
        board[r][c + 1] == curTarget) {
      record[r][c + 1] = false;
      if (helper(board, record, word, index + 1, r, c + 1)) {
        return true;
      }
      record[r][c + 1] = true;
    }

    // left
    if (c > 0 && record[r][c - 1] && board[r][c - 1] == curTarget) {
      record[r][c - 1] = false;
      if (helper(board, record, word, index + 1, r, c - 1)) {
        return true;
      }
      record[r][c - 1] = true;
    }

    // down
    if (r < board.size() - 1 && record[r + 1][c] &&
        board[r + 1][c] == curTarget) {
      record[r + 1][c] = false;
      if (helper(board, record, word, index + 1, r + 1, c)) {
        return true;
      }
      record[r + 1][c] = true;
    }

    // top
    if (r > 0 && record[r - 1][c] && board[r - 1][c] == curTarget) {
      record[r - 1][c] = false;
      if (helper(board, record, word, index + 1, r - 1, c)) {
        return true;
      }
      record[r - 1][c] = true;
    }

    return false;
  }
};