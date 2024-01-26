# Problem Explanation

The core idea of this solution is thinking about the ***edge case***

It's pretty straightforward to use two pointers and sliding window to solve this problem<br>
But have to also think about the edge case<br>

Let's walk through the thought process of solving this problem<br>
Consider the first example <br>
head = [1,2,3,4,5], n = 2<br>
We know the ouput should be [1,2,3,5]<br>
If we want to remove 4, we have to know the previous node of 4, which is 3<br>
<pre>
l
r
1 -> 2 -> 3 -> 4 -> 5 -> x
</pre>
We first assume that `l` and `r` pointer start at the beginning of the list<br>
Then we move `r` pointer to the right by `n` steps<br>
<pre>
l
          r
1 -> 2 -> 3 -> 4 -> 5 -> x
</pre>
Now, we move both `l` and `r` pointer to the end of the list<br>
<pre>
          l
                    r
1 -> 2 -> 3 -> 4 -> 5 -> x
</pre>
Now, `l` is at the 3, which is the previous node of 4<br>
Does it mean we always move `r` pointer to the position of `r->next == null`?<br>
It may seems true in this example<br>

Let's consider another example<br>
head = [1,2], n = 2<br>
We know the ouput should be [2]<br>
If we want to remove 1, we have to know the previous node of 1<br>
But what's the previous node of 1?<br>
It's dummy node<br>

***Using dummy node to deal with the edge case*** is the core idea to solve this problem<br>

Let's use dummy node to solve this example<br>
<pre>
dummy -> 1 -> 2 -> x
</pre>
Okay, but where should we put the `l` and `r` pointer?<br>
It's obvious that we should put `l` pointer at the dummy node<br>
Because we want to know the previous node of the node we want to remove<br>
Okay, let's also assume that `r` is at the dummy node<br>
<pre>
l
r
dummy -> 1 -> 2 -> x
</pre>
Move `r` pointer to the right by `n` steps<br>
<pre>
l
              r
dummy -> 1 -> 2 -> x
</pre>
Great, it seems we can use `r->next == null` to determine the end of the list<br>

So now we know four steps to solve this problem<br>
1. Use dummy node to deal with the edge case
2. Place `l` and `r` pointer at the dummy node
3. Move `r` pointer to the right by `n` steps
4. Use `r->next == null` to determine the end of the list

Let's confirm these steps also work on another edge case
head = [1], n = 1<br>
<pre>
l
r
dummy -> 1 -> x
</pre>
Move `r` pointer to the right by `n` steps<br>
<pre>
l
         r
dummy -> 1 -> x
</pre>
Now, `r` pointer is at the end of the list<br>
We can remove the node after `l` pointer<br>
Great, it works on another edge case<br>

Let's try on the first example<br>
head = [1,2,3,4,5], n = 2<br>
<pre>
l
r
dummy -> 1 -> 2 -> 3 -> 4 -> 5 -> x
</pre>
Move `r` pointer to the right by `n` steps<br>
<pre>
l
              r
dummy -> 1 -> 2 -> 3 -> 4 -> 5 -> x
</pre>
Now, `r` pointer is at the end of the list(r->next == null)<br>
<pre>
                   l
                             r
dummy -> 1 -> 2 -> 3 -> 4 -> 5 -> x
</pre>
We can remove the node after `l` pointer<br>
Great, it works on the first example<br>

Now, we're pretty sure that these four steps can solve this problem<br>

Note that we can also put `r` at the head<br>
And use `r == null` to determine the end of the list<br>
The logic is the same<br>
Just use different condition to determine the end of the list<br>
Look at the solution 2

## Another Solution: Count the length of the list
This solution is even more straightforward<br>
If we know the length of the list, we can easily calculate the position of the node we want to remove<br>
Then we can remove the node<br>

This solution also need the help of dummy node<br>
There are four steps to solve this problem<br>
1. Count the length of the list
2. Count the distance from the dummy node to the previous node we want to remove(k = length - n)
3. Place working pointer at the dummy node
4. Move the working pointer to the right by k steps
5. Now, we can remove the node after the working pointer

First example, head = [1,2,3,4,5], n = 2<br>
1. Count the length of the list<br>
   length = 5
2. Count the distance from the dummy node to the previous node we want to remove(k)<br>
    k = 5 - 2 = 3
3. Place working pointer at the dummy node<br>
<pre>
p
dummy -> 1 -> 2 -> 3 -> 4 -> 5 -> x
</pre>
4. Move the working pointer to the right by k steps<br>
<pre>
                   p
dummy -> 1 -> 2 -> 3 -> 4 -> 5 -> x
</pre>
5. Now, we can remove the node after the working pointer<br>

Second example, head = [1], n = 1<br>
1. Count the length of the list<br>
   length = 1
2. Count the distance from the dummy node to the previous node we want to remove(k)<br>
    k = 1 - 1 = 0
3. Place working pointer at the dummy node<br>
<pre>
p
dummy -> 1 -> x
</pre>
4. Move the working pointer to the right by k steps<br>
<pre>
p
dummy -> 1 -> x
</pre>
5. Now, we can remove the node after the working pointer<br>

Third example, head = [1,2], n = 2<br>
1. Count the length of the list<br>
   length = 2
2. Count the distance from the dummy node to the previous node we want to remove(k)<br>
    k = 2 - 2 = 0
3. Place working pointer at the dummy node<br>
<pre>
p
dummy -> 1 -> 2 -> x
</pre>
4. Move the working pointer to the right by k steps<br>
<pre>
p
dummy -> 1 -> 2 -> x
</pre>
5. Now, we can remove the node after the working pointer<br>

## Complexity Analysis

### Time Complexity: O(n)

### Space Complexity: O(1)