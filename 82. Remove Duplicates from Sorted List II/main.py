# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return None
        vHead = ListNode(0, head)
        curr = vHead
        while curr:
            if curr.next and curr.next.next and curr.next.next.val == curr.next.val:
                val = curr.next.val
                ptr = curr.next
                while ptr and ptr.val == val:
                    ptr = ptr.next
                curr.next = ptr
            else:
                curr = curr.next

        return vHead.next