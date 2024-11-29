# Problem Explanation
It's important to understand what's the problem is asking.<br>

Let's break down Example 2:
```
grid = [
    [3,1,2,2],  // row 0
    [1,4,4,5],  // row 1
    [2,4,2,2],  // row 2
    [2,4,2,2]   // row 3
]
```

When we look at the columns (reading vertically):
```
Column 0: [3,1,2,2]
Column 1: [1,4,4,4]
Column 2: [2,4,2,2]
Column 3: [2,5,2,2]
```

The problem asks us to find pairs of (row, column) where the row and column contain exactly the same elements in the same order. Each pair is counted separately. Let's check each match:

1. Row 0 `[3,1,2,2]` matches Column 0 `[3,1,2,2]` → This is pair #1
2. Row 2 `[2,4,2,2]` matches Column 2 `[2,4,2,2]` → This is pair #2
3. Row 3 `[2,4,2,2]` matches Column 2 `[2,4,2,2]` → This is pair #3

Even though rows 2 and 3 are identical and they both match with the same column 2, each row-column match is counted as a separate pair. This is why we get 3 pairs total.<br>

Think of it this way: we're not grouping identical arrays together first. Instead, we're counting each individual row-column match as its own pair, even if some rows or columns are identical to others.<br>


Also note that having two identical rows by themselves doesn't count. The problem specifically asks for pairs of (row, column) that are equal. <br>

Let's illustrate with a simple example:
```
[
    [1,1],  // row 0
    [1,1]   // row 1
]
```
Even though row 0 and row 1 are identical (`[1,1]`), this doesn't count as a pair because we're only looking for matches between rows and columns.

For this example, let's look at the columns:
```
Column 0: [1,1]
Column 1: [1,1]
```

In this case, we would have 4 pairs because:
1. Row 0 `[1,1]` matches Column 0 `[1,1]`
2. Row 0 `[1,1]` matches Column 1 `[1,1]`
3. Row 1 `[1,1]` matches Column 0 `[1,1]`
4. Row 1 `[1,1]` matches Column 1 `[1,1]`

The problem is specifically looking for row-to-column matches, not:
- row-to-row matches
- column-to-column matches

So in the original example, even though rows 2 and 3 are identical (`[2,4,2,2]`), this similarity by itself doesn't contribute to the count. What matters is how many times a row matches with a column.

Now, let's see how we can solve this problem using a hash table.<br>
The idea is to use two separate for-loops:
1. The first for-loop will count the occurrences of each row pattern and store them in a hash table.
2. The second for-loop will check how many times each column pattern matches the stored row patterns in the hash table.
For example, let's see how this works with Example 2:
```
grid = [
    [3,1,2,2]
    [1,4,4,5]
    [2,4,2,2]
    [2,4,2,2]
]
```
After the first for-loop, the hash table will look like this:
```
{
    "[3,1,2,2]": 1,
    "[1,4,4,5]": 1,
    "[2,4,2,2]": 2
}
```

When we go through the second for-loop, we'll check how many times each column pattern matches the stored row patterns in the hash table.<br>
First, let's look at Column 0 `[3,1,2,2]`. It matches `[3,1,2,2]`, it's count is 1, which means that Column 0 can match with Row 0 only once. So we add 1 to our answer.<br>
Next, let's look at Column 1 `[2,4,4,4]`. It doesn't match with any row, so we add 0 to our answer.<br>
Next, let's look at Column 2 `[2,4,2,2]`. It matches `[2,4,2,2]`, it's count is 2, which means that Column 2 can match with Row 2 and Row 3 twice. So we add 2 to our answer.<br>
Next, let's look at Column 3 `[2,5,2,2]`. It doesn't match with any row, so we add 0 to our answer.
In the end, we get 3 pairs, which is the answer.<br>

There's a gotcha here.<br>
At first, I thought that we can convert the rows and columns to integers and use them as keys in the hash table. This is wrong. This won't work for multiple digits.<br>
For example, `[31,22]` and `[3,122]` would be considered identical if we convert them to integers.<br>
Therefore, we need to convert the rows and columns to strings.<br>
In Go, we can use `fmt.Sprint()` to convert a slice to a string.<br>
e.g. `slice = [1 2 3]` -> `string = "[1 2 3]"`

# Complexity Analysis
## Time Complexity: O(n²)
1. First loop (rows):
   - Iterates through n rows
   - For each row, `fmt.Sprint(row)` takes O(n) to convert the row to string
   - Total: O(n²)

2. Second loop (columns):
   - Iterates through n columns
   - For each column:
     - Building the column slice takes O(n)
     - `fmt.Sprint(col)` takes O(n) to convert to string
   - Total: O(n²)

Overall time complexity: O(n²)

## Space Complexity: O(n²)
1. HashMap storage:
   - Stores up to n different row patterns
   - Each pattern is a string of length O(n)
   - Total: O(n²)

2. Column slice:
   - Single slice of length n
   - O(n)

Overall space complexity: O(n²)

The space complexity is O(n²) primarily because in the worst case, all rows could be different, and each string representation of a row/column takes O(n) space. The hashmap could therefore store up to n different strings, each of length O(n).