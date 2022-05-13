package com.test;

import javafx.util.Pair;
import org.springframework.stereotype.Service;

import java.util.Arrays;
import java.util.PriorityQueue;

@Service
class Solution {
    public int[] advantageCount(int[] nums1, int[] nums2) {
        Arrays.sort(nums1);
        
        PriorityQueue<Pair<Integer, Integer>> pq = new PriorityQueue<>((p1, p2) -> p1.getKey() > p2.getKey() ? -1 : 1);
        for (int i = 0; i < nums2.length; i++) {
            pq.add(new Pair<>(nums2[i], i));
        }
        
        int l = 0, r = nums1.length - 1;
        int[] res = new int[nums1.length];
        while (l <= r) {
            Pair<Integer, Integer> num2 = pq.poll();
            if (nums1[r] > num2.getKey()) {
                res[num2.getValue()] = nums1[r];
                r--;
            } else {
                res[num2.getValue()] = nums1[l];
                l++;
            }
        }
        
        return res;
    }
}