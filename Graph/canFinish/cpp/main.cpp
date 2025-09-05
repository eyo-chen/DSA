#include <unordered_map>
#include <vector>

using namespace std;

enum State { stateUnVisited, stateVisiting, stateVisited };

class Solution {
 public:
  bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {
    // Initialize the adjacency list
    unordered_map<int, vector<int>> adj;
    for (const vector<int>& p : prerequisites) {
      int cur = p[0], pre = p[1];
      adj[pre].push_back(cur);
    }

    // Initialize the state of each node
    vector<State> states(numCourses, stateUnVisited);

    // Check for cycles in each component of the graph
    for (int i = 0; i < numCourses; i++) {
      if (states[i] != stateUnVisited) continue;
      if (hasCycle(adj, states, i)) return false;
    }

    return true;
  }

  bool hasCycle(unordered_map<int, vector<int>>& adj, vector<State>& states,
                int node) {
    // Check if the node is being visited
    // If it is, then there is a cycle
    if (states[node] == stateVisiting) return true;

    // Check if the node has been visited
    // If it has, then there is no cycle
    if (states[node] == stateVisited) return false;

    // Mark the node as visiting
    states[node] = stateVisiting;

    // Visit all adjacent nodes
    for (int neighbor : adj[node]) {
      if (hasCycle(adj, states, neighbor)) return true;
    }

    // Mark the node as visited
    states[node] = stateVisited;

    return false;
  }
};