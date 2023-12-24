#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;

// First solution
class Solution {
 public:
  vector<vector<int>> permuteUnique(vector<int>& nums) {
    vector<vector<int>> ans;
    vector<int> temp;
    vector<bool> used(nums.size(), false);
    sort(nums.begin(), nums.end());

    helper(nums, ans, temp, used);

    return ans;
  }

  void helper(vector<int>& nums, vector<vector<int>>& ans, vector<int>& temp,
              vector<bool>& used) {
    if (temp.size() == nums.size()) {
      ans.push_back(temp);
      return;
    }

    for (int i = 0; i < nums.size(); i++) {
      if (used[i]) continue;
      if (i > 0 && nums[i] == nums[i - 1] && !used[i - 1]) continue;

      temp.push_back(nums[i]);
      used[i] = true;

      helper(nums, ans, temp, used);

      temp.pop_back();
      used[i] = false;
    }
  }
};

// Second solution
class Solution {
 public:
  vector<vector<int>> permuteUnique(vector<int>& nums) {
    vector<vector<int>> ans;
    vector<int> temp;
    unordered_map<int, int> hashMap;

    for (int n : nums) {
      hashMap[n]++;
    }

    helper(nums, ans, temp, hashMap);

    return ans;
  }

  void helper(vector<int>& nums, vector<vector<int>>& ans, vector<int>& temp,
              unordered_map<int, int>& hashMap) {
    if (temp.size() == nums.size()) {
      ans.push_back(temp);
      return;
    }

    for (const auto& [key, value] : hashMap) {
      if (value == 0) continue;

      temp.push_back(key);
      hashMap[key]--;

      helper(nums, ans, temp, hashMap);

      temp.pop_back();
      hashMap[key]++;
    }
  }
};