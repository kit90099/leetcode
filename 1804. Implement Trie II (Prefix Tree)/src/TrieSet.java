import java.util.List;
import java.util.function.Function;

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */

class TrieSet {
    private final TrieMap<Object> map;
    public TrieSet(){
        map = new CharTrieMap<>();
    }
    public TrieSet(int size, Function<Character, Integer> charToIntConvertor, Function<Integer, Character> intToCharConvertor){
        map = new CharTrieMap<Object>(size, charToIntConvertor, intToCharConvertor);
    }
    
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

