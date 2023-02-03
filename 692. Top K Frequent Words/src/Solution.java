import java.util.*;

class Solution {
    public List<String> topKFrequent(String[] words, int k) {
        Map<String, Integer> countMap = new HashMap<>();
        for (String w: words){
            countMap.putIfAbsent(w, 0);
            countMap.put(w, countMap.get(w) + 1);
        }

        PriorityQueue<Item> pq = new PriorityQueue<Item>((a, b) -> {
            if (a.count == b.count){
                return a.word.compareTo(b.word);
            }else{
                return b.count - a.count;
            }
        });

        countMap.entrySet().stream().forEach(e -> pq.add(new Item(e.getKey(), e.getValue())));
        List<String> res = new ArrayList<>();
        for (int i = 0; i< k;i++){
            res.add(pq.poll().word);
        }

        return res;
    }

    public static class Item {
        String word;
        int count;

        public Item(String word, int count) {
            this.word = word;
            this.count = count;
        }
    }

    public static void main(String... args){
        Solution s = new Solution();
        s.topKFrequent(new String[]{"i","love","leetcode","i","love","coding"}, 4);
    }
}