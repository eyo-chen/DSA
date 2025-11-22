# Problem Explanation

The core idea of this solution is very similar to `oddEvenList` problem<br>
The main difference is that we now check the value of the node instead of the index of the node<br>

## Caveats
### Caveat 1: Can't directly return `head`
In the `oddEvenList` problem, we can directly return `head` because we know that the `head` will always be the first node of the list<br>
However, in this problem, we can't guarantee that the `head` will be the first node of the list<br>
Consider the following example,
[2,1], x = 2

1.
<pre>
       ptr
        |
head -> 2 -> 1 -> x

                 lessPtr
                    |
less linked list -> x
greater linked list -> x
                       |
                    greaterPtr
                    greaterHead
</pre>

2.
<pre>
            ptr
             |
head -> 2 -> 1 -> x

                 lessPtr
                    |
less linked list -> x ->
greater linked list -> x -> 2
                       |    |
                         greaterPtr
                    greaterHead
</pre>

3.
<pre>
                 ptr
                  |
head -> 2 -> 1 -> x

                      lessptr
                         |
less linked list -> x -> 1
greater linked list -> x -> 2
                       |    |
                         greaterPtr
                    greaterHead
</pre>
Now, `ptr` is null, and then we wire up the `less` and `greater` linked list<br>
`lessPtr->next = greaterHead->next`<br>
<pre>
                 ptr
                  |
head -> 2 -> 1 -> x

less linked list -> x -> 1 -> 2
</pre>
If we return `head`, we will get the following result
<pre>
head -> 2 -> 1 -> 2
</pre>
Which is wrong

So keep in mind has to keep track of the `less` and `greater` linked list's head<br>
And return `lessHead->next`<br>

## Complexity Analysis

### Time Complexity: O(n)

### Space Complexity: O(1)