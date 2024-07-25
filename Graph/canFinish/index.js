//////////////////////////////////////////////////////
// *** Course Schedule ***
//////////////////////////////////////////////////////
/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.

Example 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take. 
To take course 1 you should have finished course 0. So it is possible.

Example 2:
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take. 
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
 
Constraints:
1 <= numCourses <= 105
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
All the pairs prerequisites[i] are unique.
*/
/**
 * @param {number} numCourses
 * @param {number[][]} prerequisites
 * @return {boolean}
 */
/*
The solution using the idea of
deteching cycle in a graph

First,
build the graph list
because the problem only give us prerequisites
we have to convert it to the graphList
 key <-> value
 pre      cur
It means pre -> cur
also means we have to finish pre, then i can finish cur
In other words, we have to touch pre, then touch cur

Second,
Iterate through all the vertex
If the vertex marked as "ed", it means it has been searched
Which means we can guarantee the path starting from this vertex has no cycle
So we don't need to any search on this vertex

If it's not makred as "ed", we have to seach from this vertex(DFS)

If DFS returns true, it means finding a cycle, we immediately return false


Note that the reason we didn't care about numCourses
The final solution does use the numCourses
Go to see that to understand why we use numCourses in that solution
The main reason we don't care about this is 
Imagine the case, 
numCourses = 5
prerequisites = [[1,4],[2,4],[3,1],[3,2]]
If numCourses = 5, all the course should be 0 ~ 4
Look at the edge in the prerequisites, there's no 0 course
Which means 0 is not connected with every one
Which means 0 is always can be taken, so we don't consider this vertext
************************************************************
Time: O(|V| + |E|)
Space: O(|V|)
*/
var canFinish = function (numCourses, prerequisites) {
  const visitTable = {};
  const graphList = {};

  buildGraphList(prerequisites, graphList);

  for (const vertex of Object.keys(graphList)) {
    if (visitTable[vertex] === 'ed') continue;

    if (findCycle(vertex, graphList, visitTable)) return false;
  }

  return true;
};

/*
It's recursive function
Base case
if vertex is marked as "ing", it means it's the vertex we've searched in this particular path, it's cycle, return true

Recursive case
Mark it as "ing", which means this vertex is currently searching
Loop through all it's edges
if the edge has been searched, then we skip it
if has not been searched, then we call the function on this vertex again

After for-loop, if it has not been returned true, it means this path from this vertex has no cycle
We marked it as "ed", which also means the path from this vertex has no cycle, and we've searched already
and return false, which means no cycle
*/
function findCycle(vertex, graphList, visitTable) {
  // this vertex is on the path we're searching right now, it's cycle
  if (visitTable[vertex] === 'ing') return true;

  // searching on the path from this vertex
  visitTable[vertex] = 'ing';

  // DFS on it's edge
  for (const neighbor of graphList[vertex]) {
    if (
      visitTable[neighbor] !== 'ed' &&
      // keep same process on the neighbor
      findCycle(neighbor, graphList, visitTable)
    )
      return true;
  }

  // finish searching, marked as searched
  visitTable[vertex] = 'ed';
  return false;
}

function buildGraphList(arr, graphList) {
  for (let i = 0; i < arr.length; i++) {
    const [cur, pre] = arr[i];
    if (!graphList[pre]) graphList[pre] = [cur];
    else graphList[pre].push(cur);

    if (!graphList[cur]) graphList[cur] = [];
  }
}

/*
This is second solution, using the idea of topological sort
*/
var canFinish = function (numCourses, prerequisites) {
  if (prerequisites.length === 0) return true;

  const graphList = {};
  const inDegree = {};
  const queue = [];

  // first build the graphList
  buildGraphList(prerequisites, graphList, inDegree);

  // push the vertex has no indegree
  for (const key in inDegree) {
    if (inDegree[key] === 0) queue.push(key);
  }

  // immediately return false if there's no zero indegree
  // which also means every course depends on others
  if (queue.length === 0) return false;

  while (queue.length > 0) {
    const course = queue.shift();

    // loop though all neighbor courses, and minus the inDegree
    for (const nextCourse of graphList[course]) {
      inDegree[nextCourse]--;

      // if it's zero degree, then push on the qeuee
      if (inDegree[nextCourse] === 0) queue.push(nextCourse);
    }
  }

  // check if there's course haven't been searched
  for (const key in inDegree) {
    if (inDegree[key] !== 0) return false;
  }
  return true;
};

function buildGraphList(prerequisites, graphList, inDegree) {
  for (let i = 0; i < prerequisites.length; i++) {
    const [cur, pre] = prerequisites[i];

    if (!graphList[pre]) {
      graphList[pre] = [cur];
      inDegree[pre] = 0;
    } else {
      graphList[pre].push(cur);
    }

    if (!graphList[cur]) {
      graphList[cur] = [];
      inDegree[cur] = 1;
    } else inDegree[cur]++;
  }
}

/*
This is third solution, using the idea of topological sort
This is one is very similar to the second one
But I think this is more accurate
because it does use the numCourses to check for the answer

There are two main difference
1. uses courseCount
=> if we take one course, we increment courseCount by 1
=> at the end of while-loop, we just need to check if courseCount === numCourses
=> In this case, we don't need  if (queue.length === 0) return false;
=> Aslo, we don't need   if (prerequisites.length === 0) return true;
=> If we take all the course, it will just return true in the end

2. buildTwoLists(inDegree, graphList, prerequisites, numCourses);
=> Go look at that function
*/
var canFinish = function (numCourses, prerequisites) {
  if (prerequisites.length === 0) {
    return true;
  }

  const graphList = {};
  const inDegree = {};
  const queue = [];
  let courseCount = 0;

  // build inDegree and graphList
  buildTwoLists(inDegree, graphList, prerequisites, numCourses);

  // put all 0 indegree into queue
  for (const degree of Object.keys(inDegree)) {
    if (inDegree[degree] === 0) {
      queue.push(degree);
    }
  }

  while (queue.length > 0) {
    const course = queue.shift();

    // loop through all the edge of current vertex(course)
    for (const nextCourse of graphList[course]) {
      // decrease it's indegree
      inDegree[nextCourse]--;

      // if it's indgree is 0, add it to the qeuue
      // also means this vertex is not depend on anyone anymore
      if (inDegree[nextCourse] === 0) {
        queue.push(nextCourse);
      }
    }

    // take the course
    courseCount++;
  }
  return courseCount === numCourses;
};

/*
In this function, the main goal is building inDegree and graphList
However, we have to think of the edge case
which is when one of the vertex is not connected with any other vertexices
like this
numCourses = 5
prerequisites = [[1,4],[2,4],[3,1],[3,2]]
If numCourses = 5, all the course should be 0 ~ 4
Look at the edge in the prerequisites, there's no 0 course
Which means 0 is not connected with every one
But we do need to add it into the graphList and inDegree

So we use courseArr and
  for (let i = 0; i < courseArr.length; i++) {
    if (!courseArr[i]) {
      graphList[i] = [];
      inDegree[i] = 0;
    }
  }
to make sure we include all the vertex
*/
function buildTwoLists(inDegree, graphList, prerequisites, numCourses) {
  const courseArr = new Array(numCourses).fill(false);

  for (const course of prerequisites) {
    // In the graph representation, cur -> next
    // so inDegree[cur] = 0, and inDegree[next] = 1
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

  // if we have any vertext hasn't been searched, just put it to the graphList and inDegree
  for (let i = 0; i < courseArr.length; i++) {
    if (!courseArr[i]) {
      graphList[i] = [];
      inDegree[i] = 0;
    }
  }
}
