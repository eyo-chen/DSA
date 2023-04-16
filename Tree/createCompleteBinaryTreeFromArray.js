'use strict';

// function TreeNode(val, left, right, next) {
//   this.val = val === undefined ? null : val;
//   this.left = left === undefined ? null : left;
//   this.right = right === undefined ? null : right;
// }

// function createCompleteBinaryTreeFromArray(arr) {
//   // [1,null,2,3]
//   let root = null;
//   let q = [];
//   let i = 0;
//   let t = arr[i] == null ? null : new TreeNode(arr[i]);
//   root = t;
//   q.push(root);
//   i++;
//   while (q.length && i < arr.length) {
//     let t1 = q.shift();
//     if (t1 != null) {
//       t1.left = arr[i] == null ? null : new TreeNode(arr[i]);
//       q.push(t1.left);
//       i++;
//       if (i >= arr.length) {
//         break;
//       }
//       t1.right = arr[i] == null ? null : new TreeNode(arr[i]);
//       q.push(t1.right);
//       i++;
//     }
//   }
//   return root;
// }

// const a = createCompleteBinaryTreeFromArray([
//   5,
//   4,
//   8,
//   11,
//   null,
//   13,
//   4,
//   7,
//   2,
//   null,
//   null,
//   null,
//   1,
// ]);
// var solve = function (board) {
//   const rowLength = board.length;
//   const colLength = board[0].length;
//   const seenSafe = {};

//   for (let i = 0; i < rowLength; i++) {
//     if (board[i][0] === 'O') findSafeVertices(board, i, 0, seenSafe);

//     if (board[i][colLength - 1] === 'O')
//       findSafeVertices(board, i, colLength - 1, seenSafe);
//   }

//   for (let i = 0; i < colLength; i++) {
//     if (board[0][i] === 'O') findSafeVertices(board, 0, i, seenSafe);

//     if (board[rowLength - 1][i] === 'O')
//       findSafeVertices(board, rowLength - 1, i, seenSafe);
//   }

//   for (let r = 0; r < rowLength; r++) {
//     for (let c = 0; c < colLength; c++) {
//       const key = `${r}-${c}`;
//       if (board[r][c] === 'O' && !seenSafe[key]) board[r][c] = 'X';
//     }
//   }
// };

// function findSafeVertices(board, row, col, seenSafe) {
//   const stack = [[row, col]];
//   let index = 0;

//   while (stack.length > 0) {
//     const [r, c] = stack.pop();
//     const key = `${r}-${c}`;
//     index++;

//     if (
//       r < 0 ||
//       r >= board.length ||
//       c < 0 ||
//       c >= board[0].length ||
//       board[r][c] === 'X' ||
//       seenSafe[key]
//     )
//       continue;

//     seenSafe[key] = true;

//     stack.push([r + 1, c]);
//     stack.push([r - 1, c]);
//     stack.push([r, c + 1]);
//     stack.push([r, c - 1]);
//   }
// }

// function deeplyFreeze(obj) {
//   if (!obj || typeof obj !== 'object') return;

//   for (const [key, val] of Object.entries(obj)) {
//     if (val && typeof val === 'object') {
//       deeplyFreeze(obj[key]);
//     }
//   }

//   return Object.freeze(obj);
// }
// const obj2 = {
//   internal: {
//     a: null,
//   },
// };

// const obj = {
//   obj1: {
//     a: 'a',
//   },
// };
// deeplyFreeze(obj);
// function cloneDeep(data) {
//   if (!data) return data;
//   if (typeof data !== 'object') return data;

//   if (Array.isArray(data)) {
//     const arr = [];

//     for (let i = 0; i < data.length; i++) {
//       if (typeof data[i] === 'object') {
//         arr.push(cloneDeep(data[i]));
//       } else {
//         arr.push(data[i]);
//       }
//     }

//     return arr;
//   }

//   const obj = {};
//   for (const key of Object.keys(data)) {
//     if (typeof data[key] === 'object') {
//       obj[key] = cloneDeep(data[key]);
//     } else {
//       obj[key] = data[key];
//     }
//   }

//   return obj;
// }
// const a = { a: 'a', c: { cc: 'c' }, d: [1, 2, 3, { aaa: 'aaa' }] };
// const aa = cloneDeep(a);

// a.d[3].aaa = 'bbb';

var findWords = function (board, words) {
  const coordinate = buildCoordinate(board);
  const res = [];

  for (const word of words) {
    const firstChar = word[0];

    if (coordinate[firstChar]) {
      const coordinateArr = coordinate[firstChar];

      for (const [row, col] of coordinateArr) {
        const key = `${row}-${col}`;
        if (DFS(board, row, col, word, 1, { [key]: true })) {
          res.push(word);
          break;
        }
      }
    }
  }

  return res;
};

function DFS(board, r, c, word, index, seen) {
  if (index === word.length) {
    return true;
  }

  if (checkValid(r + 1, c, board, seen, word, index)) {
    if (DFS(board, r + 1, c, word, index + 1, seen)) {
      return true;
    }
    seen[`${r + 1}-${c}`] = false;
  }
  if (checkValid(r - 1, c, board, seen, word, index)) {
    if (DFS(board, r - 1, c, word, index + 1, seen)) {
      return true;
    }
    seen[`${r + 1}-${c}`] = false;
  }
  if (checkValid(r, c + 1, board, seen, word, index)) {
    if (DFS(board, r, c + 1, word, index + 1, seen)) {
      return true;
    }
    seen[`${r + 1}-${c}`] = false;
  }
  if (checkValid(r, c - 1, board, seen, word, index)) {
    if (DFS(board, r, c - 1, word, index + 1, seen)) {
      return true;
    }
    seen[`${r + 1}-${c}`] = false;
  }

  return false;
}

function checkValid(row, col, board, seen, word, index) {
  const key = `${row}-${col}`;
  if (
    row < 0 ||
    row >= board.length ||
    col < 0 ||
    col >= board[0].length ||
    seen[key] ||
    board[row][col] !== word[index]
  )
    return false;

  seen[key] = true;

  return true;
}

function buildCoordinate(board) {
  const coordinate = {};

  for (let r = 0; r < board.length; r++) {
    for (let c = 0; c < board[0].length; c++) {
      const char = board[r][c];
      if (coordinate[char]) {
        coordinate[char].push([r, c]);
      } else {
        coordinate[char] = [[r, c]];
      }
    }
  }

  return coordinate;
}
const board123 = [
  ['a', 'b', 'c'],
  ['a', 'e', 'd'],
  ['a', 'f', 'g'],
];
const words123 = ['eaabcdgfa'];
// console.log(findWords(board123, words123));

function myObjectCreate(proto) {
  if (typeof proto !== 'object' || Array.isArray(proto)) {
    return new Error("Can't pass object");
  }

  function func() {}
  func.prototype = proto;

  return new func();
}
const a = { a: 'a' };
const nnn = myObjectCreate(123);
const aaa = Object.create(a);
console.log(nnn);
