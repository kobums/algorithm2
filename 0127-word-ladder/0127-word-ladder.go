func ladderLength(beginWord string, endWord string, wordList []string) int {
    wordSet := make(map[string]bool)
    for _, word := range wordList {
        wordSet[word] = true
    }

    // endWord가 wordList에 없으면 변환 불가
    if !wordSet[endWord] {
        return 0
    }

    // BFS에 사용할 큐
    queue := []string{beginWord}
    depth := 1

    for len(queue) > 0 {
        nextQueue := []string{}

        for _, word := range queue {
            if word == endWord {
                return depth
            }

            // 한 글자씩 변환
            for i := 0; i < len(word); i++ {
                for c := 'a'; c <= 'z'; c++ {
                    newWord := word[:i] + string(c) + word[i+1:]

                    // wordSet에 있으면 큐에 추가
                    if wordSet[newWord] {
                        nextQueue = append(nextQueue, newWord)
                        delete(wordSet, newWord)  // 중복 방지
                    }
                }
            }
        }

        queue = nextQueue
        depth++
    }

    return 0  // 변환할 수 없는 경우
}