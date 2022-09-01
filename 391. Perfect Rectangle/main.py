from typing import List

class Solution:
    def isRectangleCover(self, rectangles: List[List[int]]) -> bool:
        count = set()
        minX, maxX, minY, maxY = float('inf'), -float('inf'), float('inf'), -float('inf')
        area = 0
        for rec in rectangles:
            p1 = (rec[0],rec[1])
            p2 = (rec[2],rec[3])
            p3 = (rec[0],rec[3])
            p4 = (rec[2],rec[1])

            for p in [p1,p2,p3,p4]:
                if p in count:
                    count.remove(p)
                else:
                    count.add(p)

            minX = min(minX, min(rec[0], rec[2]))
            maxX = max(maxX, max(rec[0], rec[2]))
            minY = min(minY, min(rec[1], rec[3]))
            maxY = max(maxY, max(rec[1], rec[3]))

            area += (rec[2]-rec[0])*(rec[3]-rec[1])

        if not area == ((maxX-minX)*(maxY-minY)):
            return False

        if len(count) != 4:
            return False

        if not (minX, minY) in count:
            return False
        if not (maxX, maxY) in count:
            return False
        if not (maxX, minY) in count:
            return False
        if not (minX, maxY) in count:
            return False

        return True


s = Solution()
print(s.isRectangleCover([[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]))
