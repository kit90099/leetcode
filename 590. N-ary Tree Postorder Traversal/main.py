"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""

class Solution:
    def postorder(self, root: 'Node') -> List[int]:
        res = []
        def traverse(node: Optional[Node]):
            if not node:
                return
            if node.children:
                for c in node.children:
                    traverse(c)
            
            nonlocal res
            res.append(node.val)

        traverse(root)
        return res