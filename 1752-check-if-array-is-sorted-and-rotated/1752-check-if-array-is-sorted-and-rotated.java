class Solution {
    public boolean check(int[] nums) {
        int count = 0;
        int n = nums.length;
        
        for (int i = 0; i < n; i++) {
            // 현재 원소가 다음 원소보다 크면 회전이 일어난 지점입니다.
            // (i + 1) % n을 통해 마지막 원소와 첫 번째 원소를 비교합니다.
            if (nums[i] > nums[(i + 1) % n]) {
                count++;
            }
            
            // 내림차순 지점이 2번 이상 발생하면 원래 정렬된 배열일 수 없습니다.
            if (count > 1) {
                return false;
            }
        }
        
        return true;
    }
}
