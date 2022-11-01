from heapq import heappush
import heapq
from typing import List


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
                heapq.heappush(pq, PQItem([buildings[idx][1],buildings[idx][2]]))
                idx+=1
            while len(pq) > 0 and heapq.nsmallest(1, pq)[0][0] <= b:
                heapq.heappop(pq)
            max = 0
            if len(pq) != 0:
                max = heapq.nsmallest(1, pq)[0][1]

            if len(res) != 0 and res[-1][1] == max:
                continue
            res.append([b, max])
        return res
            


class PQItem(list):
    def __lt__(self, other):
        return self[1] > other[1]

s = Solution()
s.getSkyline([[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]])