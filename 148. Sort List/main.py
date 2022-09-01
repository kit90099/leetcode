# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        def mergeLists(l1: ListNode, l2: ListNode) -> ListNode:
            dummy = ListNode(0)
            ptr, t1, t2 = dummy, l1, l2
            while t1 and t2:
                if t1.val < t2.val:
                    ptr.next = t1
                    t1 = t1.next
                else:
                    ptr.next = t2
                    t2 = t2.next

                ptr = ptr.next

            if t1:
                ptr.next = t1
            elif t2:
                ptr.next = t2

            return dummy.next

        if not head:
            return None

        length = 0
        ptr = head
        while ptr:
            length += 1
            ptr = ptr.next

        dummy = ListNode(0, head)
        subLength = 1
        while subLength < length:
            prev, curr = dummy, dummy.next
            while curr:
                l1 = curr
                for i in range(1, subLength):
                    if curr.next:
                        curr = curr.next
                    else:
                        break
                l2 = curr.next
                curr.next= None
                curr = l2
                for i in range(1, subLength):
                    if curr and curr.next:
                        curr = curr.next
                    else:
                        break
                
                succ = None
                if curr:
                    succ = curr.next
                    curr.next = None

                merged = mergeLists(l1, l2)
                prev.next = merged

                while prev.next:
                    prev = prev.next
                
                curr = succ
            subLength <<= 1
        return dummy.next