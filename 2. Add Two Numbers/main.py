# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
from typing import Optional


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        curr1, curr2 = l1, l2
        addOne = False
        dummy = ListNode()
        last = dummy
        while curr1 and curr2:
            sum = curr1.val + curr2.val
            if addOne:
                sum += 1
            if sum > 9:
                addOne = True
                sum %= 10
            else:
                addOne = False
            last.next = ListNode(sum)
            last = last.next
            curr1 = curr1.next
            curr2 = curr2.next

        if curr1:
            while curr1:
                sum = curr1.val
                if addOne:
                    sum += 1
                if sum > 9:
                    addOne = True
                    sum %= 10
                else:
                    addOne = False
                last.next = ListNode(sum)
                last = last.next
                curr1 = curr1.next
        if curr2:
            while curr2:
                sum = curr2.val
                if addOne:
                    sum += 1
                if sum > 9:
                    addOne = True
                    sum %= 10
                else:
                    addOne = False
                last.next = ListNode(sum)
                last = last.next
                curr2 = curr2.next

        if addOne:
            last.next = ListNode(1)
            last = last.next

        return dummy.next