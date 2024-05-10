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

// Time complexity: O(n)
// Space complexity: O(h)
class Solution {
 public:
  int goodNodes(TreeNode* root) { return helper(root, root->val); }

  int helper(TreeNode* node, int maxVal) {
    if (node == nullptr) return 0;

    // check if the current node is a good node
    // if the current node's value is greater than or equal to maxVal,
    // it means it is a good node
    // we do two things here:
    // 1. update the current max value (for the next recursive call)
    // 2. set the isGoodNode to 1 (for the return value)
    int curMaxVal = maxVal;
    int isGoodNode = 0;
    if (node->val >= maxVal) {
      curMaxVal = node->val;
      isGoodNode = 1;
    }

    int leftCount = helper(node->left, curMaxVal);
    int rightCount = helper(node->right, curMaxVal);

    return leftCount + rightCount + isGoodNode;
  }
};