//////////////////////////////////////////////////////
// *** Populating Next Right Pointers in Each Node ***
//////////////////////////////////////////////////////
/*
You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.

Example 1:
Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.

Example 2:
Input: root = []
Output: []
 
Constraints:
The number of nodes in the tree is in the range [0, 212 - 1].
-1000 <= Node.val <= 1000
 

Follow-up:
You may only use constant extra space.
The recursive approach is fine. You may assume implicit stack space does not count as extra space for this problem.
*/
/**
 * // Definition for a Node.
 * function Node(val, left, right, next) {
 *    this.val = val === undefined ? null : val;
 *    this.left = left === undefined ? null : left;
 *    this.right = right === undefined ? null : right;
 *    this.next = next === undefined ? null : next;
 * };
 */

/**
 * @param {Node} root
 * @return {Node}
 */

/*
This is the solution I first time writing
It's very intutive, but the code is not clean
It's using BFS

The idea is quite simple
Because it's perfect binary tree
We just count what's the total node in current level
For example,
first level
=> has 2 ** 0 node

Second level
=> has 2 ** 1 node

Third level
=> has 2 ** 2 node

so on and so forth

If we traverse all the node of current level
We just increment the powerdIndex, and re-set the curNodeCount
And we need a preNode
So that we can
preNode.next = curNode

The idea is very simple

************************************************************
Time: O(n)
Space: O(h)
*/
var connect = function (root) {
  const queue = [root];
  let powerIndex = 0;
  let curNodeCount = 1;
  let preNode = null;

  while (queue.length > 0) {
    const curNode = queue.shift();

    // for the root is null
    if (curNode === null) continue;

    // this is for the root node (first level)
    // no need to wire next pointer at the first level
    if (curNodeCount !== 1) preNode.next = curNode;

    // hitting the end of current level, update indexes
    if (curNodeCount === 2 ** powerIndex) {
      // this is unncessary because all the next pointer initialize to null
      curNode.next = null;
      powerIndex++;
      curNodeCount = 1;
    } else {
      // keep updting curNodeCount
      curNodeCount++;
    }

    // add left subtree
    if (curNode.left !== null) queue.push(curNode.left);

    // add right subtree
    if (curNode.right !== null) queue.push(curNode.right);

    // update preNode
    preNode = curNode;
  }

  return root;
};

/*
BFS recursive
the logic is simple too, and the code is cleaner

For each state, we do two things
If root has left subtree, we do
root.left.next = root.right;
          4
    3         1
=> 3.next = 1

If root has right subtree, we do
root.right.next = root.next.left;
          10
     3          5
  1     4    8    11
=> 4.next = 8

However, there's one thing to note when wiring next ptr on right subtree
What it curret node is 5 (example above)
there's no next value
5.next = null
so 
root.next.left = null.left -> error
So we need another conditional check
root.next !== null
Only do the wiring if current node has next value

Note that we don't need to care about the 11 and 5 need to point to null
Because all the next pointer initialize to null

Final thing to note
This approach will first wire left subtree
Then wire right subtree

In DFS, it's going level by level

This will cause huge difference when solving second variant
************************************************************
Time: O(n)
Space: O(h)
*/
var connect = function (root) {
  // Base case
  if (root === null) return null;

  // wire next ptr of left subtree
  if (root.left !== null) {
    root.left.next = root.right;
  }

  // wire next ptr of right subtree
  if (root.right !== null && root.next !== null) {
    root.right.next = root.next.left;
  }

  root.left = connect(root.left);
  root.right = connect(root.right);

  return root;
};

/*
This is also DFS, but using iterative approach
All the logic is same
*/
var connect = function (root) {
  const stack = [root];

  while (stack.length > 0) {
    const node = stack.pop();

    if (node === null) continue;

    if (node.left !== null) {
      node.left.next = node.right;
    }

    if (node.right !== null && node.next !== null) {
      node.right.next = node.next.left;
    }

    stack.push(node.right);
    stack.push(node.left);
  }

  return root;
};

/*
This is using queue
All the logic is same

But it's wiring level by level
*/
var connect = function (root) {
  const queue = [root];

  while (queue.length > 0) {
    const node = queue.shift();

    if (node === null) continue;

    if (node.left !== null) node.left.next = node.right;

    if (node.right !== null && node.next !== null)
      node.right.next = node.next.left;

    queue.push(node.left);
    queue.push(node.right);
  }

  return root;
};

/*
This is O(1) space solution

                    10
          3                     5
   1           4          8          11
9    15    12    13   18   19    20     21

The idea is 
At 10
Wire 3 -> 5
Go left

At 3
Wire 1 -> 4 -> 8 -> 11
Go left

At 1
Wire 9 -> 15 -> 12 -> 13 -> 18 -> 19 -> 20 ->  21
Done

The inner while-loop helps us all the way to wire the next ptr

And along with the process, we need some conditional check to prevent the bug
*/
var connect = function (root) {
  let curNode = root;

  while (curNode !== null) {
    let node = curNode;

    /*
    if node.left is null, we don't need wire any next ptr
    just update the curNode, and skip the further process
    while (curNode !== null) will break out the while-loop
    
    Note that we're wiring the ptr of child node
    This is checking the situation when node is 9
    There's no child node to wire up
    */
    if (node.left === null) {
      curNode = curNode.left;
      continue;
    }

    // wire all the child node
    while (node !== null) {
      // wire left -> right
      // for example, wire 1 -> 4, 9 -> 15 ... (example above)
      node.left.next = node.right;

      /*
      wire right -> next left
      need to check if there's next to let use wire
      When we're at 5, 11 and 21, we don't need to wire it's next ptr of right subtree

      for example, wire 15 -> 12, 4 -> 8 ...
      */
      if (node.next !== null) node.right.next = node.next.left;

      // updating node (go to right node)
      node = node.next;
    }

    // updating node(go to next level)
    curNode = curNode.left;
  }

  return root;
};
