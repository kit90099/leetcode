class WordDictionary {
    public static void main(String... args){
        WordDictionary wd = new WordDictionary();
        wd.addWord("at");
        wd.addWord("and");
        wd.addWord("an");
        wd.addWord("add");
        System.out.println(wd.search("a"));
        System.out.println(wd.search(".at"));
    }
    
    private TrieSet trieSet;
    
    public WordDictionary() {
        this.trieSet = new TrieSet();
    }
    
    public void addWord(String word) {
        trieSet.add(word);
    }
    
    public boolean search(String word) {
        if(trieSet.contains(word) || trieSet.hasKeyWithPattern(word)){
            return true;
        }
        return false;
    }
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * WordDictionary obj = new WordDictionary();
 * obj.addWord(word);
 * boolean param_2 = obj.search(word);
 */