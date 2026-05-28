package main

type TrieNode struct {
	children [26]*TrieNode
	bestIdx  int
}

func stringProperties(lenA, idxA, lenB, idxB int) bool {
	if lenA != lenB {
		return lenA < lenB
	}
	return idxA < idxB
}

func insert(root *TrieNode, s string, idx int, words []string) {
	node := root
	// 루트 노드 업데이트
	if stringProperties(len(words[idx]), idx, len(words[node.bestIdx]), node.bestIdx) {
		node.bestIdx = idx
	}

	// 문자열을 역순으로 삽입
	for i := len(s) - 1; i >= 0; i-- {
		charIdx := s[i] - 'a'
		if node.children[charIdx] == nil {
			node.children[charIdx] = &TrieNode{bestIdx: idx}
		}
		node = node.children[charIdx]
		
		// 경로상에서 현재 문자열이 더 좋은 조건인지 확인 후 갱신
		if stringProperties(len(words[idx]), idx, len(words[node.bestIdx]), node.bestIdx) {
			node.bestIdx = idx
		}
	}
}

func search(root *TrieNode, s string) int {
	node := root
	for i := len(s) - 1; i >= 0; i-- {
		charIdx := s[i] - 'a'
		if node.children[charIdx] == nil {
			break
		}
		node = node.children[charIdx]
	}
	return node.bestIdx
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	root := &TrieNode{bestIdx: 0}

	// 트라이 구성
	for i, word := range wordsContainer {
		insert(root, word, i, wordsContainer)
	}

	// 쿼리별 결과 처리
	ans := make([]int, len(wordsQuery))
	for i, query := range wordsQuery {
		ans[i] = search(root, query)
	}

	return ans
}
