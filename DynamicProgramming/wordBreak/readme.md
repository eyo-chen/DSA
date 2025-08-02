## Dynamic Programming Example Walkthrough

Let's trace through `WordBreak3("leetcode", ["leet", "code"])` to see how the bottom-up approach builds the solution systematically.

**Initial Setup:**
- String: `"leetcode"` (length 8)
- Dictionary: `["leet", "code"]`
- DP table: `[true, false, false, false, false, false, false, false, false]`
  - Index represents position in string (0 to 8)
  - `dpTable[0] = true` (empty string base case)
```
["", 'l', 'e', 'e', 't', 'c', 'o', 'd', 'e']
[T,   F,   F,   F,   F,   F,   F,   F,   F]
```

**Step-by-step execution:**

**Position 0 (dpTable[0] = true):**
- Current position can be reached through valid segmentation
- Try word `"leet"` (length 4): 
  - Check if `s[0:4] == "leet"` → `"leet" == "leet"` ✓
  - Set `dpTable[4] = true`
- Try word `"code"` (length 4):
  - Check if `s[0:4] == "code"` → `"leet" == "code"` ✗
```
["", 'l', 'e', 'e', 't', 'c', 'o', 'd', 'e']
[T,   F,   F,   F,   T,   F,   F,   F,   F]
```
***Note that the index `4` in dpTable and input string(s) represent different things, but perfectly helps us to solve the problem.***<br>
***For `dpTable[4]=true`, it means we can safely reach to character `t`, so when we loop to `t` again, we know that we can try from this position***<br>
***For `s[4]="c"`, it means we can try to this speicific character to see if there's any word matches***<br>
***There's one index different between dpTable and string, and it's perfectly helps us to solve the problem.***

**Positions 1, 2, 3 (dpTable[1,2,3] = false):**
- These positions cannot be reached through valid segmentation
- Skip processing (continue to next position)

**Position 4 (dpTable[4] = true):**
- Current position can be reached (we set this in position 0)
- Try word `"leet"` (length 4):
  - Check if position 4+4=8 fits in string ✓
  - Check if `s[4:8] == "leet"` → `"code" == "leet"` ✗
- Try word `"code"` (length 4):
  - Check if position 4+4=8 fits in string ✓
  - Check if `s[4:8] == "code"` → `"code" == "code"` ✓
  - Set `dpTable[8] = true`
```
["", 'l', 'e', 'e', 't', 'c', 'o', 'd', 'e']
[T,   F,   F,   F,   T,   F,   F,   F,   T]
```

**Positions 5, 6, 7 (dpTable[5,6,7] = false):**
- These positions cannot be reached through valid segmentation
- Skip processing

**Final DP table state:** `[true, false, false, false, true, false, false, false, true]`

**Result:** `dpTable[8] = true` means the entire string `"leetcode"` can be segmented.

The DP approach systematically builds up the solution by marking which positions in the string can be reached through valid word segmentations. Each `true` value represents a position where we can successfully place the end of a valid word, creating a "stepping stone" for further segmentation.