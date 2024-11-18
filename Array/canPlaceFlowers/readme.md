# Problem Explanation
There are one key to note in this problem:<br>
We have to treat the index before the first one and after the last one is 0.<br>
For example,<br>
[0, 0, 1] -> we can place flower at index 0 because the index before the first one is 0.<br>
[1, 0, 0] -> we can place flower at index 2 because the index after the last one is 0.<br>


## Mutate the array
This solution is easy to understand<br>
We can iterate through the array,<br>
If the current value is 1, we can just update the pointer for two steps.<br>
If the current value is 0, we check the value before and after the current one.<br>
If both are 0, we can place a flower(mutate the value to 1) here and update the pointer for two steps.<br>

Let's walk through an example:<br>
[1, 0, 0, 0, 1] -> n = 1<br>
- i = 0
  - value = 1
  - update i to 2
- i = 2
  - value = 0
  - check the value before and after the current one
    - before = 0
    - after = 0
  - we can place a flower here and update the pointer for two steps
    - update value to 1
    - update i to 4
    - [1, 0, 1, 0, 1]
  - update i to 4
- i = 4
  - value = 1
  - update i to 6

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(n)


## Without mutate the array
This core logic is almost the same as the previous one, but we don't mutate the array.<br>
We just count the number of flowers we can place.<br>
If the current value is 1, we can just update the pointer for two steps.<br>
If the current value is 0, we check the value before and after the current one.<br>
If both are 0, we can increment the count and update the pointer for two steps.<br>
In the end, we just need to check if the count is greater than or equal to n.<br>

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)

## Using Math
Let's try to use some math formula to solve this problem.<br>

If the input is [0,0], how many flowers we can place?<br>
- add implicity zeros before and after the array
  - [0,0] -> [0,0,0,0]
- 1 flower
- [0,1] or [1,0]

If the input is [0,0,0], how many flowers we can place?<br>
- add implicity zeros before and after the array
  - [0,0,0] -> [0,0,0,0,0]
- 2 flowers
- [1,0,1]

If the input is [0,0,0,0], how many flowers we can place?<br>
- add implicity zeros before and after the array
  - [0,0,0,0] -> [0,0,0,0,0,0]
- 2 flowers
- [1,0,1,0] or [0,1,0,1]

If the input is [1,0,0,1], how many flowers we can place?<br>
- add implicity zeros before and after the array
  - [1,0,0,1] -> [0,1,0,0,1,0]
- zero flower

The math formula is (zeros - 1) / 2<br>
The idea is that we loop through the array<br>
If the current value is 0, we increment the zero counts<br>
If the current value is 1, we check the zero counts and use the formula to calculate the number of flowers we can place<br>
At this point, the zero counts represent the number of consecutive zeros before the current index<br>
After that, we reset the zero counts and continue the loop<br>

Note that we need to add the zero before the for loop and after the for loop<br>
Before the for loop, it represents that we add a zero before the first element<br>
After the for loop, it represents that we add a zero after the last element<br>

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)


# Gotcha About the Range Loop

```go
func main() {
	ans := []int{1, 1, 1, 5, 1, 1, 1}

	for i, f := range ans {
		fmt.Println("index", i, "value", f)

		if f > 2 {
			i++
			i++
		}
	}
}
```
Why the output is following?
```
index 0 value 1
index 1 value 1
index 2 value 1
index 3 value 5
index 4 value 1
index 5 value 1
index 6 value 1
```
I thought the index 4 and 5 will be skipped because of `i++` in the if statement.

Ah, this is a common misconception when working with Go's `range` loop. The behavior you're seeing is because **you cannot modify the loop variable `i` inside a `range` loop**. 

In a `range` loop, the loop variable `i` is actually a local variable that gets assigned a new value at the beginning of each iteration. Any modifications you make to `i` inside the loop body will be discarded when the next iteration begins.

Here's a clearer example to demonstrate:

```go
func main() {
    ans := []int{1, 1, 1, 5, 1, 1, 1}
    
    // This won't skip indices
    for i, f := range ans {
        fmt.Println("Using range:", i, f)
        i += 2  // This modification has no effect
    }
    
    fmt.Println("\n--- Compare with regular for loop ---\n")
    
    // This will skip indices
    for i := 0; i < len(ans); i++ {
        fmt.Println("Using regular for:", i, ans[i])
        if ans[i] > 2 {
            i += 2
        }
    }
}
```

If you want to skip indices, you should use a traditional `for` loop instead of `range`. The above code will output something like:

```
Using range: 0 1
Using range: 1 1
Using range: 2 1
Using range: 3 5
Using range: 4 1
Using range: 5 1
Using range: 6 1

--- Compare with regular for loop ---

Using regular for: 0 1
Using regular for: 1 1
Using regular for: 2 1
Using regular for: 3 5
Using regular for: 6 1
```

Notice how in the regular `for` loop, we actually skip indices 4 and 5 when we find the value 5.