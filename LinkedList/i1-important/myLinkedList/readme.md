# Problem Explanation

Let's look each function one by one.
Note that all nodes in the linked list are 0-indexed.

## `int get(int index)`
Get the value of the indexth node in the linked list. If the index is invalid, return -1. <br>

Suppose current linked list is 1 -> 2 -> 3 -> 4 -> 5. <br>
get(1) will return 2. <br>
get(3) will return 4. <br>
get(5) will return -1. <br>

How to check the argument is valid or not? <br>
=> If index is less than 0 or greater than or equal to the length of the linked list, then it is invalid <br>
=> get(-1), get(5) are invalid <br>

What's the core logic?<br>
=> Traverse the linked list `index` times <br>
=> If `index = 0`, it means we don't need to traverse the linked list, just return the value of the head node <br>
=> If `index = 1`, it means we need to traverse the linked list once and return the value of the next node <br>
=> If `index = 2`, it means we need to traverse the linked list twice and return the value of the next node <br>

## `void addAtHead(int val)`
Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.

Suppose current linked list is 1 -> 2 -> 3 -> 4 -> 5 <br>
addAtHead(0) will change the linked list to 0 -> 1 -> 2 -> 3 -> 4 -> 5<br>
addAtHead(6) will change the linked list to 6 -> 0 -> 1 -> 2 -> 3 -> 4 -> 5<br>

What's the core logic?<br>
=> Create a new node with value `val` <br>
=> Make the next of the new node to the head of the linked list <br>
=> Update the new node as the head of the linked list <br>
=> Increment the length of the linked list <br>

## `void addAtTail(int val)`
Append a node of value val as the last element of the linked list.

Suppose current linked list is 1 -> 2 -> 3 -> 4 -> 5 <br>
addAtTail(6) will change the linked list to 1 -> 2 -> 3 -> 4 -> 5 -> 6 <br>
addAtTail(7) will change the linked list to 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 <br>

What's the core logic?<br>
=> Create a new node with value `val`<br>
=> Traverse the linked list till the last node. (Note that is the node which next points to nullptr) <br>
=> Make the next of the last node to the new node<br>

What's the edge case?<br>
=> If the linked list is empty, then the new node will be the head of the linked list<br>

## `void addAtIndex(int index, int val)`
Add a node of value val before the indexth node in the linked list. If index equals the length of the linked list, the node will be appended to the end of the linked list. If index is greater than the length, the node will not be inserted.

Suppose current linked list is 1 -> 2 -> 3 -> 4 -> 5<br>
addAtIndex(0, 0) will change the linked list to 0 -> 1 -> 2 -> 3 -> 4 -> 5<br>
addAtIndex(3, 6) will change the linked list to 1 -> 2 -> 3 -> 6 -> 4 -> 5<br>
addAtIndex(6, 6) will not change the linked list<br>

How to check the argument is valid or not? <br>
=> If index is less than 0 or greater than the length of the linked list, then it is invalid<br>
=> addAtIndex(-1, 0), addAtIndex(7, 0) are invalid<br>
=> Note that if index is equal to the length of the linked list, then it is valid<br>
=> It means we need to append the new node to the end of the linked list<br>
=> addAtIndex(5, 6) is valid <br>

What's the core logic?<br>
=> Create a new node with value `val` <br>
=> Traverse the linked list `index - 1` times <br>
=> Because we need to add the new node before the indexth node, so we need to reference the previous node of the indexth node <br>
=> Update the next of the new node to the next of the current node <br>
=> Update the next of the current node to the new node <br>
=> Increment the length of the linked list <br>

What's the edge case?<br>
=> If index is 0, then we need to add the new node at the head of the linked list<br>
=> Simply call `addAtHead(val)` function<br>

## `void deleteAtIndex(int index)`
Delete the indexth node in the linked list, if the index is valid.

Suppose current linked list is 1 -> 2 -> 3 -> 4 -> 5<br>
deleteAtIndex(0) will change the linked list to 2 -> 3 -> 4 -> 5<br>
deleteAtIndex(3) will change the linked list to 2 -> 3 -> 4<br>
deleteAtIndex(5) will not change the linked list<br>

How to check the argument is valid or not? <br>
=> If index is less than 0 or greater than or equal to the length of the linked list, then it is invalid<br>
=> deleteAtIndex(-1), deleteAtIndex(5) are invalid<br>

What's the core logic?<br>
=> Traverse the linked list `index - 1` times<br>
=> Because we need to delete the indexth node, so we need to reference the previous node of the indexth node <br>
=> For example, if index is 3, we need to traverse the linked list 2 times to reference the `2`<br>

What's the edge case?<br>
=> If index is 0, then we need to delete the head of the linked list<br>
=> Why this is an edge case? <br>
=> Because we don't have a previous node to reference <br>
=> Look at the code, `ptr->next` will cause the error <br>