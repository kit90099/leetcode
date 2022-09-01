# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        ## find mid
        fast = slow = head
        while not fast == None and not fast.next == None:
            fast = fast.next.next
            slow = slow.next

        l1 = head
        l2 = slow.next
        slow.next = None

        ## reverse
        ptr = l2
        prev = None
        while ptr:
            next = ptr.next
            ptr.next = prev
            prev = ptr
            ptr = next

        ## merge
        l2 = prev
        while l1 and l2:
            t1 = l1.next
            t2 = l2.next

            l1.next = l2
            l1 = t1

            l2.next = l1
            l2 = t2

        return 


            


s = Solution()
s.reorderList(ListNode(1))