# Problem Explanation

Input: l1 = [7,2,4,3], l2 = [5,6,4]<br>
Output: [7,8,0,7]
<pre>
7 -> 2 -> 4 -> 3
     5 -> 6 -> 4
7 -> 8 -> 0 -> 7
</pre>

The hard part of this problem is<br>

How to handle the carry?

Look at the above example<br>
If we just build the regular linked list, we can't handle the carry easily<br>
4 + 6 = 10<br>
And we need to carry 1 to the previous node<br>
That's the hard part of this problem<br>

## First Solution: Using Stack
If we can add the number from the end, we can handle the carry easily<br>
Use above example<br>
We can add 3 and 4 first to get 7<br>
Then we can add 4 and 6 to get 10, now we just need to carry 1 to the next node<br>
Then we can add 2, 5 and carry 1 to get 8<br>
Then we can add 7 and 0 to get 7<br>

So, we can use stack to help us<br>
Because stack is LIFO (Last In First Out)<br>
We can add the number to the stack first, then pop the number from the stack to get the number from the end<br>

The stpes are:<br>
1. Create a stack for l1 and l2
2. Pop the number from the stack and add them together
3. If the sum is greater than 9, we need to carry 1 to the next node
4. Create a new node to store the sum
5. Move to the next node

Let's walk through the example<br>
1. Create a stack for l1 and l2
<pre>
s1 = [7,2,4,3]
s2 = [5,6,4]
</pre>

2. Pop the number from the stack and add them together
<pre>
3 + 4 = 7
4 + 6 = 10 (carry 1)
2 + 5 + 1 = 8
7 + 0 = 7
</pre>

3. Create a new node to store the sum
4. Move to the next node
<pre>
7 -> 8 -> 0 -> 7
</pre>

One note is that we need to build the linked list backward<br>
For example,<br>
Suppose we've created the first node 7<br>
Then we need to create the second node 0 and point it to the first node 7<br>
Then we need to create the third node 8 and point it to the second node 0<br>
Then we need to create the fourth node 7 and point it to the third node 8<br>

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(n)
 - We need to use stack to store the value


## Second Solution: Without Using Stack
Let's think about<br>
If we can't add the value from the end, can we still handle the carry?<br>
The answer is no<br>

When we add the value of 4 and 6<br>
We can't cary 1 to the previous node<br>

But how about we don't care about the carry at all?<br>
Then we just add the value in the regular way<br>
Let's see<br>
7 + 0 = 7<br>
2 + 5 = 7<br>
4 + 6 = 10<br>
3 + 4 = 7<br>

Okay, now we get the each value of the node<br>
How should we connect them together?<br>
First order<br>
7a -> 7b -> 10 -> 7c<br>
Second order<br>
7c -> 10 -> 7b -> 7a<br>

Let's look at the output<br>
7 -> 8 -> 0 -> 7<br>

Let's try the first order<br>
7a -> 7b -> 10 -> 7c<br>
We know the carry 1 need to be added to the 7b<br>
But we can't do that<br>
When we access 10, we can't back to 7b<br>

Let's try the second order<br>
7c -> 10 -> 7b -> 7a<br>
We know the carry 1 need to be added to the 7b<br>
We can do that<br>
When we access 10, 7b is the next node<br>
That's easy to handle the carry<br>
For this order, we need to reverse the linked list to get the correct output<br>

So, after we get the value of each node<br>
We build the linked list backward<br>
It's similar to the first solution<br>
For example,<br>
Suppose we've created the first node 7<br>
Then we need to create the second node 7, and point it to the first node 7<br>
Then we need to create the third node 10, and point it to the second node 7<br>
Then we need to create the fourth node 7, and point it to the third node 10<br>

Note that in order to add the value in the regular way<br>
We need to know the length of the linked list<br>
And which linked list is longer and shorter<br>
Then we can add the value in the regular way<br>

Let's summarize the steps<br>
1. Get the length of l1 and l2
2. Find the longer and shorter linked list
3. Add the value in the regular way, and build the linked list backward without considering the carry
4. Reverse the linked list, and handle the carry

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(1)
 - We don't need to use extra space to store the value


## Caveat
Both solution have one caveat<br>
Which is we need to handle the edge case<br>
When we build the final linked list<br>

Let's see one example<br>
l1 = [5]<br>
l2 = [5]<br>
The output should be [0,1]<br>

The first iteration is 5 + 5 = 10<br>
We know the carry 1 need to be added to the next node<br>
But there is no next node<br>
So we just end the iteration<br>

If we don't handle this case<br>
We will just return [0]<br>

So, we need to handle this case<br>
When we build the final linked list<br>
If there is a carry 1<br>
We need to create a new node to store the carry 1<br>
And connect it<br>
