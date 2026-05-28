class Solution {
    // Trie 노드 정의
    static class TrieNode {
        TrieNode[] children = new TrieNode[26];
        int bestIndex = -1;
        int bestLength = Integer.MAX_VALUE;
    }

    public int[] stringIndices(String[] wordsContainer, String[] wordsQuery) {
        TrieNode root = new TrieNode();
        
        // 1. 공통 접미사가 전혀 없을 때 사용할 글로벌 기본값 찾기
        int bestGlobalIdx = 0;
        for (int i = 1; i < wordsContainer.length; i++) {
            if (isBetter(wordsContainer[i].length(), i, wordsContainer[bestGlobalIdx].length(), bestGlobalIdx)) {
                bestGlobalIdx = i;
            }
        }

        // 2. wordsContainer의 단어들을 뒤집어서 Trie에 삽입
        for (int i = 0; i < wordsContainer.length; i++) {
            String word = wordsContainer[i];
            int len = word.length();
            TrieNode node = root;

            // 루트 노드 조건 업데이트
            if (isBetter(len, i, node.bestLength, node.bestIndex)) {
                node.bestIndex = i;
                node.bestLength = len;
            }

            // 단어를 뒤에서부터 탐색하여 Trie에 삽입
            for (int j = len - 1; j >= 0; j--) {
                int charIdx = word.charAt(j) - 'a';
                if (node.children[charIdx] == null) {
                    node.children[charIdx] = new TrieNode();
                }
                node = node.children[charIdx];

                // 현재 노드를 공유하는 가장 최적의 단어 정보 갱신
                if (isBetter(len, i, node.bestLength, node.bestIndex)) {
                    node.bestIndex = i;
                    node.bestLength = len;
                }
            }
        }

        // 3. wordsQuery 처리
        int[] result = new int[wordsQuery.length];
        for (int i = 0; i < wordsQuery.length; i++) {
            String query = wordsQuery[i];
            TrieNode node = root;
            int lastValidIndex = bestGlobalIdx;

            // 쿼리 단어를 뒤에서부터 탐색 (괄호 추가 완료)
            for (int j = query.length() - 1; j >= 0; j--) {
                int charIdx = query.charAt(j) - 'a';
                if (node.children[charIdx] == null) {
                    break;
                }
                node = node.children[charIdx];
                lastValidIndex = node.bestIndex;
            }
            result[i] = lastValidIndex;
        }

        return result;
    }

    // 우선순위 비교 헬퍼 메소드 (길이가 짧은 것 우선, 같으면 인덱스가 작은 것 우선)
    private boolean isBetter(int newLen, int newIdx, int oldLen, int oldIdx) {
        if (newLen != oldLen) {
            return newLen < oldLen;
        }
        return newIdx < oldIdx;
    }
}
