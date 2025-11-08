#include <algorithm>
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

// Using Queue
class Solution {
 public:
  int minDepth(TreeNode* root) {
    if (root == nullptr) return 0;

    int min = 1;
    queue<TreeNode*> q;
    q.push(root);

    while (!q.empty()) {
      int size = q.size();

      for (int i = 0; i < size; i++) {
        TreeNode* curNode = q.front();
        q.pop();

        if (curNode->left == nullptr && curNode->right == nullptr) return min;

        // we choose to do the guard clause here to avoid pushing nullptr into
        // the queue
        if (curNode->left != nullptr) q.push(curNode->left);
        if (curNode->right != nullptr) q.push(curNode->right);
      }
      min++;
    }

    // In normal case, this line will not be executed
    return min;
  }
};

// Using Queue with struct
class Solution {
 public:
  int minDepth(TreeNode* root) {
    if (root == nullptr) return 0;

    struct NodeWithHeight {
      TreeNode* node;
      int height;
    };

    queue<NodeWithHeight> q;
    q.push({root, 1});

    while (!q.empty()) {
      NodeWithHeight current = q.front();
      q.pop();

      if (current.node == nullptr) continue;
      if (current.node->left == nullptr && current.node->right == nullptr)
        return current.height;

      q.push({current.node->left, current.height + 1});
      q.push({current.node->right, current.height + 1});
    }

    return 0;
  }
};

// Using Recursion
class Solution {
 public:
  int minDepth(TreeNode* root) {
    if (root == nullptr) return 0;

    int left = minDepth(root->left);
    int right = minDepth(root->right);

    // Both left and right are 0, then it is a leaf node
    if (left == 0 && right == 0) return 1;

    // If either left or right is 0, then we need to consider the non-zero
    if (left == 0) return right + 1;
    if (right == 0) return left + 1;

    // If both left and right are non-zero, then we need to consider the minimum
    return min(left, right) + 1;
  }
};