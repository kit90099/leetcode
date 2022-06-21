import java.util.HashMap;
import java.util.Stack;

class Solution {

    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        int[] nums2Greater = nextGreaterElement(nums2);

        HashMap<Integer, Integer> map = new HashMap<>();
        for(int i = 0;i< nums2Greater.length;i++){
            map.put(nums2[i], nums2Greater[i]);
        }

        int[] res = new int[nums1.length];
        for (int i = 0;i< nums1.length;i++){
            res[i] = map.get(nums1[i]);
        }

        return res;
    }

    int[] nextGreaterElement(int[] nums) {
        int n = nums.length;
        int[] res = new int[n];

        Stack<Integer> s = new Stack<>();
        for (int i = n - 1; i >= 0; i--) {
            while(!s.isEmpty() && s.peek() <=nums[i]){
                s.pop();
            }

            res[i] = s.isEmpty() ? -1 : s.peek();
            s.push(nums[i]);
        }

        return res;
    }
}