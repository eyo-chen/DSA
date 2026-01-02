# Heap

In simple sense, we can conceptually think of Heap as a complete binary tree. A complete binary tree is a binary tree in which all the levels are completely filled except possibly the last level and the last level has all keys as left as possible.<br>
- In a max-heap, the parent node is always greater than or equal to its child nodes. 
- In a min-heap, the parent node is always less than or equal to its child nodes.
So, the root node of a max-heap contains the maximum value, while the root node of a min-heap contains the minimum value.<br>

In reality, heaps are often implemented using arrays for efficient storage and access. The relationship between the indices of parent and child nodes can be easily calculated using simple arithmetic operations.<br>
It's very easy to figure out the formula to find parent, left child, and right child nodes using the following example.<br>

Let's see one example of min-heap implemented using array:<br>
```
                0
           1        2
        3     4  5    6
```
In the above min-heap, we can represent it using an array as follows:<br>
```
[0,1,2,3,4,5,6]
```

How can we find the parent node?<br>
-> For any node at index `i`, the parent node can be found at index `(i-1)/2`<br>

How can we find the left child node?<br>
-> For any node at index `i`, the left child node can be found at index `2*i + 1`<br>

How can we find the right child node?<br>
-> For any node at index `i`, the right child node can be found at index `2*i + 2`<br>


## Properties of Heap
1. A heap is a complete binary tree.
2. The height of a heap is log(n), where n is the number of nodes in the heap.

## Insertion in Heap
To insert a new element in a heap, we follow these steps:
1. Add the new element at the end of the heap(array) (maintaining the complete tree property).
2. Compare the added element with its parent; if the added element is greater (in max-heap) or smaller (in min-heap) than its parent, swap them.
3. Repeat step 2 until the heap property is restored.

Let's use above example of min-heap to illustrate the insertion process.<br>
Suppose we want to insert the value `-1` into the min-heap represented by the array `[0,1,2,3,4,5,6]`.<br
1. We add `-1` at the end of the array: `[0,1,2,3,4,5,6,-1]`
```
                0
           1        2
        3     4  5    6
      -1
```

2. We compare `-1` with its parent `3`. Since `-1` is smaller, we swap them: `[0,1,2,-1,4,5,6,3]`
```
                0
           1        2
        -1    4  5    6
      3
```

3. We compare `-1` with its new parent `1`. Since `-1` is smaller, we swap them: `[0,-1,2,1,4,5,6,3]`
```
                0
           -1       2
        1     4  5    6
      3
```

4. We compare `-1` with its new parent `0`. Since `-1` is smaller, we swap them: `[-1,0,2,1,4,5,6,3]`
```
               -1
           0        2
        1     4  5    6
      3
```

### Complexity Analysis
#### Time Complexity O(log n)
- Insertion takes O(log n) time in the worst case because we may need to traverse the height of the tree to restore the heap property.  

## Deletion(Get the max/min value) in Heap
To delete or get the root element (the maximum in max-heap or minimum in min-heap), we follow these steps:
1. Store the root element to return later.(The root value is the max/min value)
2. Replace the root element with the last element in the heap(array).(Basically just swap the root and last element)
3. Remove the last element from the heap(array).
4. Compare the new root element with its children; if it is smaller (in max-heap) or larger (in min-heap) than either of its children, swap it with the larger (in max-heap) or smaller (in min-heap) child. (Note that we swap with only one child, the larger/smaller one)
5. Repeat step 4 until the heap property is restored.

Let's use above example of min-heap to illustrate the deletion process.<br>
Suppose we want to delete the root element from the min-heap represented by the array `[-1,0,2,1,4,5,6,3]`.<br
1. We store the root element `-1` to return later.
2. We replace the root element with the last element `3`: `[3,0,2,1,4,5,6,-1]`
```
               3
           0        2
        1     4  5    6
      -1
```

3. We remove the last element from the array: `[3,0,2,1,4,5,6]`
```
               3
           0        2
        1     4  5    6
```

4. We compare `3` with its children `0` and `2`. Since `0` is the smaller child, we swap them: `[0,3,2,1,4,5,6]`
```
               0
           3        2
        1     4  5    6
```

5. We compare `3` with its new children `1` and `4`. Since `1` is the smaller child, we swap them: `[0,1,2,3,4,5,6]`
```
               0
           1        2
        3     4  5    6
```

6. Now, `3` has no children, so the heap property is restored.


#### Complexity Analysis
#### Time Complexity O(log n)
- Deletion takes O(log n) time in the worst case because we may need to traverse the height of the tree to restore the heap property.
