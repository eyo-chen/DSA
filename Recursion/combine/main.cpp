#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<int>> combine(int n, int k) {
    vector<vector<int>> ans;
    vector<int> temp;

    helper(ans, temp, 1, n, k);

    return ans;
  }

  void helper(vector<vector<int>>& ans, vector<int>& temp, int index, int n,
              int k) {
    if (temp.size() == k) {
      ans.push_back(temp);
      return;
    }

    for (int i = index; i <= n; i++) {
      temp.push_back(i);
      helper(ans, temp, i + 1, n, k);
      temp.pop_back();
    }
  }
};