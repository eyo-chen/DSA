#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<int>> permute(vector<int>& nums) {
    vector<vector<int>> ans;
    vector<int> tmp;
    vector<bool> used(nums.size(), false);

    helper(nums, ans, tmp, used);

    return ans;
  }

  void helper(vector<int>& nums, vector<vector<int>>& ans, vector<int>& tmp,
              vector<bool>& used) {
    if (nums.size() == tmp.size()) {
      ans.push_back(tmp);
      return;
    }

    for (int i = 0; i < nums.size(); i++) {
      if (used[i]) continue;

      tmp.push_back(nums[i]);
      used[i] = true;

      helper(nums, ans, tmp, used);

      tmp.pop_back();
      used[i] = false;
    }
  }
};