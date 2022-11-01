import java.util.Arrays;
import java.util.PriorityQueue;

class Solution {
    public int maxEvents(int[][] events) {
        Arrays.sort(events, (a, b) -> a[0] == b[0] ? a[1] - b[1] : a[0] - b[0]);
        PriorityQueue<int[]> pq = new PriorityQueue<int[]>((a, b) -> a[1] - b[1]);
        int a = 0, day = 1;
        int p = 0;
        while (p < events.length || !pq.isEmpty()) {
            while (p < events.length && (events[p][0] <= day || pq.isEmpty())) {
                pq.add(new int[]{events[p][0], events[p][1]});
                p++;
            }
            while (!pq.isEmpty() && pq.peek()[1] < day) {
                pq.poll();
            }
            if (!pq.isEmpty()) {
                a++;
                int[] d = pq.poll();
                day = Math.max(day, d[0]);
                day++;
            }
        }
        return a;
    }

    public static void main(String... args){
        Solution s = new Solution();
        s.maxEvents(new int[][]{new int[]{1,2},new int[]{2,3},new int[]{3,4}});
    }
}