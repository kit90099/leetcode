import java.util.List;

class MapSum {
    private CharTrieMap<Integer> map = new CharTrieMap<Integer>(26, (c)->c-'a', (i)->(char)(i.intValue()+'a'));
    
    public MapSum() {
        
    }
    
    public void insert(String key, int val) {
        map.put(key, val);
    }
    
    public int sum(String prefix) {
        List<String> keys = map.keysWithPrefix(prefix);
        return keys.stream().mapToInt((k)->map.get(k)).sum();
    }
}

/**
 * Your MapSum object will be instantiated and called as such:
 * MapSum obj = new MapSum();
 * obj.insert(key,val);
 * int param_2 = obj.sum(prefix);
 */