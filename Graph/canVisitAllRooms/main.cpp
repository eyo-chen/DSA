#include <queue>
#include <stack>
#include <unordered_set>

using namespace std;

// using BFS(queue)
class Solution {
 public:
  bool canVisitAllRooms(vector<vector<int>>& rooms) {
    queue<int> q;
    unordered_set<int> isRoomVisited = {0};
    q.push(0);

    while (!q.empty()) {
      int key = q.front();
      q.pop();

      for (int curRoom : rooms[key]) {
        // check if the room is already visited
        // if (isRoomVisited.contains(curRoom)) continue; // C++20
        // returns the number of elements matching specific key
        if (isRoomVisited.count(curRoom) != 0) continue;

        // if the room is not visited, push it to the queue
        q.push(curRoom);

        // mark the room as visited
        isRoomVisited.insert(curRoom);

        // if all rooms are visited, return true
        // this is the short circuit to avoid unnecessary iteration
        if (isRoomVisited.size() == rooms.size()) return true;
      }
    }

    return isRoomVisited.size() == rooms.size();
  }
};

// using DFS(stack)
class Solution {
 public:
  bool canVisitAllRooms(vector<vector<int>>& rooms) {
    stack<int> s;
    unordered_set<int> isRoomVisited = {0};
    s.push(0);

    while (!s.empty()) {
      int key = s.top();
      s.pop();

      for (int curRoom : rooms[key]) {
        // check if the room is already visited
        // if (isRoomVisited.contains(curRoom)) continue; // C++20
        // returns the number of elements matching specific key
        if (isRoomVisited.count(curRoom) != 0) continue;

        // if the room is not visited, push it to the stack
        s.push(curRoom);

        // mark the room as visited
        isRoomVisited.insert(curRoom);

        // if all rooms are visited, return true
        // this is the short circuit to avoid unnecessary iteration
        if (isRoomVisited.size() == rooms.size()) return true;
      }
    }

    return isRoomVisited.size() == rooms.size();
  }
};