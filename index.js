// 'use strict';
function TreeNode(val, left, right, next) {
  this.val = val === undefined ? null : val;
  this.left = left === undefined ? null : left;
  this.right = right === undefined ? null : right;
}
var isSubtree = function (root, subRoot) {
  if (root === null || root.left === null || root.right === null) {
    return false;
  }

  if (recursiveHelper(root, subRoot)) {
    return true;
  }

  return isSubtree(root.right, subRoot) || isSubtree(root.left, subRoot);
};

function recursiveHelper(root, subRoot) {
  if (root === null && subRoot === null) {
    return true;
  }

  if (root === null || subRoot === null) {
    return false;
  }

  if (root.val !== subRoot.val) {
    return false;
  }

  return (
    recursiveHelper(root.right, subRoot.right) &&
    recursiveHelper(root.left, subRoot.left)
  );
}

function createCompleteBinaryTreeFromArray(arr) {
  // [1,null,2,3]
  let root = null;
  let q = [];
  let i = 0;
  let t = arr[i] == null ? null : new TreeNode(arr[i]);
  root = t;
  q.push(root);
  i++;
  while (q.length && i < arr.length) {
    let t1 = q.shift();
    if (t1 != null) {
      t1.left = arr[i] == null ? null : new TreeNode(arr[i]);
      q.push(t1.left);
      i++;
      if (i >= arr.length) {
        break;
      }
      t1.right = arr[i] == null ? null : new TreeNode(arr[i]);
      q.push(t1.right);
      i++;
    }
  }
  return root;
}
const a = createCompleteBinaryTreeFromArray([1]);
const b = createCompleteBinaryTreeFromArray([1]);
console.log(isSubtree(a, b));
