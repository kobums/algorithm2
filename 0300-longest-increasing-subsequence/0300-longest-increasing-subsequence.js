/**
 * @param {number[]} nums
 * @return {number}
 */
var lengthOfLIS = function(nums) {
    if (nums.length === 0) return 0;

    // 1. 모든 위치에서의 LIS 최소 길이는 1 (자기 자신)
    const dp = new Array(nums.length).fill(1);
    let maxLen = 1;

    // 2. i는 현재 끝점, j는 i 이전의 모든 원소 탐색
    for (let i = 1; i < nums.length; i++) {
        for (let j = 0; j < i; j++) {
            // 증가하는 관계라면
            if (nums[i] > nums[j]) {
                // 이전 길이 + 1과 현재 기록된 길이 중 큰 값 선택
                dp[i] = Math.max(dp[i], dp[j] + 1);
            }
        }
        // 전체 최대 길이 갱신
        maxLen = Math.max(maxLen, dp[i]);
    }

    return maxLen;
};