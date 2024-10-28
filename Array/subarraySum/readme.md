# Problem Explanation

## Brute Force
The idea is to iterate through the array and for each element, calculate the sum of all subarrays starting from that element and ending at any other element. If the sum equals k, increment the count.<br>

[1, 0, -1, 1], k = 0<br>
For the first element 1, the subarrays are [1], [1, 0], [1, 0, -1], [1, 0, -1, 1]. The sum of these subarrays are 1, 1, 0, 1. We've found the sum of [1, 0, -1] equals k, so we increment the count.<br>
For the second element 0, the subarrays are [0], [0, -1], [0, -1, 1]. The sum of these subarrays are 0, -1, 0. There are two subarrays that sum up to 0, so we increment the count by 2.<br>
For the third element -1, the subarrays are [-1], [-1, 1]. The sum of these subarrays are -1, 0. We've found the sum of [-1, 1] equals k, so we increment the count.<br>
Therfore, the total count is 4.<br>

### Complexity Analysis
#### Time Complexity O(n^2)
#### Space Complexity O(1)

## Hash Table
The idea to optimize the brute force solution is a little tricky.<br>

Can we sort the array first?<br>
No, because the problem requires the subarray, so the order of elements is important.<br>

The idea is to iterate through the array, calculate the current sum of subarray.<br>
Then, we ask two question:<br>
"What's the difference between the current sum and k?"<br>
"If that difference exists in some previous subarray we've seen?"<br>
If it does, that means we find a subarray that its sum equals k if we exclude the subarray we've seen.<br>

Let's use an simple example to explain the idea:<br>
[1, 0, -1, 1], k = 0<br>
Suppose we're at the last element 1, the current sum is 1<br>
"What's the difference between the current sum and k?"<br>
=> 1 - 0 = 1<br>
"If that difference exists in some previous subarray we've seen?"<br>
=> Yes, there are two subarrays that sum up to 1, [1] and [1, 0]<br>
What does that mean?<br>
=> If we exclude the subarray [1], the reamining subarray [0, -1, 1] equals k<br>
=> If we exclude the subarray [1, 0], the reamining subarray [-1, 1] equals k<br>

That's the core idea to solve this problem.<br>
For each iteration,<br>
We first calculate the current sum of subarray.<br>
Then, we ask two questions:<br>
"What's the difference between the current sum and k?"<br>
"If that difference exists in some previous subarray we've seen?"<br>
If it does, that means we find a subarray that its sum equals k if we exclude the subarray we've seen.<br>
After that, we store the current sum in the hash table.<br>
It represents that we've seen a subarray that has the sum of current sum.<br>

Finally, there's one thing we need to note<br>
The initial value of the hash table is {0: 1}<br>
It represents that we've seen a subarray that has the sum of 0, which is the empty subarray.<br>
For example, [1, 0, -1, 1], k = 0<br>
Suppose we're at -1, the current sum is 0<br>
"What's the difference between the current sum and k?"<br>
=> 0 - 0 = 0<br>
"If that difference exists in some previous subarray we've seen?"<br>
=> Yes, there is one subarray [ ](empty array) that sum up to 0<br>
That means if we exclude the empty subarray, the remaining subarray [1, 0, -1] equals k<br>
So that's why we need to initialize the hash table with {0: 1}<br>

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(n)
