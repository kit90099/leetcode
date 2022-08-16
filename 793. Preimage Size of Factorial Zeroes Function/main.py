import sys


class Solution:
    def preimageSizeFZF(self, k: int) -> int:
        return right(k) - left(k) + 1


def left(target: int) -> int:
    left: int = 0
    right: int = pow(10, 9)
    while left < right:
        mid: int = int(left + (right - left)/2)
        z = zeros(mid)
        if z < target:
            left = mid+1
        elif z > target:
            right = mid
        else:
            right = mid

    return left


def right(target: int) -> int:
    left: int = 0
    right: int = pow(10, 9)
    while left < right:
        mid: int = int(left + (right - left)/2)
        z = zeros(mid)
        if z < target:
            left = mid+1
        elif z > target:
            right = mid
        else:
            left = mid+1
    return left-1


def zeros(n: int) -> int:
    divisor = 5
    res: int = 0
    while divisor <= n:
        res += n / divisor
        divisor *= 5
    return int(res)
