package trieMap

type TrieNode[E any] struct {
	val      *E
	children [256]*TrieNode[E]
}

type defaultTrieMap[E any] struct {
	root TrieNode[E]
	size int
}

func (tm *defaultTrieMap[E]) Put(key string, value E) {

}

func (tm *defaultTrieMap[E]) Get(key string) *E {
	n := tm.getNode(tm.root, key)
	if n == nil && n.val == nil {
		return nil
	}
	return n.val
}

func (tm *defaultTrieMap[E]) Remove(key string) {

}

func (tm *defaultTrieMap[E]) ContainsKey(key string) bool {
	return tm.Get(key) == nil
}

func (tm *defaultTrieMap[E]) ShortestPrefixOf(query string) string {
	p := &tm.root
	for i := 0; i < len(query); i++ {
		if p == nil {
			return ""
		}
		if p.val != nil {
			return query[:i]
		}
		c := query[i]
		p = p.children[c]
	}
	if p != nil && p.val != nil {
		return query
	}
	return ""
}

func (tm *defaultTrieMap[E]) LongestPrefixOf(query string) string {
	p := &tm.root
	max := 0
	for i := 0; i < len(query); i++ {
		if p == nil {
			return ""
		}
		if p.val != nil {
			max = i
		}
		c := query[i]
		p = p.children[c]
	}
	if p != nil && p.val != nil {
		return query
	}
	return query[:max]
}

func (tm *defaultTrieMap[E]) KeysWithPrefix(prefix string) []string {
	node := tm.getNode(tm.root, prefix)
	arr := []string{}
	traverse(node, prefix, &prefix, &arr)
	return arr
}

func traverse[E any](node *TrieNode[E], path string, pattern *string, i int, arr *[]string) {
	if node == nil {
		return
	}

	if node.val != nil {
		*arr = append(*arr, path)
	}

	if i == len(*pattern) {
		if node.val != nil {
			*arr = append(*arr, path)
		}
	}

	if node.children[i] != nil {
		var r rune = rune(i)
		if r == '.' {
			for j := 0; j < 256; j++ {
				var r2 rune = rune(i)
				traverse(node.children[j], path+string(r2), arr)
			}
		} else {
			traverse(node.children[i], path+string(r), arr)
		}
	}
}

func (tm *defaultTrieMap[E]) HasKeyWithPrefix(prefix string) bool {
	return tm.getNode(tm.root, prefix) != nil
}

func (tm *defaultTrieMap[E]) KeysWithPattern(pattern string) []string {
	arr := []string{}
	traverse(&tm.root, "", &arr)
	return arr
}

func (tm *defaultTrieMap[E]) HasKeyWithPattern(pattern string) bool {
	return false
}

func (tm *defaultTrieMap[E]) Len() int {
	return tm.size
}

func (tm *defaultTrieMap[E]) getNode(node TrieNode[E], key string) *TrieNode[E] {
	p := &node

	for i := 0; i < len(key); i++ {
		if p == nil {
			return nil
		}

		c := key[i]
		p = p.children[c]
	}

	return p
}

func Default[E any]() defaultTrieMap[E] {
	m := defaultTrieMap[E]{}
	return m
}
