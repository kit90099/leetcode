# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if not head:
            return None
        size = 0
        dummy = ListNode(0, head)
        curr = dummy
        while curr.next:
            size += 1
            curr = curr.next
        
        k = k % size
        if k == 0:
            return head

        fast, slow = dummy, dummy
        counter = 0
        while fast.next:
            fast = fast.next
            if counter >= k:
                slow = slow.next
            counter += 1
        
        dummy.next = slow.next
        slow.next = None
        curr.next = head

        return dummy.next

s = Solution()
print(s.rotateRight(ListNode(0, ListNode(1, ListNode(2))), 4))