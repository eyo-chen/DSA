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
  TreeNode* deleteNode(TreeNode* root, int key) {
    // If the root is nullptr, we can return nullptr directly.
    if (root == nullptr) return nullptr;

    // When the key is greater than the root value, we need to delete the key
    // from the right subtree.
    if (key > root->val) {
      root->right = deleteNode(root->right, key);
      return root;
    }

    // When the key is less than the root value, we need to delete the key from
    // the left subtree.
    if (key < root->val) {
      root->left = deleteNode(root->left, key);
      return root;
    }

    // When the key is equal to the root value, we need to delete the root node.

    // If the root node is a leaf node, we can simply delete it.(return nullptr)
    if (root->left == nullptr && root->right == nullptr) return nullptr;

    // If the root node has only one child, we can return the non-nullptr child
    // node.
    if (root->left == nullptr) return root->right;
    if (root->right == nullptr) return root->left;

    // If the root node has two children, we need to find the leftmost node in
    // the right subtree, replace the root value with the leftmost node value,
    // and delete the leftmost node from the right subtree.
    TreeNode* leftMostNodeInRightSubTree = root->right;
    while (leftMostNodeInRightSubTree->left != nullptr) {
      leftMostNodeInRightSubTree = leftMostNodeInRightSubTree->left;
    }
    root->val = leftMostNodeInRightSubTree->val;
    root->right = deleteNode(root->right, leftMostNodeInRightSubTree->val);
    return root;
  }
};