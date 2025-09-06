#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<int>> permute(vector<int>& nums) {
    vector<vector<int>> ans;
    vector<int> temp;
    vector<bool> used(nums.size(), false);

    helper(nums, ans, temp, used);

    return ans;
  }

  void helper(vector<int>& nums, vector<vector<int>>& ans, vector<int>& temp,
              vector<bool>& used) {
    if (nums.size() == temp.size()) {
      ans.push_back(temp);
      return;
    }

    for (int i = 0; i < nums.size(); i++) {
      if (used[i]) continue;

      temp.push_back(nums[i]);
      used[i] = true;

      helper(nums, ans, temp, used);

      temp.pop_back();
      used[i] = false;
    }
  }
};