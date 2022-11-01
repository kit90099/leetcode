# Definition for singly-linked list.
from collections import deque
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dq1, dq2 = deque(), deque()
        curr = l1
        while curr:
            dq1.append(curr.val)
            curr = curr.next    
        curr = l2
        while curr:
            dq2.append(curr.val)
            curr = curr.next

        carry = 0
        res = None
        while len(dq1) > 0 or len(dq2) > 0:
            v1 = dq1.pop() if len(dq1) > 0 else 0
            v2 = dq2.pop() if len(dq2) > 0 else 0

            sum = v1 + v2 + carry
            carry = 1 if sum > 9 else 0
            node = ListNode(sum%10)
            node.next = res
            res = node

        if carry == 1:
            node = ListNode(1)
            node.next = res
            res = node

        return res