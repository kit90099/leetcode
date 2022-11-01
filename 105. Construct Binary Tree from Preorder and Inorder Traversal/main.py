# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
from collections import deque
from typing import List, Optional


class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if not preorder:
            return None
        root = TreeNode(preorder[0])
        stack = [root]
        inPtr = 0

        for i in range(1, len(preorder)):
            preOrderVal = preorder[i]
            node = stack[-1]
            if node.val != inorder[inPtr]:
                node.left = TreeNode(preorder[i])
                stack.append(node.left)
            else:
                while stack and stack[-1].val == inorder[inPtr]:
                    node = stack.pop()
                    inPtr += 1
                node.right = TreeNode(preOrderVal)
                stack.append(node.right)
        
        return root

s = Solution()
s.buildTree([1,2,4,5,3,6,7], [4,2,5,1,6,3,7])