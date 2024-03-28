# Problem Explanation

Here only explain the two iterative solutions because recursive solution is very simple<br>

We know that inorder traversal is left -> root -> right<br>
And we know we can use stack to simulate the recursive process<br>
But the problem is how do we use stack to do the inorder traversal<br>

## Using Stack and Hashmap
This is the solution I came up with initially<br>
It using additional hashmap, so it's less efficient<br>

Let's walk through the thought process<br>
Because we want to simulate the recursive approach, the first idea is just following the pattern of recursive approach<br>
go left, do something, go right<br>
But this approach has one problem, let's see the example below<br>
```
    1
   / \
  2   3
 / \
4   5
```
We first push root(1) to the stack to start the process<br>
stack: [1]<br>
result: []<br>
Then we go left, push 2 to the stack<br>

stack: [1, 2]<br>
result: []<br>
Then we go left, push 4 to the stack<br>

stack: [1, 2, 4]<br>
result: []<br>
Because there's no left to explore<br>
We pop 4 from the stack, add 4 to the result<br>

stack: [1, 2]<br>
result: [4]<br>
Now, we back to the initial iteration<br>
At this point, the current top node(n) in the stack is 2<br>
Because `2.left` is not nullptr, so we'll push 4 to the stack again<br>
Can you see the problem here?<br>
We'll add duplicate nodes to the stack in this pattern<br>
So we need a way to know if we've visited the node before<br>
That's why I came up with the hashmap solution<br>

The hashmap solution is simple<br>
We just need to add the node to the hashmap when we push it to the stack<br>
And we'll check if the node is in the hashmap before we push it to the stack<br>

## Using Stack Only
This solution is a little bit tricky<br>
To come up with this solution, we need to understand the essence of the inorder traversal<br>
It seems that the pattern is left -> root -> right<br>
But there's a hidden pattern in the process<br>
Which is ***go left as far as possible***<br>

Using the following as example
```
      1
     / \
    2   3
   / \
  4   5
     / \
    6   7
```
When we're at root(1), we go left as far as possible to 4<br>
Then we back to 2, and go right to 5<br>
As soon as we reach 5, we go left as far as possible to 6<br>
That's the hidden pattern<br>

And the concept of ***go left as far as possible*** is the key to the solution<br>
Because it implicitly tells us that we need a inner while-loop to go left as far as possible<br>
So the idea is following<br>
- We do not push the node to the stack intially(before the while-loop)<br>
- We use `curNode` outside the while-loop<br>
  - `curNode` help us to keep track the ***state*** of the process<br>
  - We can think as the arrow pointing to the current exploring node<br>
- Continue the while-loop until the stack is empty AND `curNode` is null<br>
- Inside the while-loop, the first thing we do is go left as far as possible<br>
- After finishing exploring the left, we know we push the node to the result<br>
- Update `curNode` to `curNode.right`(go right)<br>

Let's walk through the example above<br>
```
      1
     / \
    2   3
   / \
  4   5
     / \
    6   7
```
curNode = 1<br>
stack = []<br>
result = []<br>
stack is empty, but curNode is not null<br>
Go into the while-loop, and go left as far as possible<br>

curNode = nullptr (it's 4.left)<br>
stack = [1, 2, 4]<br>
result = []<br>
Break the inner while-loop<br>
Let curNode = 4<br>
Add 4 to the result<br>
Update curNode to 4.right = nullptr<br>

curNode = nullptr (it's 4.right)<br>
stack = [1, 2]<br>
result = [4]<br>
stack is not empty, but curNode is null<br>
Because curNode is null, no need to explore the left<br>
Update curNode to 2 (arrow point to 2, which tell us the current state of the process)<br>
Add 2 to the result<br>
Update curNode to 2.right = 5<br>

curNode = 5<br>
stack = [1]<br>
result = [4, 2]<br>
stack is not empty, and curNode is not null<br>
Go left as far as possible<br>

curNode = nullptr (it's 6.left)<br>
stack = [1, 5, 6]<br>
result = [4, 2]<br>
Break the inner while-loop<br>
Let curNode = 6<br>
Add 6 to the result<br>
Update curNode to 6.right = nullptr<br>

curNode = nullptr (it's 6.right)<br>
stack = [1, 5]<br>
result = [4, 2, 6]<br>
stack is not empty, but curNode is null<br>
Because curNode is null, no need to explore the left<br>
Update curNode to 5<br>
Add 5 to the result<br>
Update curNode to 5.right = 7<br>

curNode = 7<br>
stack = [1]<br>
result = [4, 2, 6, 5]<br>
stack is not empty, and curNode is not null<br>
Go left as far as possible<br>

curNode = nullptr (it's 7.left)<br>
stack = [1, 7]<br>
result = [4, 2, 6, 5]<br>
Break the inner while-loop<br>
Let curNode = 7<br>
Add 7 to the result<br>
Update curNode to 7.right = nullptr<br>

curNode = nullptr (it's 7.right)<br>
stack = [1]<br>
result = [4, 2, 6, 5, 7]<br>
stack is not empty, but curNode is null<br>
Because curNode is null, no need to explore the left<br>
Update curNode to 1<br>
Add 1 to the result<br>
Update curNode to 1.right = 3<br>

curNode = 3<br>
stack = []<br>
result = [4, 2, 6, 5, 7, 1]<br>
stack is empty, and curNode is not null<br>
Go left as far as possible<br>

curNode = nullptr (it's 3.left)<br>
stack = [3]<br>
result = [4, 2, 6, 5, 7, 1]<br>
Break the inner while-loop<br>
Let curNode = 3<br>
Add 3 to the result<br>
Update curNode to 3.right = nullptr<br>

curNode = nullptr (it's 3.right)<br>
stack = []<br>
result = [4, 2, 6, 5, 7, 1, 3]<br>
stack is empty, and curNode is null<br>
Break the outer while-loop<br>

Return the result<br>
