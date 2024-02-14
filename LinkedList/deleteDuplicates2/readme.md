# Problem Explanation

The core idea of solving this problem is similar to the first version of the problem<br>
The only difference is that we need to remove all the nodes with duplicate values<br>

Let's look at an example<br>
head = [1,2,3,3,4]<br>
We know the output should be [1,2,4]<br>

So the first problem is<br>
How to remove the duplicate node list?<br>
which is [3,3] in this case<br>
If we have the reference of `2` and `4`<br>
We can easily remove the duplicate node list<br>

So the next problem is<br>
How to get the reference of node before and after the duplicate node list?<br>
which is `2` and `4` in this case<br>
Apparently, we need to use two pointers, `p1` and `p2`<br>
`p1` is the node before the duplicate node list<br>
`p2` is the node after the duplicate node list<br>
The logic is<br>
Initially, `p1` and `p2` have one distance<br>
Then we use `p2` to find the end of the duplicate node list<br>
If `p2->val == p2->next->val`, we move `p2` to the next node<br>
Until `p2->val != p2->next->val`<br>
Then `p2` is the node after the duplicate node list<br>
And `p1` is the node before the duplicate node list<br>
Because `p1` never moves<br>

Let's consider<br>
Where should we put `p1` and `p2` at the beginning?<br>
head = [1,2,3,3,4]<br>
We put `1` and `2` or `dummy` and `1`?<br>
It's pretty abvious that we should put `dummy` and `1`<br>
Because if the given list is [1,1,1,2,3]<br>
If we put `1` and `1` at the beginning<br>
There's no way we can remove the duplicate node list ([1,1,1])<br>
So we need to put `dummy` and `1` at the beginning<br>

Let's summarize the logic<br>
1. Use two pointers
   - `p1` and `p2`
   - `p1` is the node before the duplicate node list
   - `p2` is the node after the duplicate node list
   - `p1` is initialized at `dummy`
   - `p2` is initialized at `head`
   - `p2` is working pointer
2. If `p2->val == p2->next->val`
   - We keep moving `p2` to the next node until `p2->val != p2->next->val`
   - Then `p2` is the node after the duplicate node list
   - And `p1` is the node before the duplicate node list
   - So we link `p1` to `p2` to remove the duplicate node list
3. If `p2->val != p2->next->val`
   - We move `p1` and `p2` to the next node

For example,<br>
head = [1,2,3,3,4]<br>

Use two pointers, `p1` and `p2`
<pre>
p1   p2
x -> 1 -> 2 -> 3 -> 3 -> 4
</pre>

`p2->val != p2->next->val`, so move `p1` and `p2` to the next node
<pre>
     p1   p2
x -> 1 -> 2 -> 3 -> 3 -> 4
</pre>

`p2->val != p2->next->val`, so move `p1` and `p2` to the next node
<pre>
         p1    p2
x -> 1 -> 2 -> 3 -> 3 -> 4
</pre>

`p2->val == p2->next->val`, so move `p2` to the next node until `p2->val != p2->next->val`
<pre>
         p1             p2
x -> 1 -> 2 -> 3 -> 3 -> 4
</pre>

Link `p1` to `p2`
<pre>
         p1    p2          
x -> 1 -> 2 -> 4 -> x
</pre>

Let's look at another example<br>
head = [1,1]<br>

Use two pointers, `p1` and `p2`
<pre>
p1   p2
x -> 1 -> 1
</pre>

`p2->val == p2->next->val`, so move `p2` to the next node until `p2->val != p2->next->val`
<pre>
p1            p2
x -> 1 -> 1 -> x
</pre>

Link `p1` to `p2`
<pre>
p1    p2
x -> x
</pre>


## Complexity Analysis

### Time Complexity: O(n)

### Space Complexity: O(1)