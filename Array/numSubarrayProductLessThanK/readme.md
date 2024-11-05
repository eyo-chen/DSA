# Problem Explanation

## Brute Force
The brute force solution is to find all the subarrays and check if the product of the subarray is less than k.<br>
For example, for the input `[10,5,2,6]` and `k = 100`<br>
the subarrays are `[10], [10,5], [10,5,2], [10,5,2,6], [5], [5,2], [5,2,6], [2], [2,6], [6]`.<br>
Count the product of each subarray and check if it is less than `k`.<br>

### Complexity Analysis
#### Time Complexity O(n^2)
#### Space Complexity O(1)

## Sliding Window
It's not that hard to come up with the idea of using sliding window to solve this problem.<br>
The idea is to maintain a window that contains the subarray with product less than `k`.<br>

Let's see an example. For the input `[10,5,2,6]` and `k = 100`<br>
At first, the window is empty, then we keep extending the window to the right.<br>
First updating, the window is `[10]`, the product is `10`, which is less than `k`<br>
Second updating, the window is `[10,5]`, the product is `50`, which is less than `k`<br>
Third updating, the window is `[10,5,2]`, the product is `100`, which is equal to `k`, so we need to slide the window to the right until the product is less than `k`.<br>
So, we keep shrinking the window from the left until the product is less than `k`.<br>
Now, the window is `[5,2]`, the product is `10`, which is less than `k`<br>

It's pretty obvious that how to update and shrink the window and the product.<br>
But how exactly should we calculate the number of subarrays?<br>

At first, we might intutively think that whenever we update the window(expand or shrink), and it's a valid window(subarray with product less than `k`), we increment the count by `1`.<br>
But this is wrong.<br>

Let's see an example. For the input `[10,5,2,6]` and `k = 100`<br>
When the window is `[10]`, the product is `10`, which is less than `k`, so we increment the count by `1`, `count = 1`.<br>
When the window is `[10,5]`, the product is `50`, which is less than `k`, so we increment the count by `1` again, `count = 2`.<br>
When the window is `[10,5,2]`, the product is `100`, which is equal to `k`, so we need to shrink the window from the left until the product is less than `k`.<br>
When the window is `[5,2]`, the product is `10`, which is less than `k`, so we increment the count by `1` again, `count = 3`.<br>
From these above simple process, there are a lot of wrong count.<br>
For example, when the window is `[10, 5]`
```
f   s
10  5  2  6
```
How many subarrays are there?<br>
There are actually 3 subarrays: `[10], [5], [10,5]`.<br>
However, we only have 2 counts at this point.<br>
Also, when the window is `[5,2]`<br>
```
    f  s
10  5  2  6
```
At this point, it means we have explored this subarray `[10,5,2]`, and only `[5,2]` is a valid subarray.<br>
At this point, we should have 5 valid subarrays: `[10], [5], [2], [10,5], [5,2]`.<br>
However, we only have 3 counts at this point.<br>

From this example, we can see that the intutive way of counting is wrong.<br>
The correct way to count the number of subarrays is having the following mindset:<br>
Whenever we find a valid window, how many new added subarrays are there in this window?<br>
And the formula to calculate the number of new added subarrays is `fast - slow + 1`, which is basically the length of the current window.<br>
For example,<br>
When the window is `[10]`, length = 1, so `count += 1`, count = 1<br>
It means that we add 1 new subarray `[10]` to the result.<br>
When the window is `[10,5]`, length = 2, so `count += 2`, count = 3<br>
It means that we add 2 new subarrays `[10, 5], [5]` to the result.<br>
When the window is `[5,2]`, length = 2, so `count += 2`, count = 5<br>
It means that we add 2 new subarrays `[5, 2], [2]` to the result.<br>
So on and so forth.(Check out the whole process in the javascript file)

So, the key part is to find the length of the current window, and add it to the count.<br>

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)
