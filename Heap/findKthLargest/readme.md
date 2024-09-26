# Problem Explanation

## The Most Intuitive Way
This is the most intuitive way to solve the problem.<br>
We can use nested for loop to find the kth largest element in the array.<br>
The outer for loop is to swap the largest element to the front of the array.<br>
The inner for loop is to find the largest element from (start ~ n-1)

At the end of the nested loop, the kth largest element will be the kth element in the array.

### Complexity Analysis
#### Time Complexity O(K*N)
- The outer for loop runs k times
- The inner for loop runs n times on average
- Therefore, the time complexity is O(K*N)

#### Space Complexity O(1)
- We don't use any extra space

## Using a Max Heap
The idea is to build a max heap, and then pull the root of the max heap k times.

### Complexity Analysis
#### Time Complexity O(N + KlogN)
- Building a max heap takes O(N) time
- Pulling the root of the max heap takes O(KlogN) time
  - Each pull operation takes O(logN) time
- Therefore, the time complexity is O(N + KlogN)

#### Space Complexity O(N)
- We need to store the max heap, which takes O(N) space

## Using a Min Heap
This is kind of the opposite mindset to solve the problem.<br>
We can use a min heap to store the kth largest elements in the array.<br>
The idea is to maintain a min heap of size k, so the root of the min heap is the kth largest element in the array.<br>
We can do this by inserting the elements into the min heap, and if the size of the min heap is greater than k, we pull the root of the min heap.<br>
This is working because every time we pull the root of the min heap, the root is the smallest element in the heap.<br>
So we can guarantee that the root is the kth largest element in the array.

Let's say we have the following array: [5, 3, 6, 2, 4, 1] and k = 2

1. Insert 5 into the min heap, the min heap is now [5]
```
5
```
2. Insert 3 into the min heap, the min heap is now [3, 5]
```
   3
5
```
3. Insert 6 into the min heap, the min heap is now [3, 5, 6]
```
   3
5    6
```
4. Since the size of the min heap is greater than k, we pull the root of the min heap, the min heap is now [5, 6]
```
   5
6
```
5. Insert 2 into the min heap, the min heap is now [2, 5, 6]
```
   2
5    6
```
6. Since the size of the min heap is greater than k, we pull the root of the min heap, the min heap is now [5, 6]
```
   5
6
```
7. Insert 4 into the min heap, the min heap is now [4, 5, 6]
```
   4
5    6
```
8. Since the size of the min heap is greater than k, we pull the root of the min heap, the min heap is now [5, 6]
```
   5
6
```
9. Insert 1 into the min heap, the min heap is now [1, 5, 6]
```
   1
5    6
```
10. Since the size of the min heap is greater than k, we pull the root of the min heap, the min heap is now [5, 6]
```
   5
6
```
At the end, the root of the min heap is the kth largest element in the array.

### Complexity Analysis
#### Time Complexity O(NlogK)
- It takes O(N) time to loop through the array
- Inserting an element into the min heap takes O(logK) time
- Pulling the root of the min heap takes O(KlogK) time
- Therefore, the time complexity is O(NlogK)

#### Space Complexity O(K)
- We need to store the min heap, which takes O(K) space
