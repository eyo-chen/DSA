# Problem Explanation

Given a max-heap in array representation, return the k largest elements in the heap without performing explicit removals from the max-heap (the heap is immutable).<br>

Example:<br>
Input:<br>
heap = [17, 7, 16, 2, 3, 15, 14]<br>
k = 2<br>
Output: [17, 16]<br>

Explanation:<br>
```
[17, 7, 16, 2, 3, 15, 14]
         17
        /  \
       7    16
      / \  /  \
     2  3 15  14
```
Constraints: 0 <= k <= len(heap)


The problem is asking for the k largest elements in the heap without performing explicit removals from the max-heap.<br>

How can we solve this problem?<br>
Let's walk through the process<br>
When we're given a max heap, what can we guarantee?<br>
We can guarantee that the largest element is the root of the heap.<br>

Therefore, after picking the root<br>
```
        /  \
       7    16
      / \  /  \
     2  3 15  14
```
What's the next largest element?<br>
It's either 7 or 16.<br>
This is important!!!<br>

Whenever we find the largest element in the heap, the next largest element must be either it's left or right child.<br>
This is because the heap is a complete binary tree.<br>

So, we can maintain a list of candidate indices to check.<br>
Initially, it's just the root index.<br>
After finding the largest element, we add the left and right child of the largest element to the candidate indices.<br>
We keep doing this until we have k elements in the result list.<br>

Let's summarize the process:
1. Initialize a list of candidate indices to check. Initially, it's just the root index.
2. Find the largest element in the heap using the candidate indices.
3. Add the left and right child of the largest element to the candidate indices.
4. Repeat steps 2 and 3 until we have k elements in the result list.
5. Return the result list.


There are two approaches to implement

## Using Simple Array
Let's see what's the process looks like when using simple array
1. Initialize a list of candidate indices to check. Initially, it's just the root index.
2. Find the largest element in the heap using the candidate indices.
3. Add the largest element to the result list.
4. Remove the largest element from the candidate indices.
5. Add the left and right child of the largest element to the candidate indices.
6. Repeat steps 2 and 5 until we have k elements in the result list.
7. Return the result list.

Let's walk through the process with the following heap:<br>
[17, 7, 16, 2, 3, 15, 14]
```
         17
        /  \
       7    16
      / \  /  \
     2  3 15  14
```
Before we start, we have
- candidatesIndex: [0]
- result: []

Let's start the process
- First Iteration (i = 0)
  - curMax = 17
  - maxIdx = 0
  - maxCandidateIdx = 0
  - result = [17]
  - candidatesIndex = [1, 2]

- Second Iteration (i = 1)
  - curMax = 16
  - maxIdx = 2
  - maxCandidateIdx = 0
  - result = [17, 16]
  - candidatesIndex = [1, 5, 6]

Done!

In this approach, we need to use more variables to keep track the information we need.
- `candidatesIndex`
  - it's used to keep track the indices we need to check
  - the value is the index in the input max heap array
- `curMax`
  - it's used to keep track what's the current largest element in the candidate indices
  - note that the value inside `candidatesIndex` is the index of the input max heap array
  - it's used to put into the result list
- `maxIdx`
  - it's used to keep track the index of the current largest element in the input max heap array
  - it's used to find the left and right child of the current largest element
- `maxCandidateIdx`
  - it's used to keep track the index of the current largest element in the `candidatesIndex`
  - it's used to remove the current largest element from the `candidatesIndex`

Also, note that we have to validate the left and right child index before adding them into the `candidatesIndex`<br>
If the index is out of the bound, we skip it.<br>
This is important because we'll `nums[candidatesIndex[0]]` at the beginning of each iteration.<br>
If there is invalid index, we'll get error.

### Complexity Analysis
#### Time Complexity O(K^2)
- we have nested for loop here
- the outer for loop run k times
- the inner for loop run at most k times
  - Why is that?
  - At each iteration of outer for loop, we remove one element from the `candidatesIndex`
  - Then we add 1 or 2 element to the `candidatesIndex`
  - So the inner for loop will run at most k times
- Therefore, the time complexity is O(K^2)
- `removeElement` is also a for loop, but it's O(N)

#### Space Complexity O(K)
- the maximum size of the `candidatesIndex` is K


## Using Max Heap
This approach is similar to the previous approach<br>
The only difference is that we use a max heap to keep track the largest element<br>

Let's summarize the process:
1. Initialize a max heap with the root index.
2. Find the largest element in the heap using the candidate indices.
   - In this approach, we use a max heap to keep track the largest element
   - We can simply use `Pull` method to get the largest element
3. Add the left and right child of the largest element to the candidate indices.
   - We can simply use `Insert` method to add the left and right child to the max heap
4. Repeat steps 2 and 3 until we have k elements in the result list.
5. Return the result list.

### Complexity Analysis
#### Time Complexity O(KlogK)
- The outer loop runs k times
- The `Insert` and `Pull` method has O(logK) complexity
- Therefore, the time complexity is O(KlogK)

#### Space Complexity O(K)
- The max heap will store at most K elements








