#include <queue>
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

// using queue and stack
class Solution {
 public:
  vector<vector<int>> levelOrderBottom(TreeNode* root) {
    queue<TreeNode*> q;
    stack<vector<int>> s;
    vector<vector<int>> res;

    if (root == nullptr) return res;

    q.push(root);
    while (!q.empty()) {
      int size = q.size();
      vector<int> vals;

      for (int i = 0; i < size; i++) {
        TreeNode* curNode = q.front();
        q.pop();

        vals.push_back(curNode->val);

        if (curNode->left != nullptr) q.push(curNode->left);
        if (curNode->right != nullptr) q.push(curNode->right);
      }

      // add the current level nodes to the stack
      s.push(vals);
    }

    // pop the stack to get the result
    // make sure the final result is in the bottom-up order
    while (!s.empty()) {
      res.push_back(s.top());
      s.pop();
    }

    return res;
  }
};

// using recursion to count max depth first
// then using recursion to fill the result
class Solution {
 public:
  vector<vector<int>> levelOrderBottom(TreeNode* root) {
    // get the max depth of the tree
    int depth = countDepth(root);
    vector<vector<int>> res(depth, vector<int>{});
    helper(root, res, depth - 1);

    return res;
  }

  void helper(TreeNode* node, vector<vector<int>>& res, int level) {
    if (node == nullptr) return;

    // fill the result from the bottom to the top
    res[level].push_back(node->val);
    helper(node->left, res, level - 1);
    helper(node->right, res, level - 1);
  }

  int countDepth(TreeNode* root) {
    if (root == nullptr) return 0;
    return max(countDepth(root->left), countDepth(root->right)) + 1;
  }
};