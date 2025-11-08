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

    if (root->left != nullptr) {
      if (root->right != nullptr)
        root->left->next = root->right;
      else
        root->left->next = findNextNode(root->next);
    }

    if (root->right != nullptr) {
      root->right->next = findNextNode(root->next);
    }

    root->right = connect(root->right);
    root->left = connect(root->left);

    return root;
  }

  Node* findNextNode(Node* node) {
    if (node == nullptr) return node;

    if (node->left != nullptr) return node->left;
    if (node->right != nullptr) return node->right;

    return findNextNode(node->next);
  }
};

// Iterative level by level
class Solution {
 public:
  Node* connect(Node* root) {
    if (root == nullptr) return nullptr;

    Node* leftMostNode = root;

    while (leftMostNode != nullptr) {
      Node* curNode = leftMostNode;

      while (curNode != nullptr) {
        if (curNode->left != nullptr) {
          if (curNode->right != nullptr)
            curNode->left->next = curNode->right;
          else
            curNode->left->next = findNextNode(curNode->next);
        }

        if (curNode->right != nullptr) {
          curNode->right->next = findNextNode(curNode->next);
        }

        curNode = curNode->next;
      }

      leftMostNode = findNextNode(leftMostNode);
    }

    return root;
  }

  Node* findNextNode(Node* node) {
    if (node == nullptr) return node;

    if (node->left != nullptr) return node->left;
    if (node->right != nullptr) return node->right;

    return findNextNode(node->next);
  }
};