//////////////////////////////////////////////////////
// *** 3Sum ***
//////////////////////////////////////////////////////
/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

Example 2:
Input: nums = []
Output: []

Example 3:
Input: nums = [0]
Output: []
 
Constraints:
0 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
/*
This problem is not that hard after solving the twoSum(sorted)

We have to first sorted the whole array
The main idea is iterating through all the element in the array
We use this element as our first number
Then we do exactly same thing as twoSum(sorted) to remaining array
One thing to note that we can't have duplicate triplets
Which means, [-1,0,1] and [0,-1,1] are the same

For example, nums = [-1,-1,-1,0,1,1,2,2]

1st iteration, i = 0,
right = 7
left = 1
      l               r
[-1, -1, -1, 0, 1, 1, 2, 2]
sum = nums[i] + nums[l] + num[r] = -1 - 1 + 2 = 0
=> Find it
=> What's next?
=> Can we break here?
=> No, there may be other case i + l + r = 0
=> Which means we have to keep moveing left pointer and right pointer until they overlap
=> How do we move l ptr and r ptr?
=> Because both of them are added in the output
=> And we can't have duplicate triplets
=> which means we have to move two pointers inward
=> BUT, that's not enough
=> Let's imagein we move both pointers inward once
         l            r
[-1, -1, -1, 0, 1, 1, 2, 2]
sum = nums[i] + nums[l] + num[r] = -1 + 0 + 2 = 0
=> output = [-1, -1, 2]
=> it's exactly the same as above
=> So it proved that it's not enough just move two pointers inward once
=> What we should do?
=> "We should move pointers inward until the value is different"
=> In this case, left pointer should keep moving inward until value is not -1
=> And right pointer should keep moving inward until value is not 2
=> Original pointer
      l               r
[-1, -1, -1, 0, 1, 1, 2, 2]
=> After moving
             l     r
[-1, -1, -1, 0, 1, 1, 2, 2]
sum = nums[i] + nums[l] + num[r] = -1 + 1 + 0 = 0
=> Again, we find out result
=> Now we have to move two pointers inward
=> And they will overlap, so we can break out the while-loop
=> Note that we still have to move pointers inward until the value is different

2nd iteration, i = 1
right = 7
left = 2
         l               r
[-1, -1, -1, 0, 1, 1, 2, 2]
sum = nums[i] + nums[l] + num[r] = -1 - 1 + 2 = 0
=> Can you see the problem?
=> It's duplicate triplets again
=> So after while-loop, it's not enough to increment the i once
=> Just like moving two pointers
=> We have to keep incrementing i until the value is different
=> now i = 2
right = 7
left = 3
             l            r
[-1, -1, -1, 0, 1, 1, 2, 2]

As we can see, the main hard part of this problem is  
how to avoid duplicate triplets

************************************************************
Time compelxity: O(n ^ 2)
Space comelxity: O(1) or O(n)
=> It depeneds on how do we sort the array
*/
var threeSum = function (nums) {
  const output = [];
  nums.sort((a, b) => a - b);

  for (let i = 0; i < nums.length; i++) {
    let right = nums.length - 1;
    let left = i + 1;

    while (right > left) {
      const sum = nums[i] + nums[left] + nums[right];

      // find the answer, and now we have to move two pointers inward until their value is different
      if (sum === 0) {
        // add to the output
        output.push([nums[i], nums[left], nums[right]]);

        // move left pointer until it's less than right AND the value is different
        while (right > left && nums[left + 1] === nums[left]) {
          left++;
        }
        /*
        now that after the small while-loop above the left pointer is remaing the position nums[left + 1] !== nums[left]
                            l
        For example, [1,1,1,1,2,3,4,5,5,]
        nums[left + 1] !== nums[left]
        so that we have to increment l once
        */
        left++;

        // move right pointer until it's greater than right AND the value is different
        while (right > left && nums[right] === nums[right - 1]) {
          right--;
        }
        right--;
      } else if (sum < 0) {
        left++;
      } else {
        right--;
      }
    }

    // move i until it's less than length AND the value is different
    while (i < nums.length - 1 && nums[i] === nums[i + 1]) {
      i++;
    }
  }

  return output;
};
