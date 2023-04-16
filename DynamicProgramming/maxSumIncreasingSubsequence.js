/**
 * @param {number[]} array
 * @return {number[maxSum, []]}
 */
/*
This is another variant
Returns the max sum increasing subsequence

This time we're not interested to the longest increasing subsequence(LIS), we're insterested to the max sum increasing subsequence(MIS)

************************************************************
n = the legnth of array
Time compelxity: O(n * 2)

Space comelxity: O(n)
*/
function maxSumIncreasingSubsequence(array) {
  const table = [...array];
  const res = []; // final returned array
  let maxSum = array[0]; // the global max sum
  let pathIndex = 0; // helps to keep track the final index of MIS (used in the while-loop)

  for (let i = 1; i < array.length; i++) {
    for (let j = 0; j < i; j++) {
      /*
        Small Note 
        only extend the MIS when array[i] > array[j] when it's strictly increasing
        If it's not strictly increasing, it's array[i] > array[j], which means allowed same value
        */
      if (array[i] > array[j]) {
        /*
          Note 
          even tho array[i] > array[j], that doesn't means it's necessary to extend the MIS
          When array[i] > array[j], it can
          1) extend the previous MIS
          2) start the new MIS
  
          The condition depends on what's the sum of sequence is max(not to find the longest MIS)
          */
        table[i] = Math.max(table[i], table[j] + array[i]);

        /*
          update both variable when finding the new maxSum
          again, it could be extend the previous MIS or start new MIS
          we want to update both variable in either condition
          */
        if (table[i] > maxSum) {
          maxSum = table[i];
          pathIndex = i;
        }
      }
    }
  }

  res.push(maxSum);
  res.push([]);

  /*
    now pathIndex is the final index of MIS
    backtrack to find the path
  
    maxSum could be negative, so only can use pathIndex as only condition
  
    this while-loop is upper bound to O(n)
    */
  while (pathIndex >= 0) {
    if (maxSum === table[pathIndex]) {
      // use .unshift to keep the path in increasing order
      res[1].unshift(array[pathIndex]);
      maxSum -= array[pathIndex];
    }
    pathIndex--;
  }

  return res;
}
/*
  Another approach
  Instead of backtracking at the end, store both maxSum and path inside the table
  
  ************************************************************
  n = the legnth of array
  Time compelxity: O(n * 2 * n)
  => Because of  table[i][1].unshift(...table[localMaxSumIndex][1]);
  
  Space comelxity: O(n * 2)
  => store array in each cell of table
  */
function maxSumIncreasingSubsequence1(array) {
  const table = [];
  let globalMaxSumIndex = 0;

  for (let i = 0; i < array.length; i++) {
    // intialize the [maxSum, [path]] with [value, [value]]
    table.push([array[i], [array[i]]]);

    // to keep track what's the maximum MIS gonna add in the table[i]
    let localMaxSumIndex;

    for (let j = 0; j < i; j++) {
      /*
        Tricky Part
        Here means we want to extend the prev MIS
        1) the value is greater than previous one
        2) the new MIS is greater than the old one
  
        The reason is becasue we not only want to store the max sum in table, we also want to use localMaxSumIndex to update the path array after for-loop
  
        localMaxSumIndex helps us to guarantee it's the last or max MIS 
        so we can push the array
        */
      if (array[i] > array[j] && table[j][0] + array[i] > table[i][0]) {
        table[i][0] = table[j][0] + array[i];
        localMaxSumIndex = j;
      }
    }

    /*
      Don't forget we have two choies for each cell of table
      2) start the new MIS
      Here, it represents start the new MIS, for example [-3, -2, -1]
      At i = 2, value = -1, we don't want to extend the prev MIS
      We want to start the new MIS because [-1] is the final MIS
      i = 2 won't get into the condition above because there's no MIS greater than [-1]
      But we also want to know and update the globalMaxSumIndex
  
      When the maxSum is greater than the previous MIS (table[i][0] > table[globalMaxSumIndex][0])
      We know that it's time to start the new MIS, so update the globalMaxSumIndex
  
      ----------------------------------------------------------------------
      it also represents 
      1) extend the prev MIS
      because table[i][0] has been updated in the condition above
      so even tho it's the condition of extend the prev MIS, it can still update the globalMaxSumIndex
  
      Which means this condition helps us to udpate the globalMaxSumIndex in both condition
      */
    if (table[i][0] > table[globalMaxSumIndex][0]) globalMaxSumIndex = i;

    // here is want to extend the prev MIS
    if (localMaxSumIndex !== undefined) {
      table[i][1].unshift(...table[localMaxSumIndex][1]);
    }
  }

  return table[globalMaxSumIndex];
}

// console.log(maxSumIncreasingSubsequence([1, 2, 3, 4, 5, 66, 1, 31, 321]));
// console.log(maxSumIncreasingSubsequence1([1, 2, 3, 4, 5, 66, 1, 31, 321]));
