#include <queue>

using namespace std;

// Definition for a Node.
class Node {
 public:
  int val;
  Node* left;
  Node* right;
  Node* next;

  Node() : val(0), left(nullptr), right(nullptr), next(nullptr) {}

  Node(int _val) : val(_val), left(nullptr), right(nullptr), next(nullptr) {}

  Node(int _val, Node* _left, Node* _right, Node* _next)
      : val(_val), left(_left), right(_right), next(_next) {}
};

// Using BFS
class Solution {
 public:
  Node* connect(Node* root) {
    if (root == nullptr) return root;

    queue<Node*> q;
    q.push(root);

    while (!q.empty()) {
      int size = q.size();
      Node* preNode = nullptr;

      for (int i = 0; i < size; i++) {
        Node* curNode = q.front();
        q.pop();

        if (curNode->left != nullptr) q.push(curNode->left);
        if (curNode->right != nullptr) q.push(curNode->right);

        if (preNode != nullptr) preNode->next = curNode;
        preNode = curNode;
      }
    }

    return root;
  }
};

// Using DFS
class Solution {
 public:
  Node* connect(Node* root) {
    if (root == nullptr) return nullptr;

    // wire-up the left child
    if (root->left != nullptr) {
      root->left->next = root->right;
    }

    // wire-up the right child
    if (root->right != nullptr && root->next != nullptr) {
      root->right->next = root->next->left;
    }

    root->left = connect(root->left);
    root->right = connect(root->right);

    return root;
  }
};

// Iteration
class Solution {
 public:
  Node* connect(Node* root) {
    if (root == nullptr) return nullptr;

    Node* leftMostNode = root;
    while (leftMostNode->left != nullptr) {
      Node* curNode = leftMostNode;

      while (curNode != nullptr) {
        // wire-up the left child
        curNode->left->next = curNode->right;

        // wire-up the right child
        if (curNode->next != nullptr) {
          curNode->right->next = curNode->next->left;
        }

        // move to the next node in the same level
        curNode = curNode->next;
      }

      // move to the next level
      leftMostNode = leftMostNode->left;
    }

    return root;
  }
};