public class DynamicSegmentTree {
    Node root = new Node();
    long n;

    public DynamicSegmentTree(long n){
        this.n = n;
    }

    private void updateHelper(Node curr, long idx, long l, long r, long val) {
        if (l > idx || r < idx) {
            return;
        }

        if (l == r && l == idx) {
            curr.value = val;
            return;
        }

        long mid = l - (l - r) / 2;
        long sum1 = 0, sum2 = 0;

        if (idx <= mid) {
            if (curr.left == null) {
                curr.left = new Node();
            }

            updateHelper(curr.left, idx, l, mid, val);
        } else {
            if (curr.right == null) {
                curr.right = new Node();
            }

            updateHelper(curr.right, idx, mid + 1, r, val);
        }

        if (curr.left != null) {
            sum1 = curr.left.value;
        }
        if (curr.right != null) {
            sum2 = curr.right.value;
        }

        curr.value = sum1 + sum2;
    }

    private long queryHelper(Node curr, long a, long b, long l, long r) {
        if (curr == null) {
            return 0;
        }

        if (l > b || r < a) {
            return 0;
        }

        if (l >= a && r <= b) {
            return curr.value;
        }

        long mid = l - (l - r) / 2;
        return queryHelper(curr.left, a, b, l, mid) + queryHelper(curr.right, a, b, mid + 1, r);
    }

    public void update(long idx, long val){
        updateHelper(root, idx, 1, n, val);
    }

    public long query(long l, long r){
        return queryHelper(root, l, r, 1 ,n);
    }

    public static class Node{
        boolean value = false;
        Node left;
        Node right;
    }
}
