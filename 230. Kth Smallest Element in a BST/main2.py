# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        m = {}

        def traverse(node : Optional[TreeNode]):
            if node == None:
                return 0
            
            l = traverse(node.left)
            r = traverse(node.right)
            total = l+r

            m[node] = total
            return total + 1

        traverse(root)

        def find(node, pos):
            l = m.get(node.left,-1)

            mid = l+2
            if l == -1:
                mid = 1

            if mid == pos:
                return node.val
            
            if pos < mid:
                return find(node.left, pos)
            else:
                return find(node.right, pos - mid)

        return find(root, k)

s = Solution()
print(s.kthSmallest(TreeNode(1, None, TreeNode(2)),2))
# print(s.kthSmallest(TreeNode(3, TreeNode(1, None,TreeNode(2)), TreeNode(4)),1))
##print(s.kthSmallest(TreeNode(5, TreeNode(3, TreeNode(2,TreeNode(1)),TreeNode(4)), TreeNode(6)),3))