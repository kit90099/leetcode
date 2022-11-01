public class DynamicSegmentTree {
    Node root = new Node();
    int n;

    public DynamicSegmentTree(int n){
        this.n = n;
    }

    private void updateHelper(Node curr, int idx, int l, int r, int val) {
        if (l > idx || r < idx) {
            return;
        }

        if (l == r && l == idx) {
            curr.value = val;
            return;
        }

        int mid = l - (l - r) / 2;
        int sum1 = 0, sum2 = 0;

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

    private int queryHelper(Node curr, int a, int b, int l, int r) {
        if (curr == null) {
            return 0;
        }

        if (l > b || r < a) {
            return 0;
        }

        if (l >= a && r <= b) {
            return curr.value;
        }

        int mid = l - (l - r) / 2;
        return queryHelper(curr.left, a, b, l, mid) + queryHelper(curr.right, a, b, mid + 1, r);
    }

    public void update(int idx, int val){
        updateHelper(root, idx, 1, n, val);
    }

    public int query(int l, int r){
        return queryHelper(root, l, r, 1 ,n);
    }
}
