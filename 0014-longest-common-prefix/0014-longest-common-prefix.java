class Solution {
    public String longestCommonPrefix(String[] strs) {
        if (strs == null || strs.length == 0) {
            return "";
        }
        
        // 1. 문자열 배열 사전순 정렬
        Arrays.sort(strs);
        
        // 2. 가장 첫 번째와 마지막 문자열 가져오기
        String s1 = strs[0];
        String s2 = strs[strs.length - 1];
        
        int idx = 0;
        // 3. 두 문자열의 처음부터 글자 비교
        while (idx < s1.length() && idx < s2.length()) {
            if (s1.charAt(idx) == s2.charAt(idx)) {
                idx++;
            } else {
                break;
            }
        }
        
        // 4. 일치하는 길이만큼 잘라서 반환
        return s1.substring(0, idx);
    }
}