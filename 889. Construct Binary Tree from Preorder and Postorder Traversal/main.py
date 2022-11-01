# Definition for a binary tree node.
from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def constructFromPrePost(self, preorder: List[int], postorder: List[int]) -> Optional[TreeNode]:
        postMap = {}
        def buildTree(preOrder: List[int], preStart: int, preEnd: int, postorder: List[int], postStart: int, postEnd: int) -> Optional[TreeNode]:
            if preStart > preEnd or postStart > postEnd:
                return None
            if preStart == preEnd:
                return TreeNode(preOrder[preStart])
            root = TreeNode(preOrder[preStart])
            leftRootVal = preOrder[preStart + 1]
            nonlocal postMap
            leftRootIdxInPost = postMap[leftRootVal]
            leftSize = leftRootIdxInPost - postStart + 1

            root.left = buildTree(preOrder, preStart+1, preStart + leftSize, postorder, postStart, leftRootIdxInPost)
            root.right = buildTree(preOrder, preStart + leftSize+1, preEnd, postorder, leftRootIdxInPost + 1, postEnd - 1)
            return root

        for (i, x) in enumerate(postorder):
            postMap[x] = i

        return buildTree(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)