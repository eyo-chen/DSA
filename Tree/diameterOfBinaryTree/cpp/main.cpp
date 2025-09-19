#include <algorithm>

using namespace std;

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
  int diameterOfBinaryTree(TreeNode* root) {
    int maxValue = 0;
    helper(root, maxValue);
    return maxValue;
  }

  int helper(TreeNode* node, int& maxValue) {
    if (node == nullptr) return 0;

    int left = helper(node->left, maxValue);
    int right = helper(node->right, maxValue);

    maxValue = max(left + right, maxValue);

    return max(left, right) + 1;
  }
};