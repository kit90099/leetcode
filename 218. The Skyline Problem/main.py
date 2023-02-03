from heapq import heappush
import heapq
from typing import List

class PqItem:
    def __init__(self, r, h) -> None:
        self.r = r
        self.h = h
    def __lt__(self , other):
        return self.h > other.h

class Solution:
    def getSkyline(self, buildings: List[List[int]]) -> List[List[int]]:
        pq = []
        boundries = []
        for b in buildings:
            boundries.append(b[0])
            boundries.append(b[1])
        boundries.sort()

        res = []
        n = len(buildings)
        idx = 0
        for b in boundries:
            while idx < n and buildings[idx][0] <= b:
                heapq.heappush(pq, PqItem(buildings[idx][1], buildings[idx][2]))
                idx+=1
            while len(pq) > 0 and pq[0].r <= b:
                heapq.heappop(pq)
            max = 0
            if len(pq) != 0:
                max = pq[0].h
            if len(res) == 0 or res[-1][1] != max:
                res.append([b, max])
        return res
            


class PQItem(list):
    def __lt__(self, other):
        return self[1] > other[1]

s = Solution()
s.getSkyline([[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]])