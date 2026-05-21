/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
     if (strs === null || strs.length === 0) return "";
    
    // 첫 번째 문자열을 기준(prefix)으로 잡습니다.
    let prefix = strs[0];
    
    for (let i = 1; i < strs.length; i++) {
        // prefix가 현재 문자열(strs[i])의 시작 부분과 일치하지 않는 동안,
        // prefix의 마지막 글자를 하나씩 제거합니다.
        while (strs[i].indexOf(prefix) !== 0) {
            prefix = prefix.substring(0, prefix.length - 1);
            
            // 만약 prefix가 빈 문자열이 되면 공통 접두사가 없는 것입니다.
            if (prefix === "") return "";
        }
    }
    
    return prefix;
};