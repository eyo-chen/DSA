#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<int>> subsets(vector<int>& nums) {
    vector<int> temp;
    vector<vector<int>> ans;
    helper(nums, ans, temp, 0);

    return ans;
  }

  void helper(vector<int>& nums, vector<vector<int>>& ans, vector<int>& temp,
              int index) {
    if (index == nums.size()) {
      ans.push_back(temp);
      return;
    }

    int cur = nums[index];

    temp.push_back(cur);
    helper(nums, ans, temp, index + 1);
    temp.pop_back();
    helper(nums, ans, temp, index + 1);
  }
};