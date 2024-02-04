# Problem Explanation

The solution is straightforward once drawing the problem on paper<br>
I came up with this solution by myself<br>

Let's see the example<br>
input:  [1,2,3,4]<br>
output: [2,1,4,3]<br>

If we split the list into two parts<br>
1. [1,2]<br>
2. [3,4]<br>

For each parts, we just reverse the list<br>
And then we may solve the problem<br>

Let's see the first part as an example<br>
c = current node<br>
n = next node<br>
n = next next node<br>
<pre>
c    n
1 -> 2 -> 3 -> 4 -> x
</pre>
How to reverse the list?<br>
1. Reference the current next next node<br>
<pre>
c    n   nn
1 -> 2 -> 3 -> 4 -> x
</pre>
2. Make current node to point to next next node<br>
<pre>
c    n   nn
1    2 -> 3 -> 4 -> x
 \--------/
</pre>
3. Make next node to point to current node<br>
<pre>
c    n   nn
1 <- 2    3 -> 4 -> x
 \--------/
</pre>

Output: 2 -> 1 -> 3 -> 4 -> x<br>

Now, it seems straightforward that `c` has to update to `3`(`nn`)<br> 
<pre>
          c
2 -> 1 -> 3 -> 4 -> x
</pre>
Let's do the re-wiring againg

1. Reference the current next next node<br>
<pre>
          c    n   nn
2 -> 1 -> 3 -> 4 -> x
</pre>
2. Make current node to point to next next node<br>
<pre>
          c    n   nn
2 -> 1 -> 3    4 -> x
          \--------/
</pre>
3. Make next node to point to current node<br>
<pre>
          c    n   nn
2 -> 1 -> 3 <- 4    x
          \--------/
</pre>
Can you see one problem here?<br>
Though we have reversed the [3,4] to [4,3]<br>
Now, `1` still points to `3`<br>

We need some technique to make sure that `1` points to `4`<br>
How?<br>
If we have a reference on the previous node, we can easily re-wire the list<br>
We just make sure that `prev` points to `next`<br>
<pre>
     p    c    n   nn
2 -> 1    3 <- 4    x
       \  \--------/
        \------/  
</pre>
Output: 2 -> 1 -> 4 -> 3 -> x<br>
We're done!!!<br>

Let' see the complete pattern<br>
1. Reference the current next node, and next next node<br>
2. Make current node to point to next next node<br>
3. Make next node to point to current node<br>
4. Make prev node to point to next node<br>
5. Update prev to current node<br>
6. Update current node to next next node<br>

Note that we need dummy head because we need to set up `prev`<br>

Let's see the even example<br>
input:  [1,2,3,4]<br>
1. Reference the current next node, and next next node<br>
<pre>
p    c    n   nn
x -> 1 -> 2 -> 3 -> 4 -> x
</pre>
2. Make current node to point to next next node<br>
<pre>
p    c    n   nn
x -> 1    2 -> 3 -> 4 -> x
     \--------/
</pre>
3. Make next node to point to current node<br>
<pre>
p    c    n   nn
x -> 1 <- 2    3 -> 4 -> x
     \--------/
</pre>
4. Make prev node to point to next node<br>
<pre>
p    c    n   nn
x    1 <- 2    3 -> 4 -> x
   \ \--------/
    \----/
</pre>
Let's organize the diagram<br>
<pre>
p    n    c   nn
x -> 2 -> 1 -> 3 -> 4 -> x
</pre>
5. Update prev to current node<br>
<pre>
          p
x -> 2 -> 1 -> 3 -> 4 -> x
</pre>
6. Update current node to next next node<br>
<pre>
          p    c
x -> 2 -> 1 -> 3 -> 4 -> x
</pre>

Again,
1. Reference the current next node, and next next node<br>
<pre>
          p    c    n   nn
x -> 2 -> 1 -> 3 -> 4 -> x
</pre>
2. Make current node to point to next next node<br>
<pre>
          p    c    n   nn
x -> 2 -> 1 -> 3    4 -> x
                \--------/
</pre>
3. Make next node to point to current node<br>
<pre>
          p    c    n   nn
x -> 2 -> 1 -> 3 <- 4    x
                \--------/
</pre>
4. Make prev node to point to next node<br>
<pre>
          p    c    n   nn
x -> 2 -> 1    3 <- 4    x
           \    \--------/
            \------/   
</pre>
Let's organize the diagram<br>
<pre>
          p    n    c   nn
x -> 2 -> 1 -> 4 -> 3 -> x
</pre>
5. Update prev to current node<br>
<pre>
                    p
x -> 2 -> 1 -> 4 -> 3 -> x
</pre>
6. Update current node to next next node<br>
<pre>
                    p    c
x -> 2 -> 1 -> 4 -> 3 -> x
</pre>
We're done!!!<br>

Let's see the odd example<br>
input:  [1,2,3]<br>
1. Reference the current next node, and next next node<br>
<pre>
p    c    n   nn
x -> 1 -> 2 -> 3 -> x
</pre>
2. Make current node to point to next next node<br>
<pre>
p    c    n   nn
x -> 1    2 -> 3 -> x
     \--------/
</pre>
3. Make next node to point to current node<br>
<pre>
p    c    n   nn
x -> 1 <- 2    3 -> x
     \--------/
</pre>
4. Make prev node to point to next node<br>
<pre>
p    c    n   nn
x    1 <- 2    3 -> x
   \ \--------/
    \----/
</pre>
Let's organize the diagram<br>
<pre>
p    n    c   nn
x -> 2 -> 1 -> 3 -> x
</pre>
5. Update prev to current node<br>
<pre>
          p
x -> 2 -> 1 -> 3 -> x
</pre>
6. Update current node to next next node<br>
<pre>
          p    c
x -> 2 -> 1 -> 3 -> x
</pre>
We're done!!!<br>

This solution works on when input is nullptr, one node and two nodes<br>
Perfect!!!<br>

Note that Javascript has another solution, but i personally prefer this solution<br>

## Complexity Analysis
### Time Complexity: O(n)
- We have to visit all nodes in the list
### Space Complexity: O(1)
- We don't use any extra space