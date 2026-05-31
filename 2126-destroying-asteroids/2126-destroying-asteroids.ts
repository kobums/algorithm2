function asteroidsDestroyed(mass: number, asteroids: number[]): boolean {
    // 소행성을 오름차순으로 정렬합니다
    asteroids.sort((a, b) => a - b);
    
    let currentMass: number = mass;

    for (const asteroid of asteroids) {
        // 행성의 현재 질량이 소행성의 질량보다 작으면 소행성을 파괴할 수 없습니다
        if (currentMass < asteroid) {
            return false;
        }
        // 소행성을 파괴하고 그 질량을 행성에 더합니다
        currentMass += asteroid;
    }

    return true;
}
