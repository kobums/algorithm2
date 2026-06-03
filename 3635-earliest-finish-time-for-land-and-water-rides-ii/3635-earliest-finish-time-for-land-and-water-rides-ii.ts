function earliestFinishTime(
    landStartTime: number[], 
    landDuration: number[], 
    waterStartTime: number[], 
    waterDuration: number[]
): number {
    
    function getMinFinishTime(
        start1: number[], 
        dur1: number[], 
        start2: number[], 
        dur2: number[]
    ): number {
        // 1. 첫 번째 라이드들의 종료 시간 중 최소값을 찾음
        let minFirstEnd = Infinity;
        for (let i = 0; i < start1.length; i++) {
            const end = start1[i] + dur1[i];
            if (end < minFirstEnd) {
                minFirstEnd = end;
            }
        }

        // 2. 가장 빨리 끝난 시점을 기준으로 두 번째 라이드를 순회하며 최소 최종 종료 시간 계산
        let minTotalEnd = Infinity;
        for (let j = 0; j < start2.length; j++) {
            const startTime = Math.max(minFirstEnd, start2[j]);
            const finalEnd = startTime + dur2[j];
            if (finalEnd < minTotalEnd) {
                minTotalEnd = finalEnd;
            }
        }

        return minTotalEnd;
    }

    // Case 1: Land -> Water
    const case1 = getMinFinishTime(landStartTime, landDuration, waterStartTime, waterDuration);
    // Case 2: Water -> Land
    const case2 = getMinFinishTime(waterStartTime, waterDuration, landStartTime, landDuration);

    return Math.min(case1, case2);
}
