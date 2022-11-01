# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
from typing import List, Optional


class Solution:
    def buildTree(self, inorder: List[int], postorder: List[int]) -> Optional[TreeNode]:
        if not postorder or len(postorder) == 0:
            return None

        root = TreeNode(postorder[-1])
        stack = []
        stack.append(root)
        inorderIdx = len(inorder) - 1

        for i in range(len(postorder) - 2, -1, -1):
            postOrderVal = postorder[i]
            node = stack[-1]
            if node.val != inorder[inorderIdx]:
                node.right = TreeNode(postOrderVal)
                stack.append(node.right)
            else:
                while len(stack) > 0 and stack[-1].val == inorder[inorderIdx]:
                    node = stack.pop()
                    inorderIdx-=1
                node.left = TreeNode(postOrderVal)
                stack.append(node.left)

        return root