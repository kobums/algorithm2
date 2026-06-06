class Solution {
    public int[] leftRightDifference(int[] nums) {
        int totalSum = 0;
        for (int num : nums) {
            totalSum += num;
        }

        int[] ans = new int[nums.length];
        int leftSum = 0;

        for (int i = 0; i < nums.length; i++) {
            int rightSum = totalSum - leftSum - nums[i];
            
            // 두 합의 절대값 계산 후 저장
            ans[i] = Math.abs(leftSum - rightSum);
            
            // 다음 인덱스를 위해 현재 값을 왼쪽 합에 누적
            leftSum += nums[i];
        }

        return ans;
    }
}
