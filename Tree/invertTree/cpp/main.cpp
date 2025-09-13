#include <stack>

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

// using recursion
class Solution {
 public:
  TreeNode* invertTree(TreeNode* root) {
    if (root == nullptr) return root;

    TreeNode* left = invertTree(root->left);
    TreeNode* right = invertTree(root->right);

    root->left = right;
    root->right = left;

    return root;
  }
};

// using iteration(stack)
class Solution {
 public:
  TreeNode* invertTree(TreeNode* root) {
    stack<TreeNode*> s;
    s.push(root);

    while (!s.empty()) {
      TreeNode* node = s.top();
      s.pop();
      if (node == nullptr) continue;

      TreeNode* left = node->left;
      TreeNode* right = node->right;

      s.push(left);
      s.push(right);

      node->left = right;
      node->right = left;
    }

    return root;
  }
};
