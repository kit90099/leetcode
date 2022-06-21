import java.util.List;

public interface TrieMap<E> {
    public void put(String key, E val);
    public void remove(String key);
    public E get(String key);
    public boolean containsKey(String key);
    public String shortestPrefixOf(String query);
    public String longestPrefixOf(String query);
    public List<String> keysWithPrefix(String prefix);
    public boolean hasKeyWithPrefix(String prefix);
    public List<String> keysWithPattern(String pattern);
    public boolean hasKeyWithPattern(String pattern);
    public int size();
}
