//////////////////////////////////////////////////////
// *** Keys and Rooms ***
//////////////////////////////////////////////////////
/*
There are n rooms labeled from 0 to n - 1 and all the rooms are locked except for room 0. Your goal is to visit all the rooms. However, you cannot enter a locked room without having its key.

When you visit a room, you may find a set of distinct keys in it. Each key has a number on it, denoting which room it unlocks, and you can take all of them with you to unlock the other rooms.

Given an array rooms where rooms[i] is the set of keys that you can obtain if you visited room i, return true if you can visit all the rooms, or false otherwise.

Example 1:
Input: rooms = [[1],[2],[3],[]]
Output: true
Explanation: 
We visit room 0 and pick up key 1.
We then visit room 1 and pick up key 2.
We then visit room 2 and pick up key 3.
We then visit room 3.
Since we were able to visit every room, we return true.

Example 2:
Input: rooms = [[1,3],[3,0,1],[2],[0]]
Output: false
Explanation: We can not enter room number 2 since the only key that unlocks it is in that room.
 
Constraints:
n == rooms.length
2 <= n <= 1000
0 <= rooms[i].length <= 1000
1 <= sum(rooms[i].length) <= 3000
0 <= rooms[i][j] < n
All the values of rooms[i] are unique.
*/
/**
 * @param {number[][]} rooms
 * @return {boolean}
 */
/*
BFS

This problem is very straightforward

This problem can be thought as graph problems
Each room is just a node(vertex) (each small array in the big array)
Each element(key) in small array is edge

For example,
[[1,3],[3,0,1],[2],[0]]
There are four nodes (four rooms)
First node(room) has two edges(key) [1, 3]
Second node(room) has three edges(key) [3, 0, 1]
Third node(room) has one edge(key) [2]
Fourth node(room) has one edge(key) [0]

Just see to implementation of code

************************************************************
Time: O(|V| + |E|)
Space: O(|V|)
*/
var canVisitAllRooms = function (rooms) {
  const seen = { 0: true };
  const queue = [0];

  while (queue.length > 0) {
    const key = queue.shift();
    const room = rooms[key];

    for (let i = 0; i < room.length; i++) {
      const keyIndex = room[i];
      // only add the room haven't been seen
      if (!seen[keyIndex]) {
        queue.push(keyIndex);
        seen[keyIndex] = true;
      }
    }
  }

  // return true if the length of seen rooms is eaul to the length original rooms
  return Object.keys(seen).length === rooms.length;
};

/*
DFS
*/
var canVisitAllRooms = function (rooms) {
  const seen = { 0: true };
  const stack = [0];

  while (stack.length > 0) {
    const key = stack.pop();
    const room = rooms[key];

    for (let i = 0; i < room.length; i++) {
      const keyIndex = room[i];
      if (!seen[keyIndex]) {
        stack.push(keyIndex);
        seen[keyIndex] = true;
      }
    }
  }

  return Object.keys(seen).length === rooms.length;
};
