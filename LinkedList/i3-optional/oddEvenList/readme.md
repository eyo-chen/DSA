# Problem Explanation

## Iterative Process
### First Solution: use index to separately determine odd and even nodes
Suppose input is [1,2,3,4,5]<br/>
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x<br/>
**Note that `head` is just a pointer to the first node**<br/>

`ListNode* ptr = head;`<br/>
```
       ptr
        |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x
```
It means make `ptr` to point to whatever `head` is pointing to<br/>

`ListNode* oddList = new ListNode();`<br/>
`ListNode* evenList = new ListNode();`<br/>
```
       ptr
        |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x
             oddList
                | 
odd linklist -> x

              evenList
                 |
even linklist -> x
```
`oddList` and `evenList` are just two pointers which point to nothing<br/>
Note that "odd linklist" is just the name of this linklist<br/>
Same for "even linklist"<br/>
And `oddList` and `evenList` are two pointers variable<br/>

`ListNode* evenHead = evenList;`<br/>
```
       ptr
        |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x

             oddList
                | 
odd linklist -> x

              evenList
                 |
even linklist -> x
                 |
             evenHead   
```
It means make `evenHead` to point to whatever `evenList` is pointing to<br/>

index = 1<br/>
`oddList->next = ptr;`<br/>
```
       ptr
        |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x

             oddList
                | 
odd linklist -> x -> 1

              evenList
                 |
even linklist -> x
                 |
             evenHead   
```
`oddList->next` is as same as `(*oddList).next`<br/>
`(*oddList)` means dereference the pointer `oddList`<br/>
Simply said, let me know what's `oddList` pointing to<br/>
Which is `x`<br/>
`(*oddList).next` means let me know what's `oddList` pointing to, and then access the `next` member<br/>
**So it makes the `x` of `oddList` to point to what `ptr` is pointing to**<br/>
Which is `1`<br/>

`oddList = ptr;`<br/>
```
       ptr
        |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x

                  oddList
                     | 
odd linklist -> x -> 1

              evenList
                 |
even linklist -> x
                 |
             evenHead  
```
It means make `oddList` to point to whatever `ptr` is pointing to<br/>
**So it makes the `oddList` to point to `1`**<br/>

`ptr = ptr->next;`<br/>
```
            ptr
             |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x

                  oddList
                     | 
odd linklist -> x -> 1

              evenList
                 |
even linklist -> x
                 |
             evenHead  
```
It means make `ptr` to point to whatever `ptr->next` is pointing to<br/>

index = 2<br/>
`evenList->next = ptr;`<br/>
```
            ptr
             |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x

                  oddList
                     | 
odd linklist -> x -> 1

              evenList
                 |
even linklist -> x -> 2
                 |
             evenHead
```
It means let me know what `eventList` pointing to, and then make that thing's `next` to point to whatever `ptr` is pointing to<br/>
**So it makes the `x` of `eventList` to point to what `ptr` is pointing to**<br/>
Which is `2`<br/>

`eventList = ptr;`<br/>
```
            ptr
             |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x

                  oddList
                     | 
odd linklist -> x -> 1

                   evenList
                      |
even linklist -> x -> 2
                 |
             evenHead
```
It means make `eventList` to point to whatever `ptr` is pointing to<br/>
**So it makes the `eventList` to point to `2`**<br/>
Note that `eventHead` is still pointing to `x`<br/>

`ptr = ptr->next;`<br/>
```
                 ptr
                  |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x

                  oddList
                     | 
odd linklist -> x -> 1

                   evenList
                      |
even linklist -> x -> 2
                 |
             evenHead
```
It means make `ptr` to point to whatever `ptr->next` is pointing to<br/>

index = 3<br/>
`oddList->next = ptr;`<br/>
`oddList = ptr;`<br/>
`ptr = ptr->next;`<br/>
```
                      ptr
                       |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x

                      oddList
                          | 
odd linklist -> x -> 1 -> 3

                   evenList
                      |
even linklist -> x -> 2
                 |
             evenHead
```

index = 4<br/>
`evenList->next = ptr;`<br/>
`evenList = ptr;`<br/>
`ptr = ptr->next;`<br/>
```
                           ptr
                            |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x

                      oddList
                          | 
odd linklist -> x -> 1 -> 3

                       evenList
                           |
even linklist -> x -> 2 -> 4
                 |
             evenHead
```

index = 5<br/>
`oddList->next = ptr;`<br/>
`oddList = ptr;`<br/>
`ptr = ptr->next;`<br/>
```
                                ptr
                                 |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x

                            oddList
                               | 
odd linklist -> x -> 1 -> 3 -> 5

                       evenList
                           |
even linklist -> x -> 2 -> 4
                 |
             evenHead
```

Break the while loop because `ptr` is pointing to `nullptr`<br/>

