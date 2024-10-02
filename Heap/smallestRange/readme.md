# Problem Explanation

The core idea to solve this problem is starting with the first element of each list and then iterating through the lists to find the smallest range that contains at least one element from each list.<br>

Let's summarize the step:
1. Initialize a min heap with the first element of each list.
2. Find the minimum value in the heap.
3. Update the range if the current range is smaller than the previous range.
4. If the current list has more elements, add the next element to the heap, also update the global max.
5. Repeat steps 2-4 until the length of heap is not equal to the number of lists.

Let's walk through the example:

```
[
  [4,10,15,24,26],
  [0,9,12,20],
  [5,18,22,30]
]
```
1. list = [4,0,5]
   - range = [0,5]
   - the minimum = 0
   - the maximum = 5
   - remove 0, and add 9 
2. list = [4,9,5]
   - range = [4,9]
   - the minimum = 4
   - the maximum = 9
   - remove 4, and add 10
3. list = [10,9,5]
   - range = [5,10]
   - the minimum = 5
   - the maximum = 10
   - remove 5, and add 18
....

So on and so forth, until we have iterated through all the lists.
    
# Complexity Analysis
## Time Complexity O(n * log(k))
- n is the total number of elements in all lists
- k is the number of lists
- We have to iterate through all the element in the list, and we do O(log(k)) for each insertion and deletion in the heap.

## Space Complexity O(k)
- We use a min heap to store the first element of each list, and the size of the heap is the number of lists.
