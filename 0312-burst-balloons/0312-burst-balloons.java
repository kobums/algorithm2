class Solution {
    public int maxCoins(int[] nums) {
        int n = nums.length;
        
        // 양 끝에 1 추가
        int[] arr = new int[n + 2];
        arr[0] = arr[n + 1] = 1;
        for (int i = 0; i < n; i++) {
            arr[i + 1] = nums[i];
        }
        
        // dp[i][j] = i~j 범위에서 최대 코인
        int[][] dp = new int[n + 2][n + 2];
        
        // 구간 길이를 늘려가며 계산
        for (int len = 1; len <= n; len++) {
            for (int left = 1; left <= n - len + 1; left++) {
                int right = left + len - 1;
                
                for (int k = left; k <= right; k++) {
                    int coins = arr[left - 1] * arr[k] * arr[right + 1];
                    int total = dp[left][k - 1] + coins + dp[k + 1][right];
                    
                    dp[left][right] = Math.max(dp[left][right], total);
                }
            }
        }
        
        return dp[1][n];
    }
}