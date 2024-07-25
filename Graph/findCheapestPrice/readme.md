# Problem Explanation

At first, this problem should be easy to come up with a solution.<br>

Let's walk through the thought process of the problem.<br>
Suppose we temporarily ignore the K stops and minimum price constraints.<br>
We just simply need to find the path from src to dst.<br>
In this case, we can use the DFS or BFS to find the path.<br>
However, there might be multiple paths from src to dst.<br>

Constraints 1: minimum price<br>
- Along with the search, we just need to keep track of the price.<br>
- Once we find the path, we can compare the price with the minimum price.<br>

Constraints 2: K stops<br>
- We have two options.<br>
  1. Similar to the minimum price, we can keep track of the stops.<br>
     - Once the stops is greater than K, we can stop the search.<br>
  2. We can use the BFS to find the path.<br>
     - To be more specific, we do the BFS layer by layer.<br>
     - At most, we can only search K+1 layers.<br>


Let's summarize how to solve this problem.<br>
(we choose to use to keep track of the stops)<br>
1. Create a adjacency list to store the adjacent flights and the price.<br>
2. Create a queue to store the current location, price, and stops.<br>
3. Do the BFS (Don't need to search layer by layer since we keep track of the stops)<br>
   (1) Pop the current location, price, and stops from the queue.<br>
   (2) If the current stops is greater than K, skip the process.<br>
   (3) If the current price is greater than the minimum price, skip the process.<br>
   (4) If the current location is the destination, update the minimum price.<br>
   (5) Otherwise, push all the adjacent flights to the queue.<br>


This solution is correct, but it is not efficient.<br>
And it will cause the time limit exceeded.<br>

Let's consider the following example.<br>
(We intentionally ignore the stops constraint)<br>
```
                 0(src)
             /       \
            /         \ 
            1 -------> 2
            \          /
             \        /
                  3(dst)
```
0 -> 1, 0 -> 2<br>
1 -> 2, 1 -> 3<br>
2 -> 3<br>

We know that the dst(3) is only connected with 1 and 2.<br>
If we try to find the minimum price from 0 to 3, <br>
What's the information we need to find that?
1. The price from 1 to 3
2. The price from 2 to 3
3. The shortest path to 1
4. The shortest path to 2

The third and fourth information is the key to solve this problem.<br>
We only need to know the shortest path to the location.<br>

There are two paths to 2 from 0.<br>
0 -> 1 -> 2<br>
0 -> 2<br>

Suppose
0 -> 1 -> 2 costs 300<br>
0 -> 2 costs 200<br>

And we first explore the path 0 -> 2, then we explore the path 0 -> 1 -> 2.<br>
After we explore the path 0 -> 2, we know that the minimum price to 2 is 200.<br>
Later, when we explore the path 0 -> 1 -> 2, we know that the price is greater than the existing minimum price(200).<br>
There's no point we ever consider this path.<br>
Therefore, we don't need to explore the path 0 -> 1 -> 2 -> 3.<br>

That's the optimization to solve this problem.<br>
We need to create an array to store the minimum price to each location.<br>
When we hit a location, we only consider the path if the current price is less than the existing minimum price.<br>

Let's summarize the solution.<br>
1. Create a adjacency list to store the adjacent flights and the price.<br>
2. Create a queue to store the current location, price, and stops.<br>
3. Create an array to store the minimum price to each location.<br>
4. Do the BFS (Don't need to search layer by layer since we keep track of the stops)<br>
   - Pop the current location, price, and stops from the queue.<br>
   - If the current location is the destination, update the minimum price.<br>
   - Loop through the adjacent flights.<br>
     - If the current price is greater than the minimum price, skip the process.<br>
     - If the current price is less than the existing minimum price
       - Update the minimum price.<br>
       - Push the adjacent flight to the queue.<br>

Let's consider the following example.<br>
```
                 0(src)
             /       \
      100   /         \   200
            1 --200---> 2
            \          /
      400    \        /   100
                  3(dst)
```
```
0 -> [1, 100], [2, 200]
1 -> [2, 200], [3, 400]
2 -> [3, 100]
```
stops = 2<br>

- Starting point
  - prices = [inf, inf, inf, inf]
  - minimum price = inf
  - queue = [(0, 0, 0)]

- First iteration
  - Pop (0, 0, 0) - current location, price, stops
  - queue = [(1, 100, 1), (2, 200, 1)]
  - prices = [0, 100, 200, inf]

- Second iteration
  - Pop (1, 100, 1)
  - queue = [(2, 200, 1), (3, 500, 2)]
  - prices = [0, 100, 200, 500]
  - Because the path 0 -> 1 -> 2 is greater than the existing minimum price(prices[2]), we skip the process.

- Third iteration
  - Pop (2, 200, 1)
  - queue = [(3, 500, 2), (2, 300, 3)]
  - prices = [0, 100, 200, 300]

- Fourth iteration
  - Pop (3, 500, 2)
  - queue = [(3, 300, 2)]
  - prices = [0, 100, 200, 300]
  - 3 is the destination, and the price is 500, which is less than the minimum price.<br>
  - minimum price = 500

- Fifth iteration
  - Pop (3, 300, 2)
  - queue = []
  - prices = [0, 100, 200, 300]
  - 3 is the destination, and the price is 300, which is less than the minimum price.<br>
  - minimum price = 300

The minimum price is 300.<br>


## Bellman Ford Algorithm
There's another way to solve this problem.<br>
It's called the Bellman Ford Algorithm.<br>

The idea is to explore all the paths from the source to the destination k + 1 times.<br>
In each iteration, we keep updating the minimum price to each location.<br>
After k + 1 iterations, we will have the minimum price to the destination.<br>
Also, in each iteration, we need to create a new array to store the current minimum price.<br>

Let's summarize the solution.<br>
1. Create an array to store the minimum price to each location.<br>
2. Initialize the minimum price to the source as 0.<br>
3. Loop through k + 1 times.<br>
   (1) Create a new array to store the current minimum price.<br>
   (2) Loop through the flights.<br>
       (a) If the current source flight is not reachable, skip the process.<br>
           - We want to find the minimum price to the destination.<br>
           - If the current source flight is not reachable, we don't need to consider the path.<br>
       (a) If the current price is less than the existing minimum price
           - Update the minimum price.<br>

Let's consider the following example.<br>
```
                 0(src)
             /       \
      100   /         \   200
            1 --200---> 2
            \          /
      400    \        /   100
                  3(dst)
```
k = 2<br>

- Starting point
  - prices = [0, inf, inf, inf]

- First iteration
  - prices = [0, inf, inf, inf]
  - temp = [0, inf, inf, inf]
  - First flight: 0 -> 1, price = 100
    - prices[0] is reachable (prices[0] != inf)
    - temp[1] = min(temp[1], prices[0] + 100) = 100
    - temp = [0, 100, inf, inf]
  - Second flight: 0 -> 2, price = 200
    - prices[0] is reachable (prices[0] != inf)
    - temp[2] = min(temp[2], prices[0] + 200) = 200
    - temp = [0, 100, 200, inf]
  - Third flight: 1 -> 2, price = 200
    - prices[1] is NOT reachable (prices[1] == inf), skip the process
  - Fourth flight: 1 -> 3, price = 400
    - prices[1] is NOT reachable (prices[1] == inf), skip the process
  - Fifth flight: 2 -> 3, price = 100
    - prices[2] is NOT reachable (prices[2] == inf), skip the process

- Second iteration
  - prices = [0, 100, 200, inf]
  - temp = [0, 100, 200, inf]
  - First flight: 0 -> 1, price = 100
    - prices[0] is reachable (prices[0] != inf)
    - temp[1] = min(temp[1], prices[0] + 100) = 100
    - temp = [0, 100, 200, inf]
  - Second flight: 0 -> 2, price = 200
    - prices[0] is reachable (prices[0] != inf)
    - temp[2] = min(temp[2], prices[0] + 200) = 200
    - temp = [0, 100, 200, inf]
  - Third flight: 1 -> 2, price = 200
    - prices[1] is reachable (prices[1] != inf)
    - temp[2] = min(temp[2], prices[1] + 200) = 200
    - temp = [0, 100, 200, inf]
  - Fourth flight: 1 -> 3, price = 400
    - prices[1] is reachable (prices[1] != inf)
    - temp[3] = min(temp[3], prices[1] + 400) = 500
    - temp = [0, 100, 200, 500]
  - Fifth flight: 2 -> 3, price = 100
    - prices[2] is reachable (prices[2] != inf)
    - temp[3] = min(temp[3], prices[2] + 100) = 300
    - temp = [0, 100, 200, 300]

- Third iteration
  - prices = [0, 100, 200, 300]
  - temp = [0, 100, 200, 300]
  - First flight: 0 -> 1, price = 100
    - prices[0] is reachable (prices[0] != inf)
    - temp[1] = min(temp[1], prices[0] + 100) = 100
    - temp = [0, 100, 200, 300]
  - Second flight: 0 -> 2, price = 200
    - prices[0] is reachable (prices[0] != inf)
    - temp[2] = min(temp[2], prices[0] + 200) = 200
    - temp = [0, 100, 200, 300]
  - Third flight: 1 -> 2, price = 200
    - prices[1] is reachable (prices[1] != inf)
    - temp[2] = min(temp[2], prices[1] + 200) = 200
    - temp = [0, 100, 200, 300]
  - Fourth flight: 1 -> 3, price = 400
    - prices[1] is reachable (prices[1] != inf)
    - temp[3] = min(temp[3], prices[1] + 400) = 300
    - temp = [0, 100, 200, 300]
  - Fifth flight: 2 -> 3, price = 100
    - prices[2] is reachable (prices[2] != inf)
    - temp[3] = min(temp[3], prices[2] + 100) = 300
    - temp = [0, 100, 200, 300]

The minimum price is 300.<br>

We might ask what's the point of using temp array in each iteration.<br>
Let's consider the following example.<br>
```
                 0(src)
             /       
      100   /            
            1 --200---> 2
            \          /
      400    \        /   100
                  3(dst)
```
stops = 1<br>
There's no 0 -> 2 flight.<br>

- Starting point
  - prices = [0, inf, inf, inf]

- First Iteration
  - prices = [0, inf, inf, inf]
  - First flight: 0 -> 1, price = 100
    - prices[0] is reachable (prices[0] != inf)
    - prices[1] = min(prices[1], prices[0] + 100) = 100
    - prices = [0, 100, inf, inf]
  - Second flight: 1 -> 2, price = 200
    - prices[1] is reachable (prices[1] != inf)
    - prices[2] = min(prices[2], prices[1] + 200) = 300
    - prices = [0, 100, 300, inf]
  - Third flight: 1 -> 3, price = 400
    - prices[1] is reachable (prices[1] != inf)
    - prices[3] = min(prices[3], prices[1] + 400) = 500
    - prices = [0, 100, 300, 500]
  - Fourth flight: 2 -> 3, price = 100
    - prices[2] is reachable (prices[2] != inf)
    - prices[3] = min(prices[3], prices[2] + 100) = 400
    - prices = [0, 100, 300, 400]

After the first iteration, the prices array is [0, 100, 300, 400].<br>
And we know there's a problem<br>
Because the stops is 1, the only path from 0 to 3 is 0 -> 1 -> 3.<br>
And the price of that path is 500.<br>
However, the prices array is [0, 100, 300, 400].<br>
That's the reason why we need to use the temp array in each iteration.<br>

# Complexity Analysis
## Time Complexity O(E⋅(k+1))
where E represents the number of flights and k represents the number of stops.

1. **Building the Adjacency List:**
   - The loop iterating over `flights` has a time complexity of \(O(E)\), where \(E\) is the number of flights.
   - Each flight is processed once to build the adjacency list.

2. **Initializing Prices Array:**
   - Initializing the `prices` array has a time complexity of \(O(V)\), where \(V\) is the number of nodes (cities).

3. **Breadth-First Search (BFS):**
   - The BFS loop runs while there are elements in the queue. In the worst case, every edge could be processed up to \(k+1\) times, leading to a time complexity of O(E⋅(k+1)).
   - For each flight processed, we may add new flights to the queue, leading to at most \(O(E)\) operations in total for the queue manipulations.

## Space Complexity O(E⋅(k+1))
where E represents the number of flights and k represents the number of stops.

1. **Adjacency List:**
   - The adjacency list stores all the flights, so it requires \(O(E)\) space.

2. **Prices Array:**
   - The `prices` array has a size of \(n\), requiring \(O(V)\) space.

3. **Queue:**
   - In the worst case, the queue could store all the nodes multiple times, but it will typically contain at most \(O(V \cdot (k+1))\) elements. However, considering each element in the queue represents a unique path rather than unique nodes, the space complexity can be approximated as \(O(E \cdot (k+1))\), similar to the time complexity.

