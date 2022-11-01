# Definition for a binary tree node.
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        pos = 0
        res = 0
        def traverse(root: Optional[TreeNode], k: int):
            if not root:
                return
            
            traverse(root.left, k)

            nonlocal pos
            pos += 1
            if k == pos:
                nonlocal res
                res = root.val
                return

            traverse(root.right, k)
        traverse(root, k)
        return res
