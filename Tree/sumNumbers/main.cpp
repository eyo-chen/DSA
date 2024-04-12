#include <cmath>
#include <queue>
#include <string>

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

// Using DFS(recursion) with string path
class Solution {
 public:
  int sumNumbers(TreeNode* root) { return helper(root, ""); }

  int helper(TreeNode* root, string path) {
    if (root == nullptr) return 0;

    // add current node value to the path
    // e.g. path = "123" and root->val = 4
    // then curPath = "1234"
    string curPath = path + to_string(root->val);

    // explore
    int left = helper(root->left, curPath);
    int right = helper(root->right, curPath);

    // if left and right are both 0, then we are at the leaf node
    if (left == 0 && right == 0) {
      return sum(curPath);
    }

    int total = 0;
    if (left != 0) total += left;
    if (right != 0) total += right;

    return total;

    /*
    The abvoe code can be simplified as follows:
    return left + right;
    */
  }

  int sum(string str) {
    int result = 0;
    int mul = str.length() - 1;

    for (char s : str) {
      // convert char to int
      // e.g. '1' - '0' = 1
      // By subtracting the ASCII value of '0' (which is 48) from the ASCII
      // value of char, you get the integer value.
      result += (s - '0') * pow(10, mul);
      mul--;
    }

    return result;
  }
};

// Using DFS(recursion) with int sum
class Solution {
 public:
  int sumNumbers(TreeNode* root) { return helper(root, 0); }

  int helper(TreeNode* root, int sum) {
    if (root == nullptr) return 0;

    int curSum = sum * 10 + root->val;
    int left = helper(root->left, curSum);
    int right = helper(root->right, curSum);

    if (left == 0 && right == 0) {
      return curSum;
    }

    int total = 0;
    if (left != 0) total += left;
    if (right != 0) total += right;

    return total;
    /*
    The abvoe code can be simplified as follows:
    return left + right;
    */
  }
};

// Using BFS(queue) with string path
class Solution {
 public:
  int sumNumbers(TreeNode* root) {
    if (root == nullptr) return 0;

    struct nodeWithPath {
      TreeNode* node;
      string path;
    };

    queue<nodeWithPath> q;
    int result = 0;
    q.push({root, to_string(root->val)});

    while (!q.empty()) {
      nodeWithPath n = q.front();
      q.pop();

      if (n.node->left == nullptr && n.node->right == nullptr) {
        result += sum(n.path);
      }

      if (n.node->left != nullptr)
        q.push({n.node->left, n.path + to_string(n.node->left->val)});
      if (n.node->right != nullptr)
        q.push({n.node->right, n.path + to_string(n.node->right->val)});
    }

    return result;
  }

  int sum(string str) {
    int result = 0;
    int mul = str.length() - 1;

    for (char s : str) {
      result += (s - '0') * pow(10, mul);
      mul--;
    }

    return result;
  }
};

// Using BFS(queue) with int sum
class Solution {
 public:
  int sumNumbers(TreeNode* root) {
    if (root == nullptr) return 0;

    struct nodeWithPath {
      TreeNode* node;
      int sum;
    };

    queue<nodeWithPath> q;
    int result = 0;
    q.push({root, root->val});

    while (!q.empty()) {
      nodeWithPath n = q.front();
      q.pop();

      if (n.node->left == nullptr && n.node->right == nullptr) {
        result += n.sum;
      }

      if (n.node->left != nullptr)
        q.push({n.node->left, n.sum * 10 + n.node->left->val});
      if (n.node->right != nullptr)
        q.push({n.node->right, n.sum * 10 + n.node->right->val});
    }

    return result;
  }
};