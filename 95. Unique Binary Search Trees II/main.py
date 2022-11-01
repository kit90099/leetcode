# Definition for a binary tree node.
from distutils.command.build import build
from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def generateTrees(self, n: int) -> List[Optional[TreeNode]]:

        def gen(lo: int, hi: int) -> List[TreeNode]:
            res = []

            if lo > hi:
                res.append(None)
                return res

            for i in range(lo, hi + 1):
                l = gen(lo, i-1)
                r = gen(i+1, hi)

                for left in l:
                    for right in r:
                        root = TreeNode(i)
                        root.left = left
                        root.right = right
                        res.append(root)

            return res
        
        if n == 0:
            return []
        return build(1, n)