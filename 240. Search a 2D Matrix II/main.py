from typing import List


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        x, y = len(matrix) - 1, 0
        while x >=0 and y < len(matrix[0]):
            if matrix[x][y] == target:
                return True
            if matrix[x][y] > target:
                x -= 1
            else:
                y += 1
        return False

s = Solution()
s.searchMatrix([[1,1]],0)