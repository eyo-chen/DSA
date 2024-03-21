//////////////////////////////////////////////////////
// *** Binary Tree Postorder Traversal ***
//////////////////////////////////////////////////////
/*
Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.
 
Example 1:
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].

Example 2:
Input: root = [1,2]
Output: 1
 
Constraints:
The number of nodes in the tree is in the range [1, 104].
-100 <= Node.val <= 100
*/
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
/*
I wrote the solution after watching the video

The key of this problem is to understand the diameter
=> The diameter of a binary tree is the length of the longest path between any two nodes in a tree
=> This path may or may not pass through the root
=> which means the diameter can be anywhere in the tree
=> Example 1:
                              5
                3                     6
        1            2                     
    4   
=> diameter: [4,1,3,5,6]
=> Pass through the root
=> Example 2:
                              5
                3                     6
        1            2                     
    4                    6
5                            9
=> diameter: [5,4,1,3,2,6,8]
=> Not pass through the root

This is the main hard part of this problem

Now the problem may be how can we find the diameter for each node?
=> For each node, we want to know their max height of left subtree and right subtree
=> For example, in example 1, we wanna ask 5(root),
   What's your max height of left subtree and right subtree
   It's 3 and 1
   3 + 1 = diameter
   in example 2, we wanna ask 3(root),
   What's your max height of left subtree and right subtree
   It's 3 and 3
   3 + 3 = diameter
=> Above, we just ask to the node is the root of diameter
=> But in the code, we have to ask every node the same question

Now we know that we have to first find the max height of left subtree and right subtree
So that we can know the diameter
The problem becomes simpler
Just find the max height of left subtree and right subtree for each node
When we know the answer, then we check if the sum of them(new diameter) is greater than the old diameter
aka, do we need to update the diameter?

The process, res = diameter 
                              5
                     3                 6
               1            2                     
           4                    6
      11

Stack 1
Passing node: 5(root)
What's the max height of left subtree and right subtree?
I don't know, but I go left first to find the answer

Stack 2
Passing node: 3
What's the max height of left subtree and right subtree?
I don't know, but I go left first to find the answer

Stack 3
Passing node: 1
What's the max height of left subtree and right subtree?
I don't know, but I go left first to find the answer

Stack 4
Passing node: 4
What's the max height of left subtree and right subtree?
I don't know, but I go left first to find the answer

Stack 5
Passing node: 11
What's the max height of left subtree and right subtree?
I don't know, but I go left first to find the answer

Stack 6
Passing node: null
What's the max height of left subtree and right subtree?
I know the answer
I'm null, so means the height is -1

Stack 5
Passing node: 11
What's the max height of left subtree and right subtree?
I know my max height of left subtree is 1 + (-1) = 0
Go right to find the answer
=> same thing as before, it's 0
heightLeft = 0
heightRight = 0
res = Math.max(res, heightLeft + heightRight) = 0
=> Now I'm gonna return the max height of left and right
=> return 0

Stack 4
Passing node: 4
What's the max height of left subtree and right subtree?
heightLeft = 1 + 0 = 1
heightRight = 0 (same thing before, just - 1 + 1)
res = Math.max(res, heightLeft + heightRight) = 1
=> at the momenet, the diameter is 1 (11 -> 4)
=> Now I'm gonna return the max height of left and right
=> return 1

Stack 3
Passing node: 1
What's the max height of left subtree and right subtree?
heightLeft = 1 + 1 = 2
heightRight = 0 (same thing before, just - 1 + 1)
res = Math.max(res, heightLeft + heightRight) = 2
=> at the momenet, the diameter is 2 (11 -> 4 -> 1)
=> Now I'm gonna return the max height of left and right
=> return 2

Stack 2
Passing node: 3
What's the max height of left subtree and right subtree?
heightLeft = 1 + 2 = 3
heightRight = ??
=> go right to find the anser

Stack 3
Passing node: 2
What's the max height of left subtree and right subtree?
I don't know, but I go left first to find the answer

Stack 4
Passing node: null
What's the max height of left subtree and right subtree?
I know the answer
I'm null, so means the height is -1

Stack 3
Passing node: 2
What's the max height of left subtree and right subtree?
heightLeft = 1 + (-1) = 0
heightRight = ??
=> go right to find the anser

Stack 4
Passing node: 2
What's the max height of left subtree and right subtree?
heightLeft = 0
heightRight = 0
=> won't update the diameter
=> return 0

Stack 3
Passing node: 2
What's the max height of left subtree and right subtree?
heightLeft = 1 + (-1) = 0
heightRight = 1 + 0 = 1
=> res = Math.max(res, heightLeft + heightRight);
=> res = Max(2, 1 + 0) = 2
=> won't update the diameter
=> return 1

Stack 2
Passing node: 3
What's the max height of left subtree and right subtree?
heightLeft = 1 + 2 = 3
heightRight = 1 + 1 = 2
=> res = Math.max(res, heightLeft + heightRight);
=> res = Max(2, 3 + 2) = 5
=> diameter: [11 -> 4 -> 1 -> 3 -> 2 -> 6]
=> return 3

Stack 1
Passing node: 5(root)
What's the max height of left subtree and right subtree?
heightLeft = 1 + 3 = 4
heightRight = 1 + 0 = 1(skip the process)
=> res = Math.max(res, heightLeft + heightRight);
=> res = Max(5, 3 + 2) = 5
=> diameter: [11 -> 4 -> 1 -> 3 -> 2 -> 6]
=> it's same, won't update

Done

Hope the process is clear
If it's still not, go wathcing this https://www.youtube.com/watch?v=bkxqA8Rfv04&t=622s

Final two things to note
=> When hitting null, we have to return -1
=> For example, 
         5
     1
null
=> Hitting null, return -1
=> after returning, we will do this 1 + recursiveHelper(node.left);
=> 1 + (-1) = 0
=> it means at node 1, the heightLeft is 0, it's correct
=> that's the reason we need to return -1

Second things
=> We need a global res outside the recursive function
=> In the process of recursion, we keep counting the max heightLeft and heightRight
=> Also, update the res
=> In other solution, use Class to wrapper
=> Or use array to wrap the answer, and pass into recursive function

************************************************************
Time: O(n)
Space: O(h)
*/
var diameterOfBinaryTree = function (root) {
  let res = 0;

  function recursiveHelper(node) {
    // Base case
    if (node === null) {
      return -1;
    }

    // find the max height
    const heightLeft = 1 + recursiveHelper(node.left);
    const heightRight = 1 + recursiveHelper(node.right);

    // update diameter
    res = Math.max(res, heightLeft + heightRight);

    return Math.max(heightLeft, heightRight);
  }

  recursiveHelper(root);

  return res;
};
