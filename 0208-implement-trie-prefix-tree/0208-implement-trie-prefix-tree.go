type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}

type Trie struct {
    root *TrieNode
}

func Constructor() Trie {
    return Trie{
        root: &TrieNode{
            children: make(map[rune]*TrieNode),
        },
    }
}

func (t *Trie) Insert(word string) {
    node := t.root
    
    for _, char := range word {
        if _, exists := node.children[char]; !exists {
            node.children[char] = &TrieNode{
                children: make(map[rune]*TrieNode),
            }
        }
        node = node.children[char]
    }
    
    node.isEnd = true
}

func (t *Trie) Search(word string) bool {
    node := t.traverse(word)
    return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
    return t.traverse(prefix) != nil
}

func (t *Trie) traverse(str string) *TrieNode {
    node := t.root
    
    for _, char := range str {
        if _, exists := node.children[char]; !exists {
            return nil
        }
        node = node.children[char]
    }
    
    return node
}