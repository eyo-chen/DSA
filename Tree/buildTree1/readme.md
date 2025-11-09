# Problem Explanation

The solution follows a similar approach to the "Construct Binary Tree from Preorder and Inorder Traversal" problem, with a few key differences:

**Key Difference #1: Root Location**
- In the preorder version, the root is at the **first** index of the preorder array
- In this postorder version, the root is at the **last** index of the postorder array

**Key Difference #2: Construction Order**
- We must construct the **right subtree before the left subtree**
- This is because postorder traversal visits nodes in the order: left → right → root
- Working backwards from the end, we encounter: root → right subtree → left subtree

**Shared Logic**
- Use the inorder array to determine left and right subtree boundaries
- Create a hash map for O(1) lookup of root positions in the inorder array
- Recursively build subtrees using index boundaries instead of array slicing

**Algorithm Steps:**
1. Find the root value at the last position of postorder
2. Locate the root in inorder to split into left and right subtrees
3. Build the right subtree first (elements just before root in postorder)
4. Build the left subtree (remaining elements at the beginning of postorder)
5. Return the constructed root node