/**
 * @param {string} s
 * @param {string} t
 * @return {string}
 */
var minWindow = function(s, t) {
    if (!s || !t || s.length < t.length) return "";

    // 1. Map the frequency of characters needed from 't'
    const targetFreq = {};
    for (const char of t) {
        targetFreq[char] = (targetFreq[char] || 0) + 1;
    }

    const windowFreq = {};
    let left = 0, right = 0;
    let matchedCount = 0; // Number of unique characters that meet target frequency
    const requiredCount = Object.keys(targetFreq).length;
    
    let minLen = Infinity;
    let minWindow = "";

    // 2. Expand the window using the right pointer
    while (right < s.length) {
        const char = s[right];
        windowFreq[char] = (windowFreq[char] || 0) + 1;

        // If this character's count in window matches its count in target
        if (targetFreq[char] && windowFreq[char] === targetFreq[char]) {
            matchedCount++;
        }

        // 3. Contract the window from the left once all characters are matched
        while (matchedCount === requiredCount) {
            const currentLen = right - left + 1;
            if (currentLen < minLen) {
                minLen = currentLen;
                minWindow = s.substring(left, right + 1);
            }

            const leftChar = s[left];
            // If the character being removed is part of the target
            if (targetFreq[leftChar] && windowFreq[leftChar] === targetFreq[leftChar]) {
                matchedCount--;
            }
            windowFreq[leftChar]--;
            left++;
        }
        right++;
    }

    return minWindow;
};