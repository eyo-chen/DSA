# Problem Explanation

## First Solution (Using Extrac Space)
The idea is to store the nodes of the linked list in an array, then we can reverse the group of k nodes easily by manipulating the array indices. Finally, we reconstruct the linked list from the modified array.

### Complexity Analysis
#### Time Complexity O(n)
- We traverse the linked list once to store the nodes in an array, which takes O(n) time.
- We then reverse the nodes in groups of k, which also takes O(n) time.

#### Space Complexity O(n)
- We use an array to store the nodes of the linked list, which takes O(n) space.


## Second Solution (Recursion)
The idea is to reverse the first k nodes of the linked list using recursion. We reverse the first k nodes and then recursively call the function for the next k nodes.

Let's see one example to understand this better:<br>
Consider the linked list: 1 -> 2 -> 3 -> 4 -> 5 and k = 2<br>
First, we reverse the first 2 nodes: <br>
After reversing, we got the following linked list<br>
```
   head  pre cur
x <- 2 <- 1   3
```
- `head` points to the new tail of the reversed group
- `pre` points to the new head of the reversed group
- `cur` points to the node after the reversed group

Now, we can recursively call the function for the next k nodes (starting from `cur`).<br>
And we connect the `head` of the first reversed group to the result of the recursive call.<br>
`head.next = reverseKGroup(cur, k)`<br>
Finally, we return `pre` as the new head of the reversed linked list.

### Complexity Analysis
#### Time Complexity O(n)
- We traverse the linked list once to reverse the nodes in groups of k, which takes O(n) time.

#### Space Complexity O(n/k)
- We use recursion, and the maximum depth of the recursion stack will be n/k, where n is the number of nodes in the linked list. Hence, the space complexity is O(n/k).


## Third Solution (Iterative In-Place Reversal)
The idea is to reverse the nodes of the linked list in groups of k using an iterative approach. We maintain pointers to keep track of the previous group's tail, the current group's head, and the next group's head.<br>

Let's see one example to understand this better:<br>
Consider the linked list: 1 -> 2 -> 3 -> 4 -> 5 and k = 2<br>

We first initialize two variables:
- `dummy`: A dummy node that points to the head of the linked list. In the end, we will return `dummy.next` as the new head of the reversed linked list.
- `groupPrev`: This is the node before the current group of k nodes. Initially, it points to `dummy`.
  - For example, `groupPrev -> 1 -> 2 -> 3 -> 4 -> 5`, where `groupPrev` points to the very first node.

Now, we use helper function to get the kth node from the current position. If there are less than k nodes remaining, we stop the process.<br>
```
                 kth
groupPrev -> 1 -> 2 -> 3 -> 4 -> 5
```
Now, we are about to reverse this part: `groupPrev -> 1 -> 2`<br>
We know that the ultimate result is going to be something like this:
```
groupPrev -> 2 -> 1 -> 3 -> 4 -> 5
```
We can achieve this by setting up two pointers:
- `prev := kth.next` (which is `3` in this case), so that we can set 1's next to `3` after reversal.
- `curr := groupPrev.next` (which is `1` in this case)

After reversing, we need to connect the previous part of the list to the newly reversed group and also connect the tail of the reversed group to the next part of the list.<br>
Draw it out to visualize it better:

### Complexity Analysis
#### Time Complexity O(n)
- We traverse the linked list once to reverse the nodes in groups of k, which takes O(n) time.

#### Space Complexity O(1)
- We use only a constant amount of extra space for pointers, so the space complexity is O(1).