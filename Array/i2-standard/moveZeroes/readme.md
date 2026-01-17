# Problem Explanation

## Use Another Array
The idea is to use another temporary array to store the non-zero elements in order, and then append the zero elements to the end of the array.<br>
In the end, we just need to copy the elements from the temporary array to the original array.<br>

For example, input is [0, 1, 0, 3, 12]<br>
We first iterate through the array, and keep track two things:
- Count how many zero elements are there
- Store the non-zero elements in order to a temporary array

After iteration, we have the zeroCounts is 2, and the temporary array is [1, 3, 12]<br>
Then, we add two zeros to the end of the temporary array, so it becomes [1, 3, 12, 0, 0]<br>
Finally, we copy the elements from the temporary array to the original array, so the input array becomes [1, 3, 12, 0, 0]<br>

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(n)

## Keep Track of the Last Place of Non-Zero
The idea is to keep track of the what's the last place of non-zero elements, and also move the non-zero elements to the front of the array.<br>
For example, input is [0, 1, 0, 3, 12]<br>
We first iterate through the array, and when we see a non-zero element, we put it in the front of the array, and increment the pointer.<br>
Before iteration:
```
 p
[0, 1, 0, 3, 12]
```

After iteration:
```
           p
[1, 3, 12, 3, 12]
```
Now, it means that the element before pointer `p` is all non-zero elements in correct order<br>
So, we just need to fill the rest of the array with zeros<br>

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)

## Swap Non-Zero Elements to the Front
The idea is to keep swapping the non-zero elements to the front of the array<br>
We only need a variable `swapIndex` which represents the pointer that should be swapped with when we find a non-zero element<br>
For example, input is [0, 1, 0, 3, 12]<br>
At first, `swapIndex` is 0<br>
```
 p
[0, 1, 0, 3, 12]
```
Now, `swapIndex` is at 0 index, which means that when we find a non-zero element, we swap it with the 0 index<br>

First iteration:
```
 i
 p
[0, 1, 0, 3, 12]
```
current element is 0, so we skip it

Second iteration:
```
 p  i
[0, 1, 0, 3, 12]
```
current element is 1, so we swap it with the element at `swapIndex` index, and then increment `swapIndex` by 1<br>
```
    i
    p
[1, 0, 0, 3, 12]
```
Now, `swapIndex` is at 1 index, which means that when we find a non-zero element, we swap it with the 1 index<br>

Third iteration:
```
    p  i
[1, 0, 0, 3, 12]
```
current element is 0, so we skip it

Fourth iteration:
```
    p     i
[1, 0, 0, 3, 12]
```
current element is 3, so we swap it with the element at `swapIndex` index, and then increment `swapIndex` by 1<br>
```
       p  i
[1, 3, 0, 0, 12]
```

Fifth iteration:
```
       p     i
[1, 3, 0, 0, 12]
```
current element is 12, so we swap it with the element at `swapIndex` index, and then increment `swapIndex` by 1<br>
```
          p  i
[1, 3, 12, 0, 0]
```

We're done, because we've iterated through the whole array

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)
