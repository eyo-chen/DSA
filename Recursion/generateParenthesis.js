//////////////////////////////////////////////////////
// *** Generate Parentheses ***
//////////////////////////////////////////////////////
/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:
Input: n = 1
Output: ["()"]
 
Constraints:
1 <= n <= 8
*/
/*
This is the solution wrote by myself
But the code is not clean and verbose

The idea is that we first need four arrays
1. right -> store "("
2. left -> store ")"
3. validArr -> make sure the combination of parentheses are valid
   => If add "(", we push one element
   => If add ")", we pop one element
   => So, we need to check if the length of validArr is 0 when hitting base case
   => If it's 0, it means the parentheses is balanced and valid
4. res

This is the backtracking problem
For each call stack, we make two choices
1. Add "("
2. Add ")"

But also need several constraints
1. can only add "(" or ")" when it's available
   => if n = 4, we can't have 5 ")" or "("
2. when about to add ")", we need to consider two things
   (1) if the length of working tmp is greater than 0, and now the length of validArr is 0
   => What does this means?
   => It means it's the case like, "(())", "()", "()()"
   => As we could see, the legnth of working tmp is greater than 0
   => and now the length of validArr is 0, which means parentheses is balanced and valid
   => Now, we CAN'T add ")"
   => Because we don't have "(" at the left side
   => We can't add ")" if there's no remaining "(" at the left side
   => whenenver we add ")" at those cases, the later parentheses guarantee is invalid

   (2) if ")" is the first adding element into tmp
   => it's immediate invalid,
   => Again, We can't add ")" if there's no remaining "(" at the left side

In addition to choice and constraint, we also need to be careful about the backtracking part

Choices
=> Add ")" or "("

Constaint
=> Check if ")" and "(" are available
=> We can't add ")" if there's no remaining "(" at the left side

Goal
=> the length of right and left array is 0

************************************************************
n = the input n
Time complexity: O(4 ^ n)
=> this is roughly upperbound
=> branching factor 2
=> the deepeset recirsive tree is 2 * n
=> O(2 ^ (2n))
=> O(4 ^ n)

Space complexity: O(n)
*/
/**
 * @param {number} n
 * @return {string[]}
 */
var generateParenthesis = function (n) {
  const right = new Array(n).fill('(');
  const left = new Array(n).fill(')');
  const validArr = [];
  const res = [];

  recursiveHelper(right, left, validArr, [], res);

  return res;
};

function recursiveHelper(right, left, validArr, tmp, res) {
  // Base case (Goal)
  if (left.length === 0 && right.length === 0) {
    // check if the parentheses is valid
    if (validArr.length === 0) res.push(tmp.join(''));

    return;
  }

  if (right.length > 0) {
    tmp.push(right.pop());
    validArr.push(true);

    recursiveHelper(right, left, validArr, tmp, res);

    // backtrack
    tmp.pop();
    right.push('(');
    validArr.pop();
  }

  if (left.length > 0) {
    // constaint
    if (tmp.length > 0 && validArr.length === 0) return;

    tmp.push(left.pop());

    // constaint
    if (tmp.length === 1 && tmp[0] === ')') return;

    validArr.pop();

    recursiveHelper(right, left, validArr, tmp, res);

    // backtrack
    tmp.pop();
    left.push(')');
    validArr.push(true);
  }
  return;
}

/*
This one is cleaner
We only need to remember this one
We can't add ")" if there's no remaining "(" at the left side
=> right > left
=> when right > left, it's the case like this (n = 3)
=> "(()" right 2, left 1
=> "()(" right 2, left 1
=> "(((" right 3, left 0
=> as we could see, when right > left, it means there's always remaining "(" on the left hand side
=> it means we're allowed to add ")"
=> "(())" right 1, left 1
=> "()" right 2, left 2
=> In this case right === left, it means there's no remaining "(" on the left hand side

Also we don't need array, only need to use string instead
*/
var generateParenthesis1 = function (n) {
  const res = [];

  recursiveHelper1(n, n, n, '', res);

  return res;
};

function recursiveHelper1(n, left, right, tmp, res) {
  // Base case 1
  /*
  We can have another base case
  if (right === 0 && left === 0)
  */
  if (tmp.length === 2 * n) {
    res.push(tmp);
    return;
  }

  // if "(" is available, we can keep adding
  if (left > 0) recursiveHelper1(n, left - 1, right, tmp + '(', res);

  // We can't add ")" if there's no remaining "(" at the left side
  if (right > 0 && right > left)
    recursiveHelper1(n, left, right - 1, tmp + ')', res);

  return;
}
