//////////////////////////////////////////////////////
// *** Fruit Into Baskets ***
//////////////////////////////////////////////////////
/*
You are visiting a farm that has a single row of fruit trees arranged from left to right. The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.

You want to collect as much fruit as possible. However, the owner has some strict rules that you must follow:

You only have two baskets, and each basket can only hold a single type of fruit. There is no limit on the amount of fruit each basket can hold.
Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right. The picked fruits must fit in one of your baskets.
Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
Given the integer array fruits, return the maximum number of fruits you can pick.

Example 1:
Input: fruits = [1,2,1]
Output: 3
Explanation: We can pick from all 3 trees.

Example 2:
Input: fruits = [0,1,2,2]
Output: 3
Explanation: We can pick from trees [1,2,2].
If we had started at the first tree, we would only pick from trees [0,1].

Example 3:
Input: fruits = [1,2,3,2,2]
Output: 4
Explanation: We can pick from trees [2,3,2,2].
If we had started at the first tree, we would only pick from trees [1,2].
 
Constraints:
1 <= fruits.length <= 105
0 <= fruits[i] < fruits.length
*/
/**
 * @param {number[]} fruits
 * @return {number}
 */
/*
This problem is a little hard to understand
In short, we're given an array
The problem asks us to return the longest length of subarray containing two distinct integer
[1,2,3,2,2]
the longest length of subarray is 4 [2,3,2,2]

This is the solution I came up with by myself
It seems the code is kind of verbose
But I think it's a great solution

The mina idea is using sliding window(two pointers)
We need several variables
1. hashTable
=> To help us to remember what fruit(integer) we've counted
=> Also the last index for each integer occured(important)
=> Why do we need the last index for each integer?
=> For example, [1,2,1,2,2,3,3,3,3,3,3,3,3,3]
=> Imagine the left still at index 0, and right is at index 4(2)
=> When right pointer updates, it moves to index 5, and the value is 3
=> We know that the type of fruit is 3, so we have to move the left pointers
=> But where should this left pointers move to?
=> It should move to the last index of 2, which is index 3
=> Note that we CAN'T move to index 1 because that will also contains a 1, which can't have only two types

2. count
=> the length between two pointers

3. type
=> How many different type of fruit?

4. the max count or length
=> It's the output we're gonna return

The main idea is 
1. moving the right pointer, and get the curFruit
2. If we haven't hold the curFruit, then we need to do the following
   (using hashTable to have O(1) check if we've hold the curFruit)
   (1) Check if type is greater than 3
       If yes, 
       a. caculate the max
       b. make sure type is back to 2
       c. move left pointer to the last index of fruits[right - 1]
       It may look weird, but it's very important
       Why we need to move the left point to the last index of fruits[rigjt - 1]
       Again, fruits = [1,2,1,2,2,3,3,3,3,3,3,3,3,3]
       Imagine right pointer at index 5, and left pointer at index 0
       We know that we have 3 type when right pointer at index 5
       So that we have to move left pointer
       And we know that 3 is definitely new type
       And index(5 - 1) is the one of the old type
       and it's 2, so we have to move left to the last index of 2
       Which is index 3
       This is the most important part, and that's the main reason why we use hashTable to keep tracking the last index of each fruit
       d. If fruit is not fruits[right - 1], then we have to set this fruit to false
       Again, this is kind of tricky
       Note that in the whole process, we can only have two types
       Which means hashTable can only have two index is NOT false
       [1,2,1,2,2,3,3,3,3,3,3,3,3,3], when right pointer at index 5, and left pointer at index 0
       We know that hashTable[1] is not false
       And we also know that curFruit and fruits[right - 1] is definitely our cur type to hold on
       So any fruit is not fruits[right - 1], we have to set back to false
       Why we don't need to care curFruit
       Because it's new type, hashTable[curFruit] is 100% false at this point
       We set to right pointer as index later
       e. aslo need to decrease the length of two pointers
    (2) Because curFruit is new type
        So that we just set this index in our hashTable
3. If we've hold the curFruit, then we need to do the following
   (1) check if curFruit is equal to fruits[right - 1]
   If it's not equal, then we need to update the last index of curFruit
   For example, fruits = [1,2,1,2,2,3,3,3,3,3,3,3,3,3]
   When we fitst hit index 1, we set hashTable[2] = 1
   Then when we hit index 3, we saw that 2 has new last index
   Because we can't set 1 as last index anymore,
   the new last index is 3
   So we set hashTable[2] = 3
   Later, index 4, hashTable[right - 1] is still 2, so that we don't need to set new last index

4. Always update right pointer and the count


Note
This problem is kind of tricky, but this is great problem
And I solve this by myself, hope the logic is clear
************************************************************
n = the legnth of array
Time compelxity: O(n)
Space comelxity: O(n)
*/
var totalFruit = function (fruits) {
  const hashTable = new Array(fruits.length).fill(false);
  let left = 0;
  let right = 0;
  let count = 0;
  let type = 0;
  let max = 0;

  while (right < fruits.length) {
    // get the current type of fruit
    const curFruit = fruits[right];

    // we haven't hold this type of fruit
    if (hashTable[curFruit] === false) {
      // update type, and also check if it's greater than 2 at the same time
      if (++type === 3) {
        // before adjusting sliding window, update the max
        max = Math.max(max, count);

        // make sure type is back to 2
        type--;

        // move the left pointer all the way to the last index of fruits[right - 1]
        while (left !== hashTable[fruits[right - 1]]) {
          // if we're at the fruit is not in the current two type, set it back to false
          if (fruits[left] !== fruits[right - 1]) {
            hashTable[fruits[left]] = false;
          }

          // update left and count
          left++;
          count--;
        }
      }

      // because it's new type of fruit, just set right as last index
      hashTable[curFruit] = right;
    } else {
      // check if it's needed to update last index
      if (curFruit !== fruits[right - 1]) {
        hashTable[curFruit] = right;
      }
    }

    // update both index
    right++;
    count++;
  }

  // also have to update max in the end
  max = Math.max(max, count);

  return max;
};

