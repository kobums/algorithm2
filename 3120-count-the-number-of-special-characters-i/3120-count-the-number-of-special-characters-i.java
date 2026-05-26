class Solution {
    public int numberOfSpecialChars(String word) {
        int lowercaseMask = 0;
        int uppercaseMask = 0;
        
        // 1. 문자열을 순회하며 소문자와 대문자의 등장을 비트마스크에 기록
        for (int i = 0; i < word.length(); i++) {
            char ch = word.charAt(i);
            if (ch >= 'a' && ch <= 'z') {
                lowercaseMask |= (1 << (ch - 'a'));
            } else if (ch >= 'A' && ch <= 'Z') {
                uppercaseMask |= (1 << (ch - 'A'));
            }
        }
        
        // 2. 두 마스크를 AND 연산하여 공통으로 등장한 알파벳 추출
        int specialMask = lowercaseMask & uppercaseMask;
        
        // 3. 켜진 비트(1)의 개수를 세어 결과 반환
        return Integer.bitCount(specialMask);
    }
}
