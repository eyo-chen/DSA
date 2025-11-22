# Problem Explanation

## First Solution: Using Hash Set

This idea is very straightforward, we just need to iterate through the linked list and store the node in a hash set<br>
If we encounter a node that is already in the hash set, then we know that there is a cycle in the linked list<br>

### Complexity Analysis

#### Time Complexity: O(n)

#### Space Complexity: O(n)
- We need to store all the nodes in the hash set

## Second Solution: Using Two Pointers

This solution is tricky, but it is very elegant<br>

Using two pointers, `slow` and `fast`<br>
`slow` moves one step at a time, and `fast` moves two steps at a time<br>
If there is a cycle in the linked list, then `slow` and `fast` will meet at some point<br>

Why is that?<br>
Let's say there's a cycle in the linked list<br>
And the length of the cycle is `10`<br>
`slow` and `fast` start at the same node<br>
`slow` moves one step at a time, and `fast` moves two steps at a time<br>

Before moving, we know the difference that `slow` and `fast` is gonna meet next time is `10`<br>
At each movement, the difference between `slow` and `fast` is reduced by `1`<br>
`slow` moving forward by 1, `fast` moving forward by 2<br>
So the difference is reduced by `1`<br>

Look at the following example<br>
At first move, `slow` at 1, `fast` at 2<br>
At second move, `slow` at 2, `fast` at 4<br>
At third move, `slow` at 3, `fast` at 6<br>
At fourth move, `slow` at 4, `fast` at 8<br>
At fifth move, `slow` at 5, `fast` at 10<br>
At sixth move, `slow` at 6, `fast` at 2<br>
At seventh move, `slow` at 7, `fast` at 4<br>
At eighth move, `slow` at 8, `fast` at 6<br>
At ninth move, `slow` at 9, `fast` at 8<br>
At tenth move, `slow` at 10, `fast` at 10<br>

So we know that `slow` and `fast` will meet at the tenth move<br>

### Complexity Analysis

#### Time Complexity: O(n)

#### Space Complexity: O(1)