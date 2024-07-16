func generateParenthesis(n int) []string {
    result := make([]string, 0)   
    generate(&result, "(", 1, 0, n)
    return result
}

func generate(result *[]string, s string, open int, close int, max int) {
    if (open + close) == (max * 2) {
        *result = append(*result, s)
        return
    }
    if open < max {
        generate(result, s+"(", open+1, close, max)
    }
    if close < max && open > close {
        generate(result, s+")", open, close+1, max)
    }
}