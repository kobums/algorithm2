/**
 * @param {number} n
 * @return {number[]}
 */
var countBits = function(n) {
    const result = [];
    
    for (let i = 0; i <= n; i++) {
        let count = 0;
        let num = i;
        
        while (num !== 0) {
            num = num & (num - 1);  // 1비트 제거
            count++;
        }
        
        result.push(count);
    }
    
    return result;
};