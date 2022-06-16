import java.util.ArrayList;
import java.util.List;

public class TrieMap<E> {
    private int size = 0;
    private TrieNode<E> root = null;
    private int childrenSize = 256;
    
    public TrieMap(){}
    public TrieMap(int childrenSize) {
        this.childrenSize = childrenSize;
    }
    
    public void put(String key, E value) {
        if (!containsKey(key)) {
            size++;
        }
        
        put(root, key, value, 0);
    }
    
    private TrieNode<E> put(TrieNode<E> node, String key, E value, int i) {
        if (node == null) {
            node = new TrieNode<>();
        }
        
        if (i == key.length()) {
            node.value = value;
            return node;
        }
        
        char c = key.charAt(i);
        node.children[c] = put(node.children[c], key, value, i + 1);
        return node;
    }
    
    public void remove(String key) {
    
    }
    
    private TrieNode<E> remove(TrieNode<E> node, String key, int i) {
        if (node == null) {
            return null;
        }
        
        if (i == key.length()) {
            node.value = null;
        } else {
            char c = key.charAt(i);
            node.children[c] = remove(node, key, i + 1);
        }
        
        for (int j = 0; j < node.children.length; j++) {
            if (node.children[j] != null) {
                return node;
            }
        }
        return null;
    }
    
    public E get(String key) {
        TrieNode<E> node = getNode(root, key);
        if (node == null || node.value == null) {
            return null;
        }
        return node.value;
    }
    
    public boolean containsKey(String key) {
        return get(key) == null;
    }
    
    public String shortestPrefixOf(String key) {
        TrieNode<E> p = root;
        for (int i = 0; i < key.length(); i++) {
            if (p == null) {
                return "";
            }
            
            if (p.value != null) {
                return key.substring(0, i);
            }
            
            char c = key.charAt(i);
            p = p.children[c];
        }
        
        if (p != null && p.value != null) {
            return key;
        }
        return "";
    }
    
    public String longestPrefixOf(String key) {
        TrieNode<E> p = root;
        int length = 0;
        for (int i = 0; i < key.length(); i++) {
            if (p == null) {
                return "";
            }
            
            if (p.value != null) {
                length = i;
            }
            
            char c = key.charAt(i);
            p = p.children[c];
        }
        
        if (p != null && p.value != null) {
            return key;
        }
        
        return key.substring(0, length);
    }
    
    public List<String> keysWithPrefix(String prefix) {
        TrieNode node = getNode(root, prefix);
        List<String> res = new ArrayList<>();
        if (node == null) {
            return res;
        }
        traverse(node, res, prefix);
        return res;
    }
    
    private void traverse(TrieNode<E> node, List<String> res, String path) {
        if (node == null) {
            return;
        }
        
        if (node.value != null) {
            res.add(path);
        }
        
        for (char c = 0; c < node.children.length; c++) {
            traverse(node.children[c], res, path + c);
        }
    }
    
    public boolean hasKeyWithPrefix(String prefix) {
        return getNode(root, prefix) != null;
    }
    
    public List<String> keysWithPattern(String pattern) {
        List<String> res = new ArrayList<>();
        traverse(root, res, pattern, "", 0);
        return res;
    }
    
    public void traverse(TrieNode node, List<String> res, String pattern, String path, int i) {
        if (node == null) {
            return;
        }
        
        if (i == pattern.length()) {
            if (node.value != null) {
                res.add(path);
            }
            return;
        }
        
        char c = path.charAt(i);
        if (c == '.') {
            for (char j = 0; j < node.children.length; j++) {
                traverse(node.children[j], res, pattern, path + j, i + 1);
            }
        } else {
            traverse(node, res, pattern, path + c, i + 1);
        }
    }
    
    public boolean hasKeyWithPattern(String pattern) {
        return hasKeyWithPattern(root, pattern, 0);
    }
    
    private boolean hasKeyWithPattern(TrieNode node, String pattern, int i) {
        if (node == null) {
            return false;
        }
        
        if (i == pattern.length()) {
            return node != null;
        }
        
        char c = pattern.charAt(i);
        if (c != '.') {
            return hasKeyWithPattern(node.children[c], pattern, i + 1);
        }
        for (char j = 0; j < node.children.length; j++) {
            if (hasKeyWithPattern(node.children[j], pattern, i + 1)) {
                return true;
            }
        }
        return false;
    }
    
    public int size() {
        return size;
    }
    
    private TrieNode<E> getNode(TrieNode<E> node, String key) {
        TrieNode<E> p = node;
        for (int i = 0; i < key.length(); i++) {
            if (p == null) {
                return null;
            }
            
            char c = key.charAt(i);
            p = p.children[c];
        }
        
        return p;
    }
    
    public static class TrieNode<E> {
        E value;
        TrieNode<E>[] children;
    }
}
