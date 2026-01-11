# Solution Explanation

## Two Pointer With nonZeroCount
The idea is very simple<br>
- Init two pointers, `left` and `right` for calculating the size of window
- When the current value is non-zero, update `nonZeroCount` and `right` pointer
  - `nonZeroCount` helps us to keep track how many non-zero value we've seen so far(or how many non-zero in the current window)
  - So that when we hit zero in the future, we know how to shrink the window size
- When the current value is zero, we want to do two thing
  - Shrink the window size to make sure there's no non-zero value in the current window(`nonZeroCount` helps us)
  - Update the answer


## Two Pointer
This solution is similar to the above one, but without `nonZeroCount`
- Init two pointers, `left` and `right` for calculating the size of window
- When the current value is non-zero, we just simply update `right` pointer
- When the current value is zero, we
  - update the `left` pointer to the position of `right` pointer
  - Keep updating `right` pointer until it's non-zero value
    - For each loop, we update the answer


## Clever Approach
This solution doesn't use two pointer approach, instead it uses a `zeroCount` to count how many zero we've seen so far<br>
- When value is zero, we update the `zeroCount` and answer
- When value is non-zero, we just reset the `zeroCount` 