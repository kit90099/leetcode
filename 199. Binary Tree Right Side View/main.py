    # Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
from typing import List, Optional


class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        res = []
        dic = {}
        def dfs(root: Optional[TreeNode], level: int):
            if not root:
                return None
            if level not in dic:
                res.append(root.val)
                dic[level] = True
            dfs(root.right, level+1)
            dfs(root.left, level+1)

        dfs(root, 1)
        return res