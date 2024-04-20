#include <vector>

using namespace std;

//  Definition for a binary tree node.
struct TreeNode {
  int val;
  TreeNode* left;
  TreeNode* right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode* left, TreeNode* right)
      : val(x), left(left), right(right) {}
};

// Correct Solution
class Solution {
 public:
  vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
    vector<vector<int>> res;
    vector<int> curPath;
    helper(root, res, curPath, targetSum);

    return res;
  }

  void helper(TreeNode* root, vector<vector<int>>& res, vector<int>& curPath,
              int targetSum) {
    if (root == nullptr) {
      return;
    }

    int remainingSum = targetSum - root->val;
    curPath.push_back(root->val);

    if (root->left == nullptr && root->right == nullptr) {
      if (remainingSum == 0) {
        res.push_back(curPath);
      }

      curPath.pop_back();
      return;
    }

    helper(root->left, res, curPath, remainingSum);
    helper(root->right, res, curPath, remainingSum);
    curPath.pop_back();
  }
};

// Wrong Solution
class Solution {
 public:
  vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
    vector<vector<int>> res;

    if (targetSum == 0) {
      return res;
    }

    vector<int> curPath;
    helper(root, res, curPath, targetSum);

    return res;
  }

  void helper(TreeNode* root, vector<vector<int>>& res, vector<int>& curPath,
              int targetSum) {
    if (targetSum == 0) {
      res.push_back(curPath);
      return;
    }

    if (root == nullptr) {
      return;
    }

    curPath.push_back(root->val);
    helper(root->left, res, curPath, targetSum - root->val);
    helper(root->right, res, curPath, targetSum - root->val);
    curPath.pop_back();
  }
};