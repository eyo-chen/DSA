//////////////////////////////////////////////////////
// *** Clone Graph ***
//////////////////////////////////////////////////////
/*
Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

class Node {
    public int val;
    public List<Node> neighbors;
}
 

Test case format:

For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.

An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.


Example 1:
Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
Explanation: There are 4 nodes in the graph.
1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).

Example 2:
Input: adjList = [[]]
Output: [[]]
Explanation: Note that the input contains one empty list. The graph consists of only one node with val = 1 and it does not have any neighbors.

Example 3:
Input: adjList = []
Output: []
Explanation: This an empty graph, it does not have any nodes.

Constraints:

The number of nodes in the graph is in the range [0, 100].
1 <= Node.val <= 100
Node.val is unique for each node.
There are no repeated edges and no self-loops in the graph.
The Graph is connected and all nodes can be visited starting from the given node.
*/
/**
 * Definition for a Node.
 * function Node(val, neighbors) {
 *    this.val = val === undefined ? 0 : val;
 *    this.neighbors = neighbors === undefined ? [] : neighbors;
 * };
 */
/**
 * @param {Node} node
 * @return {Node}
 */
/*
This problem should be fairly easy after understanding the logic
We use BFS, queue and hashTable
There are two reasons to use hashTable
1. we don't want to traverse(search) the vertex has been seen
2. Because while-loop and for-loop are both loop through the input graph
   when we try to connect the edge between two verteices of clone graph
   we can use hashTable to access the verterx
   In short, key -> input graph vertex, value -> clone graph vertext


The idea is sth like this, 
1. Do the DFS on input graph
2. Get the current vertex
3. Loop through all it's neighbors(vertices)
4. If neighbor vertex is NOT in the hashTable, which means we haven't seen this
   also means we haven't create this clone vertext
   So we wanna
   (1) create the clone vertex
   (2) add it to the hashTable
   (3) add it to the queue(this is just part of the process of BFS)
5. Then, it's time to connect the edge between two vertices
   Note that we always loop through input graph
   Which means the only way we can access clone graph, we have to use hashTable
   So we do
   give me the current clone vertex (hashTable[vertices.val])
   also give me the current neighbor clone vertex (hashTable[neighbor.val])
   current clone vertext.neighbors.push(current neighbor clone vertex)
   make the edge

Note that we use vertex.val as key in hashTable
Because the val won't be duplicate
We also can use entire vertex object as key
But we have to use map
************************************************************
Time: O(|V| + |E|)
Space: O(|V|)
*/
var cloneGraph = function (node) {
  if (node === null) return node;

  // put start(first) vertex in the queue
  const queue = [node];
  const hashTable = {};

  // set key(input vertex), value(clone vertex)
  hashTable[node.val] = new Node(node.val);

  // DFS
  while (queue.length > 0) {
    // get the current input vertex
    const vertices = queue.shift();

    // get the current clone vertex
    const cloneVertices = hashTable[vertices.val];

    // loop through all the neighbors(edges) of input vertex
    // again, note that we're looping through input vertex
    for (const neighbor of vertices.neighbors) {
      // If neighbor hasn't been seen or neighbor clone vertex hasn't been create
      if (!hashTable[neighbor.val]) {
        // create the neighbor clone vertex
        const node = new Node(neighbor.val);

        // add it to the hashTable
        hashTable[node.val] = node;

        // add it to the queue
        queue.push(neighbor);
      }

      // make the connections (edge)
      cloneVertices.neighbors.push(hashTable[neighbor.val]);
    }
  }

  return hashTable[node.val];
};
