# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def detectCycle(self, head: ListNode) -> ListNode:
        fast, slow = head, head
        while fast and fast.next:
            fast = fast.next
            slow = slow.next
            if not fast.next:
                return None
            else:
                fast = fast.next

            if fast.val == slow.val:
                ptr = head
                while not ptr.val == slow.val:
                    slow = slow.next
                    ptr = ptr.next
                return ptr
        return None

s = Solution()
l = [-1, -7, 7, -4, 19, 6, -9, -5, -2, -5]
node = None
curr = node
for n in l:
    if curr:
        curr.next = ListNode(n)
        curr = curr.next
    else:
        node = ListNode(n)
        curr = node

curr.next = curr

s.detectCycle(node)