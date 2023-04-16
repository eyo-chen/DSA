//////////////////////////////////////////////////////
// *** Selection Sort ***
//////////////////////////////////////////////////////
/*
The logic of selection sort is
find the minimum number, and swap that to the correct position(the first index in the first iteration)
(1) iterate the array from 0 to n (i)
(2) assume the first elemenent(min) is the smallest
(3) start at i + 1 to iterate the array again (j)
(4) if jth element is smaller than minth element, change j to min
    (This process is to find the index of smallest element)
(5) if the index min never change, it means min is equal to i. Then we don't need to change
(6) if not, swap the smallest number to the right position

The problem of selection sort is not only it needs to take O(n^2), but also it still needs do the same amount of time even tho the array has already been sorted

************************************************************
Time: O(n ^ 2)
Space: O(1)
*/
function selectionSort(arr) {
  for (let i = 0; i < arr.length; i++) {
    // assume the smallest number is just i
    let min = i;
    for (let j = i + 1; j < arr.length; j++) {
      // find the smallest number
      if (arr[min] > arr[j]) {
        min = j;
      }
    }

    swap(arr, i, min);
  }

  return arr;
}

function swap(arr, i, j) {
  const tmp = arr[i];
  arr[i] = arr[j];
  arr[j] = tmp;
}
