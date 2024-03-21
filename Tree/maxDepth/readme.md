# Problem Explanation

The problem is quite simple, so the following walk through the process of DFS and BFS using iterative approach<br>

## Using Stack (DFS)
Because we have to record the height of each node, using `TreeNode` inside stack is not enough<br>
We create a struct `HeightNode` to store the height of the node and the node itself<br>

Suppose the input is following
```
    3
   / \
  9  20
    /  \
   15   7
```
The initial stack is `[{n3, h1}]`<br>
(n means node, h means height)<br>
Let's going to the while-loop<br>

Pop the stock,<br>
current node is 3, height is 1<br>
Compare the current height(logic)<br>
Push the left and right child into the stack<br>
Current stack is 
`[{n9, h2}, {n20, h2}]`<br>

Pop the stack,<br>
current node is 9, height is 2<br>
Compare the current height(logic)<br>
Push the left and right child into the stack<br>
Current stack is 
`[{nullptr, h3}, {nullptr, h3}, {n20, h2}]`<br>

Pop the stack,<br>
current node is nullptr, height is 3<br>
Continue<br>
Current stack is 
`[{nullptr, h3}, {n20, h2}]`<br>

Pop the stack,<br>
current node is nullptr, height is 3<br>
Continue<br>
Current stack is 
`[{n20, h2}]`<br>

Pop the stack,<br>
current node is 20, height is 2<br>
Compare the current height(logic)<br>
Push the left and right child into the stack<br>
Current stack is 
`[{n15, h3}, {n7, h3}]`<br>

Pop the stack,<br>
current node is 15, height is 3<br>
Compare the current height(logic)<br>
Push the left and right child into the stack<br>
Current stack is 
`[{nullptr, h4}, {nullptr, h4}, {n7, h3}]`<br>

Pop the stack,<br>
current node is nullptr, height is 4<br>
Continue<br>
Current stack is 
`[{nullptr, h4}, {n7, h3}]`<br>

Pop the stack,<br>
current node is nullptr, height is 4<br>
Continue<br>
Current stack is 
`[{n7, h3}]`<br>

Pop the stack,<br>
current node is 7, height is 3<br>
Compare the current height(logic)<br>
Push the left and right child into the stack<br>
Current stack is 
`[{nullptr, h4}, {nullptr, h4}]`<br>

Pop the stack,<br>
current node is nullptr, height is 4<br>
Continue<br>
Current stack is 
`[{nullptr, h4}]`<br>

Pop the stack,<br>
current node is nullptr, height is 4<br>
Continue<br>
Current stack is 
`[]`<br>

The stack is empty, break the while-loop, and return 3<br>

There're two things to note about using stack to implement DFS<br>
1. The reason using stack can implement DFS is that the stack is LIFO, so the child node added later will be poped first<br>
- In the code, we push the left child first, then the right child<br>
- So the left child node will be poped first, and keep adding new child node into the stack<br>
- After exploring all the left child node, we start to explore the right child node<br>

2. We're allowed to push `nullptr` into the stack<br>
- The reason is because we put the base case(`if (d->node) == nullptr)`) ***BEFORE*** comparing the height<br>
- We can also check the node is `nullptr` before pushing the left and right child into the stack<br>

## Using Queue (BFS)
Unlike using stack, we don't need to record the height of each node<br>
Because BFS is naturally level-order traversal<br>
After exploring the tree using BFS, we can get the height of the tree<br>

Let's walk through the process of using queue to implement BFS<br>
Suppose the input is following
```
    3
   / \
  9  20
    /  \
   15   7
```
The initial queue is `[3]`<br>
Let's going to the while-loop<br>

Current size of queue is 1<br>
First iteration, <br>
Pop the queue, current node is 3<br>
Push the left and right child into the queue<br>
Current queue is 
`[9, 20]`<br>

Current size of queue is 2<br>
First iteration, <br>
Pop the queue, current node is 9<br>
Don't push the child node into the queue because the left and right child are `nullptr`<br>
Second iteration, <br>
Pop the queue, current node is 20<br>
Push the left and right child into the queue<br>
Current queue is 
`[15, 7]`<br>

Current size of queue is 2<br>
First iteration, <br>
Pop the queue, current node is 15<br>
Don't push the child node into the queue because the left and right child are `nullptr`<br>
Second iteration, <br>
Pop the queue, current node is 7<br>
Don't push the child node into the queue because the left and right child are `nullptr`<br>

The queue is empty, break the while-loop, and return 3<br>

There're two things to note about using queue to implement BFS<br>
1. The reason using queue can implement BFS is that the queue is FIFO, so the child node added first will be poped first<br>
- Look at the queue when it's `[9,20]`
- If 9 has child node, the child node will be added to the end of the queue, like this `[20, left_child, right_child]`<br> 
- Then 9 will be poped, and the left and right child of 20 will be added to the end of the queue<br>
- ***So that make sure the current level is explored before the next level***<br>

2. We can't allow to push `nullptr` into the queue<br>
- Why?
- Because we're poping the queue in the for-loop, and adding height is at the while-loop<br>
- Once we adding `nullptr` into the queue, the while-loop will add the height