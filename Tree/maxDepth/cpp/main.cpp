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

class Solution {
 public:
  int maxDepth(TreeNode* root) {
    if (root == nullptr) return 0;

    int left = maxDepth(root->left) + 1;
    int right = maxDepth(root->right) + 1;

    return max(left, right);
  }
};

class Solution {
 public:
  int maxDepth(TreeNode* root) {
    if (root == nullptr) return 0;

    int maxHeight = 0;
    struct HeightNode {
      TreeNode* node;
      int height;
    };

    stack<HeightNode*> s;

    s.push(new HeightNode{root, 1});

    while (!s.empty()) {
      HeightNode* d = s.top();
      s.pop();

      if (d->node == nullptr) continue;

      maxHeight = max(maxHeight, d->height);

      HeightNode* left = new HeightNode{d->node->left, d->height + 1};
      HeightNode* right = new HeightNode{d->node->right, d->height + 1};

      s.push(left);
      s.push(right);
    }

    return maxHeight;
  }
};

class Solution {
 public:
  int maxDepth(TreeNode* root) {
    if (root == nullptr) return 0;

    int height = 0;
    queue<TreeNode*> q;
    q.push(root);

    while (!q.empty()) {
      height++;

      int size = q.size();

      for (int i = 0; i < size; i++) {
        TreeNode* n = q.front();
        q.pop();

        if (n->left != nullptr) q.push(n->left);
        if (n->right != nullptr) q.push(n->right);
      }
    }

    return height;
  }
};