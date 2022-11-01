# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
from typing import List, Optional


class Solution:
    def binaryTreePaths(self, root: Optional[TreeNode]) -> List[str]:
        ans = []
        def traverse(root: Optional[TreeNode], order: str):
            if not root:
                return
            if not root.left and not root.right:
                order += "->" + str(root.val)
                order = order[2:]
                nonlocal ans
                ans.append(order)
            
            traverse(root.left, order + "->" + str(root.val))
            traverse(root.right, order + "->" + str(root.val))
        traverse(root, "")
        return ans


s = Solution()
s.binaryTreePaths(TreeNode(1, TreeNode(2, TreeNode(5)), TreeNode(3)))