# Problem Explanation

This idea is actualy not that hard, the main point is to handle the edga case<br/>
When dealing with the problem of matrix(2D array), be careful when the row and col is out of the bound<br/>
- row = 0
- col = 0
- row = matrix.length
- col = matrix[0].length

Especially the last two<br/>
Be careful about these edge cases along with the process

The idea of solution is 
1. Traverse all the cell of input board
2. Find the first matching character (with the first character of input word)
3. If found one, we know that we can start the recursion on this character(row, col)
4. If not found one, we can just return false after nested for-loop
   => If we can't even find the first matching character, there's no way finding correct word search

For the recursive helper, it's quite easy<br/>
We just need to go four direction to keep the process of recursion<br/>
BUT, we need to check several conditions before do the further recursion <br/>
1. Row is within the bound
2. Col is within the bound
3. Character has not been searched
4. Character is equal to the target character

If match those conditions, then we can keep searching until we find the word or we can't find the word

Another point is how to mark the character as searched<br/>
We can
- Create a new board(2D array) to mark the character as searched
- Mark the character as searched by changing the character to another character


## Choices and Constraints

- **Choice:** Go top, right, bottom, left
- **Constraint:** 
  - row and col have to be within the bound
  - cannot reuse the same position in the board (the path cannot be duplicate)
  - the character has to be matching
- **Goal:** Find all the character of input word

# Complexity Analysis

r = the length of row of input board<br/>
c = the length of col of input board<br/>
n = the length of input word

## Time Complexity: O(r * c * 3^n)
- Ar worst case, we need to do the recursion for every cell of the board
  - r * c mean traverse all the cell of the board
- Branching Factor = 3
  - We can never move backwards when searching for the next character
- Depth = n
  - We need to find all the character of input word
- Each call stack = O(1)
    - Because we only do constant work

## Space Complexity: O(n)
- The deepest call stack is n