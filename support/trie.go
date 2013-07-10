package trie

type trieChildren map[string]*Trie

type Trie struct {
	Leaf     bool
	Entry    interface{}
	Children trieChildren
}

func (t *Trie) Get(path []string) (entry interface{}, ok bool) {
	...
}

func (t *Trie) Set(path []string, value interface{}) {
	...
}
