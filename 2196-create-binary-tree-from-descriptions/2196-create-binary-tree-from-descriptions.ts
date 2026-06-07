function createBinaryTree(descriptions: number[][]): TreeNode | null {
    const nodes = new Map<number, TreeNode>();
    const hasParent = new Set<number>();

    // 1. 노드 생성 및 부모-자식 연결 구성
    for (const [pVal, cVal, isLeft] of descriptions) {
        if (!nodes.has(pVal)) nodes.set(pVal, new TreeNode(pVal));
        if (!nodes.has(cVal)) nodes.set(cVal, new TreeNode(cVal));

        const parent = nodes.get(pVal)!;
        const child = nodes.get(cVal)!;

        if (isLeft === 1) {
            parent.left = child;
        } else {
            parent.right = child;
        }

        hasParent.add(cVal);
    }

    // 2. 부모가 없는 단 하나의 루트(Root) 노드 찾기
    for (const [pVal] of descriptions) {
        if (!hasParent.has(pVal)) {
            return nodes.get(pVal)!;
        }
    }

    return null;
}
