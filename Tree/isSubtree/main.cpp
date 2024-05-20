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
  bool isSubtree(TreeNode* root, TreeNode* subRoot) {
    if (root == nullptr) return false;

    // at the current state of subtree, ask
    // if the current subtree is same as subRoot
    // if it is, return true
    // if not, go to left and right subtree and keep asking the same question
    if (isSame(root, subRoot)) return true;

    return isSubtree(root->left, subRoot) || isSubtree(root->right, subRoot);
  }

  bool isSame(TreeNode* a, TreeNode* b) {
    if (a == nullptr && b == nullptr) return true;
    if (a == nullptr || b == nullptr) return false;

    if (a->val != b->val) return false;
    return isSame(a->left, b->left) && isSame(a->right, b->right);
  }
};