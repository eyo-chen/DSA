#include <stack>
#include <unordered_map>
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

// Recursion
class Solution {
 public:
  vector<int> inorderTraversal(TreeNode* root) {
    vector<int> res;
    helper(root, res);
    return res;
  }

  void helper(TreeNode* node, vector<int>& res) {
    if (node == nullptr) return;

    helper(node->left, res);
    res.push_back(node->val);
    helper(node->right, res);
  }
};

// Using Stack and HashMap
class Solution {
 public:
  vector<int> inorderTraversal(TreeNode* root) {
    vector<int> res;
    if (root == nullptr) return res;

    stack<TreeNode*> s;
    unordered_map<TreeNode*, bool> hashMap;

    s.push(root);

    while (!s.empty()) {
      TreeNode* n = s.top();
      if (n->left != nullptr && !hashMap[n->left]) {
        s.push(n->left);
        hashMap[n->left] = true;
        continue;
      }

      res.push_back(n->val);
      s.pop();

      if (n->right != nullptr && !hashMap[n->right]) {
        s.push(n->right);
        hashMap[n->right] = true;
      }
    }

    return res;
  }
};

// Using Stack
class Solution {
 public:
  vector<int> inorderTraversal(TreeNode* root) {
    vector<int> res;

    if (root == nullptr) return res;

    stack<TreeNode*> s;
    TreeNode* curNode = root;

    while (!s.empty() || curNode != nullptr) {
      while (curNode != nullptr) {
        s.push(curNode);
        curNode = curNode->left;
      }

      curNode = s.top();
      s.pop();
      res.push_back(curNode->val);

      curNode = curNode->right;
    }

    return res;
  }
};
// ```
//       1
//      / \
//     2   3
//    / \
//   4   5
//      / \
//     6   7
// ```