func numberOfSpecialChars(word string) int {
    // Arrays to track presence of lowercase and uppercase letters
    lower := make([]bool, 26)
    upper := make([]bool, 26)
    
    for _, char := range word {
        if char >= 'a' && char <= 'z' {
            lower[char-'a'] = true
        } else if char >= 'A' && char <= 'Z' {
            upper[char-'A'] = true
        }
    }
    
    count := 0
    for i := 0; i < 26; i++ {
        if lower[i] && upper[i] {
            count++
        }
    }
    
    return count
}