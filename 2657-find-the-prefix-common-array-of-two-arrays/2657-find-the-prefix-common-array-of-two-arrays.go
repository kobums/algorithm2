func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A)
    ans := make([]int, n)
    vis := make([]int, n+1)
    commonCount := 0
    
    for i := 0; i < n; i++ {
        // Increment visit count for A[i]
        vis[A[i]]++
        if vis[A[i]] == 2 {
            commonCount++
        }
        
        // Increment visit count for B[i]
        vis[B[i]]++
        if vis[B[i]] == 2 {
            commonCount++
        }
        
        ans[i] = commonCount
    }
    
    return ans
}