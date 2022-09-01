# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        i = 1
        ptr = ListNode(0, head)
        while i < left:
            ptr = ptr.next
            i+=1

        reversePtr = ptr.next
        last = ptr
        nextPtr = ptr
        while i <= right:
            nextPtr = reversePtr.next
            reversePtr.next = last
            last = reversePtr
            reversePtr = nextPtr
            i+=1
        
        ptr.next.next = nextPtr
        ptr.next = last

        return head

s = Solution()
l = ListNode(1,ListNode(2, ListNode(3, ListNode(4, ListNode(5)))))
s.reverseBetween(l,2,4)
print("")