/*
Instead of using hashTable, we use couple pointers because the problem only aksed at most 2 integer

The idea is a liitle bit similar to the above one, but even tricker

firstFruit
=> the first fruit of current holding type

secondFruit
=> the second fruit of current holding type

firstFruitCount
=> the total count of secondFruit type
=> Why we need this variable?
   For example, fruits = [1,1,2,2,2,3,3]
   When we're at index 4, we know that at this point
   firstFruit = 2
   firstFruitCount = 3
   secondFruit = 1
   In the next iteration, we'll hit the new type (it's the type !== firstType and secondType)
   We know that we have to get rid of the 1 type
   But we still need to know the count after getting rid of the 1 type
   So the new count shold be firstFruitCount + 1
   where 1 indicates that new type of fruit
   which is 3 + 1 = 4
   It's correct, this four indicates [2,2,2,3]

1. Iterate through each element(fruit) in the fruits
2. If curFruit is either firstFruit or secondFruit
=> update the count
3. If curFruit is new type, it's not equal to firstFruit and secondFruit
=> set count = firstFruitCount + 1; (see above)
4. If curFruit is equal to firstFruit, just updating firstFruitCount
   For example, [1,1,2,2,2,3,3]
   When curFruit is index 3 and 4, just keep updating firstFruitCount
5. If it's not, 
   For example, [1,1,2,2,2,3,3], curFruit is index 5, fruit 3
   firstFruit = 2, secondFruit = 1
   1. set firstFruitCount to 1
   Because we're now have new firstFruit(3)
   2. secondFruit = firstFruit
      seconFruit = 2
   3. firstFruit = curFruit
      firstFruit = 3
6. always update the max variable


This solution is really clever, but i think it's not meant for this problem
I reference this solution from discuss section

************************************************************
n = the legnth of array
Time compelxity: O(n)
Space comelxity: O(1)
*/
var totalFruit = function (fruits) {
  let count = 0;
  let max = 0;
  let firstFruit = -1;
  let secondFruit = -1;
  let firstFruitCount = 0;

  for (const fruit of fruits) {
    if (fruit === firstFruit || fruit === secondFruit) {
      count++;
    } else {
      count = firstFruitCount + 1;
    }

    if (fruit === firstFruit) {
      firstFruitCount++;
    } else {
      firstFruitCount = 1;
      secondFruit = firstFruit;
      firstFruit = fruit;
    }

    max = Math.max(max, count);
  }

  return max;
};
