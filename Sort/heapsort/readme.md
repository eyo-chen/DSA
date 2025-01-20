# Problem Explanation

The core idea of heap sort is to build a max heap and then sort the array by repeatedly extracting the maximum element from the heap.<br>

The core idea of heap sort involves two main phases:
1. Building a max heap from the input array.
2. Extracting the maximum element from the heap and placing it at the end of the array, then heapifying the remaining elements to maintain the heap property.


How to build a max heap?<br>
1. Conceptually convert the array into a complete binary tree.
2. Start from the last non-leaf node(middle of the array) and move towards the root of the tree.
3. For each node, check if it is greater than its children. If not, swap it with the larger child. After swapping, do the same thing for the next tree level.
4. Continue this process until the entire array is converted into a max heap.

Let's see an example:<br>
[4,5,8,10,1]<br>
Convert the array into a complete binary tree.<br>
```
             4
        5        8
      10   1
```
Start from the last non-leaf node(middle of the array) and move towards the root of the tree.<br>
Why do we start from the middle of the array?<br>
Because if we look at all the leaf nodes as binary trees, they are already max heaps<br>
The definition of a heap is that the parent node should be greater than its children.<br>
Since the leaf nodes have no children, they are already in a max heap state.<br>
So we can start from the middle of the array and move towards the root of the tree.<br>

node = 5, check if it is greater than its children.<br>
No, it is not greater than its children. So we swap it with the larger child.(swap 5 and 10)
```
             4
        10        8
      5   1
```
arr becomes [4,10,8,5,1]

node = 4, check if it is greater than its children.<br>
No, it is not greater than its children. So we swap it with the larger child.(swap 4 and 10)
```
             10
        4        8
      5   1
```
arr becomes [10,4,8,5,1]<br>
Now, the array is a valid max heap.<br>

Let's focus on the process of maintaining the heap property.<br>
Basically, when we're given a node(index), all we want to do is to make sure it's at the correct position where it's greater than its children.<br>
If it's a large value, we probably don't need to swap it with its children.<br>
If it's a small value, we might need to move this all the way down to the bottom of the tree.<br>
This process is called heapify.<br>


Next, let's look at the process of sorting the array.<br>
The idea is pretty straightforward.<br>
Once we have a valid max heap, the largest value is at the root of the tree.<br>
So, we can swap the root with the last element of the array.<br>
After swapping, two things happen:<br>
1. We for sure know that the last element is the largest value in the array(which means it's sorted)
2. The heap property is violated because the root is no longer the largest value.

To fix this, we can call heapify on the root of the tree.<br>
Which is the same process as we discussed earlier.<br>
We try to move the root value down to the correct position where it's greater than its children.<br>

After doing this, the size of the heap is reduced by 1 because the last element is already sorted.<br>
We repeat this process until the heap is empty.<br>

# Complexity Analysis
## Time Complexity O(nlogn)
- We can break down the time complexity into two parts:
  - Building a max heap: O(n)
    - Start from the last non-leaf node and apply the `heapify` operation to ensure that each subtree satisfies the max heap property.
    - Heapify Operation: For a single node, the `heapify` function may traverse from the node down to the leaf nodes. In the worst case, this traversal cost is proportional to the height of the heap, which is O(log n) for a binary heap.
    - Number of Nodes to Heapify: Only the non-leaf nodes need to be heapified. In a binary heap represented as an array of size n, there are approximately n/2 non-leaf nodes.
    - While it might seem that heapifying n/2 nodes each costing O(log n) would result in O(n log n) time, the actual time complexity is O(n).
    - Most of the nodes are near the bottom of the heap and require fewer operations. The deeper a node is in the heap, the fewer levels it has to traverse during heapification.
    - The total number of operations across all nodes sums up to a linear function of n.
  - Extracting the maximum element from the heap and placing it at the end of the array: O(nlogn)
    - Iteratively remove the maximum element (root of the heap) and rebuild the heap with the remaining elements.
    - Number of Extract Operations: There are n elements to extract.
    - Heapify After Each Extraction: After removing the root, the heapify function is called to maintain the heap property, which costs O(log n) time per extraction.
    - Total Time: n extractions O(log n) heapify time = O(n log n)

## Space Complexity O(1)
- An in-place algorithm transforms the input using a constant amount of extra space.
- Heap Sort's In-Place Nature:
  - Heap Sort organizes the input array to represent a heap without needing additional arrays or data structures.
  - The sorting process involves swapping elements within the original array to move the maximum elements to their correct positions.
- Conclusion: Heap Sort is an in-place sorting algorithm.



