# Definition for a binary tree node.
from collections import deque
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def widthOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        """ q = deque()
        q.append([root, 1])
        res = 1

        while len(q) > 0:
            tmp = deque()
            for node, idx in q:
                if node.left:
                    tmp.append([node.left, idx*2])
                if node.right:
                    tmp.append([node.right, idx*2+1])
            res = max(res, tmp[-1][1] - tmp[0][1] + 1)
            q = tmp
        return res """

        res = 1
        arr = [[root, 1]]
        while arr:
            tmp = []
            for node, index in arr:
                if node.left:
                    tmp.append([node.left, index * 2])
                if node.right:
                    tmp.append([node.right, index * 2 + 1])
            res = max(res, arr[-1][1] - arr[0][1] + 1)
            arr = tmp
        return res

s =Solution()
print(s.widthOfBinaryTree(TreeNode(1, TreeNode(3, TreeNode(5), TreeNode(3)), TreeNode(2, None, TreeNode(9)))))