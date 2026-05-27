function numberOfSpecialChars(word: string): number {
    const firstUpper = new Map<string, number>();
    const lastLower = new Map<string, number>();

    // 각 문자의 첫 번째 대문자 인덱스와 마지막 소문자 인덱스 기록
    for (let i = 0; i < word.length; i++) {
        const char = word[i];
        if (char >= 'a' && char <= 'z') {
            lastLower.set(char, i);
        } else if (char >= 'A' && char <= 'Z') {
            const lowerChar = char.toLowerCase();
            if (!firstUpper.has(lowerChar)) {
                firstUpper.set(lowerChar, i);
            }
        }
    }

    let count = 0;

    // 대문자와 소문자가 모두 존재하고, 마지막 소문자 인덱스가 첫 대문자 인덱스보다 작은 경우 카운트
    for (const [char, firstUpperIndex] of firstUpper) {
        if (lastLower.has(char)) {
            if (lastLower.get(char)! < firstUpperIndex) {
                count++;
            }
        }
    }

    return count;
}
