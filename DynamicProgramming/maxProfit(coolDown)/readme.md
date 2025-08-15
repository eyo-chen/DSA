# Recursive Approach

To solve this problem, we can model this as a decision problem: "What are the choices at any given day?"<br>
Then, we realize that our available choices depend on the current state.<br>

For any give day, we only have two states:<br>
1. **Holding stock**:
   - When holding a stock, we can either (1) keep holding it or (2) sell it.
2. **Not holding stock**
   - When not holding a stock, we can either (1) buy it or (2) do nothing.

Let's look more closely.<br>
- When we are in **Holding stock** state
  - Keep holding it -> stay in same state
  - Sell it -> transition to **Not holding stock** state, and get the profit, but with ***cooldown***, which means that we need to skip the next day.
- When we are in **Not holding stock** state
  - Do nothing -> stay in same state
  - Buy Stock -> transtition to **Holding stock**, and subtract the cost

## Complexity Analysis
### Time Complexity O(2^n)
- In worst case, we explore 2 choices at each day for n days

### Space Complexity O(n)
- Maximum recursion depth is n (the length of prices array)



# DP Approach
Let's see how to solve this problem using dynamic programming.<br>

Think of this problem like we're a trader who has a very specific rulebook. At any moment in time, we can only be in one of three situations:<br>
1. **You own stock** (we call this "hold" state)<br>
2. **You just sold stock today** (we call this "sold" state)<br>
3. **You don't own stock and didn't sell today** (we call this "rest" state)<br>

The key insight is that the cooldown rule creates a forced sequence: if we sell today, we MUST rest tomorrow. We cannot go directly from selling to buying.<br>

We might wonder: why not just track "have stock" vs "don't have stock"?<br>
The reason is the cooldown constraint creates two different types of "not having stock":<br>
- **Just sold**: We don't have stock, but we're in cooldown (can't buy tomorrow)
- **Resting**: We don't have stock, but we're free to buy if we want<br>
This distinction is crucial because it determines what actions are available to us<br>


Let's explain each transition rule by thinking through what decisions lead to each state:
## Getting to "Hold" State
```go
hold[i] = max(hold[i-1], rest[i-1] - prices[i])
```
To hold stock on day i, we have exactly two options:
- **Option 1**: We already held stock yesterday, so we just keep holding it. Our profit stays the same: `hold[i-1]`
- **Option 2**: We were resting yesterday (not in cooldown), so we can buy today. Our new profit becomes: `rest[i-1] - prices[i]` (previous profit minus what we pay for the stock)

We choose whichever option gives us more profit.

## Getting to "Sold" State  
```go
sold[i] = hold[i-1] + prices[i]
```
This one is simpler. To sell today, we must have held stock yesterday. There's only one way to get here: take yesterday's holding profit and add today's selling price.

## Getting to "Rest" State
```go
rest[i] = max(rest[i-1], sold[i-1])
```
To be resting today, we have two options:
- **Option 1**: We were already resting yesterday, so we continue resting. Profit stays: `rest[i-1]`
- **Option 2**: We sold yesterday, so today we're forced into cooldown (which is a form of resting). Our profit is: `sold[i-1]`

Again, we get the maximum of these two scenarios.

## Walking Through the Example Step by Step

Let's trace through `prices = [1, 2, 3, 0, 2]` like we're living through each day:

**Day 0 (price = 1):**
- We start with no money and no stock
- Our only option is to buy or do nothing
- If we buy: hold = -1, sold = 0 (impossible), rest = 0
- Best strategy: We can either buy for -1 profit or rest for 0 profit

**Day 1 (price = 2):**
Now we have choices based on yesterday's state:
- **To hold today**: Either keep holding (-1) OR buy today after resting (0 - 2 = -2). Best: -1
- **To sell today**: We can only sell if we held yesterday: -1 + 2 = 1
- **To rest today**: Either keep resting (0) OR cooldown after selling (impossible since we didn't sell yesterday). Best: 0

So after day 1: hold = -1, sold = 1, rest = 0

**Day 2 (price = 3):**
- **To hold today**: Keep holding (-1) OR buy after resting (0 - 3 = -3). Best: -1
- **To sell today**: Sell the stock we held: -1 + 3 = 2
- **To rest today**: Keep resting (0) OR cooldown after yesterday's sale (1). Best: 1

After day 2: hold = -1, sold = 2, rest = 1

**Day 3 (price = 0):**
- **To hold today**: Keep holding (-1) OR buy after resting (1 - 0 = 1). Best: 1 (buy today!)
- **To sell today**: Sell yesterday's holding: -1 + 0 = -1 (terrible idea)
- **To rest today**: Keep resting (1) OR cooldown after yesterday's sale (2). Best: 2

After day 3: hold = 1, sold = -1, rest = 2

**Day 4 (price = 2):**
- **To hold today**: Keep holding (1) OR buy after resting (2 - 2 = 0). Best: 1
- **To sell today**: Sell yesterday's holding: 1 + 2 = 3
- **To rest today**: Keep resting (2) OR cooldown after yesterday's sale (-1). Best: 2

Final result: hold = 1, sold = 3, rest = 2

The maximum profit is 3 (from the sold state), which corresponds to: buy at 1, sell at 3 (profit +2), cooldown, buy at 0, sell at 2 (profit +2), but wait... that's 4, not 3.

Let me recalculate more carefully: buy at 1, sell at 2 (profit +1), cooldown, buy at 0, sell at 2 (profit +2). Total: 3. That matches!

## Complexity Analysis
### Time Complexity O(n)
### Space Complexity O(n)
