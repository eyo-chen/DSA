# Problem Explanation

## Using Hash Table
The idea is pretty simple, we use a hash table to count the frequency of each number, then we iterate through the hash table to fill the array with the correct numbers.<br>

For example, if the input array is [2,0,2,1,1,0]<br>
Then, the hash table will be [2,1,2], which means there are 2 zeros, 1 one, and 2 twos.<br>
Then we iterate through the hash table to fill the array with the correct numbers.<br>
We just said there are 2 zeros, so we fill the first two spaces with 0, then there is 1 one, so we fill the next space with 1, then there are 2 twos, so we fill the next two spaces with 2.<br>

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)
- Even though we use a hash table, the space complexity is O(1) because the hash table will only have 3 elements (0, 1, 2)

## Loop through the array three times
The idea is that we loop through the array three times, each time we aim to put the correct number at the correct position.<br>
What does that mean?<br>
Because we're guaranteed that the input array only contains 0, 1, 2, so the first time we loop through the array, we aim to put all the 0s at the beginning of the array.<br>
The second time we loop through the array, we aim to put all the 1s at the beginning of the array, but after the 0s.<br>
The third time we loop through the array, we aim to put all the 2s at the beginning of the array, but after the 0s and 1s.<br>

For example, if the input array is [2,0,2,1,1,0], and we init a pointer at the beginning of the array<br>
The first iteration, our goal is to move all the 0s to the beginning of the array<br>
When we find first 0, we swap it with the current pointer, which is 0<br>
So, the array becomes [0,2,2,1,1,0] (swap index 0 and 1)<br>
Then, we update the pointer to the next position, which is 1. It represents position 0 is already correct<br>
Later, when we find the second 0, we swap it with the current pointer, which is 1<br>
So, the array becomes [0,0,2,1,1,2] (swap index 1 and 5)<br>
Then, we update the pointer to the next position, which is 2. It represents position 0 and 1 is already correct<br>
Now, we continue next iteration, and the goal this time is to move all the 1s to the beginning of the array, but after the 0s<br>
When we find first 1, we swap it with the current pointer, which is 2<br>
So, the array becomes [0,0,1,2,1,2] (swap index 2 and 3)<br>
Then, we update the pointer to 3.<br>
When we find the second 1, we swap it with the current pointer, which is 3<br>
So, the array becomes [0,0,1,1,2,2] (swap index 3 and 4)<br>
Then, we update the pointer to 4.<br>
Now, we continue the next iteration, and the goal this time is to move all the 2s to the beginning of the array, but after the 0s and 1s<br>

This is the whole process<br>

### Complexity Analysis
#### Time Complexity O(n)
- Although we have nested loops, the outer loop will only run 3 times, so it's seen as O(1)
- Ultimately, we just loop through the array three times, so the time complexity is O(n)
#### Space Complexity O(1)
- We don't use any extra space, so the space complexity is O(1)

## Using Two Pointers
This solution involves a lot of edge cases, so it's a bit tricky<br>
The idea is that we use two pointers to track the position of 0 and 2, `left` and `right`<br>
Also, we use a pointer to iterate through the array, let's call it `ptr`<br>
If the number is 0, we swap it with the pointer `left`<br>
If the number is 2, we swap it with the pointer `right`<br>
If the number is 1, we just move on<br>
That's the overall idea, but there are somethings we need to pay attention to<br>

When should we stop the loop?<br>
When `ptr` is greater than `right`, we should stop the loop<br>
Because `right` represents the position of 2, if `ptr` is greater than `right`, that means we've already sorted the array<br>

When should we update `left` and `right`?<br>
We should update `left` when we find 0<br>
We should update `right` when we find 2<br>

When should we update `ptr`?<br>
We should update `ptr` when we find 0 or 1<br>
But why not update `ptr` when we find 2?<br>
Let's consider this example<br>
[1,2,0], left = 0, right = 2, ptr = 0<br>
```
p
l
      r
1  2  0
```

When `ptr` is 0, we find 1, so we update `ptr` to 1<br>
```
   p
l
      r
1  2  0
```

When `ptr` is 1, we find 2, so we swap it with `right`, and update `right` and `ptr`
```
      p
l
   r
1  0  2
```
As we can see, `ptr` is greater than `right`, so we should stop the loop<br>
But the array is not sorted, which is [1,0,2]<br>

Therefore, we should only update `ptr` when we find 0 or 1<br>  
When `ptr` is 1, we find 2, so we swap it with `right`, and only update `right`<br>
```
   p
l
   r
0  1  2
```

When `ptr` is 1, we find 1, so we swap it with `left`, and update both `left` and `ptr`
```
      p
l
  r
0  1  2
```
Now, `ptr` is 2, which is greater than `right`, so we should stop the loop<br>
And the array is sorted, which is [0,1,2]<br>


Let's see an example<br>
[2,0,2,1,1,0]<br>
left = 0, right = 5, ptr = 0<br>
```
p
l              r
2  0  2  1  1  0
```

When ptr is 0, we find 2, so we swap it with right, and decrement right<br>
```
p
l           r
0  0  2  1  1  2
```
Note that we didn't update `ptr` because we don't know what number we swapped with, so we need to check it again in the next iteration<br>
Therefore, after we find 2 and swap it with right, the only thing we need to do is to decrement right<br>
Because we're only sure that the value of right pointer is definitely 2<br>
But we're not sure about if the value at left pointer is 0 or 1<br>

When ptr is 0, we find 0, so we swap it with left, and increment both left and ptr<br>
```
   p
   l        r
0  0  2  1  1  2
```
Note that we update both `left` and `ptr` because we're sure that the value at left pointer is 0<br>

When ptr is 1, we find 0, so we swap it with left, and increment both left and ptr<br>
```
      p
      l     r
0  0  2  1  1  2
```

When ptr is 2, we find 2, so we swap it with right, and decrement right<br>
```
      p
      l  r
0  0  1  1  2  2
```

When ptr is 2, we find 1, so we just increment ptr<br>
```
         p
         l
         r
0  0  1  1  2  2
```

When ptr is 3, we find 1, so we just increment ptr<br>
```
            p
            l
         r
0  0  1  1  2  2
```
Now, ptr is 4, which is greater than right, so we break the loop<br>


### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)
