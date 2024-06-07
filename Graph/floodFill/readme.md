# Problem Explanation

The problem is quite simple. We have three approaches to solve this problem.<br>
However, the core logic is the same, we just use different data structures or way to implement the logic.<br>

The core logic is
1. Put the starting point into the data structure.
2. While the data structure is not empty, do the following:
    1. Pop the point from the data structure.
    2. If the point is not valid, continue.
    3. If the point is valid, change the color of the point and add the neighbors of the point to the data structure.
3. Return the image.

Note that there's another way to implement the core logic
1. Put the starting point into the data structure.
2. While the data structure is not empty, do the following:
    1. Pop the point from the data structure.
    2. If the point is not valid, continue.
    3. Add the neighbors of the point to the data structure.
    4. Change the color of the point.
3. Return the image.

The main difference is that we just put the point into the data structure at first<br>
Then we check if the point is valid in the future iteration.<br>
(I implemented the second approach in the code)<br>

The point is not valid when
1. The point is out of the image.
2. The point has been explored.
3. The point is not the same color as the starting point.

We can use a stack or a queue to implement the data structure.<br>
Or we can use a recursive function to implement the data structure.<br>

# Complexity Analysis
## Time Complexity: O(r * c)
- where r is the number of rows and c is the number of columns in the image.
- assume that `set.find` is O(1) time complexity.
- at worst, we might need to visit every point in the image.

## Space Complexity: O(r * c)
- where r is the number of rows and c is the number of columns in the image.
- the space complexity is O(r * c) because we might need to store every point in the image in the data structure and the set.