class Solution {
    public boolean asteroidsDestroyed(int mass, int[] asteroids) {
        // 소행성 크기순으로 오름차순 정렬
        Arrays.sort(asteroids);
        
        long currentMass = mass;
        
        for (int asteroid : asteroids) {
            // 행성의 질량이 소행성보다 작으면 파괴 불가
            if (currentMass < asteroid) {
                return false;
            }
            // 소행성을 파괴하고 질량 추가
            currentMass += asteroid;
        }
        
        return true;
    }
}
