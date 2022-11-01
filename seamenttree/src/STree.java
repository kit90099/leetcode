import java.util.ArrayList;

public class STree {
    private int[] tree;
    private int n;

    public STree(int[] arr) {
        n = arr.length;
        tree = new int[2 * n];

        for (int i = 0; i < n; i++) {
            tree[n + i] = arr[i];
        }
        for (int i = n - 1; i > 0; i--) {
            tree[i] = tree[2 * i] + tree[2 * i + 1];
        }
    }

    public void update(int idx, int val) {
        tree[n + idx] = val;

        for (int i = idx + n; i > 1; i >>= 1) {
            tree[i >> 1] = tree[i] + tree[i ^ 1];
        }
    }

    public int query(int left, int right) {
        int res = 0;

        for (int l = left + n, r = right + n; l < r; l >>= 1, r >>= 1) {
            if ((l & 1) > 0) {
                res += tree[l++];
            }
            if ((r & 1) > 0) {
                res += tree[--r];
            }
        }

        return res;
    }
}
