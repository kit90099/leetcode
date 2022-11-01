# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrderBottom(self, root: Optional[TreeNode]) -> List[List[int]]:
        res = deque()
        if not root:
            return list(res)
        
        queue = deque()
        queue.append(root)
        while len(queue) > 0:
            size = len(queue)
            level = []
            for i in range(0, size):
                node = queue.popleft()
                level.append(node.val)
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            res.appendleft(level)
        
        return list(res)