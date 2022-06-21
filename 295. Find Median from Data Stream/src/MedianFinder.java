import java.util.PriorityQueue;

class MedianFinder {
    private PriorityQueue<Integer> smaller;
    private PriorityQueue<Integer> larger;

    public MedianFinder() {
        smaller = new PriorityQueue<>();
        larger = new PriorityQueue<>((a, b) ->
            b-a
        );
    }

    public void addNum(int num) {
        if(smaller.size() >= larger.size()){
            smaller.offer(num);
            larger.offer(smaller.poll());
        }else{
            larger.offer(num);
            smaller.offer(larger.poll());
        }
    }

    public double findMedian() {
        if (larger.size() < smaller.size()) {
            return smaller.peek();
        } else if (smaller.size() < larger.size()) {
            return larger.peek();
        } else {
            return (smaller.peek().doubleValue() + larger.peek().doubleValue()) / 2;
        }
    }
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * MedianFinder obj = new MedianFinder();
 * obj.addNum(num);
 * double param_2 = obj.findMedian();
 */