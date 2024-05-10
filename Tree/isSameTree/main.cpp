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
  bool isSameTree(TreeNode* p, TreeNode* q) {
    // if both are nullptr, it means they are same
    if (p == nullptr && q == nullptr) return true;

    // if one of them is nullptr, it means they are not same
    if (p == nullptr || q == nullptr) return false;

    // if value is different, it means they are not same
    if (p->val != q->val) return false;

    // have to make sure left and right are same
    return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
  }
};