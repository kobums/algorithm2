class Solution {
    public int earliestFinishTime(int[] landStartTime, int[] landDuration, int[] waterStartTime, int[] waterDuration) {
        long plan1 = findEarliestFinish(landStartTime, landDuration, waterStartTime, waterDuration);
        long plan2 = findEarliestFinish(waterStartTime, waterDuration, landStartTime, landDuration);
        return (int) Math.min(plan1, plan2); // Cast the final minimum long value to int
    }

    private long findEarliestFinish(int[] firstStart, int[] firstDuration, int[] secondStart, int[] secondDuration) {
        long minFinish = Long.MAX_VALUE;
        
        for (int i = 0; i < firstStart.length; i++) {
            long finishTime = (long) firstStart[i] + firstDuration[i];
            minFinish = Math.min(minFinish, finishTime);
        }

        long earliestFinishTime = Long.MAX_VALUE;
        for (int i = 0; i < secondStart.length; i++) {
            long startSecond = Math.max((long) secondStart[i], minFinish);
            long finishSecond = startSecond + secondDuration[i];
            earliestFinishTime = Math.min(earliestFinishTime, finishSecond);
        }

        return earliestFinishTime;
    }
}
