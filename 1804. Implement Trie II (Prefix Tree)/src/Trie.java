import java.util.List;

class Trie {
    public static void main(String... args){
        Trie trie = new Trie();
        trie.insert("apple");               // Inserts "apple".
        trie.insert("apple");               // Inserts another "apple".
        System.out.println(trie.countWordsEqualTo("apple"));    // There are two instances of "apple" so return 2.
        System.out.println(trie.countWordsStartingWith("app")); // "app" is a prefix of "apple" so return 2.
        trie.erase("apple");                // Erases one "apple".
        System.out.println(trie.countWordsEqualTo("apple"));    // Now there is only one instance of "apple" so return 1.
        System.out.println(trie.countWordsStartingWith("app")); // return 1
        trie.erase("apple");                // Erases "apple". Now the trie is empty.
        System.out.println(trie.countWordsStartingWith("app")); // return 0
    }
    
    private CharTrieMap<Integer> map = new CharTrieMap<Integer>(26, (c) -> c.charValue() - 'a', (i) -> (char) (i.intValue() + 'a'));
    
    public Trie() {}
    
    public void insert(String word) {
        Integer count = map.get(word);
        if (count != null && count != 0) {
            map.put(word, count + 1);
        } else {
            map.put(word, 1);
        }
    }
    
    public int countWordsEqualTo(String word) {
        Integer count = map.get(word);
        return count == null ? 0 : count;
    }
    
    public int countWordsStartingWith(String prefix) {
        List<String> keys = map.keysWithPrefix(prefix);
        return keys.stream().mapToInt((k)-> map.get(k)).sum();
    }
    
    public void erase(String word) {
        Integer count = map.get(word);
        if(count == null || count == 0){
            return;
        }
        if(count == 1){
            map.remove(word);
            return;
        }
        map.put(word, count - 1);
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * int param_2 = obj.countWordsEqualTo(word);
 * int param_3 = obj.countWordsStartingWith(prefix);
 * obj.erase(word);
 */