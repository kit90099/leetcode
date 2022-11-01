# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
from math import floor
from typing import List, Optional


class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        def generateNode(nums: List[int]) -> Optional[TreeNode]:
            if not nums or len(nums) == 0:
                return None
            mid = floor(len(nums) / 2)
            root = TreeNode(val=nums[mid])
            root.left = generateNode(nums[:mid])
            root.right = generateNode(nums[mid+1:])
            return root
        
        return generateNode(nums)

s = Solution()
s.sortedArrayToBST([3,5,8])