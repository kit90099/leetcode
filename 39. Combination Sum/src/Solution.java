import java.util.*;
import java.util.stream.Collectors;

public class Solution {
    public static void main(String[] args) {
        System.out.println(new Solution().combinationSum(new int[]{2, 3, 5}, 8));
    }

    private List<List<Integer>> answer = new ArrayList<>();

    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<Integer> arr = Arrays.stream(candidates).boxed().collect(Collectors.toList());
        backTrack(arr, new LinkedList<Integer>(), 0, target, 0);

        return answer;
    }

    public void backTrack(List<Integer> nums, LinkedList<Integer> track, int start, int target, int sum) {
        if(sum == target){
            answer.add(new LinkedList<>(track));
            return;
        }

        if(sum > target){
            return;
        }

        for (int i = start; i < nums.size(); i++) {
            int num = nums.get(i);

            track.addLast(num);
            sum += num;

            backTrack(nums, track, i, target, sum);

            sum -= num;
            track.removeLast();
        }
    }
}
