class Solution {
    public int earliestFinishTime(int[] landStartTime, int[] landDuration, int[] waterStartTime, int[] waterDuration) {
        int minTotalFinishTime = Integer.MAX_VALUE;
        int n = landStartTime.length;
        int m = waterStartTime.length;

        // 1. 지상(i) -> 수상(j) 순서로 타는 모든 조합 확인
        for (int i = 0; i < n; i++) {
            int landEndTime = landStartTime[i] + landDuration[i]; // 지상 놀이기구 종료 시간
            for (int j = 0; j < m; j++) {
                // 지상 기구가 끝난 시간과 수상 기구 시작 시간 중 더 늦은 시간에 탑승 시작
                int waterEndTime = Math.max(landEndTime, waterStartTime[j]) + waterDuration[j];
                minTotalFinishTime = Math.min(minTotalFinishTime, waterEndTime);
            }
        }

        // 2. 수상(j) -> 지상(i) 순서로 타는 모든 조합 확인
        for (int j = 0; j < m; j++) {
            int waterEndTime = waterStartTime[j] + waterDuration[j]; // 수상 놀이기구 종료 시간
            for (int i = 0; i < n; i++) {
                // 수상 기구가 끝난 시간과 지상 기구 시작 시간 중 더 늦은 시간에 탑승 시작
                int landEndTime = Math.max(waterEndTime, landStartTime[i]) + landDuration[i];
                minTotalFinishTime = Math.min(minTotalFinishTime, landEndTime);
            }
        }

        return minTotalFinishTime;
    }
}
