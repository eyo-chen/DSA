class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def insertIntoBST(self, root, val):
        """
        :type root: Optional[TreeNode]
        :type val: int
        :rtype: Optional[TreeNode]
        """

        if root == None:
            return TreeNode(val=val)

        if root.val > val:
            root.left = self.insertIntoBST(root.left, val)
        else:
            root.right = self.insertIntoBST(root.right, val)

        return root


class Solution(object):
    def insertIntoBST(self, root, val):
        """
        :type root: Optional[TreeNode]
        :type val: int
        :rtype: Optional[TreeNode]
        """

        node = TreeNode(val=val)
        if root == None:
            return node

        cur_node = root

        while cur_node != None:
            if cur_node.val > val:
                if cur_node.left == None:
                    cur_node.left = node
                    break
                cur_node = cur_node.left
                continue

            if cur_node.right == None:
                cur_node.right = node
                break
            cur_node = cur_node.right

        return root
