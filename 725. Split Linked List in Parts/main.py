# Definition for singly-linked list.
from math import floor
from typing import List, Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def splitListToParts(self, head: Optional[ListNode], k: int) -> List[Optional[ListNode]]:
        if not head:
            return []
        length = 0
        curr = head
        while curr:
            length += 1
            curr = curr.next

        listSize = 0
        listSize = floor(length / k)
        remainder = 0
        if length > listSize:
            remainder = length % k

        res = []
        while head:
            curr = head
            counter = listSize
            if remainder > 0:
                counter += 1
                remainder -= 1
            while counter > 1:
                curr = curr.next
                counter -= 1
            res.append(head)
            next = curr.next
            curr.next = None
            head = next

        while len(res) < k:
            res.append([])

        return res

s = Solution()
s.splitListToParts(ListNode(1,ListNode(2, ListNode(3))), 2)