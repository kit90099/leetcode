package trieMap

type TrieMap[E any] interface {
	Put(key string, value E)
	Get(key string) E
	Remove(key string)
	ContainsKey(key string) bool
	ShortestPrefixOf(query string) string
	LongestPrefixOf(query string) string
	KeysWithPrefix(prefix string) []string
	HasKeyWithPrefix(prefix string) bool
	KeysWithPattern(pattern string) []string
	HasKeyWithPattern(pattern string) bool
	Len() int
}
