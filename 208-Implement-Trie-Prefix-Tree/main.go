package main

type Trie struct {
	size int
	root *TrieNode
}

type TrieNode struct {
	Val      *int
	children []*TrieNode
}

func getNode(node *TrieNode, key string) *TrieNode {
	p := node
	for i := 0; i < len(key); i++ {
		if p == nil {
			return nil
		}

		c := key[i]
		p = p.children[c]
	}
	return p
}

func (this *Trie) get(key string) *int {
	p := getNode(this.root, key)
	if p == nil || p.Val == nil {
		return nil
	}
	return p.Val
}

func (this *Trie) contains(key string) bool {
	p := getNode(this.root, key)
	if p == nil || p.Val == nil {
		return false
	}
	return true
}

func (this *Trie) hasKeysWithPrefix(key string) bool {
	p := getNode(this.root, key)
	return p != nil
}

func (this *Trie) shortestPrefixOf(key string) string {
	p := this.root
	for i := 0; i < len(key); i++ {
		if p == nil {
			return ""
		}

		if p.Val != nil {
			return key[:i]
		}
		c := key[i]
		p = p.children[c]
	}
	if p != nil && p.Val != nil {
		return key
	}
	return ""
}

func (this *Trie) longestPrefixOf(key string) string {
	p := this.root
	max := 0
	for i := 0; i < len(key); i++ {
		if p == nil {
			return ""
		}

		if p.Val != nil {
			max = i
		}
		c := key[i]
		p = p.children[c]
	}
	if p != nil && p.Val != nil {
		return key
	}
	return key[:max]
}

func (this *Trie) keysWithPrefix(prefix string) []string {
	res := []string{}
	x := getNode(this.root, prefix)

	var traverse func(*TrieNode, string)
	traverse = func(node *TrieNode, path string) {
		if node == nil {
			return
		}

		if node.Val != nil {
			res = append(res, path)
		}

		var c rune
		for c = 0; c < 256; c++ {
			s := path + string(c)
			traverse(node.children[c], s)
		}
	}

	traverse(x, "")
	return res
}

func (this *Trie) keysWithPattern(pattern string) []string {
	res := []string{}

	var traverse func(node *TrieNode, path string, level int)
	traverse = func(node *TrieNode, path string, level int) {
		if node == nil {
			return
		}

		if level == len(pattern) {
			if node.Val != nil {
				res = append(res, path)
			}
			return
		}

		c := pattern[level]
		if c == '.' {
			var i rune
			for i = 0; i < 256; i++ {
				traverse(node.children[i], path+string(i), level+1)
			}
		} else {
			traverse(node.children[c], path+string(c), level+1)
		}
	}

	traverse(this.root, "", 0)
	return res
}

func (this *Trie) hasKeyWithPattern(pattern string) bool {
	var traverse func(node *TrieNode, i int) bool
	traverse = func(node *TrieNode, i int) bool {
		if node == nil {
			return false
		}
		if i == len(pattern) {
			return node.Val == nil
		}
		c := pattern[i]
		if c != '.' {
			return traverse(node.children[i], i+1)
		}
		for j := 0; j < 256; j++ {
			if traverse(node.children[j], i+1) {
				return true
			}
		}
		return false
	}
	return traverse(this.root, 0)
}

func (this *Trie) put(key string, val int) {
	if !this.contains(key) {
		this.size++
	}

	var traverse func(node *TrieNode, i int) *TrieNode
	traverse = func(node *TrieNode, i int) *TrieNode {
		if node == nil {
			node = &TrieNode{
				children: make([]*TrieNode, 256),
			}
		}
		if i == len(key) {
			node.Val = &val
			return node
		}
		c := key[i]
		node.children[c] = traverse(node.children[c], i+1)
		return node
	}
	this.root = traverse(this.root, 0)
}

func (this *Trie) remove(key string) {
	if !this.contains(key) {
		return
	}

	var traverse func(node *TrieNode, i int) *TrieNode
	traverse = func(node *TrieNode, i int) *TrieNode {
		if node == nil {
			return nil
		}
		if i == len(key) {
			node.Val = nil
		} else {
			c := key[i]
			node.children[c] = traverse(node.children[c], i+1)
		}

		if node.Val != nil {
			return node
		}
		for j := 0; j < 256; j++ {
			if node.children[j].Val != nil {
				return node
			}
		}
		return nil
	}
}
