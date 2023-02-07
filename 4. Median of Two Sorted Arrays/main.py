from typing import List


""" class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        merged = []
        m,n=len(nums1),len(nums2)
        i,j=0,0
        while i < m or j < n:
            if i == m:
                merged.append(nums2[j])
                j+=1
            elif j == n:
                merged.append(nums1[i])
                i+=1
            elif nums1[i] < nums2[j]:
                merged.append(nums1[i])
                i+=1
            else:
                merged.append(nums2[j])
                j+=1

        total = m+n
        if total % 2 == 0:
            return (merged[total//2] + merged[total//2-1]) / 2
        else:
            return merged[total//2] """


class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        m, n = len(nums1), len(nums2)
        if m > n:
            return self.findMedianSortedArrays(nums2, nums1)

        iMin, iMax = 0, m
        while iMin <= iMax:
            i = (iMin + iMax)//2
            j = (m + n + 1) // 2 - i

            if j != 0 and i != m and nums2[j-1] > nums1[i]:
                iMin = i + 1
            elif i != 0 and j != n and nums1[i-1] > nums2[j]:
                iMax = i - 1
            else:
                maxLeft = 0
                if i == 0:
                    maxLeft = nums2[j - 1]
                elif j == 0:
                    maxLeft = nums1[i - 1]
                else:
                    maxLeft = max(nums1[i - 1], nums2[j-1])
                if (m + n) % 2 == 1:
                    return maxLeft

                minRight = 0
                if i == m:
                    minRight = nums2[j]
                elif j == n:
                    minRight = nums1[i]
                else:
                    minRight = min(nums1[i], nums2[j])

                return (maxLeft + minRight) / 2
        return 0


s = Solution()
s.findMedianSortedArrays([1, 2, 3], [4, 5])
