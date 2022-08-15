from typing import List

class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key=lambda x: (x[0], -x[1]))

        start ,end = intervals[0][0], intervals[0][1]
        res = []
        for i in range(1, len(intervals)):
            interval = intervals[i]
            if interval[0] > end:
                res.append([start,end])
                start = interval[0]
                end = interval[1]
            elif interval[0] <= end and interval[1] > end:
                end = interval[1]
        res.append([start,end])
        return res



s = Solution()
print(s.merge([[1,4],[4,5]]))