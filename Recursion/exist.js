//////////////////////////////////////////////////////
// *** Word Search ***
//////////////////////////////////////////////////////
/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example 1:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true

Example 2:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true

Example 3:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false
 
Constraints:
m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board and word consists of only lowercase and uppercase English letters.
*/
/*
I came up the solution by myself

This idea is actualy not that hard, the main point is to handle the edga case
When dealing with the problem of matrix(2D array), be careful when the row and col is out of the bound
like, row = 0, col = 0, row = matrix.length, col = matrix[0].length
Especially the last two one
Truly be careful about these edge cases along with the process

The idea of solution is 
1. traverse all the cell of input board
2. Find the first matching character (with the first character of input word)
3. If found one, we know that we can start the recursion on this character(row, col)
4. If not found one, we can just return false after nested for-loop
   => If we can't even find the first matching character, there's no way finding correct word search

For the recursive helper, it's quite easy
We just need to go four direction to keep the process of recursion
BUT, we need to check several conditions before do the further recursion
1. row is within the bound
2. col is within the bound
3. character has not been searched (it's not equal to ".")
4. character is equal to the next character of word

If match those conditions, then we can do the recursion

Choose
=> go top, go right, go bottom, go left

Constraint
=> row and col have to be within the bound
=> cannot reuse the same character in the board (the path cannot be duplicate)
=> the character has to be matching

Goal
=> find all the character of input word

************************************************************
r = the length row of input board, c = the length of col of input board, w = the length of input word
Time complexity: O(r * c * (3 ^ w))
=> At worst case, we need to do the recursion on each character in the board
=> r * n -> traverse all the cell of board
=> 4 ^ w -> for each recursive function, the worst branching factor is 3 (because three dimension), 
   and the deepest recursive tree is w, have to at least find w length of word
=> why branching factor is always 3?
=> Because we can't never go back, once choosing going up, we won't choose to go down ever again


Space complexity: O(w)
=> we'll have charPath to keep tracking the searched character
=> the longest length of this array is O(w)
=> also has the call stack, the deepest recursive tree is also w 
*/
/**
 * @param {character[][]} board
 * @param {string} word
 * @return {boolean}
 */
var exist = function (board, word) {
  const charPath = [];

  for (let row = 0; row < board.length; row++) {
    for (let col = 0; col < board[0].length; col++) {
      // find the first matching character
      if (word[0] === board[row][col]) {
        // choose the character
        charPath.push(board[row][col]);

        // mark the cell as searched
        board[row][col] = '.';

        // explore (index = 1, means find the first character)
        if (recursiveHelper(row, col, board, 1, word, charPath)) return true;

        // unChoose, also re-mark the the cell to the original character
        board[row][col] = charPath.pop();
      }
    }
  }
  return false;
};

function recursiveHelper(row, col, board, index, word, charPath) {
  // base case
  if (index === word.length) return true;

  // go up
  if (
    row >= 1 &&
    board[row - 1][col] !== '.' &&
    board[row - 1][col] === word[index]
  ) {
    // choose the character
    charPath.push(board[row - 1][col]);

    // mark the cell as searched
    board[row - 1][col] = '.';

    // explore
    if (recursiveHelper(row - 1, col, board, index + 1, word, charPath))
      return true;

    // unChoose, also re-mark the the cell to the original character
    board[row - 1][col] = charPath.pop();
  }

  // go right
  if (
    col < board[0].length - 1 &&
    board[row][col + 1] !== '.' &&
    board[row][col + 1] === word[index]
  ) {
    charPath.push(board[row][col + 1]);
    board[row][col + 1] = '.';

    if (recursiveHelper(row, col + 1, board, index + 1, word, charPath))
      return true;

    board[row][col + 1] = charPath.pop();
  }

  // go down
  if (
    row < board.length - 1 &&
    board[row + 1][col] !== '.' &&
    board[row + 1][col] === word[index]
  ) {
    charPath.push(board[row + 1][col]);
    board[row + 1][col] = '.';

    if (recursiveHelper(row + 1, col, board, index + 1, word, charPath))
      return true;

    board[row + 1][col] = charPath.pop();
  }

  // go left
  if (
    col >= 1 &&
    board[row][col - 1] !== '.' &&
    board[row][col - 1] === word[index]
  ) {
    charPath.push(board[row][col - 1]);
    board[row][col - 1] = '.';

    if (recursiveHelper(row, col - 1, board, index + 1, word, charPath))
      return true;

    board[row][col - 1] = charPath.pop();
  }

  return false;
}

/*
This solution has same time and space compexlity
but the code is shorted, and non-verbose
*/
var exist1 = function (board, word) {
  const charPath = [];

  for (let row = 0; row < board.length; row++) {
    for (let col = 0; col < board[0].length; col++) {
      /*
      Instead of finding first match character,
      we just simply try every each character in the board
      If the first character is not matching, it's still find
      because the second condition of recursion will immediately return false anyway
      */
      if (recursiveHelper1(row, col, board, 0, word, charPath)) return true;
    }
  }
  return false;
};

function recursiveHelper1(row, col, board, index, word, charPath) {
  // base case
  if (index === word.length) return true;

  // base case 2
  /*
  return false, if
  1. row and col are out of the bound
  2. the character is not matching
  3. the character has been searched
  */
  if (
    row < 0 ||
    col < 0 ||
    row >= board.length ||
    col >= board[0].length ||
    board[row][col] !== word[index] || // the character is not matching
    board[row][col] === '.' // the character has been searched
  )
    return false;

  // choose
  charPath.push(board[row][col]);
  board[row][col] = '.';

  // explore all four dimension
  /*
  It's okay row and col is out of the bound here
  because the second base case will check all the condition together
  */
  if (
    recursiveHelper1(row + 1, col, board, index + 1, word, charPath) ||
    recursiveHelper1(row - 1, col, board, index + 1, word, charPath) ||
    recursiveHelper1(row, col + 1, board, index + 1, word, charPath) ||
    recursiveHelper1(row, col - 1, board, index + 1, word, charPath)
  )
    return true;

  // unchoose
  board[row][col] = charPath.pop();

  return false;
}

console.log(
  exist(
    [
      ['A', 'B', 'C', 'E'],
      ['S', 'F', 'C', 'S'],
      ['A', 'D', 'E', 'E'],
    ],
    'SEE'
  )
);
console.log(
  exist1(
    [
      ['A', 'B', 'C', 'E'],
      ['S', 'F', 'C', 'S'],
      ['A', 'D', 'E', 'E'],
    ],
    'SEE'
  )
);
