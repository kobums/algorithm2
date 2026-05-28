class TrieNode {
    children: Map<string, TrieNode>;
    bestIndex: number;
    bestLength: number;

    constructor() {
        this.children = new Map();
        this.bestIndex = -1;
        this.bestLength = Infinity;
    }
}

// 두 인덱스의 우선순위를 비교하는 헬퍼 함수 (길이가 짧은 것 우선, 같으면 인덱스가 작은 것 우선)
function isBetter(
    newLen: number, newIdx: number,
    oldLen: number, oldIdx: number
): boolean {
    if (newLen !== oldLen) {
        return newLen < oldLen;
    }
    return newIdx < oldIdx;
}

function stringIndices(wordsContainer: string[], wordsQuery: string[]): number[] {
    const root = new TrieNode();

    // 1. 전체 container에서 기본 fallback으로 사용할 가장 최적의 인덱스 찾기
    let bestGlobalIdx = 0;
    for (let i = 1; i < wordsContainer.length; i++) {
        if (isBetter(wordsContainer[i].length, i, wordsContainer[bestGlobalIdx].length, bestGlobalIdx)) {
            bestGlobalIdx = i;
        }
    }

    // 2. 단어들을 뒤집어서 Trie에 삽입
    for (let i = 0; i < wordsContainer.length; i++) {
        const word = wordsContainer[i];
        const len = word.length;
        let node = root;

        // 루트 노드 업데이트 (공통 접미사가 전혀 없는 경우의 대비책)
        if (isBetter(len, i, node.bestLength, node.bestIndex)) {
            node.bestIndex = i;
            node.bestLength = len;
        }

        // 뒤에서부터 문자 삽입 (접미사 매칭을 위해)
        for (let j = word.length - 1; j >= 0; j--) {
            const char = word[j];
            if (!node.children.has(char)) {
                node.children.set(char, new TrieNode());
            }
            node = node.children.get(char)!;

            // 해당 노드를 거쳐가는 단어 중 가장 최적의 값 저장
            if (isBetter(len, i, node.bestLength, node.bestIndex)) {
                node.bestIndex = i;
                node.bestLength = len;
            }
        }
    }

    // 3. 쿼리 처리
    const result: number[] = [];
    for (const query of wordsQuery) {
        let node = root;
        let lastValidIndex = bestGlobalIdx;

        // 쿼리 문자열을 뒤에서부터 탐색
        for (let i = query.length - 1; i >= 0; i--) {
            const char = query[i];
            if (!node.children.has(char)) {
                break;
            }
            node = node.children.get(char)!;
            lastValidIndex = node.bestIndex;
        }
        result.push(lastValidIndex);
    }

    return result;
}
