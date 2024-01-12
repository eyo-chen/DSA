#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<int>> combinationSum3(int k, int n) {
    vector<vector<int>> ans;
    vector<int> temp;

    helper(ans, temp, k, n, 1);

    return ans;
  }

  void helper(vector<vector<int>>& ans, vector<int>& temp, int k, int n,
              int i) {
    if (n == 0 && temp.size() == k) {
      ans.push_back(temp);
      return;
    }

    // if the current size is equal to k, and n is still greater than 0, then
    // there is no need to continue
    if (temp.size() == k && n > 0) return;

    // if n is less than 0, then there is no need to continue
    if (n < 0) return;

    // if i is greater than 9, time to stop
    if (i > 9) return;

    temp.push_back(i);
    helper(ans, temp, k, n - i, i + 1);

    temp.pop_back();
    helper(ans, temp, k, n, i + 1);
  }
};

class Solution {
 public:
  vector<vector<int>> combinationSum3(int k, int n) {
    vector<vector<int>> ans;
    vector<int> temp;

    helper(ans, temp, k, n, 1);

    return ans;
  }

  void helper(vector<vector<int>>& ans, vector<int>& temp, int k, int n,
              int i) {
    if (n == 0 && temp.size() == k) {
      ans.push_back(temp);
      return;
    }

    if (temp.size() == k && n > 0) return;
    if (n < 0) return;
    if (i > 9) return;

    for (int j = i; j <= 9; j++) {
      temp.push_back(j);
      helper(ans, temp, k, n - j, j + 1);
      temp.pop_back();
    }
  }
};