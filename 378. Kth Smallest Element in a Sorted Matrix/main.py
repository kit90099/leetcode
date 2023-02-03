from ast import List
import heapq


class item:
    def __init__(self, i) -> None:
        self.i = i

    def __lt__(self, o):
        return self.i[0] < o.i[0]


class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        n = len(matrix)
        pq = []

        for i in range(n):
            heapq.heappush(pq, item([matrix[i][0], i, 0]))

        res = 0
        while pq and k > 0:
            temp = heapq.heappop(pq)
            res = temp[1]
            k -= 1

            i = temp[1]
            j = temp[2]
            if j+1 < n:
                heapq.heappush(pq, item([matrix[i][j+1], i, j+1]))
        return res
