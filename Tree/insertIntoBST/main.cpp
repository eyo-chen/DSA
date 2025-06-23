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

// Recursive Solution
class Solution {
 public:
  TreeNode* insertIntoBST(TreeNode* root, int val) {
    if (root == nullptr) return new TreeNode(val);

    if (val > root->val) {
      root->right = insertIntoBST(root->right, val);
    } else {
      root->left = insertIntoBST(root->left, val);
    }

    return root;
  }
};

// Iterative Solution1
class Solution {
 public:
  TreeNode* insertIntoBST(TreeNode* root, int val) {
    if (root == nullptr) return new TreeNode(val);

    TreeNode* cur = root;
    TreeNode* pre = nullptr;

    while (cur != nullptr) {
      pre = cur;

      if (val > cur->val) {
        cur = cur->right;
      } else {
        cur = cur->left;
      }
    }

    if (val > pre->val) {
      pre->right = new TreeNode(val);
    } else {
      pre->left = new TreeNode(val);
    }

    return root;
  }
};

// Iterative Solution2
class Solution {
 public:
  TreeNode* insertIntoBST(TreeNode* root, int val) {
    if (root == nullptr) return new TreeNode(val);

    TreeNode* cur = root;

    while (true) {
      // go right
      if (val > cur->val) {
        if (cur->right == nullptr) {
          cur->right = new TreeNode(val);
          break;
        }

        cur = cur->right;
      }
      // go left
      else {
        if (cur->left == nullptr) {
          cur->left = new TreeNode(val);
          break;
        }

        cur = cur->left;
      }
    }
    return root;
  }
};