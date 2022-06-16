import java.util.List;

public class TrieSet {
    private final TrieMap<Object> map = new TrieMap<>();
    
    public void add(String key) {
        map.put(key, new Object());
    }
    
    public void remove(String key) {
        map.remove(key);
    }
    
    public boolean contains(String key) {
        return map.containsKey(key);
    }
    
    public String shortestPrefixOf(String query) {
        return map.shortestPrefixOf(query);
    }
    
    public String longestPrefixOf(String query) {
        return map.longestPrefixOf(query);
    }
    
    public List<String> keysWithPrefix(String prefix) {
        return map.keysWithPrefix(prefix);
    }
    
    public boolean hasKeyWithPrefix(String prefix) {
        return map.hasKeyWithPrefix(prefix);
    }
    
    public List<String> keysWithPattern(String pattern) {
        return map.keysWithPattern(pattern);
    }
    
    public boolean hasKeyWithPattern(String pattern) {
        return map.hasKeyWithPattern(pattern);
    }
    
    public int size() {
        return map.size();
    }
}
