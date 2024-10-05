# Problem Explanation

In order to satisfy the LRU cache, we need a data structure to handle the order of the elements.<br>
For example,<br>
```
least -> 1 -> 2 -> 3 -> most
```
Suppose we get or update the value of 2, the order will be like this:<br>
```
least -> 1 -> 3 -> 2 -> most
```
How can we implement this?<br>
The idea is that we<br>
1. remove the node from the current position
2. add the node to the most position

In order to remove the node from the current position in O(1) time, we use double linked list.<br>
For example,<br>
```
head <-> 1 <-> 2 <-> 3 <-> tail
```
If we want to remove node 2, the process is following:<br>
1. 1.next = 3
2. 3.prev = 1
That's it.


Also, as long as we keep the head and tail of the double linked list, we can easily insert and remove the node from the head and tail.<br>
For the constant access time, we use hash table.

Now, let's see how to implement the LRU cache.<br>
- Get:
  - If the key is in the hash table, 
    1. remove the node from the current position
    2. add the node to the most position(head)
    3. return the value
  - Otherwise, we return -1.
- Put:
  - If the key is in the hash table,
    1. update the value
    2. remove the node from the current position
    3. add the node to the most position(head)
  - Otherwise, we check the capacity of the cache. If the cache is full, 
    1. remove the tail node
    2. remove the tail node from the hash table
  - Add the new node to the hash table
  - Add the new node to the most position(head)


# Complexity Analysis
## Time Complexity
- Get: O(1)
- Put: O(1)

## Space Complexity
- O(N) where N is the capacity of the cache
