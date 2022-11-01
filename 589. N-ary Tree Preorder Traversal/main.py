"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""

class Solution:
    def preorder(self, root: 'Node') -> List[int]:
        res = []
        def traverse(root: Optional[Node]):
            if not root:
                return
            
            nonlocal res
            res.append(root.val)

            if root.children:
                for c in root.children:
                    traverse(c)

        traverse(root)
        return res