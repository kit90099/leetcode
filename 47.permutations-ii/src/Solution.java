import java.util.*;
import java.util.stream.Collectors;

public class Solution {
    public static void main(String[] args) {
        System.out.println(new Solution().permuteUnique(new int[]{1, 2, 3}));
    }

    private List<List<Integer>> answer = new ArrayList<>();

    public List<List<Integer>> permuteUnique(int[] nums) {
        List<Integer> sorted = Arrays.stream(nums).sorted().boxed().collect(Collectors.toList());
        backTrack(sorted, new LinkedList<>());

        return answer;
    }

    public void backTrack(List<Integer> nums, LinkedList<Integer> track) {
        if(0 == nums.size()){
            answer.add(new LinkedList<>(track));
        }

        for (int i = 0; i < nums.size(); i++) {
            int num = nums.get(i);
            if (i > 0 && num == nums.get(i-1)) {
                continue;
            }
            track.addLast(num);
            nums.remove(i);

            backTrack(nums, track);

            track.removeLast();
            nums.add(i, num);
        }
    }
}
