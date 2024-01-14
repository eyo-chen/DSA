/*
At the beginning, i try to solve this problem by myself
But the solution i wrote only can pass when n = 4
which means it's failed
So I went watching video, and re-wrote the solution by myself again, and it worked

This problem is all about the core idea of backtracking
There are two keys to solve this question
1. Understand how to check if it's valid placement, especially the diagonal case
   => How to know if it's diagonally attack?
   => If current queens is (2,1)(row, col), and it's 4x4 board
   => Diagonal placement: (3,0), (1,2), (0,3), (1,0), (3,2)
   => How to find the pattern to find all diagonal placement?
   => There are two diagonal lines
   => One is positive diagonal lines(from left to right), 
      the other one is negative diagonal lines(from right to left)
   => positive diagonal lines: curRow + curCol = row + col
      curRow + curCol = 3
      (3,0), (1,2), (0,3) => row + col = 3
   => negative diagonal lines: curRow - curCol = row - col 
      curRow - curCol = 1
      (1,0), (3,2) => row - col = 1
   => Use this formula/pattern to find if it's diagonal case
2. For each row, there's always one queen
   => In other words, once placing one queen in one row, we immediately go to next row
   => For example, if it's 4x4 board
   => Once we first choose the (0,0) to place queen, we don't need to try further possibilities to place the next queen on (0,1), (0,2), (0,3), NO, this is immediate and obvious invalid, all we need is go to next row, and try (1,0), (1,1), (1,2), (1,3)
   => If we use recursive tree to think of this process, each state/call stack frame represent each rows, and the branching factor of each call stack frame represent the each columns
                                        top
                    (0,0)                    (0,1)           (0,2)           (0,3)
        (1,0)   (1,1)   (1,1)   (1,2)  
(2,0)(2,1)(2,2)(2,3)
   => It's abvious we won't have perfect tree like this, which means the branching factor won't always be 4, because we have the constraint, can't be same row and col, also diagonal case
   => But this is just overview of the process

Hope now the solution is more clear
Choose: try all columns on each rows
Constaint: can't be on same row, colum and diagonal case
Goal: Have placed queen on each row(aka, row === n)
      When row = n, it means we have tried row from 0 ~ n - 1, which means we have placed one queen in each row

One thing to note that
At first, I misunderstand the return type
I return 2D array, outer array represent row, and inner array represent col
But it's wrong
The problem ask us to return 1D array
And for each cell of 1D array, use string to represent

But the solution I wrote, I first use 1D array to represent each row
When i push to the tmp, i convert it to string
The reason i did this is because it's easy to mark to "Q" amd backtrack

Final note
It actually have a more clever way to construct the position array
In my solution, i try to store an array in position array
For example, the position array would be like this
[[row,col], [row,col], [row,col]]
But actually we don't need to do this
We can just use 1D array
[col, col, col, col]
The value of each cell represent the col
And index represent the row
1row 2row 3row 4row
[col, col, col, col]
************************************************************
n = the input n
Time complexity: O(n! * n)
=> Imagine the recursive tree, each call stack fram represent the each row, which means the deepeset height of recursive tree is n
=> and each branching factor represent each column
=> so we can roughly upper bound to O(n ^ n)
=> But this is not enough
=> As we can see, first row we can have n choices
=> But second row, we only have at least n - 2 choices
=> In general, it could have even n - 3 choices
=> For example, if we first put queen at (0,0)
=> Then the next possible placement in second row, is (1,2), (1,3)
=> Because (1,0) and (1,1) is invalid
=> But if we first put queen at (0,1)
=> Then the next possible placement in second row, is (1,3)
=> Because (1,0), (1,1) and (1,2) is invalid
=> As we could see, we can very at least lower bound to O(n!)
=> for each call stack frame, we have 
    tmp.push(curRow.join(''));
    checkValidPlacement()
    also base case have res.push([...tmp]);
=> all of them are O(n) work

Space complexity: O(n ^ 2)
=> the deepest height of recursive tree is n
=> for each call stack frame, we use array to represent row
=> const curRow = new Array(n).fill('.');
*/
/**
 * @param {number} n
 * @return {string[][]}
 */
var solveNQueens = function (n) {
  const res = [];

  recursiveHelper(n, 0, [], [], res);

  return res;
};

// pos array represent the previously placed queens
function recursiveHelper(n, row, pos, tmp, res) {
  /*
  Base case (Goal)
  All rows have been placed (0 ~ (n - 1))
  */
  if (row === n) {
    res.push([...tmp]);
    return;
  }

  /*
  Use array to represent each row
  Easy to mark "Q" and backtrack to "."
  */
  const curRow = new Array(n).fill('.');

  // For each row, try all possibilities of column(branching factor)
  for (let i = 0; i < n; i++) {
    /*
    Constraint
    Only keep recursing further if it's valid placement(row, col, diagonal)
    */
    if (checkValidPlacement(row, i, pos)) {
      // Choose
      curRow[i] = 'Q';
      pos.push([row, i]);
      tmp.push(curRow.join(''));

      // Explore
      recursiveHelper(n, row + 1, pos, tmp, res);

      // Unchoose
      curRow[i] = '.';
      pos.pop();
      tmp.pop();
    }
  }

  return;
}

function checkValidPlacement(row, col, pos) {
  /*
  Try all position array
  Each element in position array represent the previously placed queens 
  */
  for (let i = 0; i < pos.length; i++) {
    const [curRow, curCol] = pos[i];

    if (
      curRow === row || // row
      curCol === col || // col
      curRow + curCol === row + col || // diagonal
      curRow - curCol === row - col // diagonal
    )
      return false;
  }

  return true;
}


/*
This is second solution
It's optimize space complexity
because we don't need to keep tracking each row (new Array(n))
But we need to do O(n ^ 2) work when hitting base case
Because of generating the final output array

************************************************************
n = the input n
Time complexity: O(n! * (n ^ 2))
=> Because now generateBoard have (n ^ 2) work

Space complexity: O(n)
*/
var solveNQueens1 = function (n) {
  const res = [];

  recursiveHelper1(n, 0, [], res);

  return res;
};

function recursiveHelper1(n, row, pos, res) {
  if (n === row) {
    res.push(generateBoard(pos));
    return;
  }

  for (let i = 0; i < n; i++) {
    if (checkValidPlacement1(row, i, pos)) {
      pos.push(i);
      recursiveHelper1(n, row + 1, pos, res);
      pos.pop();
    }
  }

  return;
}

// use placed position array to construct the returned board array
function generateBoard(pos) {
  const board = [];

  for (let i = 0; i < pos.length; i++) {
    let row = '';
    const col = pos[i];

    for (let j = 0; j < pos.length; j++) {
      if (j === col) row += 'Q';
      else row += '.';
    }

    board.push(row);
  }

  return board;
}

function checkValidPlacement1(row, col, pos) {
  for (let i = 0; i < pos.length; i++) {
    if (
      row === i ||
      col === pos[i] ||
      row + col === i + pos[i] ||
      row - col === i - pos[i]
    )
      return false;
  }

  return true;
}

console.log(solveNQueens(5));
console.log(solveNQueens1(5));
