# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def insertionSortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode()
        curr = head

        while curr:
            prev = dummy

            while prev.next and prev.next.val < curr.val:
                prev = prev.next

            next = curr.next
            curr.next = prev.next
            prev.next = curr

            curr = next

        return dummy.next


s = Solution()
s.insertionSortList(ListNode(1, ListNode(
    2, ListNode(4, ListNode(5, ListNode(3, ListNode(6)))))))
