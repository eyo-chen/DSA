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
