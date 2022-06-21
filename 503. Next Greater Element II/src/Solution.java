import java.util.Stack;

class Solution {
    public static void main(String... args) {
        int[] ans = new Solution().nextGreaterElements(new int[]{1, 2, 3, 4});
        System.out.println(":");
    }

    public int[] nextGreaterElements(int[] nums) {
        int n = nums.length;
        int[] res = new int[n];
        Stack<Integer> s = new Stack<>();

        for (int i = n * 2 - 1; i >= 0; i--) {
            while (!s.isEmpty() && s.peek() <= nums[i % n]) {
                s.pop();
            }

            if (s.isEmpty()) {
                res[i % n] = -1;
            } else {
                res[i % n] = s.peek();
            }

            s.push(nums[i % n]);
        }

        return res;
    }
}