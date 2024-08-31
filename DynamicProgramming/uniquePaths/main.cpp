#include <string>
#include <unordered_map>
#include <vector>
using namespace std;

// Brute Force Approach
class Solution {
 public:
  int uniquePaths(int m, int n) { return helpers(m, n, 0, 0); }

  int helpers(int m, int n, int r, int c) {
    if (r < 0 || c < 0 || r >= m || c >= n) return 0;
    if (r == m - 1 && c == n - 1) return 1;

    int right = helpers(m, n, r + 1, c);
    int left = helpers(m, n, r, c + 1);

    return right + left;
  }
};

// Memoization Approach
class Solution {
 public:
  int uniquePaths(int m, int n) {
    unordered_map<string, int> memo;
    return helpers(m, n, 0, 0, memo);
  }
  int helpers(int m, int n, int r, int c, unordered_map<string, int>& memo) {
    if (r < 0 || c < 0 || r >= m || c >= n) return 0;
    if (r == m - 1 && c == n - 1) return 1;

    string key = to_string(r) + "," + to_string(c);
    if (memo.find(key) != memo.end()) return memo[key];

    int right = helpers(m, n, r + 1, c, memo);
    int left = helpers(m, n, r, c + 1, memo);

    memo[key] = right + left;

    return right + left;
  }
};

// Dynamic Programming Approach (Bottom-Up)
class Solution {
 public:
  int uniquePaths(int m, int n) {
    vector<vector<int>> table(m, vector<int>(n, 0));

    for (int i = 0; i < m; i++) {
      table[i][0] = 1;
    }

    for (int i = 0; i < n; i++) {
      table[0][i] = 1;
    }

    for (int i = 1; i < m; i++) {
      for (int j = 1; j < n; j++) {
        table[i][j] = table[i - 1][j] + table[i][j - 1];
      }
    }

    return table[m - 1][n - 1];
  }
};
