#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<int>> subsets(vector<int>& nums) {
    vector<int> tmp;
    vector<vector<int>> ans;
    helper(nums, ans, tmp, 0);

    return ans;
  }

  void helper(vector<int>& nums, vector<vector<int>>& ans, vector<int>& tmp,
              int index) {
    if (index == nums.size()) {
      ans.push_back(tmp);
      return;
    }

    int cur = nums[index];

    tmp.push_back(cur);
    helper(nums, ans, tmp, index + 1);
    tmp.pop_back();
    helper(nums, ans, tmp, index + 1);
  }
};