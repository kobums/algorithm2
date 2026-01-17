/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {
    if (s.length < 2) return s;

    const n = s.length;
    // DP 테이블 초기화 (false로 채움)
    const dp = Array.from(Array(n), () => Array(n).fill(false));
    
    let start = 0;
    let maxLen = 1;

    // 길이 1인 경우 (대각선) 모두 true
    for (let i = 0; i < n; i++) {
        dp[i][i] = true;
    }

    // 길이를 2부터 n까지 늘려가며 확인
    for (let len = 2; len <= n; len++) {
        // i는 시작 인덱스
        for (let i = 0; i <= n - len; i++) {
            // j는 끝 인덱스
            let j = i + len - 1;
            
            // 양 끝 문자가 같을 때
            if (s[i] === s[j]) {
                // 길이가 2이거나, 내부 문자열(dp[i+1][j-1])이 팰린드롬인 경우
                if (len === 2 || dp[i + 1][j - 1]) {
                    dp[i][j] = true;
                    // 최대 길이 갱신
                    if (len > maxLen) {
                        start = i;
                        maxLen = len;
                    }
                }
            }
        }
    }

    return s.substring(start, start + maxLen);
};