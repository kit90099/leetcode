import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

class Trie {
    public static void main(String... args) throws FileNotFoundException {
        File commandFile = new File("command");
        Scanner scanner = new Scanner(commandFile);
        String command = scanner.nextLine();
        String word = scanner.nextLine();
        scanner.close();
        String[] commands = command.split(",");
        String[] words = word.split(",");
        for (int i = 0; i < words.length; i++) {
            words[i] = words[i].replace("[", "").replace("]", "");
        }
        File expectedFile = new File("expected");
        scanner = new Scanner(expectedFile);
        String expected = scanner.nextLine();;
        scanner.close();
        String[] expecteds = expected.split(",");
        
        Trie trie = new Trie();
        List<String> res = new ArrayList<>();
        for (int i = 0; i < commands.length; i++) {
            if(commands[i].equals("Trie")){
                res.add("null");
            }
            if(commands[i].equals("insert")){
                trie.insert(words[i]);
                res.add("null");
            }
            if(commands[i].equals("countWordsEqualTo")){
                res.add(String.valueOf(trie.countWordsEqualTo(words[i])));
            }
            if(commands[i].equals("countWordsStartingWith")){
                res.add(String.valueOf(trie.countWordsStartingWith(words[i])));
            }
            if(commands[i].equals("erase")){
                trie.erase(words[i]);
                res.add("null");
            }
        }
        
        for(int i = 0;i<commands.length;i++){
            if(!res.get(i).equals(expecteds[i])){
                System.out.println(i + " " + commands[i] + " " + res.get(i) + " " + expecteds[i]);
            }
        }
        
        System.out.println("");
    }
    
    private CharTrieMap<Integer> map = new CharTrieMap<Integer>(26, (c) -> c.charValue() - 'a', (i) -> (char) (i.intValue() + 'a'));
    
    public Trie() {
    }
    
    public void insert(String word) {
        Integer count = map.get(word);
        if (count != null) {
            map.put(word, count + 1);
        } else {
            map.put(word, 1);
        }
    }
    
    public int countWordsEqualTo(String word) {
        Integer count = map.get(word);
        if(count == null){
            return 0;
        }
        return count;
    }
    
    public int countWordsStartingWith(String prefix) {
        List<String> keys = map.keysWithPrefix(prefix);
        return keys.stream().mapToInt((k) -> map.get(k)).sum();
    }
    
    public void erase(String word) {
        Integer count = map.get(word);
        if (count == null || count == 0) {
            return;
        }
        if (count == 1) {
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