class Solution {
    public long totalWaviness(long num1, long num2) {
        return countUpTo(num2) - countUpTo(num1 - 1);
    }

    private long countUpTo(long num) {
        if (num < 100) {
            return 0;
        }

        String s = Long.toString(num);
        int n = s.length();

        // dp[idx][prev1][prev2][isLess][isStart]
        // prev1, prev2는 0~9(미지정 시 10) 총 11가지 상태 필요
        long[][][][][] dp = new long[n][11][11][2][2];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < 11; j++) {
                for (int k = 0; k < 11; k++) {
                    for (int l = 0; l < 2; l++) {
                        Arrays.fill(dp[i][j][k][l], -1);
                    }
                }
            }
        }

        return dfs(0, 10, 10, false, true, s, dp);
    }

    private long dfs(int idx, int prev1, int prev2, boolean isLess, boolean isStart, String s, long[][][][][] dp) {
        int n = s.length();
        if (idx == n) {
            return 0;
        }

        int il = isLess ? 1 : 0;
        int is = isStart ? 1 : 0;

        if (dp[idx][prev1][prev2][il][is] != -1) {
            return dp[idx][prev1][prev2][il][is];
        }

        int limit = isLess ? 9 : (s.charAt(idx) - '0');
        long res = 0;

        for (int i = 0; i <= limit; i++) {
            boolean nextLess = isLess || (i < limit);
            boolean nextStart = isStart && (i == 0);

            // 피크(Peak) 및 밸리(Valley) 체크
            long waviness = 0;
            if (!isStart && prev1 != 10 && prev2 != 10) {
                if (prev1 > prev2 && prev1 > i) { // Peak
                    waviness = 1;
                } else if (prev1 < prev2 && prev1 < i) { // Valley
                    waviness = 1;
                }
            }

            int n1, n2;
            if (nextStart) {
                n1 = 10;
                n2 = 10;
            } else {
                n1 = i;
                n2 = prev1;
            }

            // 현재 자릿수의 waviness 조합 수 계산 후 재귀 호출 누적
            long combinations = countCombinations(n, idx + 1, nextLess, s);
            res += waviness * combinations + dfs(idx + 1, n1, n2, nextLess, nextStart, s, dp);
        }

        return dp[idx][prev1][prev2][il][is] = res;
    }

    // 이후 자릿수에서 발생할 수 있는 유효 숫자 조합의 수를 반환하는 보조 함수
    private long countCombinations(int length, int idx, boolean isLess, String s) {
        if (idx == length) {
            return 1;
        }
        if (isLess) {
            long res = 1;
            for (int i = idx; i < length; i++) {
                res *= 10;
            }
            return res;
        }
        return Long.parseLong(s.substring(idx)) + 1;
    }
}
