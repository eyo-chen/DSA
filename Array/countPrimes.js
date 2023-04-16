//////////////////////////////////////////////////////
// *** Count Primes ***
//////////////////////////////////////////////////////
/*
Given an integer n, return the number of prime numbers that are strictly less than n.

Example 1:
Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.

Example 2:
Input: n = 0
Output: 0

Example 3:
Input: n = 1
Output: 0

Constraints:
0 <= n <= 5 * 106
*/
/**
 * @param {number} n
 * @return {number}
 */
/*
Bruth force solution
Just iterate through all the num, and check if it's is prime number

************************************************************
Time compelxity: O(n ^ n)
Space comelxity: O(1)
*/
function countPrimes(n) {
  let count = 0;

  for (let i = 0; i < n; i++) {
    if (isPrime(i)) count++;
  }

  return count;
}

function isPrime(num) {
  if (num < 2) return false;

  for (let i = 2; i < num; i++) {
    if (num % i === 0) return false;
  }

  return true;
}

/*
Optimize solution
The idea is loop through 2 to n, and if the number(i) is prime (!isNotPrime[i])
We want to mark all the i + i + i + i.... to not prime
For example, n = 20
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
F F F F F F F F F F F  F  F  F  F  F  F  F  F  F  F  

We don't need to care 0 and 1, just start at 2 because it's first prime number
i = 2
,!isNotPrime[i] = true, which means it's prime number, mark all i + i + i... to true
Because all of them is not prime
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
F F F F T F T F T F T  F  T  F  T  F  T  F  T  F  T

i = 3
,!isNotPrime[i] = true, which means it's prime number, mark all i + i + i... to true
Because all of them is not prime
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
F F F F T F T F T T T  F  T  F  T  T  T  F  T  F  T

i = 4, !isNotPrime[i] = false, which means it's not the prime number, skip it
so on and so forth....

************************************************************
Time compelxity:  https://leetcode.com/problems/count-primes/discuss/473021/Time-Complexity-O(log(log(n)
Space comelxity: O(1)
*/
var countPrimes = function (n) {
  const isNotPrime = new Array(n).fill(false);
  let count = 0;

  for (let i = 2; i < n; i++) {
    if (!isNotPrime[i]) {
      for (let j = i + i; j < n; j += i) {
        isNotPrime[j] = true;
      }
      count++;
    }
  }

  return count;
};
/*
This is another solution, the only difference is the inner for-loop

For example, n = 20
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
F F F F F F F F F F F  F  F  F  F  F  F  F  F  F  F  

We don't need to care 0 and 1, just start at 2 because it's first prime number
i = 2
mark all 2 * 2, 2 * 3, 2 * 4, 2 * 5.... to true
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
F F F F T F T F T F T  F  T  F  T  F  T  F  T  F  T

i = 3
mark all 3 * 2, 3 * 3, 3 * 4, 3 * 5.... to true
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
F F F F T F T F T T T  F  T  F  T  T  T  F  T  F  T

Basiaclly same idea as above
*/
var countPrimes = function (n) {
  const isNotPrime = new Array(n).fill(false);
  let count = 0;

  for (let i = 2; i < n; i++) {
    if (!isNotPrime[i]) {
      for (let j = 2; i * j < n; j++) {
        isNotPrime[i * j] = true;
      }

      count++;
    }
  }

  return count;
};
