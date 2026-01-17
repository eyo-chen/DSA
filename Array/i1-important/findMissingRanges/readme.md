# Problem Description
You're given a sorted array of unique integers nums and a range defined by [lower, upper]. Your task is to find all the missing numbers in this range that are not present in nums, and represent them as a list of ranges.<br>

The problem asks you to identify which numbers between lower and upper (inclusive) are missing from the array nums. Instead of listing each missing number individually, you need to group consecutive missing numbers into ranges represented as [start, end].<br>

For example:<br>
If nums = [0, 1, 3, 50, 75] and the range is [0, 99]<br>
Missing numbers are: 2, 4-49, 51-74, 76-99<br>
The output would be: [[2, 2], [4, 49], [51, 74], [76, 99]]<br>