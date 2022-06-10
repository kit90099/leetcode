import java.util.HashMap;
import java.util.LinkedHashSet;

class LFUCache {
    private int capacity;
    private HashMap<Integer, Integer> keyToVal = new HashMap<>();
    private HashMap<Integer, Integer> keyToFreq = new HashMap<>();
    private HashMap<Integer, LinkedHashSet<Integer>> freqToKey = new HashMap<>();
    private int minFreq = 0;
    
    public LFUCache(int capacity) {
        this.capacity = capacity;
    }
    
    public int get(int key) {
        if (!keyToVal.containsKey(key)) {
            return -1;
        }
        
        increaseFreq(key);
        return keyToVal.get(key);
    }
    
    public void put(int key, int value) {
        if(keyToVal.containsKey(key)){
            keyToVal.put(key, value);
            increaseFreq(key);
        }
        
        if(capacity == keyToVal.size()){
            removeLeast();
        }
        
        keyToVal.put(key, value);
        increaseFreq(key);
        
    }
    
    private void increaseFreq(int key) {
        int freq = keyToFreq.get(key);
        
        keyToFreq.put(key, freq + 1);
        
        freqToKey.get(freq).remove(key);
        freqToKey.putIfAbsent(freq + 1, new LinkedHashSet<>());
        freqToKey.get(freq + 1).add(key);
        
        if (freqToKey.get(freq).isEmpty()) {
            freqToKey.remove(freq);
            if (freq == minFreq){
                minFreq++;
            }
        }
    }
    
    private void removeLeast(){
        LinkedHashSet<Integer> freqToKeySet = freqToKey.get(minFreq);
        int toBeRemoved = freqToKeySet.iterator().next();
        
        keyToVal.remove(toBeRemoved);
        keyToFreq.remove(toBeRemoved);
        freqToKeySet.remove(toBeRemoved);
        
        if(freqToKeySet.isEmpty()){
            freqToKeySet.remove(minFreq);
        }
    }
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * LFUCache obj = new LFUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */