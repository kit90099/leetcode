import java.util.*;

public class Solution {
    // https://leetcode.cn/problems/rectangle-area-ii/solution/by-capital-worker-7efv/
    private static final int IN = 1;
    private static final int OUT = -1;
    private static final int MOD = (int) 1e9 + 7;

    public int rectangleArea(int[][] rectangles) {
        TreeSet<Integer> ySet = new TreeSet<>();
        Map<Integer, Integer> y2Index = new HashMap<>();
        Map<Integer, Integer> index2y = new HashMap<>();

        List<int[]> xList = new ArrayList<>();
        for (int[] rect : rectangles) {
            xList.add(new int[]{rect[0], rect[1], rect[3], IN});
            xList.add(new int[]{rect[2], rect[1], rect[3], OUT});
            ySet.add(rect[1]);
            ySet.add(rect[3]);
        }

        // discretization
        int count = 1;
        for (int y : ySet) {
            y2Index.put(y, count);
            index2y.put(count, y);
            count++;
        }

        xList.sort(Comparator.comparingInt(x -> x[0]));

        SegmentTree segmentTree = new SegmentTree();
        long ans = 0;
        // how many to scan
        int n = xList.size();
        for (int i = 0; i < n - 1; i++) {
            int[] curr = xList.get(i);
            // left, right is range of segment tree
            // convert to index
            int left = y2Index.get(curr[1]);
            int right = y2Index.get(curr[2]);
            segmentTree.update(left, right - 1, curr[3], index2y);
            // xList.get(i + 1)[0] - curr[0]: distance for this vertical line(x) to next vertical line(x) i.e. x
            // query(): get y2 - y1 on vertical line i.e. y
            // area of rectangle = x * y =  area of rectangle between curr and next line
            ans += (long) segmentTree.query() * (xList.get(i + 1)[0] - curr[0]);
        }

        return (int) (ans % MOD);
    }

    public static class SegmentTree {
        private static final int maxNode = 201;

        public SegmentTree() {
            root = new TreeNode();
        }

        private TreeNode root;

        private class TreeNode {
            public TreeNode left;
            public TreeNode right;
            public int coverLen;
            public int cover;
        }


        public void update(int left, int right, int value, Map<Integer, Integer> index2y) {
            update(root, 1, maxNode, left, right, value, index2y);
        }

        public int query() {
            return root.coverLen;
        }

        private void update(TreeNode root, int start, int end, int left, int right, int value, Map<Integer, Integer> index2y) {
            createNode(root);
            if (left <= start && end <= right) {
                root.cover += value;
                pushUp(root, start, end, index2y);
                return;
            }

            int mid = (start) + (end - start) /2;
            if (left <= mid) {
                update(root.left, start, mid, left, right, value, index2y);
            }
            if (mid < right) {
                update(root.right, mid + 1, end, left, right, value, index2y);
            }
            pushUp(root, start, end, index2y);
        }

        // start and end is index of y1 and y2, end - start should be the difference between y2 and y1
        private void pushUp(TreeNode root, int start, int end, Map<Integer, Integer> index2y) {
            if (root.cover > 0) {
                // fully covered eg 1 - 2 is fully covered by range is 1 - 10
                root.coverLen = index2y.get(end + 1) - index2y.get(start);
            } else if (start != end) {
                // not fully covered eg 1-5 when range is 4-6
                root.coverLen = root.left.coverLen + root.right.coverLen;
            } else {
                // not covered when start = end, they are same and not a range
                root.coverLen = 0;
            }
        }

        private void createNode(TreeNode root) {
            if (root.left == null) {
                root.left = new TreeNode();
            }
            if (root.right == null) {
                root.right = new TreeNode();
            }
        }
    }
}
