func isPalindrome(s string) bool {
    str := processString(s)

    for i := 0; i<len(str); i++ {

        if str[i] != str[len(str)-1-i] {
            return false
        }
    }

    return true

}

func processString(s string) string {
    var builder strings.Builder
    for _, ch := range s {
        if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
            builder.WriteRune(unicode.ToLower(ch))
        }
    }
    return builder.String()
}