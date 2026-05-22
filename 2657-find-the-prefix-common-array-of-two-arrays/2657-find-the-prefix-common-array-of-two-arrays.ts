function findThePrefixCommonArray(A: number[], B: number[]): number[] {
    const n = A.length;
    const ans: number[] = new Array(n);
    const freq: number[] = new Array(n + 1).fill(0);
    
    let commonCount = 0;
    
    for (let i = 0; i < n; i++) {
        freq[A[i]]++;
        if (freq[A[i]] === 2) {
            commonCount++;
        }
        
        freq[B[i]]++;
        if (freq[B[i]] === 2) {
            commonCount++;
        }
        
        ans[i] = commonCount;
    }
    
    return ans;

};