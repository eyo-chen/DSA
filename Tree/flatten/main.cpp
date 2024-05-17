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

// using vector to store the nodes in pre-order
class Solution {
 public:
  void flatten(TreeNode* root) {
    vector<TreeNode*> nodes;
    genPreOrderNodes(root, nodes);

    // reconstruct the tree
    for (int i = 0; i < nodes.size(); i++) {
      TreeNode* node = nodes[i];

      // always set left node to nullptr
      node->left = nullptr;

      // if the current node is the last node, set right node to nullptr
      // otherwise, set right node to the next node
      TreeNode* nextNode = i == nodes.size() - 1 ? nullptr : nodes[i + 1];
      node->right = nextNode;
    }
  }

  void genPreOrderNodes(TreeNode* node, vector<TreeNode*>& nodes) {
    if (node == nullptr) return;

    nodes.push_back(node);
    genPreOrderNodes(node->left, nodes);
    genPreOrderNodes(node->right, nodes);
  }
};

// using helper function to return the tail node of the current subtree
class Solution {
 public:
  void flatten(TreeNode* root) { helper(root); }

  TreeNode* helper(TreeNode* root) {
    if (root == nullptr) return nullptr;

    TreeNode* leftTail = helper(root->left);
    TreeNode* rightTail = helper(root->right);

    // when the leftTail is not nullptr, we need to do the rewiring
    if (leftTail != nullptr) {
      leftTail->right = root->right;
      root->right = root->left;
      root->left = nullptr;
    }

    if (rightTail) return rightTail;
    if (leftTail) return leftTail;
    return root;
  }
};