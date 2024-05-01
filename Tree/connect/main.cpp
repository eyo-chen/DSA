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

class Solution {
 public:
  Node* connect(Node* root) {
    if (root == nullptr) return nullptr;

    if (root->left != nullptr) {
      root->left->next = root->right;
    }

    if (root->right != nullptr && root->next != nullptr) {
      root->right->next = root->next->left;
    }

    root->left = connect(root->left);
    root->right = connect(root->right);

    return root;
  }
};

class Solution {
 public:
  Node* connect(Node* root) {
    if (root == nullptr) return nullptr;

    Node* leftMostNode = root;
    while (leftMostNode->left != nullptr) {
      Node* curNode = leftMostNode;

      while (curNode != nullptr) {
        curNode->left->next = curNode->right;

        if (curNode->next != nullptr) {
          curNode->right->next = curNode->next->left;
        }

        curNode = curNode->next;
      }

      leftMostNode = leftMostNode->left;
    }

    return root;
  }
};