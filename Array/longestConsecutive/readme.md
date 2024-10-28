# Problem Explanation

## Brute Force
This is the solution I came up with first. It's not efficient but it's a good starting point.<br>
The idea is to put all the numbers in a hash table.<br>
Then, loop through each number, and check if the next number(n + 1) is in the hash table.<br>
If it is, then we keep incrementing the value and the counter until we don't find the next number.<br>
Note that we do this for each number, and return the maximum count.<br>

### Complexity Analysis
#### Time Complexity O(n^2)
- We loop through each number, and for each number, we loop through the next numbers until we don't find the next number.

#### Space Complexity O(n)
- We use a hash table to store the numbers, so the space complexity is O(n).

## Finding Starting Point Of Sequence
In previous solution, we loop through each number, and for each number, we loop through the next numbers until we don't find the next number.<br>
This is not efficient, so we need to find a better way to do this.<br>

If the input is [100, 4, 200, 1, 3, 2], we can model this as following:<br>
[1,2,3,4............, 100, 101, ....... 200]<br>
What's the point of thinking about this?<br>
It means there are only three starting points to find the longest consecutive sequence.<br>
1. 1
2. 100
3. 200

That means we only need to start at these three starting points to find the longest consecutive sequence.<br>
We don't need to start at every number in the array.<br>

Now, how can we find these starting points?<br>
We know that a starting point is a number that has no number at it's left side.<br>
So, we can check if the previous number(n - 1) is in the hash table.<br>
If it is, then it's not a starting point because it means there's a number at it's left side.<br>
If it's not, then it's a starting point.<br>

Let's summarize the steps:
1. Put all numbers in a hash table
2. Loop through each number, and check if the previous number(n - 1) is in the hash table.
3. If it is, then it's not a starting point.
4. If it's not, then it's a starting point.
5. From this starting points, count the longest consecutive sequence.

### Complexity Analysis
#### Time Complexity O(n)
- We loop through each number, and for each number, we check if the previous number is in the hash table.

The key insight is that while we have a nested loop structure, each number in the array is only visited a constant number of times. Let's break it down:

1. First loop to build the hash table: O(n)
```go
for _, n := range nums {
    hashTable[n] = true
}
```

2. The second part looks like it might be O(n²) at first glance, but it's actually O(n) because:
```go
for _, n := range nums {
    // This check is crucial for O(n) complexity
    if hashTable[n-1] {
        continue
    }
    // ...
}
```

The important part is the condition `if hashTable[n-1]`. This ensures that we only start counting sequences from the smallest number in each sequence. For example:

- If we have sequence [1,2,3,4]:
  - When n=1: we count the full sequence (1→2→3→4)
  - When n=2: we skip (because 1 exists)
  - When n=3: we skip (because 2 exists)
  - When n=4: we skip (because 3 exists)

This means that each number is only visited:
- Once during the hash table creation
- At most once when checking if it's the start of a sequence
- At most once when it's part of a sequence count

Therefore, even though we have nested loops, each element is processed a constant number of times, making the overall time complexity O(n).

This is much more efficient than a naive approach that would check every possible sequence from every number, which would be O(n²).

#### Space Complexity O(n)
- We use a hash table to store the numbers, so the space complexity is O(n).
