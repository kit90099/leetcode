from uuid import getnode


class Node:
    def __init__(self, val: int, prev=None, next=None):
        self.val = val
        self.prev = prev
        self.next = next


class MyLinkedList:

    def __init__(self):
        self.dummyHead = Node(0)
        self.dummyTail = Node(0)
        self.size = 0

    def getNode(self, index: int) -> Node:
        if index < 0 or index >= self.size:
            return None
        if index < self.size/2:
            curr = self.dummyHead.next
            counter = 0
            while counter < index:
                curr = curr.next
                counter += 1
            return curr
        else:
            curr = self.dummyTail.prev
            counter = self.size - 1
            while counter > index:
                curr = curr.prev
                counter -= 1
            return curr

    def get(self, index: int) -> int:
        node = self.getNode(index)
        if not node:
            return -1
        return node.val

    def addAtHead(self, val: int) -> None:
        if self.size == 0:
            node = Node(val)
            node.next = self.dummyTail
            node.prev = self.dummyHead
            self.dummyHead.next = node
            self.dummyTail.prev = node
        else:
            node = Node(val)
            node.next = self.dummyHead.next
            node.prev = self.dummyHead
            self.dummyHead.next.prev = node
            self.dummyHead.next = node
        self.size += 1

    def addAtTail(self, val: int) -> None:
        if self.size == 0:
            node = Node(val)
            node.next = self.dummyTail
            node.prev = self.dummyHead
            self.dummyHead.next = node
            self.dummyTail.prev = node
        else:
            node = Node(val)
            node.next = self.dummyTail
            node.prev = self.dummyTail.prev
            self.dummyTail.prev.next = node
            self.dummyTail.prev = node
        self.size += 1

    def addAtIndex(self, index: int, val: int) -> None:
        if index == 0:
            self.addAtHead(val)

        prev = self.getNode(index - 1)
        if not prev:
            return
        node = Node(val)
        node.prev = prev
        node.next = prev.next
        prev.next = node
        node.next.prev = node
        self.size += 1

    def deleteAtIndex(self, index: int) -> None:
        if index < 0 or index >= self.size:
            return
        # del at tail
        if index == self.size - 1:
            self.dummyTail.prev = self.dummyTail.prev.prev
            self.dummyTail.prev.next = self.dummyTail
        else:
            prev = None
            if index == 0 and self.size > 0:
                prev = self.dummyHead
            else:
                prev = self.getNode(index - 1)
                if not prev:
                    return
            next = prev.next.next
            next.prev = prev
            prev.next = next

        self.size -= 1

    def printList(self):
        curr = self.dummyHead.next
        print("head:")
        while not curr == self.dummyTail:
            print(str(curr.val) + " ")
            curr = curr.next
        curr = self.dummyTail.prev
        print("tail:")
        while not curr == self.dummyHead:
            print(str(curr.val) + " ")
            curr = curr.prev


l = MyLinkedList()
cList = ["addAtHead","deleteAtIndex","addAtHead","addAtHead","addAtHead","addAtHead","addAtHead","addAtTail","get","deleteAtIndex","deleteAtIndex"]
numList = [[2],[1],[2],[7],[3],[2],[5],[5],[5],[6],[4]]

for i in range(0, len(cList)):
    c = cList[i]
    n = numList[i]

    if c == "get":
        print(l.get(n[0]))
    elif c == "addAtHead":
        l.addAtHead(n[0])
        l.printList()
        print(l.get(n[0]))
    elif c == "addAtTail":
        l.addAtTail(n[0])
        l.printList()
        print(l.get(n[0]))
    elif c == "addAtIndex":
        l.addAtIndex(n[0], n[1])
        l.printList()
        print(l.get(n[0]))
    elif c == "deleteAtIndex":
        l.deleteAtIndex(n[0])
        l.printList()