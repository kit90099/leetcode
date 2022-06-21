import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

class Solution {
    public static void main(String... args) {
        ArrayList a = new ArrayList() {{
            add("a");
            add("b");
            add("c");
        }};
        System.out.println(new Solution().replaceWords(a, "aadsfasf absbs bbab cadsfafs"));
    }
    
    public String replaceWords(List<String> dictionary, String sentence) {
        TrieSet trieSet = new TrieSet(26, (c)->c.charValue()-'a', (i)->(char)(i.intValue()+'a'));
        for (String word : dictionary) {
            trieSet.add(word);
        }
        
        String[] words = sentence.split(" ");
        List<String> res = new ArrayList<>();
        for (String w : words) {
            String shortestPrefix = trieSet.shortestPrefixOf(w);
            if (shortestPrefix != null && !shortestPrefix.equals("")) {
                res.add(shortestPrefix);
            } else {
                res.add(w);
            }
        }
        
        return res.stream().collect(Collectors.joining(" "));
    }
}