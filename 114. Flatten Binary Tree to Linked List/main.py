# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
from typing import Optional


class Solution:
    def flatten(self, root: Optional[TreeNode]) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        def innerFlat(node: TreeNode):
            if not node:
                return
            
            if node.left:
                innerFlat(node.left)
                temp = node.right
                node.right = node.left
                curr = node.right
                while curr and curr.right:
                    curr = curr.right
                curr.right = temp
                node.left = None

            curr = node
            while curr:
                if curr.left:
                    innerFlat(curr)
                curr = curr.right
                
        innerFlat(root)

s = Solution()
s.flatten(TreeNode(1, TreeNode(2,TreeNode(3),TreeNode(4)), TreeNode(5, None, TreeNode(6))))