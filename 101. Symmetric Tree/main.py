# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        def check(a, b):
            if a == None and b == None:
                return True
            if a == None or b == None:
                return False
            return a.val == b.val and check(a.left, b.right) and check(a.right, b.left)
        return check(root, root)