//////////////////////////////////////////////////////
// *** Lowest Common Ancestor of a Binary Tree ***
//////////////////////////////////////////////////////
/*
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Example 1:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.

Example 2:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.

Example 3:
Input: root = [1,2], p = 1, q = 2
Output: 1
 
Constraints:
The number of nodes in the tree is in the range [2, 105].
-109 <= Node.val <= 109
All Node.val are unique.
p != q
p and q will exist in the tree.
*/
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
/*
This is the solution I wrote by myself
The idea is not thar hard, but it's not easy neither

The idea of this solution is
For each state(node), there are only three situations
1. I'm eqaul to one of the target node(p & q)
=> I may be the Lowest Common Ancestor
=> Not sure 100%, but maybe

2. the target node may be on my left subtree
=> Go left to find

3. the target node may be on my right subtree
=> Go right to find

These three are the general situations in each node
However, only knowing this cannot solve the problem
We need to further think how do we handle when we get the answer
For example,
How do I handle when the curret node is equal to the target node
How do I handle when one of my subtree having the target node
How do I handle when none of my subtree having the target node

                              5
                3                     6
        1            2         7            8
    4                                  11      19
                                   15              20

1. If none of my subtree having the target node
=> Just return false
=> It means, hey my caller, I can't find any target node in my subtree
=> For example, p = 11, q = 7, and the current node is 3
=> Hi, I'm node 3, I've exlplored my right and left subtree, but I can't find the target node
=> So I just return false

2. If I'm eqaul to one the target node
=> This situation is kind of interesting
=> There are two different situations
=> For example, p = 8, q = 15
=> (1) current node = 8, it's obvious that it's the Lowest Common Ancestor, 
       return myself(node)
   (2) current node = 15, I'm one of the target node, but not the Lowest Common Ancestor
       It's okay, I also return myself(node)
=> As we can see, how I handle this situation is
=> Store true or false in a variable(cur)
const cur = root === p || root === q;
=> If one of the subtree having target node, then I can 100% make sure I'm Lowest Common Ancestor
=> For example, p = 8, q = 15, current node = 8
=> I just hold the cur to true, and keep exploring the sub tree
=> If one of the subtree having the target node
=> I can 100% make sure I'm Lowest Common Ancestor, return myself(node)
  if (cur && (goLeft || goRight)) {
    return root;
  }

3. I'm not target node, but my both left and right subtree having the target node
=> I can 100% make sure I'm Lowest Common Ancestor, return myself(node)
=> For example, p = 7, q = 15, current node = 6
=> After exploring right and left subtree, 6 will find the it's the Lowest Common Ancestor
  if (goLeft && goRight) {
    return root;
  }

These are basically are the main logic of this solution, but final detail to note
In the whole process of recursion,
we either return boolean or node(object)
When do we return boolean?
=> (1) return false, hitting the base case
=> (2) return true, I'm one of the target node, but not sure if it's Lowest Common Ancestor

When do we return node(object)?
=> When we're 100% sure we find the Lowest Common Ancestor
=> Once find the Lowest Common Ancestor, we're gonna bubble up the node all the way up
=> That's why this code
  if (typeof goLeft === 'object') {
    return goLeft;
  }

  if (typeof goRight === 'object') {
    return goRight;
  }
=> After exploring both left and right, if one of them is node(object)
=> I know that I've found the Lowest Common Ancestor, just return it

  if (cur && (goLeft || goRight)) {
    return root;
  }
=> I'm the target node, so cur = true
=> One of the right and left is also target node
=> So it returns true
=> Then I'm 100% the  Lowest Common Ancestor

return cur || goLeft || goRight
=> Only return false, if none of these are true
=> For example, p = 11, q = 19, current node = 3
=> It means I'm not the target node, and I've finished exploring the sub tree
=> And none of them are the target node, so I return false

************************************************************
Time: O(n)
Space: O(h)
*/
var lowestCommonAncestor = function (root, p, q) {
  if (root === null) {
    return false;
  }

  const cur = root === p || root === q;
  const goLeft = lowestCommonAncestor(root.left, p, q); // go left
  const goRight = lowestCommonAncestor(root.right, p, q); // go right

  // Once find the Lowest Common Ancestor, we're gonna bubble up the node all the way up
  if (typeof goLeft === 'object') {
    return goLeft;
  }

  // Once find the Lowest Common Ancestor, we're gonna bubble up the node all the way up
  if (typeof goRight === 'object') {
    return goRight;
  }

  // I'm the target node, and one of my subtree is also target node
  // I'm 100% the Lowest Common Ancestor
  if (cur && (goLeft || goRight)) {
    return root;
  }

  // I'm not the target node, but two of my subtree are target node
  // I'm 100% the Lowest Common Ancestor
  if (goLeft && goRight) {
    return root;
  }

  return cur || goLeft || goRight;
};

/*
The above solution is okay, but kind of verbose
We can improve the logic, so does make sure code shorter

In the second situation, If I'm eqaul to one the target node
Acutally, we can directly return the node at this point
Why?
                              5
                3                     6
        1            2         7            8
    4                                  11      19
                                   15              20
Situation 1
For example, p = 1, q = 6
=> When hitting 1, immediately return itself, return 1(node)
=> When hitting 6, immediately return itself, return 6(node)
=> So when 5(root) hit this line of code, it will return itself
  if (goLeft && goRight) {
    return root;
  }

Situation 2
For example, p = 7, q = 19
=> 5(root) will first explore it's left sub tree
=> None of node in the left subtree is equal to p or q
=> So false will bubble up to this line of code const goLeft = lowestCommonAncestor(root.left, p, q);
=> goLeft = false
=> It means the Lowest Common Ancestor must be in the left subtree
=> keep going right
=> when hitting 7, return itself, bubble up to the state of current node is 6
=> At the sate of current node is 6, const goLeft = lowestCommonAncestor(root.left, p, q);
=> goLeft = 7
=> keep going right
=> hitting 19, return itself, bubble up to the state of current node is 6
=> At the sate of current node is 6, const goRight = lowestCommonAncestor(root.right, p, q);
=> goRight = 19
  if (goLeft && goRight) {
    return root;
  }
=> Hitting this line of code, return 6

Situation 3
For example, p = 8, q = 19
=> 5(root) will first explore it's left sub tree
=> None of node in the left subtree is equal to p or q
=> So false will bubble up to this line of code const goLeft = lowestCommonAncestor(root.left, p, q);
=> goLeft = false
=> It means the Lowest Common Ancestor must be in the left subtree
=> Hitting 8, immediately return itself, return itself, bubble up iself to the root, and return the answer
=> It doesn't matter the other target node is where
=> Because, this is key point
=> If I'm equal to one of the target node, there are only two situations
=> I'm the Lowest Common Ancestor(this case), so return myself
=> One of my parent is Lowest Common Ancestor(the case above), also need to return myself
=> Eventually, my parent(Lowest Common Ancestor) will hit this line of code
  if (goLeft && goRight) {
    return root;
  }

These logic make the code shorter
But the main logic is as same as previous one
Go watching this video https://www.youtube.com/watch?v=13m9ZCB8gjw
*/
var lowestCommonAncestor = function (root, p, q) {
  if (root === null) {
    return false;
  }

  if (root === p || root === q) {
    return root;
  }

  const goLeft = lowestCommonAncestor(root.left, p, q);
  const goRight = lowestCommonAncestor(root.right, p, q);

  if (goLeft && goRight) {
    return root;
  }

  return goLeft || goRight;
};