`evenList->next = nullptr;`<br/>
```
                                ptr
                                 |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x

                            oddList
                               | 
odd linklist -> x -> 1 -> 3 -> 5

                       evenList
                           |
even linklist -> x -> 2 -> 4 -> x
                 |
             evenHead
```
This is the key step<br/>
Before this step, `oddList` is pointing to `5`<br/>
Without this step, the end result will be something like following<br/>
```
head -> 1 -> 3 -> 5 -> 2 -> 4
                  |  <-  <- |
```
There is a cycle<br/>

`oddList->next = evenHead->next;`<br/>
Look at the following picture<br/>
```
                                ptr
                                 |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x

                            oddList
                               | 
odd linklist -> x -> 1 -> 3 -> 5

                       evenList
                           |
even linklist -> x -> 2 -> 4 -> x
                 |
             evenHead
```
`evenHead` is still pointing to `x`<br/>
And `x.next` is pointing to `2`<br/>
After while-loop, `oddList` is pointing to `5`<br/>
So `oddList->next = evenHead->next;` makes `5.next` to point to `x.next`<br/>
Which means `5.next` is pointing to `2`<br/>
```
head -> 1 -> 3 -> 5 -> 2 -> 4 -> x
```

Because `head` is never changed, so it's still pointing to `1`<br/>
And just return `head`<br/>

### Second Solution: use two pointers to wire up the odd and even list
Instead of using two linklist and index, we can use two pointers to wire up the odd and even list along with the process<br/>
This solution is not as intuitive as the first solution<br/>
But it's good example to illustrate how to wire up linklist<br/>

`ListNode* oddPtr = head;`<br/>
`ListNode* evenPtr = head->next;`<br/>
`ListNode* evenHead = evenPtr`<br/>
```
      oddPtr
        |
head -> 1 -> 2 -> 3 -> 4 -> 5 -> x
             | 
          evenPtr
          evenHead
```

`oddPtr->next = evenPtr->next;`
```
      oddPtr
        |   ----- \
head -> 1  /  2 -> 3 -> 4 -> 5 -> x
              | 
           evenPtr
           evenHead
```
`evenPtr->next;` means let me know what `evenPtr` is referencing to, and then check its `next` member<br/>
`evenPtr` is referencing to `2`<br/>
And `2` is pointing to `3`<br/>
So `evenPtr->next` is pointing to `3`<br/>
Again, `oddPtr` is referencing to `1`<br/>
`oddPtr->next = evenPtr->next;` means makes `1.next` to point to `3`<br/>

`oddPtr = oddPtr->next;`
```
                 oddPtr
                   | 
            ----- \
head -> 1  /  2 -> 3 -> 4 -> 5 -> x
              | 
           evenPtr
           evenHead
```
Let `oddPtr` referencing to whatever `oddPtr->next` is referencing to<br/>
`oddPtr->next` is referencing to `3`<br/>

`evenPtr->next = oddPtr->next;`
```
                 oddPtr
                   | 
            ----- \
head -> 1  /  2  \ 3 -> 4 -> 5 -> x
              |   ----- /
           evenPtr
           evenHead
```
`oddPtr` is referencing to `3`<br/>
`oddPtr->next` is referencing to `4`<br/>
So `evenPtr->next` is pointing to `4`<br/>

`evenPtr = evenPtr->next;`
```
                 oddPtr
                   | 
            ----- \
head -> 1  /  2  \ 3 -> 4 -> 5 -> x
              |   ----- /
              |      evenPtr
           evenHead
```
Let `evenPtr` referencing to whatever `evenPtr->next` is referencing to<br/>
`evenPtr->next` is referencing to `4`<br/>
Note that `evenHead` is still referencing to `2`<br/>

`oddPtr->next = evenPtr->next;`<br/>
`oddPtr = oddPtr->next;`
```
                          oddPtr
                            | 
            ----- \   ----- \
head -> 1  /  2  \ 3 / 4 -> 5 -> x
              |   ---- /
              |      evenPtr
           evenHead
```
First it makes `3.next` to point to `4.next`<br/>
Then it makes `oddPtr` to point to `4.next`<br/>
And `4.next` is pointing to `5`<br/>

`evenPtr->next = oddPtr->next;`<br/>
`evenPtr = evenPtr->next;`
```
                         oddPtr
                           | 
            ----- \   ---- \
head -> 1  /  2  \ 3 / 4 \ 5 -> x
              |   ---- / ----- /
              |              evenPtr
           evenHead
```
First it makes `4.next` to point to `5.next`<br/>
Then it makes `evenPtr` to point to `5.next`<br/>
And `5.next` is pointing to `x`<br/>

Because `evenPtr` is pointing to `x`, so break the while loop<br/>

`oddPtr->next = evenHead;`
```
head -> 1 -> 3 -> 5 -> 2 -> 4 -> x
```
Make `5.next` to point to `evenHead`<br/>


## Complexity Analysis
### Time Complexity: O(n)

### Space Complexity: O(1)