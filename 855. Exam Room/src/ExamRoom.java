import java.util.Comparator;
import java.util.HashMap;
import java.util.Map;
import java.util.TreeSet;

class ExamRoom {
    int n;
    TreeSet<int[]> pq;
    private Map<Integer, int[]> startMap;
    private Map<Integer, int[]> endMap;

    public ExamRoom(int n) {
        this.n = n;
        pq = new TreeSet<int[]>((a, b) -> {
            int distanceA = distance(a);
            int distanceB = distance(b);

            if (distanceA == distanceB) {
                return b[0] - a[0];
            }

            return distanceA - distanceB;
        });
        startMap = new HashMap<>();
        endMap = new HashMap<>();

        addInterval(new int[]{-1, n});
    }

    public int seat() {
        int[] longest = pq.last();
        int seat;

        if (longest[0] == -1) {
            seat = 0;
        } else if (longest[1] == n) {
            seat = n - 1;
        } else {
            seat = (longest[0] + longest[1]) / 2;
        }

        int[] left = new int[]{longest[0] , seat};
        int[] right = new int[]{seat , longest[1]};

        removeInterval(longest);
        addInterval(left);
        addInterval(right);

        return seat;
    }

    public void leave(int p) {
        int[] right = startMap.get(p);
        int[] left = endMap.get(p);

        int[] newInterval = new int[]{left[0], right[1]};
        removeInterval(left);
        removeInterval(right);
        addInterval(newInterval);
    }

    private int distance(int[] interval) {
        int x = interval[0];
        int y = interval[1];

        if (x == -1){
            return y-1;
        }
        if (y == n){
            return n-x-2;
        }
        return (y-x)/2;
    }

    private void addInterval(int[] interval) {
        pq.add(interval);
        startMap.put(interval[0], interval);
        endMap.put(interval[1], interval);
    }

    private void removeInterval(int[] interval) {
        pq.remove(interval);
        startMap.remove(interval[0]);
        endMap.remove(interval[1]);
    }
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * ExamRoom obj = new ExamRoom(n);
 * int param_1 = obj.seat();
 * obj.leave(p);
 */