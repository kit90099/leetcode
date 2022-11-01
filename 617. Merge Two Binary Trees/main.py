# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def mergeTrees(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> Optional[TreeNode]:
        if not root1:
            return root2
        if not root2:
            return root1
        
        q = deque()
        q1 = deque()
        q2 = deque()
        q1.append(root1)
        q2.append(root2)
        merged = TreeNode(root1.val + root2.val)
        q.append(merged)
        
        while len(q1) > 0 or len(q2) > 0:
            node = q.popleft()
            node1 = q1.popleft()
            node2 = q2.popleft()
            
            l1 = node1.left
            r1 = node1.right
            l2 = node2.left
            r2 = node2.right
            
            if l1 or l2:
                if l1 and l2:
                    left = TreeNode(l1.val + l2.val)
                    node.left = left
                    q.append(left)
                    q1.append(l1)
                    q2.append(l2)
                elif l1:
                    node.left = l1
                elif l2:
                    node.left = l2
            
            if r1 or r2:
                if r1 and r2:
                    right = TreeNode(r1.val + r2.val)
                    node.right = right
                    q.append(right)
                    q1.append(r1)
                    q2.append(r2)
                elif r1:
                    node.right = r1
                elif r2:
                    node.right = r2
                    
        return merged