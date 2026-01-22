var Trie = function() {
    this.root = {};
};

Trie.prototype.insert = function(word) {
    let node = this.root;
    
    for (const char of word) {
        if (!node[char]) {
            node[char] = {};
        }
        node = node[char];
    }
    
    node.isEnd = true;
};

Trie.prototype.search = function(word) {
    const node = this._traverse(word);
    return node !== null && node.isEnd === true;
};

Trie.prototype.startsWith = function(prefix) {
    return this._traverse(prefix) !== null;
};

// 헬퍼: prefix까지 탐색
Trie.prototype._traverse = function(str) {
    let node = this.root;
    
    for (const char of str) {
        if (!node[char]) {
            return null;
        }
        node = node[char];
    }
    
    return node;
};