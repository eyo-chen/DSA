#include <queue>
#include <unordered_map>

using namespace std;

// Definition for a Node.
class Node {
 public:
  int val;
  vector<Node*> neighbors;
  Node() {
    val = 0;
    neighbors = vector<Node*>();
  }
  Node(int _val) {
    val = _val;
    neighbors = vector<Node*>();
  }
  Node(int _val, vector<Node*> _neighbors) {
    val = _val;
    neighbors = _neighbors;
  }
};

// using BFS
class Solution {
 public:
  Node* cloneGraph(Node* node) {
    if (node == nullptr) return nullptr;

    // create the new root node
    Node* rootNode = new Node(node->val);

    // create a queue to traverse the graph
    // note that the queue contains the original nodes
    queue<Node*> q({node});

    // create a hashmap to store the new nodes
    // it helps to reference the new nodes(pointer) to avoid creating duplicate
    // nodes
    unordered_map<int, Node*> hashMap = {{node->val, rootNode}};

    while (q.size() > 0) {
      Node* n = q.front();
      q.pop();

      // get the new node from the hashmap
      // note that the queue contains the original nodes
      // hashMap contains the new copy nodes
      // so we can get the new copy node from the hashMap
      Node* copyNode = hashMap[n->val];

      // iterate over the neighbors of the original node
      for (Node* ne : n->neighbors) {
        Node* neighborNode;

        // if the neighbor node is not in the hashmap, which means it is not
        // visited yet then
        // 1. create a new node
        // 2. add the new node to the hashmap
        // 3. add the new node to the queue
        if (hashMap.find(ne->val) == hashMap.end()) {
          neighborNode = new Node(ne->val);
          hashMap[ne->val] = neighborNode;
          q.push(ne);
        } else {
          neighborNode = hashMap[ne->val];
        }

        // add the neighbor node to the neighbors of the copy node
        copyNode->neighbors.push_back(neighborNode);
      }
    }

    return rootNode;
  }
};

// using BFS (using pointer as hashmap key)
class Solution {
 public:
  Node* cloneGraph(Node* node) {
    if (node == nullptr) return nullptr;

    // create the new root node
    Node* rootNode = new Node(node->val);

    // create a queue to traverse the graph
    // note that the queue contains the original nodes
    queue<Node*> q({node});

    // create a hashmap to store the new nodes
    // it helps to reference the new nodes(pointer) to avoid creating duplicate
    // nodes
    // note that the key is the pointer of the original node
    unordered_map<Node*, Node*> hashMap = {{node, rootNode}};

    while (q.size() > 0) {
      Node* n = q.front();
      q.pop();

      // get the new node from the hashmap
      // note that the queue contains the original nodes
      // hashMap contains the new copy nodes
      // so we can get the new copy node from the hashMap
      Node* copyNode = hashMap[n];

      // iterate over the neighbors of the original node
      for (Node* ne : n->neighbors) {
        Node* neighborNode;

        // if the neighbor node is not in the hashmap, which means it is not
        // visited yet then
        // 1. create a new node
        // 2. add the new node to the hashmap
        // 3. add the new node to the queue
        if (hashMap.find(ne) == hashMap.end()) {
          neighborNode = new Node(ne->val);
          hashMap[ne] = neighborNode;
          q.push(ne);
        } else {
          neighborNode = hashMap[ne];
        }

        // add the neighbor node to the neighbors of the copy node
        copyNode->neighbors.push_back(neighborNode);
      }
    }

    return rootNode;
  }
};

// using DFS
class Solution {
 public:
  Node* cloneGraph(Node* node) {
    if (node == nullptr) return nullptr;

    // create the new root node
    Node* rootNode = new Node(node->val);

    // create a hashmap to store the new nodes
    // it helps to reference the new nodes(pointer) to avoid creating duplicate
    // nodes
    unordered_map<int, Node*> hashMap = {{node->val, rootNode}};

    DFS(node, hashMap);

    return rootNode;
  }

  void DFS(Node* node, unordered_map<int, Node*>& hashMap) {
    if (node == nullptr) return;

    // get the new node from the hashmap
    // note that the parameter node is the original node
    // hashMap contains the new copy nodes
    // so we can get the new copy node from the hashMap
    Node* copyNode = hashMap[node->val];

    // iterate over the neighbors of the original node
    for (Node* ne : node->neighbors) {
      Node* neighborNode;

      // if the neighbor node is not in the hashmap, which means it is not
      // visited yet then
      // 1. create a new node
      // 2. add the new node to the hashmap
      // 3. call the dfs function recursively with the neighbor node(original
      // node)
      if (hashMap.find(ne->val) == hashMap.end()) {
        neighborNode = new Node(ne->val);
        hashMap[ne->val] = neighborNode;
        DFS(ne, hashMap);
      } else {
        neighborNode = hashMap[ne->val];
      }

      // add the neighbor node to the neighbors of the copy node
      copyNode->neighbors.push_back(neighborNode);
    }
  }
};