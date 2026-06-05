function totalWaviness(num1: number, num2: number): number {
    // JS/TS 안전 정수 범위를 벗어날 수 있으므로 계산은 BigInt로 진행합니다.
    const b1 = BigInt(num1);
    const b2 = BigInt(num2);
    
    return Number(countUpTo(b2) - countUpTo(b1 - 1n));
}

function countUpTo(num: bigint): bigint {
    if (num < 100n) {
        return 0n;
    }

    const s = num.toString();
    const n = s.length;

    // DP 상태 개수 최적화 및 1차원 평탄화 배열 크기 계산
    // idx(최대 16) * prev1(11) * prev2(11) * isLess(2) * isStart(2)
    const STATE_SIZE = n * 11 * 11 * 2 * 2;
    const dp = new BigInt64Array(STATE_SIZE).fill(-1n);

    // 5차원 인덱스를 1차원 배열 인덱스로 매핑해주는 헬퍼 함수
    function getStateIndex(idx: number, prev1: number, prev2: number, isLess: number, isStart: number): number {
        return ((((idx * 11 + prev1) * 11 + prev2) * 2 + isLess) * 2 + isStart);
    }

    function dfs(idx: number, prev1: number, prev2: number, isLess: boolean, isStart: boolean): bigint {
        if (idx === n) {
            return 0n;
        }

        const il = isLess ? 1 : 0;
        const is = isStart ? 1 : 0;
        const stateIdx = getStateIndex(idx, prev1, prev2, il, is);

        if (dp[stateIdx] !== -1n) {
            return dp[stateIdx];
        }

        const limit = isLess ? 9 : parseInt(s[idx], 10);
        let res = 0n;

        for (let i = 0; i <= limit; i++) {
            const nextLess = isLess || i < limit;
            const nextStart = isStart && i === 0;

            // 피크(Peak) 및 밸리(Valley) 체크 조건식
            let waviness = 0n;
            if (!isStart && prev1 !== 10 && prev2 !== 10) {
                if (prev1 > prev2 && prev1 > i) {       // Peak
                    waviness = 1n;
                } else if (prev1 < prev2 && prev1 < i) { // Valley
                    waviness = 1n;
                }
            }

            let n1 = 10, n2 = 10;
            if (!nextStart) {
                n1 = i;
                n2 = prev1;
            }

            // 현재 자릿수의 유효 숫자 조합 수 계산 후 재귀 호출 누적
            const combinations = countCombinations(n, idx + 1, nextLess, s);
            res += waviness * combinations + dfs(idx + 1, n1, n2, nextLess, nextStart);
        }

        dp[stateIdx] = res;
        return res;
    }

    return dfs(0, 10, 10, false, true);
}

// 해당 자릿수 뒤로 나타날 수 있는 유효 숫자 조합의 총 개수를 구하는 보조 함수
function countCombinations(length: number, idx: number, isLess: boolean, s: string): bigint {
    if (idx === length) {
        return 1n;
    }
    if (isLess) {
        let res = 1n;
        for (let i = idx; i < length; i++) {
            res *= 10n;
        }
        return res;
    }
    return BigInt(s.substring(idx)) + 1n;
}
