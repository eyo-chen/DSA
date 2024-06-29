#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<int>> getAncestors(int n, const vector<vector<int>>& edges) {
    vector<vector<int>> adj(n);
    vector<vector<int>> result(n);

    // convert edges to adjacency list
    for (const vector<int>& edge : edges) {
      int from = edge[0];
      int to = edge[1];
      adj[from].push_back(to);
    }

    // loop through each node to find ancestors, treat each node as root node
    for (int i = 0; i < n; ++i) {
      vector<bool> visited(n, false);
      findAncestors(adj, result, visited, i, i);
    }

    return result;
  }

 private:
  void findAncestors(const vector<vector<int>>& adj,
                     vector<vector<int>>& result, vector<bool>& visited,
                     int cur, int parent) {
    visited[cur] = true;

    for (int edge : adj[cur]) {
      if (visited[edge]) continue;

      result[edge].push_back(parent);
      findAncestors(adj, result, visited, edge, parent);
    }
  }
};
