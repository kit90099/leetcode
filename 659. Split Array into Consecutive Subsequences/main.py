from typing import List


class Solution:
    def isPossible(self, nums: List[int]) -> bool:
        freq, need = {}, {}
        for x in nums:
            freq.setdefault(x, 0)
            freq[x] += 1
            need[x] = 0

        for x in nums:
            if freq[x] == 0:
                continue

            if x in need and need[x] > 0:
                freq[x] -= 1
                need[x] -= 1
                need[x+1] = need.get(x+1, 0) + 1
            elif freq[x] > 0 and freq.get(x+1,0) > 0 and freq.get(x+2,0) > 0:
                freq[x] -= 1
                freq[x + 1] -= 1
                freq[x + 2] -= 1
                need[x + 3] = need.get(x+3, 0) + 1
            else:
                return False

        return True


s = Solution()
print(s.isPossible([1, 2, 3, 4, 4, 5]))
