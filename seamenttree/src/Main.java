public class Main {
    public static void main(String... args) {
        /*STree segmentTree = new STree(new int[]{1, 2, 3, 4, 5, 6});
        int res = segmentTree.query(0, 3);*/
        DynamicSegmentTree dst = new DynamicSegmentTree(10);
        dst.update(1, 10);
        dst.update(3, 5);
        System.out.println(dst.query(2, 8));
        System.out.println(dst.query(1, 10));
        System.out.println("");
    }
}
