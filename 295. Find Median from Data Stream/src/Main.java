import java.util.PriorityQueue;

public class Main {
    public static void main(String... args){
        MedianFinder obj = new MedianFinder();
        obj.addNum(1);
        obj.addNum(2);
        System.out.println(obj.findMedian());
        obj.addNum(3);
        System.out.println(obj.findMedian());

        /*PriorityQueue<Integer> smaller;
        PriorityQueue<Integer> larger;
        smaller = new PriorityQueue<>();
        larger = new PriorityQueue<>((a, b) ->
                b-a
        );

        smaller.offer(1);
        smaller.offer(2);
        smaller.offer(3);

        larger.offer(1);
        larger.offer(2);
        larger.offer(3);

        while (!smaller.isEmpty()){
            System.out.println(smaller.poll());
        }

        while (!larger.isEmpty()){
            System.out.println(larger.poll());
        }*/
    }
}
