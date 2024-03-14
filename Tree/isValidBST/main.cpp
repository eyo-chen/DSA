#include <limits>

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

class Solution {
 public:
  bool isValidBST(TreeNode* root) {
    if (root == nullptr) return true;

    // using long to be larger than the max value of problem description
    long max = numeric_limits<long>::max();
    long min = numeric_limits<long>::min();
    return helper(root, min, max);
  }

  bool helper(TreeNode* node, long min, long max) {
    if (node == nullptr) return true;
    if (node->val <= min) return false;
    if (node->val >= max) return false;

    return helper(node->left, min, node->val) &&
           helper(node->right, node->val, max);
  }
};