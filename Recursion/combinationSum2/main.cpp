#include <iostream>
#include <map>
#include <vector>

using namespace std;

// Use hashTable and unique vector
class Solution {
 public:
  vector<vector<int>> combinationSum2(vector<int>& candidates, int target) {
    map<int, int> hashTable;
    for (const auto& e : candidates) {
      hashTable[e]++;
    }

    // sort the vector and remove the duplicate elements
    sort(candidates.begin(), candidates.end());
    auto last = unique(candidates.begin(), candidates.end());
    candidates.erase(last, candidates.end());

    vector<vector<int>> ans;
    vector<int> temp;

    helper(hashTable, candidates, ans, temp, target, 0);

    return ans;
  }

  void helper(map<int, int>& hashTable, vector<int>& candidates,
              vector<vector<int>>& ans, vector<int>& temp, int target,
              int index) {
    if (target == 0) {
      ans.push_back(temp);
      return;
    }

    for (int i = index; i < candidates.size(); i++) {
      int num = candidates[i];

      if (target < num) break;
      if (hashTable[num] == 0) continue;

      temp.push_back(num);
      hashTable[num]--;
      helper(hashTable, candidates, ans, temp, target - num, i);
      temp.pop_back();
      hashTable[num]++;
    }
  }
};

// Without hashTable and unique vector
class Solution {
 public:
  vector<vector<int>> combinationSum2(vector<int>& candidates, int target) {
    vector<vector<int>> ans;
    vector<int> temp;
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
      int curNum = candidates[i];
      if (i > index && curNum == candidates[i - 1]) continue;
      if (target < curNum) break;

      temp.push_back(curNum);
      helper(candidates, ans, temp, target - curNum, i + 1);
      temp.pop_back();
    }
  }
};