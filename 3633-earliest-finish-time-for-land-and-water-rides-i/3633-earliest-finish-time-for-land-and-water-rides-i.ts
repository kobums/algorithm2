function earliestFinishTime(
    landStartTime: number[],
    landDuration: number[],
    waterStartTime: number[],
    waterDuration: number[]
): number {
    let minFinishTime = Infinity;
    const n = landStartTime.length;
    const m = waterStartTime.length;

    // 시나리오 1: Land ride -> Water ride
    for (let i = 0; i < n; i++) {
        for (let j = 0; j < m; j++) {
            // land ride가 끝나는 시간 = (시작 가능 시간 + 소요 시간)
            const landEnd = Math.max(landStartTime[i], landStartTime[i]) + landDuration[i];
            // water ride는 land ride가 끝난 후 또는 워터 라이드가 오픈했을 때 탈 수 있음
            const totalEnd = Math.max(landEnd, waterStartTime[j]) + waterDuration[j];
            minFinishTime = Math.min(minFinishTime, totalEnd);
        }
    }

    // 시나리오 2: Water ride -> Land ride
    for (let i = 0; i < n; i++) {
        for (let j = 0; j < m; j++) {
            // water ride가 끝나는 시간
            const waterEnd = Math.max(waterStartTime[j], waterStartTime[j]) + waterDuration[j];
            // land ride는 water ride가 끝난 후 또는 랜드 라이드가 오픈했을 때 탈 수 있음
            const totalEnd = Math.max(waterEnd, landStartTime[i]) + landDuration[i];
            minFinishTime = Math.min(minFinishTime, totalEnd);
        }
    }

    return minFinishTime;
}
