/**
 * @param {number[]} nums
 * @return {number}
 */
var maxCoins = function(nums) {
    // 양 끝에 가상의 1 추가
    const arr = [1, ...nums, 1];
    const n = arr.length;
    
    // dp[i][j] = i~j 범위(exclusive)에서 최대 코인
    const dp = Array.from({ length: n }, () => Array(n).fill(0));
    
    // 구간 길이를 늘려가며 계산
    for (let len = 1; len <= nums.length; len++) {
        for (let left = 1; left <= nums.length - len + 1; left++) {
            const right = left + len - 1;
            
            // k를 마지막으로 터뜨리는 경우
            for (let k = left; k <= right; k++) {
                const coins = arr[left - 1] * arr[k] * arr[right + 1];
                const leftCoins = dp[left][k - 1];
                const rightCoins = dp[k + 1][right];
                
                dp[left][right] = Math.max(
                    dp[left][right],
                    leftCoins + coins + rightCoins
                );
            }
        }
    }
    
    return dp[1][nums.length];
};