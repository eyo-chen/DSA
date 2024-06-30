//////////////////////////////////////////////////////
// *** Nearest Exit from Entrance in Maze ***
//////////////////////////////////////////////////////
/*
You are given an m x n matrix maze (0-indexed) with empty cells (represented as '.') and walls (represented as '+'). You are also given the entrance of the maze, where entrance = [entrancerow, entrancecol] denotes the row and column of the cell you are initially standing at.

In one step, you can move one cell up, down, left, or right. You cannot step into a cell with a wall, and you cannot step outside the maze. Your goal is to find the nearest exit from the entrance. An exit is defined as an empty cell that is at the border of the maze. The entrance does not count as an exit.

Return the number of steps in the shortest path from the entrance to the nearest exit, or -1 if no such path exists.

Example 1:
Input: maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance = [1,2]
Output: 1
Explanation: There are 3 exits in this maze at [1,0], [0,2], and [2,3].
Initially, you are at the entrance cell [1,2].
- You can reach [1,0] by moving 2 steps left.
- You can reach [0,2] by moving 1 step up.
It is impossible to reach [2,3] from the entrance.
Thus, the nearest exit is [0,2], which is 1 step away.

Example 2:
Input: maze = [["+","+","+"],[".",".","."],["+","+","+"]], entrance = [1,0]
Output: 2
Explanation: There is 1 exit in this maze at [1,2].
[1,0] does not count as an exit since it is the entrance cell.
Initially, you are at the entrance cell [1,0].
- You can reach [1,2] by moving 2 steps right.
Thus, the nearest exit is [1,2], which is 2 steps away.

Example 3:
Input: maze = [[".","+"]], entrance = [0,0]
Output: -1
Explanation: There are no exits in this maze.
*/
/**
 * @param {character[][]} maze
 * @param {number[]} entrance
 * @return {number}
 */
/*
This problem should be pretty easy if modeling as graph problem
Because the question asks us shortest path
We use BFS, which is exactly perfect in this case
So we just use normal BFS
And using path as counter to count the current path
Once finding the exit, immediately return the path

Note that we have to use a for-loop inside the while-loop
Because we go level by level
We can only go to next level after we finish searching in this current level

The most hard part is the logic of isExit
row !== eRow || col !== eCol 
=> if both condition are false, it means it's the starting position
=> We want to immediately return false because starting position can't be as exit

&&
    (row === 0 ||
      row === arr.length - 1 ||
      col === 0 ||
      col === arr[0].length - 1)
Then we keep testing other conditions
If one of them is true, it means it's a exit
just return true, and say we found the exit
************************************************************
r = row, c = col
Time: O(r * c)
Space: O(r * c)
*/
var nearestExit = function (maze, entrance) {
  const queue = [entrance];
  const seen = {};
  let path = 0;

  while (queue.length > 0) {
    const len = queue.length;

    for (let i = 0; i < len; i++) {
      const [row, col] = queue.shift();
      const key = `${row}-${col}`;

      // only continue if it's valid vertices
      if (!isValid(maze, row, col, seen, key)) continue;

      // test if it's exit
      if (isExit(maze, row, col, entrance[0], entrance[1])) {
        return path;
      }

      seen[key] = true;

      /*
      Just directly go top, bottom, right and left
      isValid will help us to validate
      */
      queue.push([row + 1, col]);
      queue.push([row - 1, col]);
      queue.push([row, col + 1]);
      queue.push([row, col - 1]);
    }

    path++;
  }

  return -1;
};

function isValid(arr, row, col, seen, key) {
  return (
    row >= 0 &&
    row <= arr.length - 1 &&
    col >= 0 &&
    col <= arr[0].length - 1 &&
    arr[row][col] !== '+' &&
    !seen[key]
  );
}

function isExit(arr, row, col, eRow, eCol) {
  return (
    (row !== eRow || col !== eCol) &&
    (row === 0 ||
      row === arr.length - 1 ||
      col === 0 ||
      col === arr[0].length - 1)
  );
}
