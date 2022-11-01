import java.util.*;

class RangeModule {
    private TreeMap<Integer, Integer> map;

    public RangeModule() {
        map = new TreeMap<Integer, Integer>((a, b) -> {
            return a - b;
        });
    }

    public void addRange(int left, int right) {
        Map.Entry<Integer, Integer> entry = map.higherEntry(left);
        if (entry != map.firstEntry()) {
            Map.Entry<Integer, Integer> start = entry != null ? map.lowerEntry(entry.getKey()) : map.lastEntry();
            // l <= r <= left <= right
            if (start != null && start.getValue() >= right) {
                return;
            }
            if (start != null && start.getValue() >= left) {
                left = start.getKey();
                map.remove(start.getKey());
            }
        }
        while (entry != null && entry.getKey() <= right) {
            right = Math.max(right, entry.getValue());
            map.remove(entry.getKey());
            entry = map.higherEntry(entry.getKey());
        }
        map.put(left, right);
    }

    public boolean queryRange(int left, int right) {
        Map.Entry<Integer, Integer> entry = map.higherEntry(left);
        if (entry == map.firstEntry()) {
            return false;
        }
        entry = entry != null ? map.lowerEntry(entry.getKey()) : map.lastEntry();
        return entry != null && right <= entry.getValue();
    }

    public void removeRange(int left, int right) {
        Map.Entry<Integer, Integer> entry = map.higherEntry(left);
        if (entry != map.firstEntry()) {
            Map.Entry<Integer, Integer> start = entry != null ? map.lowerEntry(entry.getKey()) : map.lastEntry();
            // l <= left <= right <= r
            if (start != null && start.getValue() >= right){
                int ri = start.getValue();
                if (left == start.getKey()){
                    map.remove(left);
                }else{
                    map.put(start.getKey(), left);
                }
                if (right != ri){
                    map.put(right, ri);
                }
                return;
            }else if (start != null && start.getValue() > left){
                // l <= left <= r <= right
                if (start.getKey() == left){
                    map.remove(start.getKey());
                }else{
                    map.put(start.getKey(), left);
                }
            }
        }
        // second interval
        while (entry!=null && entry.getKey() < right){
            if (entry.getValue() <= right){
                map.remove(entry.getKey());
                entry = map.higherEntry(entry.getKey());
            }else {
                map.put(right, entry.getValue());
                map.remove(entry.getKey());
                break;
            }
        }
    }
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * RangeModule obj = new RangeModule();
 * obj.addRange(left,right);
 * boolean param_2 = obj.queryRange(left,right);
 * obj.removeRange(left,right);
 */