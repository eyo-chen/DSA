/*
This problem is kinda similar to the problem yesterday solved(permutateUnique)
and also to the combinationSum 1 problem which is in the DP section

But use different trick

The part is similar to combinationSum 1 is
Use index to prevent loop from 0 ~ the length of candidates every time
For example, if candidates = [1,2,3]
Once choosing 2, we don't back to choose 1 again in the later iteration

The part is simialr to permutateUnique is
if (i > index && candidates[i] === candidates[i - 1]) continue;
In permuateUnique, we also use usedArr, but we don't need that here
I think the most tricky part is i > index, what does this means?
Remeber from that problem, we said choosing order DOES matter when it comes to how to deal with duplicate case
[1a, 1b, 2] is as same as [1b, 1a, 2]
If we don't care oder, we'll ouput those two arrays
For example, candidates = [1,2,2,2,5], target = 5
                                               5
                            1                  2a                   2b           2c          
                            (4)                (2)                 (2)           (2)
                  2a         2b       2c               
                  (2)       (2)       (2) 
        2b        2c        ***       ****
        (0)       (0)
                  **
** part
=> Can we choose 2c in here?
=> Once choosing, the output is [1,2a,2c]
=> It's as same as first output [1,2a,2b]
=> (i > index && candidates[i] === candidates[i - 1]) helps us to avoid this case
=> Note that in the call stack of ** part, the iteration is from 2 ~ 4, why?
=> because after choosing 2a, we pass i + 1 down as our initial index for next call stack
=> here, i is the postion of 2a, which is 1
=> so we we loop from 2
=> for 2b, 2 > 2 is false, what does this mean?
=> it means it's the very first element in this call stack, it's okay to choose it no matter what
=> for 2c, 2 > 3 && candidate[3] === candidate[2], what does this means?
=> 2c is not the very first element in this call stack, and it's as same as previous one
=> don't choose it
=> it's the same logic as *** and **** part

The best way to deeply understand this trick is writing the recursive tree on the paper

************************************************************
n = the length of candiate t = target
Time complexity: O(2 ^ n)
=> In the worst case, our algorithm will exhaust all possible combinations from the input array. Again, in the worst case, let us assume that each number is unique. The number of combination for an array of size n would be 2 ^ n, i.e. each number is either included or excluded in a combination.
=> The number of combinations of a list of size N is 2 ^ N
=> because we can either choose the element or not choose it


Space complexity: O(n)
=> the max length of tmp array is determined by n, which is O(n)
=> the deepest recursive tree is also O(n)

Note: we did not take into account the space needed to hold the final results of combination in the above analysis.
*/
var combinationSum2 = function (candidates, target) {
  const res = [];
  const sortedCandidates = candidates.sort((a, b) => a - b);

  recursiveHelper(sortedCandidates, target, 0, [], res);

  return res;
};

function recursiveHelper(candidates, target, index, tmp, res) {
  // base case
  if (target === 0) {
    res.push([...tmp]);
    return;
  }

  // start from index
  for (let i = index; i < candidates.length; i++) {
    // avoid choosing duplicate element
    if (i > index && candidates[i] === candidates[i - 1]) continue;

    // entirely break the for-loop because we've sorted the array before
    // once this element is too large, later element is definitely invalid
    if (target - candidates[i] < 0) break;

    // choose
    tmp.push(candidates[i]);

    // explore (pass index as i + 1)
    recursiveHelper(candidates, target - candidates[i], i + 1, tmp, res);

    // unchoose
    tmp.pop();
  }

  return;
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*
Solution 2
This solution uses the similar approach to permutateUnique

we bundle the candidates as hashtable
key -> each candidates
value -> frequencey of each candidates

go to look permutateUnique to see the recursive tree

************************************************************
n = the length of candiate t = target
Time complexity: O(2 ^ n)
=> same as above

Space complexity: O(n)
=> same as above
=> although we additionally use O(n) hashTable
=> overall, it's still O(n)
*/
var combinationSum21 = function (candidates, target) {
  const res = [];
  const hashTable = {};
  // build hashTable
  for (let i = 0; i < candidates.length; i++) {
    if (hashTable[candidates[i]]) hashTable[candidates[i]]++;
    else hashTable[candidates[i]] = 1;
  }

  // the array is gonna iterate in the recursive function
  const newCandidates = Object.keys(hashTable);

  recursiveHelper1(newCandidates, target, hashTable, 0, [], res);

  return res;
};

function recursiveHelper1(candidates, target, hashTable, index, tmp, res) {
  if (target === 0) {
    res.push([...tmp]);
    return;
  }

  for (let i = index; i < candidates.length; i++) {
    if (hashTable[candidates[i]] === 0) continue;
    if (target - candidates[i] < 0) break;

    // choose (frequency--)
    tmp.push(candidates[i]);
    hashTable[candidates[i]]--;

    // explore
    recursiveHelper1(
      candidates,
      target - candidates[i],
      hashTable,
      i,
      tmp,
      res
    );

    // unchoose (frequencey++)
    tmp.pop();
    hashTable[candidates[i]]++;
  }

  return;
}
