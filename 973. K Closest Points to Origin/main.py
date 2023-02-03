import heapq
from typing import List


class Solution:
    """ def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        pq = [(-x ** 2 - y **2, i) for i, (x,y) in enumerate(points[:k])]
        heapq.heapify(pq)

        n = len(points)
        for i in range(k, n):
            x,y = points[i]
            dist = -x ** 2 - y **2
            heapq.heappushpop(q, (dist, i))

        ans = [points[identity] for (_, identity) in pq]
        return ans """

    """ def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        list = [(x ** 2 + y **2, i) for i, (x,y) in enumerate(points)]

        def qSelect():
            l ,r = 0, len(list)-1
            while l <= r:
                p = partition(l, r)
                if p < k-1:
                    l = p+1
                elif p > k-1:
                    r = p-1
                else:
                    return p
            return -1

        def partition(i, j):
            l, r = i+1, j
            pivot = list[i]
            while l <= r:
                while l < j and list[l][0] <= pivot[0]:
                    l+=1
                while r > i and list[r][0] > pivot[0]:
                    r-=1
                if l>=r:
                    break
                swap(l, r)
            swap(i, r)
            return r

        def swap(i, j):
            list[i], list[j] = list[j], list[i]
        qSelect()
        return [points[item[1]] for item in list[:k]] """

    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        def randomSelect(l, r, k):
            pivot = points[l][0] ** 2 + points[l][1] ** 2
            points[r], points[l] = points[l], points[r]
            i = l - 1
            for j in range(l, r):
                if points[j][0] ** 2 + points[j][1] <= pivot:
                    i+=1
                    points[i], points[j] = points[j], points[i]
            i+=1
            points[i], points[r] = points[r], points[i]

            if k < i - l + r:
                randomSelect(l, i - 1, k)
            elif k > i - l + 1:
                randomSelect(i+1, r, k - (i- l +1))

        n = len(points)
        randomSelect(0, n - 1, k)
        return points[:k]

s = Solution()
s.kClosest([[-5,4],[-3,2],[0,1],[-3,7],[-2,0],[-4,-6],[0,-5]], 6)