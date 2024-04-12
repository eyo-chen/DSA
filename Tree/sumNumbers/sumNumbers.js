//////////////////////////////////////////////////////
// ***  Sum Root to Leaf Numbers ***
//////////////////////////////////////////////////////
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
This is my first solution I wrote by myself
It worked, but the code it's not good

Use res to store each path
So res will store multiple arrays
After recursion, use reduce to count the sum of each path

************************************************************
Time: O(n)
Space: O(h)
*/
var sumNumbers = function (root) {
  const res = [];
  let sum = 0;
  dfs(root, res, []);

  res.forEach(path => {
    sum += pathToSum(path);
  });

  return sum;
};

function dfs(root, res, tmp) {
  if (root === null) return;

  // it's the leaf
  if (root.left === null && root.right === null) {
    // add the path, and current value
    res.push([...tmp, root.val]);
    return;
  }

  tmp.push(root.val); // choose
  dfs(root.left, res, tmp); // go left
  dfs(root.right, res, tmp); // go right
  tmp.pop(); // unchoose
  return;
}

// to add the sum of each path
function pathToSum(arr) {
  const powerIndex = arr.length - 1;
  return arr.reduce((acc, cur, index) => {
    return (acc += cur * 10 ** (powerIndex - index));
  }, 0);
}

/*
The main difference is that we don't use array
Instead, use string
After finding the leaf, return the sum

Note that it's perferct to use string here because it helps us to add the val at the end of string
For example, one of the path is
             6
        4
    3
str = ""
str = "6"
str = "64"
str = "643"
Finally, easily convert str to number

************************************************************
Time: O(n)
Space: O(h)
*/
var sumNumbers = function (root) {
  return dfs(root, '');
};

function dfs(node, str) {
  // base case
  if (node === null) {
    return 0;
  }

  // find the leaf
  if (node.left === null && node.right === null) {
    // first add the value at the end of string
    str += node.val;

    // return the sum after converting to number
    return Number(str);
  }

  // go left and right
  return dfs(node.left, str + node.val) + dfs(node.right, str + node.val);
}

/*
The main difference is that we don't use string
Instead, use number

But how can we correctly add the number as sum?
For example, one of the path is
             6
        4
    3

sum = 0
sum = 0 * 10 + 6 = 6
sum = 6 * 10 + 4 = 64
sum = 64 * 10 + 3 = 643

As we can see, if we first mutiply to ten, and then add the value
We can get the correct sum of each path
*/
var sumNumbers = function (root) {
  return dfs(root, 0);
};

function dfs(node, sum) {
  // base case
  if (node === null) {
    return 0;
  }

  // find the leaf
  if (node.left === null && node.right === null) {
    // mutiply to 10, then add the value
    return sum * 10 + node.val;
  }

  return (
    dfs(node.left, sum * 10 + node.val) + dfs(node.right, sum * 10 + node.val)
  );
}
