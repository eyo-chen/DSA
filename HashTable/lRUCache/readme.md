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


# Thought Process

## Intuition(Just walk through the simple case)
It's pretty obvious that we need to use a hash table to store the key-value pairs.<br>
If capacity is 2<br>
Put(a, 1) -> {a: 1}<br>
Put(b, 2) -> {a: 1, b: 2}<br>
Put(c, 3)<br>
Now, the cache is full, how we can remove the least used key-value pair?<br>

## First Problem: How to find the least used key-value pair?
Before answering this question, the problem we have to answer is<br>
How to find the least used key-value pair?<br>
We can't know that from the hash table<br>
Therefore, we need to use another data structure to help us keep that information.<br>

### First Solution: Use another hash table
One solution is that we can use another hash table to store the key-frequency pairs.<br>
For example, if the operation is following:<br>
Put(a, 11)<br>
Put(b, 22)<br>
Put(c, 32)<br>
The frequency hash table will be like this:<br>
{1: a, 2: b, 3: c}<br>
But if we store value as key, why not just use array(slice)?

### Second Solution: Use array(slice)
If we use array(slice), the structure will be like this:<br>
[a, b, c]<br>
We know that the first element is the least used key<br>
And the last element is the most used key<br>

So, now back to our original example<br>
If capacity is 2<br>
Put(a, 1) -> {a: 1}<br>
Put(b, 2) -> {a: 1, b: 2}<br>
The array will be like this:<br>
[a, b]<br>

Now, the user wants to put(c, 3)<br>
The array will be like this:<br>
[b, c]<br>

What we need to do to convert [a,b] to [b,c]?<br>
(1) remove the first element from the array<br>
(2) add the new element to the last position of the array<br>

This works well, but both of the operations are O(N) time, which is not efficient.<br>

What data structure can help us to remove and add the element in O(1) time while helps us to maintain the order of the elements?<br>
That is linked list.

### Third Solution: Use linked list
We know that we can remove and add the element in O(1) time from the linked list.<br>
(If we can reference the node we want to remove or add)<br>

There are two types of linked list, singly linked list and double linked list.<br>
Which one should we use?<br>

It's pretty obvious that we should use double linked list.<br>
Why is that?<br>
Because we need to reference the previous node to remove the node from the list.<br>
Suppose we have the following linked list:<br>
head <-> 1 <-> 2 <-> 3 <-> tail<br>
If we want to remove the node 2, we need to reference the node 1.<br>
In singly linked list, we can't reference the previous node.<br>
Therefore, we need to traverse the list to find the previous node.<br>
This will take O(N) time, which is not efficient.<br>

Back to our original problem<br>
If capacity is 2<br>
Put(a, 1) -> {a: 1}<br>
Put(b, 2) -> {a: 1, b: 2}<br>
The linked list will be like this:<br>
head <-> a <-> b <-> tail<br>

Now, the user wants to put(c, 3)<br>
We can<br>
(1) remove the least used node a from the list(head.next)<br>
head <-> b <-> tail<br>
(2) add the node c to the last position of the list(tail.prev)<br>
head <-> b <-> c <-> tail<br>
It works!!!

However, the next problem is that how can we reference the node we want to remove or add?<br>

## Second Problem: How to reference the node we want to remove or add?
Suppose the current linked list is like this:<br>
head <-> a <-> b <-> c <-> tail<br>

If user wants to get(b), we need to reference the node b.<br>
We know that we have to do two things<br>
(1) remove the node b from the list<br>
(2) add the node b to the most position of the list<br>

However, if we can't reference the node b, we still need to traverse the list to find the node b.<br>
This will take O(N) time, which is not efficient.<br>

At the very beginning, we know that we need to use hash table to store the key-value pairs.<br>
At first, we only store values as values in the hash table.<br>
But we can actually store the node itself as the value in the hash table.<br>
Therefore, when user wants to get(b), we can directly reference the node b from the hash table in O(1) time.<br>

## Conclusion
Let's recap how we can implement the LRU cache.<br>
1. Use hash table to store the key-value pairs
2. Use double linked list to maintain the order of the elements
3. When we want to update the value to the most used position, we need to do two things<br>
   - (1) remove the node from the current position
   - (2) add the node to the most position
4. Store the node itself as the value in the hash table, so that we can reference the node in O(1) time.





