from typing import List


class Solution:
    def intervalIntersection(self, firstList: List[List[int]], secondList: List[List[int]]) -> List[List[int]]:
        joined = []
        for i in firstList:
            joined.append(i)
        for i in secondList:
            joined.append(i)

        joined.sort(key=lambda x: (x[0], -x[1]))
        res = []
        start, end = joined[0][0], joined[0][1]
        for i in range(1, len(joined)):
            interval = joined[i]
            if interval[0] <= end and interval[1] > end:
                res.append([interval[0], end])
                start, end = interval[0], interval[1]
            elif interval[0] <= end and interval[1] <= end:
                res.append([interval[0], interval[1]])
            else:
                start, end = interval[0], interval[1]

        return res


s = Solution()
print(s.intervalIntersection([[0, 2], [5, 10], [13, 23], [24, 25]],
                             [[1, 5], [8, 12], [15, 24], [25, 26]]))
