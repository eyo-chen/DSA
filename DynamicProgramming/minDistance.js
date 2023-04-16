//////////////////////////////////////////////////////
// *** Edit Distance ***
//////////////////////////////////////////////////////
/**
 * @param {string} word1
 * @param {string} word2
 * @return {number}
 */
/*
At first, this question is extremely hard
But if we can start solving this problems by small and simple inputs
And try to figure the pattern
This problem is actually quite easy to understand


Let's start with the small inputs
EX1
w1: "abc"
w2: ""
=> 3, cuz we remove three times

EX2
w1: ""
w2: "abc"
=> 3, cuz we Insert three times

EX3
w1: "abe"
w2: "cde"
=> we try to see this problem from the end to the beginning
=> imagine we set two pointers and the end of the character
i = 2, k = 2
   i
"abe"
"cde"
   k
=> the last character of both inputs are the same, so we can just ignore "e", and ask the sub problem "ab" and "cd"
(updating two pointer together)
i = 1, k = 1
  i 
"abe"
"cde"
  k
=> This is the first important pattern to solve this problem
=> If the character is the same, we can just ignore that character, and do NOT need to do any operation, and keep looking at the rest of the characters
=> But now "b" and "d" are different, so we can have three choices
   1) insert -> "abd", "cd"
      => the sub problem become "ab" and "c"
      (i remains the same, update k)
      i = 1, k = 0
        i 
      "abe"
      "cde"
       k
   2) delete -> "a", "cd"
      => the sub problem become "a" and "cd"
      (k remains the same, update i)
      i = 0, k = 1
       i 
      "abe"
      "cde"
        k
   2) replace -> "ad", "cd"
      => the sub problem become "a" and "c" cuz now can ignore "d"
      (update both pointer)
      i = 0, k = 0
       i 
      "abe"
      "cde"
       k
=> This is another important pattern to solve this problem
=> 1) insert:  (i, k - 1)
   2) delte:   (i - 1, k)
   3) replace: (i - 1, k - 1)


Once we understand the patter, we can build the DP table
For examplem,
word1 = "horse", word2 = "ros"

1. Build DP table, note that extra row and column. Both of them represent the empty string, which is kinda like our base case
Also, "r" column represent subString "hor", "s" row represent subString "ros"
    ""   h   o   r   s   e
""
r
o
s

2. Initialize first row and first column with intuition
    ""   h   o   r   s   e
""   0   1   2   3   4   5
r    1
o    2
s    3
It's straightforward.
How many operation is needed to convert "hor" to "" -> Do delete three times
How many operation is needed to convert "" to "ros" -> Do insert three times

3. Again, it's very important to understand what those cells represent
    ""   h   o   r   s   e
""   0   1   2   3   4   5
r    1   *      
o    2           
s    3
This cell represent "How many operation is need to convert "h" to "r"? "

    ""   h   o   r   s   e (k)
""   0   1   2   3   4   5
r    1          
o    2           *
s    3
(i)
This cell represent "How many operation is need to convert "hor" to "ro"? "

4. How those patterns we've seen above apply to this DP table?
   Use above case as example, convert "hor" to "ro" (i, k) = (3, 2)
   1) insert: "horo", "ro"
   => now "o" are guarantee the same, so we can ignore that
   => sub problem become "hor", "r"
   => (i - 1, k) = (2, 3)

   2) delete: "ho", "ro"
   => sub problem become "ho", "ro"
   => note that it just accidently that "o" are the same, it could not be the same, so we can't just say ignore "o". We can only ignore that after insert and replace because we can guarantee to have the same character after these two operations
   => (i, k - 1) = (2, 2)

   3) replace: "hoo", "ro"
   => now "o" are guarantee the same, so we can ignore that
   => sub problem become "ho", "r"
   => (i - 1, k - 1) = (1, 2)

   Now the whole pattern and picture is very clear
            ""    h     o     r     s     e   (k)

       ""   0     1     2     3     4     5

       r    1        replace insert  

       o    2        delete   *

       s    3

5. For each cell, we just need to find the minimum of thoes three neighboring cells if both of character are different, and add 1 because it means we do one more operation

6. Don't forget we just need to check (i - 1, k - 1) if both character are the same
For example, at the point of this cell,
it represent "ho" convert to "ro" (2, 2)
    ""   h   o   r   s   e
""   0   1   2   3   4   5
r    1    
o    2       *
s    3
It's obvious that "o" are the same, so we can just ignore that, and ask the sub problem how to convert "h" to "r", which is essentially what's (1, 1) are asking
*/
function minDistance(word1, word2) {
  const table = [];

  // build the DP table
  for (let i = 0; i < word2.length + 1; i++) {
    const row = [];

    for (let k = 0; k < word1.length + 1; k++) {
      // initialize the base
      if (i === 0) {
        row.push(k);
      } else if (k === 0) {
        row.push(i);
      } else row.push(null);
    }

    table.push(row);
  }

  for (let i = 1; i < word2.length + 1; i++) {
    for (let k = 1; k < word1.length + 1; k++) {
      /*
        Note here
        Because the table include empty string
        So when the indices of table is (2,2)
        It basically means the both first character, so we need to minus one
        */
      if (word1[k - 1] === word2[i - 1]) table[i][k] = table[i - 1][k - 1];
      else {
        table[i][k] = Math.min(
          table[i - 1][k] + 1, // insert
          table[i][k - 1] + 1, // delete
          table[i - 1][k - 1] + 1 // replace
        );
      }
    }
  }

  return table[word2.length][word1.length];
}

// console.log(minDistance('horse', 'ros'));
