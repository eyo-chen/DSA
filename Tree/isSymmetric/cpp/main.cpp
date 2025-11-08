#include <stack>
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

// Solution 1: Recursive
class Solution {
 public:
  bool isSymmetric(TreeNode* root) { return helper(root->left, root->right); }

  bool helper(TreeNode* node1, TreeNode* node2) {
    if (node1 == nullptr && node2 == nullptr) return true;
    if (node1 == nullptr && node2 != nullptr) return false;
    if (node1 != nullptr && node2 == nullptr) return false;
    /*
    The above two conditions can be combined into one:
    if (node1 == nullptr || node2 == nullptr) return false;
    */

    if (node1->val != node2->val) return false;

    return helper(node1->left, node2->right) &&
           helper(node1->right, node2->left);
  }
};

// Solution 2: Iterative
class Solution {
 public:
  bool isSymmetric(TreeNode* root) {
    stack<TreeNode*> s1;
    stack<TreeNode*> s2;

    s1.push(root->left);
    s2.push(root->right);

    while (!s1.empty() || !s2.empty()) {
      TreeNode* n1 = s1.top();
      TreeNode* n2 = s2.top();

      s1.pop();
      s2.pop();

      if (n1 == nullptr && n2 == nullptr) continue;
      if (n1 == nullptr && n2 != nullptr) return false;
      if (n1 != nullptr && n2 == nullptr) return false;
      /*
      The above two conditions can be combined into one:
      if (n1 == nullptr || n2 == nullptr) return false;
      */

      if (n1->val != n2->val) return false;

      s1.push(n1->left);
      s1.push(n1->right);
      s2.push(n2->right);
      s2.push(n2->left);
    }

    return true;
  }
};