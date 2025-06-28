class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def isBalanced(self, root):
        """
        :type root: Optional[TreeNode]
        :rtype: bool
        """
        return self.helper(root) != -1

    def helper(self, node):
        if node is None:
            return 0

        left = self.helper(node.left)
        if left == -1:
            return -1

        right = self.helper(node.right)
        if right == -1:
            return -1

        if abs(left - right) > 1:
            return -1

        return max(left, right) + 1
