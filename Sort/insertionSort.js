//////////////////////////////////////////////////////
// *** Insertion Sort ***
//////////////////////////////////////////////////////
/*
The main process is like this
Each time we take out a element as insertedVal
Traverse backward to find the correct position to insert this element

Insertion sort basically only do two things
1) comparison
2) move

We first assume the first elemet has been sorted
(1) start at second element(1th), iterate the array(1 ~ ith)
(2) take the ith element out(insertedVal)
(3) iterate backward from (i - 1) to 0 (jth ~ 0)
(4) if jth element is greater than insertedVal element
=> What does this mean?
=> It means the the correct position of insertedVal is definitely in front of jth element
=> In other words, jth element should be back of the insertedVal
=> So we just simply move jth element forward one position
=> Because we know there is a position has to be for the insertedVal later
(5) move the jth element forward to (j + 1) position
=> it's okay to mutate the original element because we've already get the element out into insertedVal variable
(6) After while-loop, now jth index means the position is smaller than insertedVal element
(7) which means now we could put insertedVal element into j + 1 position, right after the element smaller than it

Note that it's just like bubble sort
The best case is O(n)
Imagine the input array is [1,2,3,4,5]
jth element is always smaller than insertedVal
So we never enter the while-loop
We only have outer for-loop works
Which is O(n)

************************************************************
Time: O(n ^ 2)
Space: O(1)
*/
/**
 * @param {Array} arr
 * @return {Array}
 */
function insertionSort(arr) {
  for (let i = 1; i < arr.length; i++) {
    // take out the element which is gonna be insered into correct position
    const insertedVal = arr[i];

    // start the while-loop from i - 1 to 0
    let j = i - 1;

    // keep while-loop until arr[j] is smaller or we over the head of array
    while (j >= 0 && arr[j] > insertedVal) {
      // move arr[j] one position forward
      arr[j + 1] = arr[j];
      j--;
    }

    // now jth element is the element smaller than insertedVal
    // we just insert insertedVal to it's next position
    arr[j + 1] = insertedVal;
  }

  return arr;
}
function swap(arr, i, j) {
  const tmp = arr[i];
  arr[i] = arr[j];
  arr[j] = tmp;
}
