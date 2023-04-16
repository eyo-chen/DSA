//////////////////////////////////////////////////////
// *** Course Schedule II ***
//////////////////////////////////////////////////////
/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

Example 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].

Example 2:
Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

Example 3:
Input: numCourses = 1, prerequisites = []
Output: [0]
 
Constraints:
1 <= numCourses <= 2000
0 <= prerequisites.length <= numCourses * (numCourses - 1)
prerequisites[i].length == 2
0 <= ai, bi < numCourses
ai != bi
All the pairs [ai, bi] are distinct.
*/
/**
 * @param {number} numCourses
 * @param {number[][]} prerequisites
 * @return {number[]}
 */
/*
This one is very similar to the final solution of canFinish
Just go to there to see the review

************************************************************
Time Complexity: O(V + E)O(V+E) where VV represents the number of vertices and EE represents the number of edges. We pop each node exactly once from the zero in-degree queue and that gives us VV. Also, for each vertex, we iterate over its adjacency list and in totality, we iterate over all the edges in the graph which gives us EE. Hence, O(V + E)O(V+E)

Space Complexity: O(V + E)O(V+E). We use an intermediate queue data structure to keep all the nodes with 0 in-degree. In the worst case, there won't be any prerequisite relationship and the queue will contain all the vertices initially since all of them will have 0 in-degree. That gives us O(V)O(V). Additionally, we also use the adjacency list to represent our graph initially. The space occupied is defined by the number of edges because for each node as the key, we have all its adjacent nodes in the form of a list as the value. Hence, O(E)O(E). So, the overall space complexity is O(V + E)O(V+E).
*/
var findOrder = function (numCourses, prerequisites) {
  const graphList = {};
  const inDegree = {};
  const queue = [];
  const res = [];

  buildTwoLists(inDegree, graphList, prerequisites, numCourses);

  for (const course of Object.keys(inDegree)) {
    if (inDegree[course] === 0) {
      queue.push(course);
    }
  }

  while (queue.length > 0) {
    const course = queue.shift();

    for (const nextCourse of graphList[course]) {
      inDegree[nextCourse]--;

      if (inDegree[nextCourse] === 0) {
        queue.push(nextCourse);
      }
    }

    // difference, keep track the course we're taking
    res.push(Number(course));
  }

  // if we can't take all the course, return [], else return res
  return res.length === numCourses ? res : [];
};

function buildTwoLists(inDegree, graphList, prerequisites, numCourses) {
  const courseArr = new Array(numCourses).fill(false);

  for (const course of prerequisites) {
    const [next, cur] = course;

    if (!graphList[cur]) {
      graphList[cur] = [next];
      inDegree[cur] = 0;
      courseArr[cur] = true;
    } else {
      graphList[cur].push(next);
    }

    if (!graphList[next]) {
      graphList[next] = [];
      inDegree[next] = 1;
      courseArr[next] = true;
    } else {
      inDegree[next]++;
    }
  }

  for (let i = 0; i < courseArr.length; i++) {
    if (!courseArr[i]) {
      graphList[i] = [];
      inDegree[i] = 0;
    }
  }
}

/*
This is second solution, which is kind of similar to the first solution of canFinish
Also, I come up with this solution after watching this video https://www.youtube.com/watch?v=Akt3glAwyfY&t=640s

The main idea is 
1. build the graphList (include non-connected vertex)
2. iterate through all the vertex 
=> the order here doesn't matter
=> so we just iterate 0 ~ numCourses - 1
3. if the vertex has been seen
=> if(seen[i])
=> then we just skip it
=> Look at the DFS function, we only add vertex into seen when it pushes into the res
4. If not been seen, then do the DFS on this vertex
5. After iteration, the res storing the reverse order
=> so we use the idea of stack to build the correct output
*/
var findOrder = function (numCourses, prerequisites) {
  const graphList = {};
  const seen = {};
  const res = [];
  const output = [];

  buildTwoLists(graphList, prerequisites, numCourses);

  for (let i = 0; i < numCourses; i++) {
    if (seen[i]) continue;
    if (!DFS(graphList, seen, i, res)) {
      return [];
    }
  }

  while (res.length > 0) {
    output.push(res.pop());
  }

  return output;
};

function DFS(graphList, seen, vertex, res, hashTable = {}) {
  /*
  Note that hashTable keep tracks the vertex we're currently searching
  So if hashTable[vertex] is true, it means we've searched this vertex before
  it's a cycle, return false
  */
  if (hashTable[vertex]) {
    return false;
  }

  /*
  seen is different
  the vertex in the seen means it has been added into the res, so just skip it
  also means there's no cycle
  */
  if (seen[vertex]) {
    return true;
  }

  // mark it as searching
  hashTable[vertex] = true;

  // loop through all the edge
  for (const edge of graphList[vertex]) {
    // again, if it's seen before, just skip it
    if (seen[edge]) continue;

    // keep DFS, if it's false, we immediately return false
    if (!DFS(graphList, seen, edge, res, hashTable)) {
      return false;
    }
  }

  // add to the res, and marked as seen
  seen[vertex] = true;
  res.push(vertex);

  // means there's no cycle
  return true;
}

function buildTwoLists(graphList, prerequisites, numCourses) {
  const courseArr = new Array(numCourses).fill(false);

  for (const course of prerequisites) {
    const [next, cur] = course;

    if (!graphList[cur]) {
      graphList[cur] = [next];
      courseArr[cur] = true;
    } else {
      graphList[cur].push(next);
    }

    if (!graphList[next]) {
      graphList[next] = [];
      courseArr[next] = true;
    }
  }

  for (let i = 0; i < courseArr.length; i++) {
    if (!courseArr[i]) {
      graphList[i] = [];
    }
  }
}
