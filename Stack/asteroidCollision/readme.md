# Problem Explanation
To solve this problem, we have to figure out the different cases<br>
- When stack is empty, we just add the current value to the stack no matter what
- When current value is positive, we just add it to the stack
  - Why?
  - When the last value in the stack is positive, the current value will never collide because both of them are moving to right
  - When the last value in the stack is negative, the current value will also never collide because the last value is moving to left and the current value is moving to right
- When the last value in the stack is negative, we just add the current value to the stack
  - Why?
  - When the current value is negative, they will never collide because they are both moving to left
  - When the current value is positive, they will also never collide because the last value is moving to left and the current value is moving to right

To sum up with the above cases, we only need to check if the collision when all the following conditions are true:
- The stack is not empty
- The current value is negative
- The last value in the stack is positive

Now, let's figure out the collision cases:
(Remember that the current value is negative and the last value in the stack is positive)
- If both of them have the same size, both of them will be destroyed and we can stop checking the next values in the stack
- If the last value in the stack is bigger, the current value will be destroyed and we can stop checking the next values in the stack
- If the last value in the stack is smaller, the last value will be destroyed and we need to check the next value in the stack

## Complexity Analysis
### Time Complexity O(n)
- We have to iterate through all the values in the array once

### Space Complexity O(1)
- We didn't consider the space used by the stack