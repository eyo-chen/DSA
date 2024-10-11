# Problem Explanation

## Compare Each Word with Dictionary
This is the solution I came up with, and it's the most straightforward approach.<br>

The idea is<br>
For each word in the sentence, we compare it with each word in the dictionary.<br>
If a word matches the word in the dictionary, we replace the word in the sentence with the word in the dictionary.<br>

For comparison logic, we just need to loop through the dictionary word.<br>
Compare character by character and if all characters are the same, we replace the word in the sentence with the word in the dictionary.<br>
Because the word in the dictionary is shorter than the word in the sentence, we only need to loop through the dictionary word.<br>

Because the problem requires to use the shortest dictionary word to replace the word in the sentence, we need to sort the dictionary by the length of the word at the beginning.<br>

### Complexity Analysis
#### Time Complexity O(N * M * L)
- N is the number of words in the dictionary
- M is the number of words in the sentence
- L is the average length of the words in the dictionary
- We have three nested loop here
  - First loop is for each word in the sentence
  - Second loop is for each word in the dictionary
  - Third loop is for each character in the dictionary word

#### Space Complexity O(1)
- We don't use any additional space


## Use Hash Table to Store Dictionary
This is another straightforward approach.<br>

The idea is<br>
We first put all the words in the dictionary into a hash table for quick lookup.<br>
Then for each word in the sentence<br>
We check all the possible prefix of the word<br>
If the prefix is in the hash table, we replace the word with the prefix

Because we loop through all the possible prefix for a word, we guarantee to use the shortest dictionary word to replace the word in the sentence.

### Complexity Analysis
#### Time Complexity O(N * M * M)
- N is the number of words in the setence
- M is the average length of the words in the sentence
- We have two nested loop here
  - First loop is for each word in the sentence
  - Second loop is for each prefix of the word
  - For the second loop, we need to create a substring of the word, so it's O(M^2)

#### Space Complexity O(D)
- D is the number of words in the dictionary
- We use a hash table to store the dictionary words
