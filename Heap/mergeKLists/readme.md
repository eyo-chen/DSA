# Problem Explanation

## Naive Approach
The idea is very simple. Just insert all values into a min heap and then pop all values from the min heap and create a new linked list.

- Insert all values into a min heap
- Pop all values from the min heap and create a new linked list

### Complexity Analysis
#### Time Complexity O(n*log(n))
- where n is the total number of nodes in all linked lists
- First loop to insert all values into a min heap takes O(n*log(n)) time
  - We loop through all values in all linked lists, so it's O(n)
  - For each value, we insert it into a min heap, so it's O(log(n))
  - Therefore, the first loop takes O(n*log(n)) time
  - Note that the heap will contains all values in all linked lists, so the heap size is O(n)
  - Inserting and removing an element from a heap takes O(log(n)) time
- Second loop to pop all values from the min heap and create a new linked list takes O(n*log(n)) time
  - We pop all values from the min heap
  - For each node, we use `Pull`, which is O(log(n))
  - Therefore, the second loop takes O(n*log(n)) time

#### Space Complexity O(n)
- We use a min heap to store all values in all linked lists, so the space complexity is O(n)

## Optimized Approach
Instead of inserting all values into a min heap, we can insert the head of each linked list into a min heap.<br>
Then, we can pop the smallest element from the min heap and insert the next element from the same linked list into the min heap.<br>
We repeat this process until the min heap is empty.

Here's the steps:
1. Insert all heads of the linked lists into a min heap
2. Pop the smallest element from the min heap
3. Insert the next element of the smallest element from the same linked list into the min heap
4. Repeat steps 2 and 3 until the min heap is empty

### Complexity Analysis
#### Time Complexity O(n*log(k))
- where n is the total number of nodes in all linked lists
- where k is the number of linked lists
- Inserting all heads of the linked lists into a min heap takes O(k*log(k)) time
  - We loop through all linked lists, so it's O(k)
  - For each linked list, we insert its head into a min heap, so it's O(log(k))
  - Therefore, the first loop takes O(k*log(k)) time
- Popping the smallest element from the min heap and inserting the next element from the same linked list into the min heap takes O(n*log(k)) time
  - We loop through all nodes in all linked lists, so it's O(n)
  - For each node, we pop it from the min heap and insert the next node from the same linked list into the min heap, so it's O(log(k))
  - Therefore, the second loop takes O(n*log(k)) time
- Therefore, the overall time complexity is O(n*log(k))

#### Space Complexity O(k)
- We use a min heap to store the heads of the linked lists, so the space complexity is O(k)

## Divide and Conquer
The idea is to divide the linked lists into two halves, sort each half, and then merge the two halves.

Let's walk through an example with 4 lists: `[list1, list2, list3, list4]`
1. Initial call: `helper(lists, 0, 3)`
   - mid = 1
   - left = helper(lists, 0, 1)
   - right = helper(lists, 2, 3)

2. For left: `helper(lists, 0, 1)`
   - This hits the `start+1 == end` case
   - Returns mergeList(list1, list2)

3. For right: `helper(lists, 2, 3)`
   - This also hits the `start+1 == end` case
   - Returns mergeList(list3, list4)

4. Back in the initial call, we now have:
   - left = merged(list1, list2)
   - right = merged(list3, list4)
   - Finally, return mergeList(left, right)

### Complexity Analysis
#### Time Complexity O(N*log(k))
- where N is the total number of nodes in all linked lists
- where k is the number of linked lists
- The `helper` function:
   - This function recursively divides the problem into halves until it reaches base cases.
   - At each level of recursion, it merges two lists.
   - The number of levels in the recursion tree is log k, where k is the number of input lists.
   - At each level, we're processing all N nodes (where N is the total number of nodes across all lists).
- Overall time complexity:
   - O(N log k), where N is the total number of nodes across all lists, and k is the number of lists.
   - This is because we're doing log k levels of merging, and at each level, we're processing N nodes in total.

#### Space Complexity O(log(k))
- The `mergeList` function:
  - Uses a constant amount of extra space (for the dummy head and current node).
  - Space complexity: O(1)
- The `helper` function:
   - The space complexity here is determined by the recursion stack.
   - The maximum depth of the recursion is log k (where k is the number of lists).
  - At each level of recursion, we're using a constant amount of extra space.
- Overall space complexity:
  - O(log k) due to the recursion stack.
  - Note that this doesn't include the space used by the input lists or the output list, as that's considered part of the input/output space.

## Iterative Approach
The idea is to merge two lists at a time until all lists are merged into one.

Let's walk through an example to illustrate:<br>
Suppose we start with `lists = [list1, list2, list3, list4]`
1. First iteration:
   - Merge `list1` and `list2` into `mergedList1`
   - First, append `mergedList1` to `lists`, so `lists` becomes `[list1, list2, list3, list4, mergedList1]`
   - Then, remove the first two lists `[list1, list2]` from `lists`
   - So, `lists` becomes `[list3, list4, mergedList1]`
2. Second iteration:
   - Merge `list3` and `list4` into `mergedList2`
   - First, append `mergedList2` to `lists`, so `lists` becomes `[list3, list4, mergedList1, mergedList2]`
   - Then, remove the first two lists `[list3, list4]` from `lists`
   - So, `lists` becomes `[mergedList1, mergedList2]`
3. Third iteration:
   - Merge `mergedList1` and `mergedList2` into `finalList`
   - First, append `finalList` to `lists`, so `lists` becomes `[mergedList1, mergedList2, finalList]`
   - Then, remove the first two lists `[mergedList1, mergedList2]` from `lists`
   - So, `lists` becomes `[finalList]`
4. The loop ends because `len(lists) == 1`
5. `return lists[0]` returns `finalList`

This approach efficiently merges the lists in a pairwise manner, reducing the number of lists by half in each iteration until only one fully merged list remains. It's a clever way to merge multiple lists without using additional data structures like a heap.

### Complexity Analysis
#### Time Complexity O(N*k)
- where N is the total number of nodes in all linked lists
- where k is the number of linked lists
- The `mergeList` function:
   - This function iterates through both input lists once, performing constant-time operations for each node.
   - Time complexity: O(n + m), where n and m are the lengths of the two lists being merged.
- The `MergeKLists3` function:
   - In each iteration of the loop, we merge two lists and remove them from the front of the slice.
   - The number of iterations is k-1, where k is the initial number of lists.
   - In each iteration, we're merging two lists, which takes O(N_i) time, where N_i is the total number of nodes in the two lists being merged.
- Overall time complexity:
   - O(N * k), where N is the total number of nodes across all lists, and k is the number of lists.
   - This is because we're doing k-1 merges, and in total, we're processing each node k/2 times on average.

#### Space Complexity O(1)
- The `mergeList` function:
  - Uses a constant amount of extra space (for the dummy head and current node).
  - Space complexity: O(1)
- The `MergeKLists3` function:
   - We're not using any additional data structures that grow with the input size.
   - The space used by the `lists` slice is constantly shrinking as we process more lists.
   - We're not creating any new nodes, just rearranging existing ones.
- Overall space complexity:
  - O(1) or constant space.
  - Note that this doesn't include the space used by the input lists or the output list, as that's considered part of the input/output space.

This solution is more straightforward and uses less space compared to the divide-and-conquer approach, but it has a higher time complexity, especially when k is large. It performs well when k is small or when space is a primary concern.

The trade-off here is between time and simplicity. This approach is easier to understand and implement, and it uses constant extra space, but it's less efficient in terms of time complexity compared to the divide-and-conquer approach when dealing with a large number of lists.