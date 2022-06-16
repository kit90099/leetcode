package triemap

type TrieSet struct {
	trieMap *TrieMap[int]
}

func (ts *TrieSet) Put(key string) {
	(*ts.trieMap).Put(key, 0)
}
func (ts *TrieSet) Remove(key string) {
	(*ts.trieMap).Remove(key)
}
func (ts *TrieSet) ContainsKey(key string) bool {
	return (*ts.trieMap).ContainsKey(key)
}
func (ts *TrieSet) ShortestPrefixOf(query string) string {
	return (*ts.trieMap).ShortestPrefixOf(query)
}
func (ts *TrieSet) LongestPrefixOf(query string) string {
	return (*ts.trieMap).LongestPrefixOf(query)
}
func (ts *TrieSet) KeysWithPrefix(prefix string) []string {
	return (*ts.trieMap).KeysWithPrefix(prefix)
}
func (ts *TrieSet) HasKeyWithPrefix(prefix string) bool {
	return (*ts.trieMap).HasKeyWithPrefix(prefix)
}
func (ts *TrieSet) KeysWithPattern(pattern string) []string {
	return (*ts.trieMap).KeysWithPattern(pattern)
}
func (ts *TrieSet) HasKeyWithPattern(pattern string) bool {
	return (*ts.trieMap).HasKeyWithPattern(pattern)
}
func (ts *TrieSet) Len() int {
	return (*ts.trieMap).Len()
}

func GetTrieSet() *TrieSet {
	x := Default[int]()
	var t TrieMap[int] = &x

	return &TrieSet{
		trieMap: &t,
	}
}
