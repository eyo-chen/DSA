# Problem Explanation
The hard part of this problem is to avoid duplicates.<br>

Let's first talk about the overall idea to solve this problem.<br>
We first sort the array, and take each element as the target, then use twoSum logic to the rest of the array.<br>
For example, nums = [-4,-1,-1,0,1,2]<br>
We first take -4 as the target, then we use twoSum logic to the rest of the array, which is [-1,-1,0,1,2]<br>
We then take -1 as the target, and use twoSum logic to the rest of the array, which is [-1,0,1,2]<br>
So on and so forth.<br>

Okay, let's say the nums = [-1,-1,-1,0,1,1,2,2]<br>
ptr -> current pointer<br>
l -> left pointer<br>
r -> right pointer<br>
```
  i  l                   r
[-1, -1, -1, 0, 1, 1, 2, 2]
```
Now, we find one answer, which is [-1, -1, 2]<br>
The problem is how should we update left and right pointer?<br>
- Can we just break here, and increment i?<br>
  - No, because we might miss other answers, for example, [-1, 0, 1]<br>
- Can we just increment left and right pointer?<br>
  - No, because we might get duplicates, for example:
```
  i       l           r
[-1, -1, -1, 0, 1, 1, 2, 2]
```
It's the same as the previous one, we will get [-1, -1, 2] again.<br>
So, we have to move both pointers inward until the value is different.<br>
We have to move left pointer until the value is not -1<br>
We have to move right pointer until the value is not 2<br>
So, after updating left and right pointer, we have:
```
  i          l     r
[-1, -1, -1, 0, 1, 1, 2, 2]
```

Now, how should we update i?<br>
- Can we just increment i?<br>
  - No, because we might get duplicates, for example:
```
      i  l               r
[-1, -1, -1, 0, 1, 1, 2, 2]
```
After only updating i, we get [-1, -1, 2] again.<br>
So, we have to update i until the value is different.<br>
In this case, we update i until the value is not -1.<br>
```
             i  l        r
[-1, -1, -1, 0, 1, 1, 2, 2]
```

# Complexity Analysis
## Time Complexity O(n^2)
- Sorting the array takes O(n log n) time
- The twoSum part takes O(n) time
- We do this for each element, so it's O(n^2)

## Space Complexity O(1)
- It depends on the sorting algorithm, if it's an in-place sorting algorithm, then it's O(1), otherwise it's O(n)
