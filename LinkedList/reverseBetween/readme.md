# Problem Explanation

The overall idea of solving this problem is very straightforward<br>
The hard part is just the implementation<br>

## First Solution
I came up with this solution<br>
The code is fine<br>

Let's walk through the thought process<br>
Consider the input and output<br>
Input: head = [1,2,3,4,5], left = 2, right = 4<br>
Output: [1,4,3,2,5]

We know we need to reverse the linked list<br>
The core problem is<br>
1. What other nodes do we need to reference?
2. Are there any other things we need to do after reversing the linked list?

First Problem<br>
It's obvious that we at least need to reference the node before the left node, which is `1` in this case<br>
Because we need to change the `next` pointer of `1` to point to the new head of the reversed linked list<br>
1 -> 4<br>

Second Problem<br>
We're not sure at this point<br>

Let's try to solve the problem<br>
Input: head = [1,2,3,4,5], left = 2, right = 4<br>
Output: [1,4,3,2,5]<br>
ns: node before the start pointer<br>
sp: start pointer<br>

1. We need to find both the `ns` and `sp`<br>
<pre>
    ns   sp
x -> 1 -> 2 -> 3 -> 4 -> 5
</pre>

2. From `sp`, we need to reverse the linked list<br>
c: current pointer<br>
p: previous pointer<br>
<pre>
    ns   sp         p    c
x -> 1 -> 2 <- 3 <- 4    5
     x <- | 
</pre>
After reversing the linked list, we know that
1. `c` is at the next pointer of the right node
2. `p` is at the right node
3. `sp` is pointing at the nullptr, which is incorrect

<pre>
    ns   sp
x -> 1 -> 2 -> x
          2 <- 3 <- 4  5
                    p  c
</pre>
Output: [1,4,3,2,5]<br>

Now, it's clear that what should we do after reversing the linked list<br>
1. We need to connect the `sp` to the `c`
2. We need to connect the `ns` to the `p`

<pre>
    ns    p        sp    c
x -> 1 -> 4 -> 3 -> 2 -> 5
</pre>

Let's summarize the implementation steps<br>
1. Find the `ns` and `sp`
 - `ns` is for the node before the left node(start pointer), so we know where to connect the new head of the reversed linked list
 - `sp` is for the start pointer, so we know where to start reversing the linked list
2. Reverse the linked list from `sp` to `right`
3. Connect the `ns` to the `p`
4. Connect the `sp` to the `c`

Let's see the final example of this implementation<br>
Input: head = [1,2,3], left = 2, right = 3<br>
Output: [1,3,2]<br>

1. Find the `ns` and `sp`<br>
<pre>
    ns   sp
x -> 1 -> 2 -> 3 -> x
</pre>

2. Reverse the linked list from `sp` to `right`<br>
<pre>
    ns   sp    p    c
x -> 1 -> 2 <- 3    x
     x <- |
</pre>

3. Connect the `ns` to the `p`<br>
4. Connect the `sp` to the `c`<br>
<pre>
    ns    p        sp   
x -> 1 -> 3 -> 2 -> x
</pre>

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(1)


## Second Solution
The second solution use a liitle bit of different approach<br>
In the previous solution, we reverse the sub-list, and re-wire the pointers after reversing<br>

In this problem, we use different approach<br>
Instead of reversing the sub-list, we just move the node step by step<br>

For example, 1 -> 2 -> 3 -> 4 -> 5 -> x, left = 2, right = 5<br>
Again, we still need to referecne the node BEFORE the head of sublist<br>
In this case, is node 1(nbr = nodeBeforeReverse)<br>
<pre>
nbr  cur
 1 -> 2 -> 3 -> 4 -> 5 -> x
</pre>
The process is sth like this<br>
In each step, we move the next node of `cur` to the head of sublist<br>
In other words, we move the next node of `cur` to in front of the sublist<br>
***In other words, we move the next node of `cur` to the next node of `nbr`***<br>

<pre>
nbr  cur  next
 1 -> 2 -> 3 -> 4 -> 5 -> x
</pre>
next = cur.next -> 3 <br>
move 3 to the next node of `nbr` <br>
we want to sth like this 1 -> 3<br>

1. 
<pre>
nbr       cur  next
 1 -> 3 -> 2 -> 4 -> 5 -> x
</pre>
move 3 to the next node of `nbr`<br>
Now the new next node is 4 because cur.next = 4 <br>

2.
<pre> 
nbr            cur  next
 1 -> 4 -> 3 -> 2 -> 5 -> x
</pre>
move 4 to the next node of `nbr`<br>
now the new next node is 5, because cur.next = 5 <br>

Note that we always want to move next node to the next node of `nbr`

Just like previous solution<br>
`dummyHead` helps us to move the ptr to correct position

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(1)

## Third Solution (2025/10/10)
The idea is very similar to the first solution, but it's more easy to understand<br>
It's obvious that all we need to do is to reference two nodes<br>
1. The node before the left node
2. The node after the right node

After we have that reference, we can reverse the linked list between left and right<br>
Before starting the reversal, we know that we can set `prev` node to the node after the right node(Look at the diagram to understand why)<br>
Then we can reverse the linked list in the standard way<br>
In the end, we just need to connect the node before the left node to the new head of the reversed linked list, which is `prev`<br>

It's very straightforward from the diagram and the code<br>