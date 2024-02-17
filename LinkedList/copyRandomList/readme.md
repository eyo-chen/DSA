# Problem Explanation

Let's see one example<br>
head = [[7,null],[13,0],[11,4],[10,2],[1,0]]<br>
<pre>
7 -> 13 -> 11 -> 10 -> 1
|    |     |     |    |
null 0     4     2    0
</pre>
It's easy to clone just the next node<br>
But How can we clone the random node?<br>

See the second node(13)<br>
We know the random node point to 7, but how can we find the cloned node of 7?<br>
That's the hard part<br>

## First Solution: Using Hash Table
If we create a hash table<br>
Key: original node<br>
Value: cloned node<br>

The core logic to wire up random node is<br>
`hashTable[originalNode].random = hashTable[originalNode.random]`<br>

Because `hashTable[originalNode]` is the cloned node of `originalNode`<br>
`hashTable[originalNode].random` -> `clonedNode.random`<br>
And `hashTable[originalNode.random]` is the cloned node of `originalNode.random`<br>

Let's summarize the steps<br>
1. Create a hash table
2. Iterate the original list, and set the key to original node and value to cloned node
3. Iterate the original list again, and wire up the next and random node

Let's see the example, and see how it works<br>
Input: head = [[1,null],[2,0],[3,1]]<br>
<pre>
1 -> 2 -> 3 -> x
| -> -> -> ->  |
| <- |    
     | <- |
</pre>

1. Create a hash table<br>
`hashTable = {}`<br>

2. Iterate the original list, and set the key to original node and value to cloned node<br>
`hashTable = {old1: new1, old2: new2, old3: new3}`<br>

3. Iterate the original list again, and wire up the next and random node<br>
`hashTable[old1].next = hashTable[old1->next]`<br>
`hashTable[old1].random = hashTable[old1->random]`<br>

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(n)
- We need O(n) space to store the hash table


## Second Solution: Using Constant Space
This solution is a bit tricky<br>
We iterate the original list three times<br>

First Pass: make the next of original node to point to cloned node<br>
<pre>
old1    old2    old3    x
  |   -- |    -- |    - |
new1 /  new2 /  new3 /  x
</pre>
It becomes like this<br>
<pre>
old1 -> new1 -> old2 -> new2 -> old3 -> new3 -> x
</pre>

Second Pass: wire up the random node<br>
The logic is<br>
`oldNode->next->random = oldNode->random->next`<br>

Because `oldNode->next` is the cloned node of `oldNode`<br>
`oldNode->next` is `cloneNode`<br>
`oldNode->next->random` is `cloneNode->random`<br>
And `oldNode->random` is the random node of `oldNode`<br>
It's next node is exactly the cloned node of `oldNode->random`<br>

Edge case: <br>
If `oldNode->random` is nullptr, we can't directly access `oldNode->random->next`<br>
We just directly set `oldNode->next->random` to nullptr<br>

Third Pass: correct the next node<br>
Because we re-wire the next node in the first pass<br>
We not only mutate the new node list, but also the original node list<br>
So we need to correct the next reference of both original and new node list<br>

The logic is<br>
`oldNode->next = newNode->next`<br>
Correct the next node of original node<br>
`newNode->next = newNode->next->next`<br>
Correct the next node of new node<br>

Edge case: <br>
Look at `old3`<br>
If `newNode->next` is nullptr, we can't directly access `newNode->next->next`<br>
We just directly set `newNode->next` to nullptr<br>

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(1)