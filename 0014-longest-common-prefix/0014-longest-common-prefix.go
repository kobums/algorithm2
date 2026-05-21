func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    // Iterate through the characters of the first string
    for i := 0; i < len(strs[0]); i++ {
        char := strs[0][i]
        
        // Compare this character with the same index in all other strings
        for j := 1; j < len(strs); j++ {
            // Check if we've reached the end of current string or characters don't match
            if i == len(strs[j]) || strs[j][i] != char {
                return strs[0][:i]
            }
        }
    }

    return strs[0]
}