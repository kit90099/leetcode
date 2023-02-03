from typing import List


class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src: int, dst: int, k: int) -> int:
        indegree = {}
        for f in flights:
            ffrom, to, price = f[0], f[1], f[2]
            if not to in indegree:
                indegree[to] = []
            indegree[to].append([ffrom, price])

        memo = []
        for i in range(0, n):
            tmp = []
            for j in range(0, k+2):
                tmp.append(-999)
            memo.append(tmp)

        def dp(s, x) -> int:
            if s == src:
                return 0
            if x == 0:
                return -1

            if memo[s][x] != -999:
                return memo[s][x]

            res = 99999999
            if s in indegree:
                for e in indegree[s]:
                    ffrom, price = e[0], e[1]
                    sub = dp(ffrom, x-1)
                    if sub != -1:
                        res = min(res, sub + price)

            if res == 99999999:
                memo[s][x] = -1
            else:
                memo[s][x] = res
            return memo[s][x]

        return dp(dst, k+1)


s = Solution()
s.findCheapestPrice(4, [[0, 1, 100], [1, 2, 100], [2, 0, 100], [
                    1, 3, 600], [2, 3, 200]], 0, 3, 1)
