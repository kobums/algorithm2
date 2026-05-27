class Solution {
    public int numberOfSpecialChars(String word) {
        // 알파벳 26개에 대해 첫 번째 대문자와 마지막 소문자의 인덱스를 저장할 배열
        int[] firstUpper = new int[26];
        int[] lastLower = new int[26];
        
        // 인덱스가 0일 수도 있으므로 초기값을 -1로 설정
        Arrays.fill(firstUpper, -1);
        Arrays.fill(lastLower, -1);
        
        // 1. 각 문자의 위치를 한 번에 탐색하며 기록
        for (int i = 0; i < word.length(); i++) {
            char c = word.charAt(i);
            
            if (Character.isLowerCase(c)) {
                int idx = c - 'a';
                lastLower[idx] = i; // 마지막 소문자 인덱스 갱신
            } else {
                int idx = c - 'A';
                if (firstUpper[idx] == -1) {
                    firstUpper[idx] = i; // 첫 번째 대문자 인덱스 기록
                }
            }
        }
        
        int specialCount = 0;
        
        // 2. 소문자와 대문자가 모두 존재하고, 모든 소문자가 대문자보다 먼저 나왔는지 확인
        for (int i = 0; i < 26; i++) {
            if (lastLower[i] != -1 && firstUpper[i] != -1) {
                if (lastLower[i] < firstUpper[i]) {
                    specialCount++;
                }
            }
        }
        
        return specialCount;
    }
}
