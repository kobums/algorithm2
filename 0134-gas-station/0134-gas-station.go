func canCompleteCircuit(gas []int, cost []int) int {
    totalTank, currTank := 0, 0
    startStation := 0
    
    for i := 0; i < len(gas); i++ {
        totalTank += gas[i] - cost[i]
        currTank += gas[i] - cost[i]
        
        // 현재 탱크가 음수면 더 이상 진행할 수 없으므로 다음 주유소로 이동
        if currTank < 0 {
            startStation = i + 1
            currTank = 0
        }
    }
    
    // 전체 합이 음수면 순회를 완료할 수 없으므로 -1을 반환
    if totalTank < 0 {
        return -1
    }
    
    return startStation
}