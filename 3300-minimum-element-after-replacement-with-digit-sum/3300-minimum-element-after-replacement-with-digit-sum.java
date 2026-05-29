class Solution {
    public int minElement(int[] nums) {
        int minSum = Integer.MAX_VALUE;
        
        for (int num : nums) {
            int digitSum = 0;
            int temp = num;
            
            // Calculate the sum of digits
            while (temp > 0) {
                digitSum += temp % 10;
                temp /= 10;
            }
            
            // Update the global minimum
            minSum = Math.min(minSum, digitSum);
        }
        
        return minSum;
    }
}
