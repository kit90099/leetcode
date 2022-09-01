# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        if not head:
            return None
        dummy = ListNode(0)
        original = ListNode(0)
        oPtr = original
        dPtr = dummy
        curr = head
        last = None
        while curr:
            if curr.val >= x:
                dPtr.next = curr
                curr = curr.next
                dPtr = dPtr.next
                dPtr.next = None
            else:
                oPtr.next = curr
                curr = curr.next
                oPtr = oPtr.next
                oPtr.next = None

        oPtr.next = dummy.next
        return original.next