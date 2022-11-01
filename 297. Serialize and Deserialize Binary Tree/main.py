# Definition for a binary tree node.
from typing import List, Optional


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        if not root:
            return "n"

        left = self.serialize(root.left)
        right = self.serialize(root.right)

        return str(root.val) + "," + left + "," + right

        

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        def de(data: List[str], rangeStart) -> Optional[TreeNode]:
            if rangeStart >= len(data):
                return None
            if data[rangeStart] == 'n':
                return None, rangeStart
            
            root = TreeNode(int(data[rangeStart]))
            root.left, end = de(data, rangeStart+1)
            root.right, end = de(data, end + 1)

            return root, end


        splitted = data.split(",")
        res, _ = de(splitted, 0)
        return res
        

# Your Codec object will be instantiated and called as such:
# ser = Codec()
# deser = Codec()
# ans = deser.deserialize(ser.serialize(root))