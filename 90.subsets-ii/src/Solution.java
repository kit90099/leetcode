import java.util.*;
import java.util.stream.Collectors;

public class Solution {
    public static void main(String[] args) {
        System.out.println(new Solution().subsetsWithDup(new int[]{1, 2, 2}));
    }

    private List<List<Integer>> answer = new ArrayList<>();

    public List<List<Integer>> subsetsWithDup(int[] nums) {
        List<Integer> sorted = Arrays.stream(nums).sorted().boxed().collect(Collectors.toList());
        backTrack(sorted, new LinkedList<Integer>(), 0);

        return answer;
    }

    public void backTrack(List<Integer> nums, LinkedList<Integer> track, int start) {
        answer.add(new LinkedList<>(track));

        for (int i = start; i < nums.size(); i++) {
            int num = nums.get(i);
            if (i > start && num == nums.get(i - 1)) {
                continue;
            }
            track.addLast(num);

            backTrack(nums, track, i + 1);

            track.removeLast();
        }
    }
}
