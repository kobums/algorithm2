type WordDictionary struct {
    sons [26]*WordDictionary
    end int
}


func Constructor() WordDictionary {
    return WordDictionary{}
}


func (this *WordDictionary) AddWord(word string)  {
    node := this
    size := len(word)
    for i := 0; i < size; i++ {
        idx := word[i] - 'a'
        if node.sons[idx] == nil {
            node.sons[idx] = &WordDictionary{}
        }

        node = node.sons[idx]
    }

    node.end++
}


func (this *WordDictionary) Search(word string) bool {
    for i, b := range word {
		if b != '.' {
			idx := b - 'a'
			if this.sons[idx] == nil {
				return false
			}

			this = this.sons[idx]
		} else {
			for _, son := range this.sons {
				if son == nil {
					continue
				}

				this = son
				if i == len(word)-1 {
					if this.end > 0 {
						return true
					}
					continue
				}

				if this.Search(word[i+1:]) {
					return true
				}
			}

			return false
		}
	}

	if this.end > 0 {
		return true
	}

	return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */