# Problem Explanation

The core idea of solving this problem is divide and conquer

Input: nums = [-10,-3,0,5,9]<br>
Output: 
<pre>
             0
          /      \
        -3        9 
        /        / 
      -10       5  
</pre>

As we can see, the root node is the middle element of the array<br>
The left subtree is the middle element of the left half of the array<br>
The right subtree is the middle element of the right half of the array<br>
That's the core idea of solving this problem<br>

Let's break down even further<br>
First, we know that the root node is the middle element of the array<br>

Then, what's the root node of left subtree?<br>
Not sure yet<br>
But what we do know is that the subtree is the middle element of the left half of the array ([-10, -3])<br>
So, what's the root node of the left subtree?<br>
It basically repeats the same logic<br>
The root node is the middle element of divided array<br>
The process is just like a recursive call<br>
So does the right subtree<br>

Let's summarize the process<br>
1. Find the middle element of the array, and make it as the root node
2. Divide the array into two halves
3. The middle element of the left half is the root node of the left subtree
4. The middle element of the right half is the root node of the right subtree

Final thing to note that is how to deal with the index<br>
In this solution, we simply use `(right + left) / 2` to find the middle element<br>
Or (`left + ((right - left) / 2)`)<br>
Two formulas are the same<br>
Let's see in detail<br>
At first, `left` is 0, and `right` is 4<br>
<pre>
-10  -3  0  5  9
 0    1  2  3  4
</pre>
The middle element is at the position of `(4 + 0) / 2 = 2`<br> 
It's 0<br>
Then, we divide the array into two halves<br>
For the left half, `left` is 0, and `right` is 1 (middle - 1)<br>
For the right half, `left` is 3 (middle + 1), and `right` is 4<br>
<pre>
[-10, -3] [5, 9]
</pre>

What's the base case?<br>
When `left` is greater than `right`, we stop the process<br>
Note that when `left` is equal to `right`, we still continue the process<br>
For example, [-10, -3]<br>
`left` is 0, and `right` is 1<br>
The middle element is at the position of `(1 + 0) / 2 = 0`<br>
It's -10<br>
Then, we divide the array into two halves<br>
For the left half, `left` is 0, and `right` is -1 (middle - 1)<br>
In this case, we stop the process<br>
For the right half, `left` is 1, and `right` is 1<br>
In this case, we still continue the process<br>
Because we still need to process -3<br>


Let's walk through the example<br>
Input: nums = [-10,-3,0,5,9]<br>
1. Find the middle element of the array, and make it as the root node<br>
<pre>
             0
</pre>

2. Divide the array into two halves<br>
<pre>
[-10, -3] [5, 9]
</pre>

3. The middle element of the left half is the root node of the left subtree<br>

3-1. Find the middle element of the array, and make it as the root node<br>
<pre>
             -10
</pre>

3-2. Divide the array into two halves<br>
<pre>
[-3] []
</pre>

Now, -10 is the root node of the left subtree<br>
<pre>
             -10
                   \ 
                   -3
</pre>
<pre>
              0
            /    
          -10
              \ 
              -3
</pre>

4. The middle element of the right half is the root node of the right subtree<br>
4-1. Find the middle element of the array, and make it as the root node<br>
<pre>
             5
</pre>

4-2. Divide the array into two halves<br>
<pre>
[] [9]
</pre>

Now, 9 is the root node of the right subtree<br>
<pre>
             5
          /      \
        NULL      9
</pre>
<pre>
              0
            /   \
          -10    5
              \   \
              -3   9
</pre>

## Complexity Analysis

### Time Complexity: O(n)

### Space Complexity: O(1)