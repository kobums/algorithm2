function numberOfSpecialChars(word: string): number {
    const lowerSet = new Set<string>();
    const upperSet = new Set<string>();

    // 소문자와 대문자를 각각의 Set에 저장
    for (const char of word) {
        if (char >= 'a' && char <= 'z') {
            lowerSet.add(char);
        } else {
            upperSet.add(char.toLowerCase());
        }
    }

    let count = 0;
    // lowerSet에 존재하면서 upperSet에도 존재하는 문자 카운트
    for (const char of lowerSet) {
        if (upperSet.has(char)) {
            count++;
        }
    }

    return count;
}
