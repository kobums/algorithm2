class Solution {
    public String longestPalindrome(String s) {
        int n = s.length();
        if (n < 2) return s;

        boolean[][] dp = new boolean[n][n];
        int start = 0;
        int maxLen = 1;

        // 길이 1 초기화
        for (int i = 0; i < n; i++) {
            dp[i][i] = true;
        }

        // 길이(len) 루프
        for (int len = 2; len <= n; len++) {
            // 시작 인덱스(i) 루프
            for (int i = 0; i <= n - len; i++) {
                int j = i + len - 1; // 끝 인덱스

                if (s.charAt(i) == s.charAt(j)) {
                    // 길이 2이거나 내부가 팰린드롬일 때
                    if (len == 2 || dp[i + 1][j - 1]) {
                        dp[i][j] = true;
                        
                        if (len > maxLen) {
                            start = i;
                            maxLen = len;
                        }
                    }
                }
            }
        }

        return s.substring(start, start + maxLen);
    }
}