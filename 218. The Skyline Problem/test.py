import heapq


class Test(list):
    def __lt__(self, other):
        return self[1] > other[1]
    
pq = []
heapq.heappush(pq, Test([1,2]))
heapq.heappush(pq, Test([2,1]))
heapq.heappush(pq, Test([3,3]))
heapq.heappush(pq, Test([4,4]))

while len(pq) > 0:
    print(heapq.nsmallest(0, pq))