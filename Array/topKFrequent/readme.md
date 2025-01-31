# Problem Explanation

# Sorting
The idea is to count the frequency of each number, and then sort the numbers by their frequency in descending order.<br>
It's very straightforward, but it's not the most efficient solution.<br>

Let's summarize the steps:
1. Create a hash table to store the frequency of each number ({number: frequency})
2. Create a slice of pairs to store the number and its frequency ([number, frequency])
3. Sort the pairs by frequency in descending order
4. Create a slice to store the result
5. Return the first `k` numbers

## Complexity Analysis
### Time Complexity O(nlog(n))
### Space Complexity O(n)

# Bucket Sort
The idea is not that straightforward, but it's easy to understand.<br>
If n is the length of the input array, that means the max frequency we can get is n.<br>
Therefore, there are n+1 possible frequencies, range from 0 to n.<br>

Based on this idea, we can create an array of size n+1<br>
The index represents the frequency, and the value is a slice of numbers that have this frequency.<br>
Note that the value of each index is a slice, because there might be multiple numbers that have the same frequency.<br>

For example,<br>
nums = [1,1,1,2,2,3], n = 6<br>
freqArr = [[], [3], [2], [1], [], [], []]<br>
frequency = 0 has no number<br>
frequency = 1 has number 3<br>
frequency = 2 has number 2<br>
frequency = 3 has number 1<br>
frequency = 4 to 6 has no number<br>

After we having the freqArr, we can iterate from the end of the array to the beginning, and add the numbers to the result slice until we have k numbers.<br>

Let's summarize the steps:
1. Create a hash table to store the frequency of each number ({number: frequency})
2. Create a freqArr of size n+1 (each index represents a frequency)
3. Iterate through the hash table, and add the number to the freqArr at the index of its frequency
4. Iterate from the end of the freqArr to the beginning, and add the numbers to the result slice until we have k numbers
5. Return the result slice

## Complexity Analysis
### Time Complexity O(n)
### Space Complexity O(n)

# Min Heap
It's pretty straightforward that we can use a heap(min or max) to solve this problem.<br>
Imagine we already have a max heap, the root of the heap is the most frequent number.<br>
We can pop the root from the heap k times, and the numbers we pop are the most frequent numbers.<br>
Every pop operation is O(log(n)), and we need to pop k times, so the time complexity is O(klog(n)).<br>
Also, because we know that k <= n, so the time complexity is a little bit better than O(nlog(n)).<br>
Note that the value in the heap is a pair of number and its frequency, so we need to define a struct to store the number and its frequency.<br>

However, how to build a max heap?<br>
We have to loop through the hash table, and insert the value into the heap.<br>
In the worst case, we need to insert n elements into the heap, and each insert operation is O(log(n)), so the time complexity is O(nlog(n)).<br>
Therefore, the overall time complexity is still O(nlog(n)).<br>

To have better performance, we can shift the mindset a little bit.<br>
We can use a min heap with a size of k, and the root of the heap is the most frequent number.<br>
Every time we insert a number into the heap, we check if the heap is overflowed.<br>
If that's the case, we pop the root from the heap<br>
Because the root is the least frequent number, so we can ensure that all the values in the heap are the most frequent k numbers.<br>

For constructing the heap,<br>
We still need to loop through n times to insert the value into the heap<br>
However, every insert and pop operation is O(log(k)) because the heap is a min heap with a size of k<br>
so the time complexity is O(nlog(k)).<br>

For generating the final result,<br>
We need to loop through k times to pop the root from the heap<br>
Each pop operation is O(log(k)), so the time complexity is O(klog(k)).<br>

Therefore, the overall time complexity is O(nlog(k)).<br>


Let's summarize the steps:
1. Create a hash table to store the frequency of each number ({number: frequency})
2. Create a min heap with a size of k
3. Iterate through the hash table, and insert the value into the heap
   - First, insert the value into the heap
   - Then, check if the heap is overflowed
     - If that's the case, pop the root from the heap
     - This ensures that the heap is always a min heap with a size of k
4. Pop the root from the heap k times, and the numbers we pop are the most frequent numbers
5. Return the result slice

## Complexity Analysis
### Time Complexity O(nlog(k))
### Space Complexity O(n)







