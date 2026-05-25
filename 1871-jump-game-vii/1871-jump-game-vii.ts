function canReach(s: string, minJump: number, maxJump: number): boolean {
    const n = s.length;
    if (s[n - 1] === '1') return false;

    const isReachable = new Array<boolean>(n).fill(false);
    isReachable[0] = true;
    
    let pre = 0;
    
    for (let i = 0; i < n; i++) {
        if (!isReachable[i]) continue;
        
        const start = Math.max(i + minJump, pre);
        const end = Math.min(i + maxJump, n - 1);
        
        for (let j = start; j <= end; j++) {
            if (s[j] === '0') {
                isReachable[j] = true;
            }
        }
        
        pre = end + 1;
    }
    
    return isReachable[n - 1];
}
