# Problem Explanation

This solution is pretty straightforward<br>

1. Use two pointers, `p1` and `p2`
2. `p2` is ahead of `p1` by one node
3. If `p1` and `p2` have the same value
  - `p2` is moved to the next node
  - `p1` is linked to `p2`
4. If `p1` and `p2` have different values
  - `p1` is moved to the next node
  - `p2` is moved to the next node

For example,<br>
head = [1,1,2,3,3]<br>

Use two pointers, `p1` and `p2`, `p2` is ahead of `p1` by one node
<pre>
p1   p2
1 -> 1 -> 2 -> 3 -> 3
</pre>

p1->val = p2->val<br>
so move p2 to the next node<br>
and link p1 to p2
<pre>
p1       p2
1 \  1 -> 2 -> 3 -> 3
   - - - /
</pre>

p1->val != p2->val<br>
so move p1 and p2 to the next node
<pre>
         p1   p2         
1 \  1 -> 2 -> 3 -> 3
   - - - /
</pre>

p1->val != p2->val<br>
so move p1 and p2 to the next node
<pre>
              p1   p2
1 \  1 -> 2 -> 3 -> 3
    - - - /
</pre>

p1->val = p2->val<br>
so move p2 to the next node<br>
and link p1 to p2
<pre>
              p1         p2
1 \  1 -> 2 -> 3 \ 3 -> x
    - - - /       - - - /  
</pre>

output: 1 -> 2 -> 3 -> x

## Complexity Analysis

### Time Complexity: O(n)

### Space Complexity: O(1)