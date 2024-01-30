# Problem Explanation

The solution of this problem is not about the process of rotating the list<br>
***It's about what's the result looks like after rotating the list***<br>

For example,<br>
[1,2,3,4,5], k = 1<br>
The result should be [5,1,2,3,4]<br>

[1,2,3,4,5], k = 2<br>
The result should be [4,5,1,2,3]<br>

[1,2,3,4,5], k = 3<br>
The result should be [3,4,5,1,2]<br>

[1,2,3,4,5], k = 4<br>
The result should be [2,3,4,5,1]<br>

[1,2,3,4,5], k = 5<br>
The result should be [1,2,3,4,5]<br>

Look at each result, and really think about what exactly we really care about<br>
- The first node of the result
- The last node of the result

That's it!<br>

If we can know the head and tail of the rotated list, we solve the problem<br>
Again, it's all about re-wiring the pointers(node)<br>

Let's first to think about how to find the head of the rotated list<br>
[1,2,3,4,5], k = 1<br>
result = [5,1,2,3,4]<br>
Head is 5<br>

[1,2,3,4,5], k = 2<br>
result = [4,5,1,2,3]<br>
Head is 4<br>

[1,2,3,4,5], k = 3<br>
result = [3,4,5,1,2]<br>
Head is 3<br>

[1,2,3,4,5], k = 4<br>
result = [2,3,4,5,1]<br>
Head is 2<br>

[1,2,3,4,5], k = 5<br>
result = [1,2,3,4,5]<br>
Head is 1<br>

We can easily find the pattern<br>
Head is the node at the position of `length - k` at the original list<br>
(`length - k` means the distance from the head to the position)<br>

How about the tail?<br>
It's easy!
Because tail is just the node before the head<br>
(Again, it's at the original list)<br>

What does that mean?<br>
It means as long as we can find the tail of new rotated list at the original list<br>
We know the head of the new rotated list too<br>
Because `newHead = newTail->next`<br>

So, now the problem is how to find the tail of the new rotated list<br>
Look at the above example,
Tail is the node at the position of `length - k - 1` at the original list<br>
(`length - k - 1` means the distance from the head to the position)<br>

Now, we can know the step to solve this problem<br>
1. Find the length of the list
  - Without length, we can't find the head and tail
  - Look above example, the new head and tail are all based on the length of the list
2. Find the tail of the new rotated list
3. Find the head of the new rotated list

However, finding tail and head is not enough<br>
The linked list won't change unless we re-wire the pointers<br>
So, we have two more steps to do (all re-wire the pointers)<br>
4. Re-wire the head and tail at the original list<br>
5. Remove the connection between the tail and the head of the new rotated list

Final thing to note is that when k is larger than the length of the list<br>
We need to use `k % length` to get the real k(rotation index)<br>

## Complexity Analysis

### Time Complexity: O(n)

### Space Complexity: O(1)