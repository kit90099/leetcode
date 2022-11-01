# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        def reverse(head: ListNode, stop: ListNode) -> ListNode:
            curr = head
            prev = None
            while not curr == stop:
                next = curr.next
                curr.next = prev
                prev = curr
                curr = next
            
            return prev

        if head == None:
            return None

        start: ListNode = head
        end: ListNode = head
        for i in range(0, k):
            if not end:
                return head
            end = end.next
        
        newHead = reverse(start, end)
        start.next = self.reverseKGroup(end, k)

        return newHead

s = Solution()
s.reverseKGroup(ListNode(1,ListNode(2, ListNode(3, ListNode(4, ListNode(5))))), 2)