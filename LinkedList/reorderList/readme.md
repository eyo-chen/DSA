# Problem Explanation

This problem is kind of hard, so does the implementation

In linked list problem, <br>
***Always compare the input and output***<br>
***See how to re-wire the node to solve the problem***<br>

Note that both solution need some complex re-wiring<br>

## First Solution
I came up with this solution by myself<br>
But the time complexity is bad<br>

Let's see the example<br>
input:  [1,2,3,4,5]<br>
output: [1,5,2,4,3]<br>

It seems that the first node has to connect to the last node<br>
And last node connects to the second node<br>
f = first node<br>
l = last node<br>
<pre>
f                   l
1    2 -> 3 -> 4 -> 5
  \ /
   5
   l
</pre>
And now we have to update the first node to `2`
And update the last node to `4`<br>

We seems to have a pattern here<br>
1. First node connects to the last node
2. Last node connects to the second node
3. Update first node to the second node
4. Update last node to the second last node

This is just the high level idea<br>
We have two details to note<br>
1. We somehow have to set the next of node before tail to `nullptr`
  - In the above example, we have to manually set `4->next` to `nullptr`
  - So that `4` becomes real tail
2. We have to reference the next node before we update the node
  - In the above example, we have to reference `2` before we update `1` to `5`
  - Without this, we will lose the reference to `2` after re-wiring

Finally, let's see one edge case<br>
1. Two nodes
<pre>
f    l
1 -> 2
</pre>
`curNextNode` is `2`<br>
`tail` is also `2`<br>
We know we have to stop here<br>

Let's see all the steps<br>
1. Find the tail, and the node before tail
2. Reference the current next node
3. Check if `curNextNode` is `tail`
4. Set `ptr->next` to `tail`
5. Set `tail->next` to `curNextNode`
6. Set `nodeBeforeTail->next` to `nullptr`
7. Update `ptr` to `curNextNode`

Input:  [1,2,3,4,5]<br>
p = ptr<br>
l = tail<br>
bl = nodeBeforeTail<br>
n = curNextNode<br>
1. Find the tail, and the node before tail
<pre>
p              bl   l
1 -> 2 -> 3 -> 4 -> 5
</pre>

2. Reference the current next node
<pre>
p    n         bl   l
1 -> 2 -> 3 -> 4 -> 5
</pre>

3. Check if `curNextNode` is `tail`
4. Set `ptr->next` to `tail`
<pre>
p    n         bl   l
1   2 -> 3 -> 4 -> 5
 \ --------------- /
</pre>

5. Set `tail->next` to `curNextNode`
<pre>
p    n         bl   l
1   2 -> 3 -> 4 -> 5
 \  /
  5
</pre>

6. Set `nodeBeforeTail->next` to `nullptr`
<pre>
p    n        bl  
1   2 -> 3 -> 4 -> x
 \  /
  5
</pre>

7. Update `ptr` to `curNextNode`
<pre>
    p
1   2 -> 3 -> 4 -> x
 \  /
  5
</pre>

Again!!<br>

1. Find the tail, and the node before tail
<pre>
          p   bl    l
1 -> 5 -> 2 -> 3 -> 4 -> x
</pre>

2. Reference the current next node
<pre>
               n   
          p    bl   l
1 -> 5 -> 2 -> 3 -> 4 -> x
</pre>

3. Check if `curNextNode` is `tail`
4. Set `ptr->next` to `tail`
<pre>
               bl   
          p    n    l
1 -> 5 -> 2   3 -> 4 -> x
          \ ------- /
</pre>

5. Set `tail->next` to `curNextNode`
<pre>
               bl   
          p    n    l
1 -> 5 -> 2   3 -> 4 -> x
          \    /
            4 
</pre>

6. Set `nodeBeforeTail->next` to `nullptr`
<pre>
               bl   
          p    n    
1 -> 5 -> 2   3 -> x
          \    /
            4 
</pre>

7. Update `ptr` to `curNextNode`
<pre>
                    p
1 -> 5 -> 2 -> 4 -> 3 -> x
</pre>
Done!!

### Complexity Analysis

#### Time Complexity: O(n ^ 2)
- We have to find the tail and the node before tail
  - So there's a nested loop

#### Space Complexity: O(1)

## Second Solution
This solution is kind of tricky<br>
Let's look at the example<br>
input:  [1,2,3,4,5]<br>
output: [1,5,2,4,3]<br>

If we cut the list in the middle, we have two lists<br>
1. [1,2,3]<br>
2. [4,5]<br>

We reverse the second list<br>
1. [1,2,3]<br>
2. [5,4]<br>

And then we merge the two lists<br>
[1,5,2,4,3]<br>
This is the answer<br>

So the high level idea is<br>
1. Find the middle of the list
2. Reverse the second list
3. Merge the two lists

Let's walk through the example<br>
input:  [1,2,3,4,5]<br>
1. Find the middle of the list
<pre>
          s
                    f
1 -> 2 -> 3 -> 4 -> 5
</pre>
`s` is at the middle<br>
Now, we have two lists<br>
[1,2,3]<br>
[4,5] (s->next is the second list)<br>

2. Reverse the second list
<pre>
          s
1 -> 2 -> 3 -> 4 

p      
5 -> 4 -> x 
</pre>
`p` is the previous node after reversing<br>
At this point, `s->next` still points to `4`<br>
So we have to set `s->next` to `nullptr`<br>

3. Merge the two lists<br>
f = firstList<br>
s = secondList<br>
fn = firstNext<br>
sn = secondNext<br>

(1) reference the next node<br>
<pre>
f   fn
1 -> 2 -> 3 -> x
5 -> 4 -> x
s    sn
</pre>

(2) re-wire the node<br>
`f->next = s`<br>
`s->next = fn`<br>
<pre>
f   fn
1   2 -> 3 -> x
|  /
5 / 4 -> x
s   sn
</pre>

(3) update the pointer<br>
<pre>
    f
1   2 -> 3 -> x
|  /
5 / 4 -> x
    s
</pre>

(1) reference the next node<br>
<pre>
    f   fn
1   2 -> 3 -> x
|  /
5 / 4 -> x
    s   sn
</pre>
  
(2) re-wire the node<br>
`fn->next = s`<br>
`s->next = fn`<br>
<pre>
    f   fn
1   2    3 -> x
|  /  |  /
5 /   4 / x
      s    sn
</pre>

We are done<br>
<pre>
1 -> 5 -> 2 -> 4 -> 3 -> x
</pre>

### Complexity Analysis

#### Time Complexity: O(n)
#### Space Complexity: O(1)

