# Problem Explanation

The overall idea to solve this problem is straightforward.<br>
We immediately know that we have to use breadth-first search (BFS) to solve this problem since the question asks us to find the ***shortest path***.<br>

There are two things to note:<br>
1. We have to use a for-loop inside the while-loop
2. Check if the location is valid or is an exit before adding it to the queue is more efficient

For the first point, why?<br>
Because we have to explore the path level by level.<br>
For example, if we have a graph like this:<br>
```
  .   x   x   x   x
  x   x   s   .   x
  .   .   .   .   x
  x   x   x   x   .
```
Suppose the start point is `s`, we know that the shortest path is 3.<br>
We can see that the path is explored level by level.<br>
```
  .   x   x   x   x
  x   x   s   1   x
  3   2   1   2   x
  x   x   x   x   .
```

That's reason we have to use a for-loop inside the while-loop.<br>
It guarantees that we explore the path level by level.<br>

For the second point, why?<br>
Use same example above.<br>
```
  .   x   x   x   x
  x   x   s   .   x
  .   .   .   .   x
  x   x   x   x   .
```
If we don't check if the location is valid or is an exit before adding it to the queue, we will have to put the following location(q) into queue:
```
  .   x   q   x   x
  x   q   s   q   x
  .   .   q   .   x
  x   x   x   x   .
```
There are two locations are not valid, but we still need to do four iterations in the next level.<br>
But if we do check beforehand, we will have to put the following location(q) into queue:
```
  .   x   x   x   x
  x   x   s   q   x
  .   .   q   .   x
  x   x   x   x   .
```
There are only two locations need to further explore.<br>

Another reason is that when we're at the following location(q):
```
  .   x   x   x   x
  x   x   s   .   x
  .   q   .   .   x
  x   x   x   x   .
```
Now, if we do (qX - 1), then we immediately find the exit.<br>
And can return the result.<br>

However, there is one thing to remember if we follow this pattern:<br>
We have to set the initial `steps` to 1 instead of 0.<br>
Why is that?<br>
For out initial idea, the pattern is (explore -> check)
1. add all the locations into the queue (explore)
2. check if the location is valid or is an exit (check)

For our new idea, the pattern is (check -> explore)
1. check if the location is valid or is an exit (check)
2. add all the valid locations into the queue (explore)

And we update the `steps` at the end of the while-loop.<br>
For our initial idea, we're indeed put the exit into the queue, then we check if the location is an exit.It means that we update the `steps` before we return the result.<br>
For our new idea, we're never put the exit into the queue because we check if the location is an exit before we add it to the queue. It means that we haven't updated the `steps` before we return the result.<br>

Although the overall time complexity is the same, the new idea is more efficient.<br>

The final thing to note is how to check if a location is an exit.<br>
We know that the exit is the border of the maze.<br>
```
isOnBorder =  r == 0 || r == row - 1 || c == 0 || c == col - 1
```

But we have to handle the case when the start point is also on the border.<br>
```
notTheStartPoint = r != (start point x) || c != (start point y)
```
We can guarantee that the current location(x, y) is not the start point if one of them is not equal to the start point.<br>

To combine these two conditions, we must use `AND` because we have to make sure that the current location is on the border and is not the start point.<br>
```
(isOnBorder && notTheStartPoint)
```
```
(r == 0 || r == row - 1 || c == 0 || c == col - 1) && (r != (start point x) || c != (start point y))
```

# Complexity Analysis
## Time Complexity (O(N))
- where n = row * col
- At most, we have to visit all the locations in the maze

## Space Complexity (O(N))
- where n = row * col
- At most, the size of the queue is n