/**
 * @param {string} s
 * @param {string} t
 * @return {number}
 */
var numDistinct = function(s, t) {
    const m = s.length;
    const n = t.length;
    
    // dp[i][j] = s[0..i-1]에서 t[0..j-1]를 만드는 경우의 수
    const dp = Array.from({ length: m + 1 }, () => Array(n + 1).fill(0));
    
    // 빈 문자열 t를 만드는 방법은 항상 1가지
    for (let i = 0; i <= m; i++) {
        dp[i][0] = 1;
    }
    
    for (let i = 1; i <= m; i++) {
        for (let j = 1; j <= n; j++) {
            if (s[i - 1] === t[j - 1]) {
                // 현재 문자 사용 + 안 사용
                dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j];
            } else {
                // 현재 문자 스킵
                dp[i][j] = dp[i - 1][j];
            }
        }
    }
    
    return dp[m][n];
};