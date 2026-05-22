/**
 * @param {number[]} A
 * @param {number[]} B
 * @return {number[]}
 */
var findThePrefixCommonArray = function(A, B) {
    const n = A.length;
    const freq = new Array(n + 1).fill(0);
    const result = new Array(n);
    let commonCount = 0;

    for (let i = 0; i < n; i++) {
        // A[i]의 빈도수 체크
        freq[A[i]]++;
        if (freq[A[i]] === 2) {
            commonCount++;
        }

        // B[i]의 빈도수 체크
        freq[B[i]]++;
        if (freq[B[i]] === 2) {
            commonCount++;
        }

        result[i] = commonCount;
    }

    return result;
};