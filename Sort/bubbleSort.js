//////////////////////////////////////////////////////
// *** Bubble Sort ***
//////////////////////////////////////////////////////
/**
 * @param {Array} arr
 * @return {Array}
 */
/*
The logic of bubble sort is
At each iteration, bubble(move) the largest number up to the final position
In other words, we start at index n, and we said that I want to find the correct value at this position in this iteration

The algorithm is based on 1)comparison 2)swap

(1) iterate the array from n to 0 (outer while-loop)
(2) focus on two element at a time
(3) compare two elements, if arr[i] is greater than arr[i + 1] (comparison)
    For example, [4,3,2,1]
    i = 0,
    arr[0] > arr[0 + 1]
    => swap(arr, i, i + 1)
(4) Swap them 
(4) after inner for-loop, the last element of array should have the correct value
(5) in the next while-loop, we DONT need to iterate all the array from 0 to last again becuase we know that the final position is correct and fixed, we don't need to swap anymore


Let's walk through the example, we'll omit the sortIndex in this example
arr = [5,3,4,2,1]
last = 4

First while-loop
arr = [5,3,4,2,1]
last = 4
=> In this iteration, I want to find the correct value to put at the last position(index 4)
=> In other words, I want to find the max value, and bubble up to the last position
Inner for-loop
i = 0
arr[0] > arr[0 + 1]
=> swap(0, 1)
=> [3,5,4,2,1]
i = 1
arr[1] > arr[1 + 1]
=> swap(1, 2)
=> [3,4,5,2,1]
i = 2
arr[2] > arr[2 + 1]
=> swap(2, 3)
=> [3,4,2,5,1]
i = 3
arr[3] > arr[3 + 1]
=> swap(3, 4)
=> [3,4,2,1,5]
Inner for-loop done
Note two things
1. We won't go to i = 4 which is the real last index because the comparison is arr[i] and arr[i + 1]
If we have i = 4, then it will become arr[4] and arr[5] which is out of the bound
So that's why our initial last is start at 4, and we do i < last
2. After this for-loop, the final position is fixed, we bubble up the largest value to the last position

Second while-loop
arr = [3,4,2,1,5]
last = 3
=> In this iteration, I want to find the correct value to put at the last position(index 3)
=> In other words, I want to find the second max value, and bubble up to the last two position
..........

Basically the same process happens again, and again ..........


There's one final thing to note
The main great part of bubble sort is that it's best case is O(n)
which means that we only have O(n) works if the input array is sorted
Why?
We use sortIndex
It represents that question 
Does this array sort yet?
Let's break down the process again
arr = [1,2,3,4,5] -> sorted
last = 4
sortIndex = false -> we assume that input array is not sorted

First while-loop
arr = [1,2,3,4,5]
last = 4
sortIndex = false
At the beginning of entering while-loop, we always set sortIndex to true
It means we just temporarily assume the array is sorted
Inner for-loop
index i will go from 0 to 3
and all arr[i] is less than arr[i + 1]
But there's no swap in this case because the array is sorted
So there's no chance to set sortIndex to false
When go into next while-loop, !sortIndex becomes false
So we just break out the while-loop
and return the array
The total amount of works is O(n)

Again, the main idea is that
Eeach time entering the while-loop, we just assume array is sorted
so set sortIndex to true
Inside the for-loop,
If we ever do the swap once or more times, we set sortIndex to false
It means "hey, we do the swap, so the array may not be sorted"
If we never do the swap, then sortIndex remain the same
It means "hey, we never do the swap, so the array is guarantee sorted, job done"

************************************************************
Time: O(n ^ 2)
Space: O(1)
*/
function bubbleSort(arr) {
  let sortIndex = false;
  let last = arr.length - 1;

  while (last >= 0 && !sortIndex) {
    // assume array is sorted
    sortIndex = true;

    for (let i = 0; i < last; i++) {
      // comparison, do we need to do the swap?
      if (arr[i] > arr[i + 1]) {
        swap(arr, i, i + 1);

        // we do the swap, so the array may not be sorted
        sortIndex = false;
      }
    }

    // we've fixed the last position, no need to do the work for last position again in next ieration
    last--;
  }

  return arr;
}

function swap(arr, i, j) {
  const tmp = arr[i];
  arr[i] = arr[j];
  arr[j] = tmp;
}
