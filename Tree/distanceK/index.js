//////////////////////////////////////////////////////
// ***  All Nodes Distance K in Binary Tree ***
//////////////////////////////////////////////////////
/*
Given the root of a binary tree, the value of a target node target, and an integer k, return an array of the values of all nodes that have a distance k from the target node.

You can return the answer in any order.

Example 1:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
Output: [7,4,1]
Explanation: The nodes that are a distance 2 from the target node (with value 5) have values 7, 4, and 1.

Example 2:
Input: root = [1], target = 1, k = 3
Output: []
 
Constraints:
The number of nodes in the tree is in the range [1, 500].
0 <= Node.val <= 500
All the values Node.val are unique.
target is the value of one of the nodes in the tree.
0 <= k <= 1000
*/
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
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
 * @param {TreeNode} target
 * @param {number} k
 * @return {number[]}
 */
/*
Using parent node + BFS
If we conceptually think this question as graph problem
it would be much easier

                                       1
                    2                                   2
            3              44              4                       3
       1      33       34    21        12    40

If target node 2
First level
=> [44, 3, 1]

Second level
=> [12, 40, 2]

Third level
=> [3, 44]

Fourth level
=> [1, 33, 34, 21]

This is exactly the pattern of DFS
So now the only problem is that we have to build the parent node by ourself
But it's not that hard

So the process is 
1. Build the parent hashTable
2. Do the BFS

One thing to note
We have use another map to store the seen node
For example, first level we store [4, 3, 1]
Later, 4 would also having a parent node 2
And 2 has been seen before
It's duplicate
So we have to use map to store the seen node

https://www.youtube.com/watch?v=nPtARJ2cYrg

************************************************************
Time: O(n)
Space: O(n)
*/
var distanceK = function (root, target, k) {
  const res = [];
  const queue = [target];
  const map = new Map();
  const hashTable = createHashTable(root);
  let level = 0;

  while (level <= k && queue.length > 0) {
    const len = queue.length;

    for (let i = 0; i < len; i++) {
      const node = queue.shift();

      map.set(node, true);

      if (node === null) continue;

      if (level === k) {
        res.push(node.val);
      }

      if (!map.get(node.left)) queue.push(node.left);

      if (!map.get(node.right)) queue.push(node.right);

      if (!map.get(hashTable.get(node))) queue.push(hashTable.get(node));
    }

    level++;
  }

  return res;
};

function createHashTable(root) {
  const map = new Map();
  const queue = [[root, null]];

  while (queue.length > 0) {
    const [node, parent] = queue.shift();

    map.set(node, parent);

    if (node.left !== null) {
      queue.push([node.left, node]);
    }

    if (node.right !== null) {
      queue.push([node.right, node]);
    }
  }

  return map;
}

/*
This is recursive approach
This is back to tree problem, but it's harder

The idea is
1. Keep finding the target node
2. Once finding, invoke the findKNode function
   => this function just use passed node as root, and search it's own subtree
   to find the k level node
   => For example,
                                          1
                    2                                   23
            3              44              4                       3
       1      33       34    21        12    40
    => target node = 2, k = 2
    => Once finding target node 2, doing the DFS to find [1, 33, 34, 21]
    => After that, return 1
    => This 1 means one level distance
    => Look at the example above, after finding 2
    => Return 1 to root node 1, so root node 1 knows that the distance between target node and I is 1

3. If hitting null, just return false
   => It means NOT found the target node

4. Go left and right to find level distance and target node

5. If distance of left or right is eqaul to k
   => Find the res node, push it to the res array
   => For example,
                                          1
                    2                                   23
            3              44              4                       3
       11      33       34    21        12    40
    => target node = 33, k = 2
    => After finding target node 11, return 1 to node 3
    => So node 3 knows that the distance between me and target node is 1
    => Then return 1 + 1 to node 2
    => So node 2 knows that the distance between me and target node is 2
    => Hi, I found the correct level node, so I push 2 to the res array

6. If left is not false
   => It means I find the target node on my left subtree
   => Then I need to do two things
   => 1) Go to my right subtree to find the correct level node
   => 2) return left distance + 1
   => For example,
                                          1
                    2                                   23
            3              44              4                       31
       11      33       34    21        12    40
    => target node = 3, k = 2
    => After finding target node 3, return 1 to it's parent node 2
    => So node 2 knows that the distance between me and target node is 1
    => Does two things
    => 1) Go to my right subtree to find the correct level node
    => Find 44 at correct level

    => 2) return left distance + 1
    => I return 1 + 1 to root node 1
    => So node 1 knows that the distance between me and target node is 2
    => Find the correct level node

7. So does right


Note these two lines of code
findKNode(root.right, left + 1, k, res);
findKNode(root.left, right + 1, k, res);
For example, 
                                          1
                    2                                   23
            3              44              4                       31
       11      33       34    21        12    40
target node = 3, k = 2
=> node 2 knows that I find the target node at my left subtree, and level distance is 1
=> So I go to right to find correct level target node
=> So I directly go to 44, and left + 1
=> + 1 is the distance between 2 and 44

Go to leetcode to see solution
************************************************************
Time: O(n)
Space: O(n)
*/
var distanceK = function (root, target, k) {
  const res = [];

  recursiveHelper(root, k, target, res);

  return res;
};

function recursiveHelper(root, k, target, res) {
  if (root === null) return false;

  if (root === target) {
    findKNode(root, 0, k, res);
    return 1;
  }

  const left = recursiveHelper(root.left, k, target, res);
  const right = recursiveHelper(root.right, k, target, res);

  if (left === k || right === k) {
    res.push(root.val);
    return false;
  }

  if (left) {
    findKNode(root.right, left + 1, k, res);
    return left + 1;
  }

  if (right) {
    findKNode(root.left, right + 1, k, res);
    return right + 1;
  }

  return false;
}

function findKNode(root, level, k, res) {
  if (root === null || level > k) return false;

  if (level === k) {
    res.push(root.val);
    return;
  }

  findKNode(root.left, level + 1, k, res);
  findKNode(root.right, level + 1, k, res);

  return;
}
