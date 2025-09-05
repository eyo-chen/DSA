# Solution Explanation

## Problem Understanding
We need to find all cells from which water can flow to **both** the Pacific Ocean (top/left edges) and Atlantic Ocean (bottom/right edges). Water flows from higher or equal height cells to lower height cells.

## Key Insight: Reverse Thinking
Instead of checking if water can flow **from** each cell to both oceans (which would be computationally expensive), we reverse the problem:
- Start **from** the ocean edges and see which cells water can reach **backwards**
- If a cell is reachable from both ocean edges, then water from that cell can flow to both oceans

## Solution Approach

### Step 1: Identify Starting Points
- **Pacific Ocean**: All cells on the **top** row and **left** column
- **Atlantic Ocean**: All cells on the **bottom** row and **right** column

### Step 2: DFS from Ocean Edges
- Run DFS from each Pacific edge cell to find all cells reachable from Pacific
- Run DFS from each Atlantic edge cell to find all cells reachable from Atlantic
- During DFS, we can only move to cells with **higher or equal** height (reverse flow)

### Step 3: Find Intersection
- A cell can flow to both oceans if and only if it's reachable from both Pacific and Atlantic edges
- Iterate through all cells and check if they exist in both hash tables

