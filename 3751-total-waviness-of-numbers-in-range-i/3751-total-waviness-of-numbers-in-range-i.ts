function totalWaviness(num1: number, num2: number): number {
    let totalWaviness = 0;
    
    for (let i = num1; i <= num2; i++) {
        totalWaviness += getWaviness(i);
    }
    
    return totalWaviness;
};

function getWaviness(num: number): number {
    const digits: number[] = [];
    let temp = num;
    
    // 숫자의 각 자릿수를 배열에 저장 (역순)
    while (temp > 0) {
        digits.push(temp % 10);
        temp = Math.floor(temp / 10);
    }
    
    // 배열을 올바른 자릿수 순서로 뒤집기
    digits.reverse();
    
    if (digits.length < 3) {
        return 0;
    }
    
    let wavinessCount = 0;
    
    // 각 자릿수를 순회하며 봉우리와 골짜기 개수 계산
    for (let i = 1; i < digits.length - 1; i++) {
        const prev = digits[i - 1];
        const current = digits[i];
        const next = digits[i + 1];
        
        // 봉우리: 현재 숫자가 양옆 숫자보다 모두 클 때
        const isPeak = current > prev && current > next;
        
        // 골짜기: 현재 숫자가 양옆 숫자보다 모두 작을 때
        const isValley = current < prev && current < next;
        
        if (isPeak || isValley) {
            wavinessCount++;
        }
    }
    
    return wavinessCount;
}
