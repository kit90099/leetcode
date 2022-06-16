package triemap

type TrieNode[E any] struct {
	val      *E
	children [256]*TrieNode[E]
}

type defaultTrieMap[E any] struct {
	root TrieNode[E]
	size int
}

func (tm *defaultTrieMap[E]) Put(key string, value E) {
	if !tm.ContainsKey(key) {
		tm.size++
	}

	tm.root = *_put(&tm.root, key, &value, 0)
}

func _put[E any](node *TrieNode[E], key string, val *E, i int) *TrieNode[E] {
	if node == nil {
		node = &TrieNode[E]{}
	}
	if i == len(key) {
		node.val = val
		return node
	}

	r := rune(key[i])
	node.children[r] = _put(node.children[r], key, val, i+1)
	return node
}

func (tm *defaultTrieMap[E]) Get(key string) *E {
	n := tm.getNode(tm.root, key)
	if n == nil || n.val == nil {
		return nil
	}
	return n.val
}

func (tm *defaultTrieMap[E]) Remove(key string) {
	if !tm.ContainsKey(key) {
		return
	}

	tm.root = *_remove(&tm.root, key, 0)
	tm.size--
}

func _remove[E any](node *TrieNode[E], key string, i int) *TrieNode[E] {
	if node == nil {
		return nil
	}

	if i == len(key) {
		node.val = nil
	} else {
		r := rune(key[i])
		node.children[r] = _remove(node.children[r], key, i+1)
	}

	if node.val != nil {
		return node
	}

	for j := 0; i < 256; j++ {
		if node.children[j].val != nil {
			return nil
		}
	}

	return nil
}

func (tm *defaultTrieMap[E]) ContainsKey(key string) bool {
	return tm.Get(key) != nil
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
	traverse4Prefix(node, prefix, &arr)
	return arr
}

func (tm *defaultTrieMap[E]) HasKeyWithPrefix(prefix string) bool {
	return tm.getNode(tm.root, prefix) != nil
}

func traverse4Prefix[E any](node *TrieNode[E], path string, arr *[]string) {
	if node == nil {
		return
	}

	if node.val != nil {
		*arr = append(*arr, path)
	}

	for i := 0; i < 256; i++ {
		if node.children[i] != nil {
			var r rune = rune(i)
			traverse4Prefix(node.children[i], path+string(r), arr)
		}
	}
}

func (tm *defaultTrieMap[E]) KeysWithPattern(pattern string) []string {
	arr := []string{}
	traverse4Pattern(&tm.root, "", pattern, 0, &arr)
	return arr
}

func (tm *defaultTrieMap[E]) HasKeyWithPattern(pattern string) bool {
	return _hasKeyWithPattern(&tm.root, pattern, 0)
}

func _hasKeyWithPattern[E any](node *TrieNode[E], pattern string, i int) bool {
	if node == nil {
		return false
	}

	if i == len(pattern) {
		return node.val != nil
	}

	r := rune(pattern[i])
	if r == '.' {
		for j := 0; j < 256; j++ {
			if node.children[j] != nil {
				if _hasKeyWithPattern(node.children[j], pattern, i+1) {
					return true
				}
			}
		}
		return false
	} else {
		return _hasKeyWithPattern(node.children[r], pattern, i+1)
	}
}

func traverse4Pattern[E any](node *TrieNode[E], path string, pattern string, i int, arr *[]string) {
	if node == nil {
		return
	}

	if node.val != nil {
		*arr = append(*arr, path)
	}

	if i == len(pattern) {
		if node.val != nil {
			*arr = append(*arr, path)
		}
	}

	if node.children[i] != nil {
		var r rune = rune(i)
		if r == '.' {
			for j := 0; j < 256; j++ {
				var r2 rune = rune(i)
				traverse4Pattern(node.children[j], path+string(r2), pattern, i+1, arr)
			}
		} else {
			traverse4Pattern(node.children[i], path+string(r), pattern, i+1, arr)
		}
	}
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
