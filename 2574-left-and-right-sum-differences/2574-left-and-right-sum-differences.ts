function leftRightDifference(nums: number[]): number[] {
    const totalSum: number = nums.reduce((acc, curr) => acc + curr, 0);
    const ans: number[] = new Array(nums.length);
    
    let leftSum: number = 0;
    
    for (let i = 0; i < nums.length; i++) {
        const rightSum: number = totalSum - leftSum - nums[i];
        
        // 두 합의 절대값 계산 후 저장
        ans[i] = Math.abs(leftSum - rightSum);
        
        // 다음 루프를 위해 현재 값을 왼쪽 합에 누적
        leftSum += nums[i];
    }
    
    return ans;
}
