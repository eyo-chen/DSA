# Problem Explanation

The most hard part of this problem is ***How to model the problem as graph problem***.<br>
If we look at the simple example,<br>
```
rooms = [[1], [2], [3], []]
```
What does this tell us?<br>
It tells us two things,<br>
1. There are 4 rooms in total.
2. Room 0 has key to room 1, Room 1 has key to room 2, Room 2 has key to room 3, Room 3 has no key.<br>

The second one is the most important one.<br>
It tells us the ***relationship between rooms***.<br>
which is the ***edges between nodes*** in graph.<br>

Let's model this as graph,<br>
```
0 -> 1 
     |
     v
3 <- 2 
```
Now, the problem is to find whether we can visit all rooms starting from room 0.<br>
This is a simple graph traversal problem.<br>

Let's model another example,<br>
```
rooms = [[1,3], [3,0,1], [2], [0]]
```
```
0 <-> 1    2
|    /
v   /
 3
```
adjacency list representation of graph,<br>
```
0 -> 1, 3
1 -> 3, 0, 1
2 -> 2
3 -> 0
```
As we can see, we can never visit 2 from 0.<br>

After modeling the problem as graph, we can use any graph traversal algorithm to solve this problem.<br>
All we need to do is to traverse the graph starting from room 0, and check whether we can visit all rooms.<br>
We can use both DFS and BFS.<br>

The core logic is as following,<br>
1. Put starting key(0) into the data structure.
2. While the data structure is not empty,
    1. Pop the key from the data structure.
    2. Check the adjacent rooms that the key can open.
    3. For each adjacent room
       1. If the room is not visited, put the key into the data structure, and mark the room as visited.
       2. If the room is already visited, do nothing.
3. If all rooms are visited, return True. Otherwise, return False.<br>

# Complexity Analysis
## Time complexity O(N + E)
- N : Number of rooms
- E : Number of keys
- The while loop runs at most N times, and for each room, the inner for-loop runs at most E times in total.
- Or we can say, the time complexity is O(|V| + |E|).
- Look at the adjacency list representation of the graph above, then we can break down the time complexity as follow
  - 0 can open 1 and 3, 1 can open 3 and 0, 2 can open 2, 3 can open 0.
  - 0room + 0keys(1,3) + 1room + 1keys(3,0,1) + 2room + 2keys(2) + 3room + 3keys(0)
  - (0rooms + 1rooms + 2rooms + 3rooms) + (0keys + 1keys + 2keys + 3keys)
  - (all the rooms) + (all the keys)
  - O(N + E)
  - Note that we treat room as vertex and key as edge in the graph.

## Space complexity O(N)
- N : Number of rooms
- We use two data structures, `visited` and `stack`, which takes O(N) space.
- In the worst case, the queue can contain all rooms, so the space used by the queue is O(N).
- The unordered set stores the rooms that have been visited. In the worst case, it will contain all rooms, so the space used by the set is O(N).