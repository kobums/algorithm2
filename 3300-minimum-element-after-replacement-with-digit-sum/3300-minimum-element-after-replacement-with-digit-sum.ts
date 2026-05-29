function minElement(nums: number[]): number {
    let minSum = Infinity;

    for (let num of nums) {
        let digitSum = 0;
        let temp = num;
        
        // 각 자릿수의 합 계산
        while (temp > 0) {
            digitSum += temp % 10;
            temp = Math.floor(temp / 10);
        }

        // 최솟값 갱신
        minSum = Math.min(minSum, digitSum);
    }

    return minSum;
}
