import heapq

class MedianFinder:

    def __init__(self):
        self.accending = []
        self.decending = []

    def addNum(self, num: int) -> None:
        sizeA, sizeD = len(self.accending), len(self.decending)
        if sizeD >= sizeA:
            temp = heapq.heappushpop(self.decending, -num)
            heapq.heappush(self.accending, -temp)
        else:
            temp = heapq.heappushpop(self.accending, num)
            heapq.heappush(self.decending, -temp)

    def findMedian(self) -> float:
        sizeA, sizeD = len(self.accending), len(self.decending)
        if sizeA < sizeD:
            return -heapq.nsmallest(1, self.decending)[0]
        elif sizeD < sizeA:
            return heapq.nsmallest(1, self.accending)[0]
        else:
            return (-heapq.nsmallest(1, self.decending)[0] + heapq.nsmallest(1, self.accending)[0]) / 2

m = MedianFinder()
m.addNum(-1)
print(m.findMedian())
m.addNum(-2)
print(m.findMedian())
m.addNum(-3)
print(m.findMedian())
m.addNum(-4)
print(m.findMedian())
m.addNum(-5)
print(m.findMedian())




# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()
