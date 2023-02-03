import random
from typing import List


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        shuffle(nums)
        lo, hi = 0, len(nums) - 1
        ps = len(nums) - k
        while lo <= hi:
            p = partition(nums, lo, hi)
            if p < ps:
                lo = p + 1
            elif p > ps:
                hi = p - 1
            else:
                return nums[p]
        partition(nums)


def partition(arr, lo, hi):
    pivot = arr[lo]
    i = lo+1
    j = hi
    while i <= j:
        while i <= hi and arr[i] <= pivot:
            i += 1
        while j > lo and arr[j] > pivot:
            j -= 1

        if i >= j:
            break

        swap(arr, i, j)
    swap(arr, j, lo)
    return j

def shuffle(arr):
    for i in range(0,len(arr)):
        rand = random.randint(0, len(arr)-1)
        swap(arr, i, rand)

def swap(arr, a, b):
    arr[a], arr[b] = arr[b], arr[a]

s = Solution()
s.findKthLargest([2,1], 1)