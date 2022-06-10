import java.util.LinkedHashMap;

class LRUCache {
    private int capacity;
    private LinkedHashMap<Integer, Integer> cache = new LinkedHashMap();
    
    
    public LRUCache(int capacity) {
        this.capacity = capacity;
    }
    
    public int get(int key) {
        if (!cache.containsKey(key)) {
            return -1;
        }
        
        makeRecent(key);
        
        return cache.get(key);
    }
    
    public void put(int key, int value) {
        if (cache.containsKey(key)){
            cache.put(key, value);
            makeRecent(key);
            return;
        }
        
        if(cache.size() == this.capacity){
            int oldestKey = cache.entrySet().iterator().next().getKey();
            cache.remove(oldestKey);
        }
        
        cache.put(key, value);
    }
    
    private void makeRecent(int key){
        int value = cache.get(key);
        cache.remove(key);
        cache.put(key, value);
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */
