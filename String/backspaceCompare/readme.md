# Problem Explanation

## Reading Backwards
The idea is pretty simple. We can read the string from the end to the beginning.

If we encounter a `#`, we increase the `skip` counter.<br>
If we encounter a character, and `skip` is greater than 0, we decrease the `skip` counter. It means this character should be ignored.<br>
If `skip` is 0, we add the character to the result.<br>

At the end, we compare the two results and return true if they are the same, false otherwise.

Note that the final result string is reversed because we build the string from the end to the beginning.<br>
However, it's okay because we are comparing the two results, and the order does not matter.

### Complexity Analysis
#### Time Complexity O(n + m)
- n is the length of the first string.
- m is the length of the second string.
- We iterate through the first and second string once.

#### Space Complexity O(n + m)
- n is the length of the first string.
- m is the length of the second string.
- For each string, we build a result string to compare.

## Using Stack
The idea is to use a stack to build the final result string.<br>

When building the stack,<br>
If the current character is not `#`, we push it to the stack.<br>
If the current character is `#`, we pop the stack. (Also, need to check if the stack is not empty before popping)

At the end, we compare the two stacks and return true if they are the same, false otherwise.

### Complexity Analysis
#### Time Complexity O(n + m)
- n is the length of the first string.
- m is the length of the second string.
- We iterate through the first and second string once.

#### Space Complexity O(n + m)
- n is the length of the first string.
- m is the length of the second string.
- For each string, we use a stack to build the result string.

### Two Pointers
This solution is a little bit tricky, but still understandable.<br>

Let's walk through the though process.<br>
Suppose I have following strings: "ab#c" and "ad#c"<br>
Let's say the pointers are both at the end of the string.
```
         p
a  b  #  c

         p
a  d  #  c
```
Both characters are the same, so we move the pointers to the left.

```
      p
a  b  #  c

      p
a  d  #  c
```
When character is "#", we need to skip the next character.
So we move the pointer to the left two times.

```
p
a  b  #  c

p
a  d  #  c
```
Now, the characters are also the same

It seems that it's straightforward to just move the pointers to the left.<br>
When we encounter "#", we just need to skip the next character.<br>
Is that correct?<br>

Let's consider another example: "ab##" and "c#d#"<br>
```
         p
a  b  #  #

         p
c  #  d  #
```
For first string, it's obvious that we need to skip the character two times.<br>
But how about second string?<br>
After we skip the 'd', the pointer is at the first '#'.<br>
We should keep skiping the next character, right?<br>
Yes!!!<br>

Let's see how we should update the pointer.<br>
When the current character is NOT '#', we just move the pointer to the left.<br>
When the current character is '#', we need to skip until<br>
(1) the current character is not '#'<br>
(2) skip variable is 0

Let's use "ab##" as example.
```
         p
a  b  #  #
```
Current character is '#', init skip to 1
- skip = 1
- move pointer to the left

```
      p
a  b  #  #
```
Current character is '#', increase skip by 1
- skip = 2
- move pointer to the left

```
   p
a  b  #  #
```
Current character is NOT '#', decrease skip by 1<br>
- skip = 1
- move pointer to the left

```
p
a  b  #  #
```
Current character is NOT '#', decrease skip by 1<br>
- skip = 0
- move pointer to the left

Now, we get the final result string.<br>

Let's see another example, "c#d#"<br>
```
         p
c  #  d  #
```
Current character is '#', init skip to 1
- skip = 1
- move pointer to the left

```
      p
c  #  d  #
```
Current character is NOT '#', decrease skip by 1.<br>
We can't stop here because skip variable is still greater than 0.<br>
- skip = 0
- move pointer to the left

```
   p
c  #  d  #
```
Current character is '#', increase skip by 1<br>
- skip = 1
- move pointer to the left

```
p
c  #  d  #
```
Current character is NOT '#', decrease skip by 1<br>
- skip = 0
- move pointer to the left

Now, we get the final result string.<br>

So, the idea to update the pointer is:<br>
When encounter '#'<br>
1. Init skip to 1<br>
   - represents that we have to skip one character<br>
2. Move pointer to the left<br>
   - it means we are done considering the current character('#')<br>
3. Keep the while loop when<br>
    - current pointer is greater than or equal to 0<br>
      - when pointer is less than 0, we don't need to consider the rest of the characters<br>
    - skip is greater than 0 OR current character is '#'<br>
      - this two conditions has to consider together<br>
      - when skip is greater than 0, that means we have to skip current character<br>
      - when current character is '#', we have to go into while-loop to increase skip variable<br>
4. Inside the while loop,<br>
    - if current character is '#', increase skip by 1<br>
    - if current character is NOT '#', decrease skip by 1<br>
      - it means skipping current character<br>
    - move pointer to the left<br>


After moving the pointer for two strings, we have the following situation:
- Both pointers are less than 0
  - It means both strings are done skipping characters
  - So, we return true
- One of the pointers is less than 0
  - It means one string is done skipping characters, but the other string still has characters to consider
  - So, we return false
- Both pointers are greater than or equal to 0, compare the characters
  - If the characters are the same, move the pointers to the left
  - If the characters are not the same, return false

### Complexity Analysis
#### Time Complexity O(n + m)
- n is the length of the first string.
- m is the length of the second string.
- We iterate through the first and second string once.

#### Space Complexity O(1)
- We only use a few variables to keep track of the pointers and the skip variable.
