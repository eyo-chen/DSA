# Trie (Prefix Tree) Implementation Notes

## 1. High Level Overview of Tries

A **Trie** (pronounced "try"), also called a prefix tree, is a tree-like data structure used to store and retrieve strings efficiently. Each node in the trie represents a single character, and paths from the root to various nodes form complete words or prefixes.

Key characteristics:
- The root node is empty and serves as the starting point
- Each edge represents a character
- Nodes can be marked to indicate the end of a valid word
- Children of a node share a common prefix (the path from root to that node)

For example, inserting "cat", "car", and "dog" would create:
```
        root
       /    \
      c      d
      |      |
      a      o
     / \     |
    t   r    g
```

## 2. Why Use Tries

Tries excel at prefix-based operations and offer several advantages:

**Efficient Prefix Matching**: Checking if any word starts with a given prefix takes O(m) time, where m is the prefix length, regardless of how many words are stored.

**Fast Lookups**: Word searches are proportional to the word length, not the number of words in the dataset.

**Space Efficiency for Common Prefixes**: Words sharing prefixes (like "cat" and "car") share nodes, reducing redundant storage.

**Common Use Cases**:
- Autocomplete and search suggestions
- Spell checkers and dictionaries
- IP routing tables
- Pattern matching in strings
- Implementing phone directories

## 3. How to Implement a Trie

The implementation consists of two main components:

### Node Structure
```go
type node struct {
    chars [26]*node  // Array of pointers to child nodes (for 'a'-'z')
    isEnd bool       // Flag to mark if this node completes a word
}
```

Each node contains:
- A fixed-size array (26 slots for lowercase letters 'a'-'z')
- A boolean flag indicating whether this node marks the end of a valid word

### Core Operations

**Insert**: Traverse the trie character by character, creating new nodes as needed. Mark the final node as a word ending.

**Search**: Follow the path corresponding to each character. Return true only if the path exists AND the final node is marked as a word ending.

**StartsWith**: Similar to search, but only checks if the prefix path exists, ignoring the `isEnd` flag.

### Key Implementation Details

- Character to index mapping: `char - 'a'` converts 'a' to 0, 'b' to 1, etc.
- Nil checks prevent accessing non-existent paths
- The `isEnd` flag distinguishes complete words from mere prefixes (e.g., "car" vs "cart")

## 4. Time and Space Complexity

### Time Complexity

| Operation | Complexity | Description |
|-----------|------------|-------------|
| **Insert** | O(m) | m = length of the word being inserted |
| **Search** | O(m) | m = length of the word being searched |
| **StartsWith** | O(m) | m = length of the prefix being checked |

All operations scale with the length of the input string, not the number of words stored in the trie.

### Space Complexity

**Per Node**: O(ALPHABET_SIZE) = O(26) for lowercase English letters

**Overall**: O(N × M × ALPHABET_SIZE) in the worst case, where:
- N = number of words
- M = average length of words
- ALPHABET_SIZE = 26

**Best Case**: When words share many common prefixes, space usage is significantly reduced since nodes are shared.

**Worst Case**: When all words have completely different prefixes, the trie approaches maximum space usage.

**Trade-off**: Tries sacrifice space for speed. The fixed-size array in each node uses memory even for unused characters, but enables O(1) child lookups.

### Optimization Notes

For sparse tries (where most character slots are empty), you could use a hashmap instead of a fixed array to reduce space complexity, though this slightly increases lookup time due to hashing overhead.