# Problem Explanation

The key point of this problem is
- Each element can only be used once
- The result can't have duplicate combinations

It's obvious that we can't use the regular recursion approach<br/>
There're multiple ways to solve this problem<br/>
Here list two of them

## Use hashTable and unique vector
This solution is straightforward<br/>
We use two techniques to avoid duplicate combinations
1. Use hashTable to keep tracking the frequency of each element
2. Use unique vector to choose from

The iteration logic is similar to the combinationSum problem<br/>
***Once we choose an element, we can't choose the element before it***

### Recursive Tree Visualization
candidates = [2,1,2], target = 5<br/>
The iterative vector is [1,2]
<pre>
                                                                                  5
                                    1                                                                        2
                                   1,2                                                                       2,2
                                  1,2,2
</pre>
Look at the position of 2 <br/>
We can't choose 1 again after we choose 2<br/>

Look at the position of 1,2<br/>
We can't choose 1 because the frequency is 0<br/>
We can choose 2 because the frequency remains 1

### Complexity Analysis

n = the legnth of candidates array<br/>
t = the target<br/>

#### Time Complexity: O(n ^ t)
- Branching Factor = n
  - At worst, we can decompose from index 0 to n - 1
- Depth = t
  - At worst, we can decompose t times (if the smallest element is 1)
- Each call stack = O(n)

#### Space Complexity: O(t)


## Without hashTable and unique vector

The key point of this solution to avoid duplicate combinations is<br/>
***If I'm not the first element of the iteration in this call stack, and I'm as same as the previous one, I know I'm a duplicate one, so I can't be chosen*** <br/>
For example, candidates is [1,2,2,2,5] and target is 5<br/>
We can see the candidates as [1, 2a, 2b, 2c, 5]<br/>
If we're at the call stack where the starting index is 2a<br/>
`for i = 1; i < candidates.length; i++`<br/>
It's okay to choose 2a to explore the next call stack<br/>
But it's not okay to choose 2b and 2c<br/>

Note that we pass `i + 1` to the next call stack<br/>
Why is different from the previous solution?<br/>
In previous solution, we use hashTable to keep tracking the frequency of each element<br/>
So we can choose the same element again and again until the frequency is 0<br/>

However, in this solution, we don't have hashTable<br/>
And we know one of the requirement is **Each element can only be used once**<br/>
So after we choose an element, we can't choose it again(pass `i + 1` to the next call stack)<br/> 

### Recursive Tree Visualization
candidates = [1,2,2,2], target = 5
<pre>
                                                                                   5
                            1                            2a                                 2b                      2c           
                  2a         2b(***)       2c(****)               
        2b            2c(**)        
</pre>
Look at the (**) part<br/>
Can we choose 2c at here?<br/>
Once choosing, the output is [1,2a,2c]<br/>
It's as same as first output [1,2a,2b]<br/>

`i > index && candidates[i] == candidates[i - 1]` helps us to avoid this case<br/>
Note that in the call stack of ** part, the iteration is from 2 ~ 4<br/>
Why?<br/>
Because after choosing 2a, we pass i + 1 down as our initial index for next call stack<br/>
So we start from 2<br/>
- For 2b, 2 > 2 is false, what does this mean?
  - It means it's the very first element in this call stack, it's okay to choose it no matter what 
- For 2c, 3 > 2 && candidate[2] == candidate[1], what does this means?
  - 2c is not the very first element in this call stack, and it's as same as previous one
  - Don't choose it
  - It's the same logic as (\***\) and (****) part

The best way to deeply understand this trick is writing the recursive tree on the paper


### Complexity Analysis

n = the legnth of candidates array<br/>
t = the target<br/>

#### Time Complexity: O(2 ^ n)
- In the worst case, our algorithm will exhaust all possible combinations from the input array. Again, in the worst case, let us assume that each number is unique. The number of combination for an array of size n would be 2 ^ n, i.e. each number is either included or excluded in a combination.

#### Space Complexity: O(t)