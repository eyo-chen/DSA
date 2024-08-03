#include <queue>
#include <unordered_map>
#include <unordered_set>

using namespace std;

// Definition for a binary tree node.
struct TreeNode {
  int val;
  TreeNode* left;
  TreeNode* right;
  TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
 public:
  vector<int> distanceK(TreeNode* root, TreeNode* target, int k) {
    unordered_map<int, vector<int>> adj;
    genAdj(root, adj);

    queue<int> q({target->val});
    unordered_set<int> seen({target->val});
    vector<int> res;

    while (q.size() > 0) {
      // if k == 0, return the current layer's node
      if (k == 0) {
        while (!q.empty()) {
          res.push_back(q.front());
          q.pop();
        }
        return res;
      }
      k--;

      // BFS
      int size = q.size();
      for (int i = 0; i < size; i++) {
        int node = q.front();
        q.pop();

        for (const int nei : adj[node]) {
          if (seen.count(nei) != 0) continue;
          q.push(nei);
          seen.insert(nei);
        }
      }
    }

    return res;
  }

  // genAdj generates adjacency list for the tree
  void genAdj(TreeNode* root, unordered_map<int, vector<int>>& adj) {
    if (root == nullptr) return;

    if (root->left != nullptr) {
      // add child node to parent node's adjacency list
      adj[root->val].push_back(root->left->val);

      // add parent node to child node's adjacency list
      adj[root->left->val].push_back(root->val);
      genAdj(root->left, adj);
    }

    if (root->right != nullptr) {
      // add child node to parent node's adjacency list
      adj[root->val].push_back(root->right->val);

      // add parent node to child node's adjacency list
      adj[root->right->val].push_back(root->val);
      genAdj(root->right, adj);
    }
  }
};
