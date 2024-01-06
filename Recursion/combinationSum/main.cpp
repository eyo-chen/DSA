#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
    vector<vector<int>> ans;
    vector<int> temp;

    // sort the vector for small optimization
    sort(candidates.begin(), candidates.end());

    helper(candidates, ans, temp, target, 0);

    return ans;
  }

  void helper(vector<int>& candidates, vector<vector<int>>& ans,
              vector<int>& temp, int target, int index) {
    if (target == 0) {
      ans.push_back(temp);
      return;
    }

    for (int i = index; i < candidates.size(); i++) {
      /*
      Because the vector is sorted, if the target is smaller than the current
      element, then the rest of the elements will be larger than the target for
      sure. There's no need to continue to explore.

      It also replace the base case -> condition of "if (target < 0)".
      */
      if (target < candidates[i]) {
        return;
      }

      temp.push_back(candidates[i]);
      helper(candidates, ans, temp, target - candidates[i], i);
      temp.pop_back();
    }
  }
};