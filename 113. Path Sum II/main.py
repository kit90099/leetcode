# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
from typing import List, Optional


class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> List[List[int]]:
        res = []
        path = []
        def dfs(root: Optional[TreeNode], targetSum: int):
            if not root:
                return

            path.append(root.val)
            targetSum -= root.val
            if not root.left and not root.right and targetSum == 0:
                res.append(path)

            dfs(root.left, targetSum)
            dfs(root.right, targetSum)

            path.pop()

        dfs(root, targetSum)