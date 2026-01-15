class Solution {
    private int[][] memo;
    private int[][] matrix;
    private int m, n;
    private int[][] dirs = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
    
    public int longestIncreasingPath(int[][] matrix) {
        this.matrix = matrix;
        m = matrix.length;
        n = matrix[0].length;
        memo = new int[m][n];
        
        int result = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                result = Math.max(result, dfs(i, j));
            }
        }
        
        return result;
    }
    
    private int dfs(int r, int c) {
        if (memo[r][c] != 0) return memo[r][c];
        
        int maxLen = 1;
        for (int[] d : dirs) {
            int nr = r + d[0];
            int nc = c + d[1];
            
            if (nr >= 0 && nr < m && nc >= 0 && nc < n 
                && matrix[nr][nc] > matrix[r][c]) {
                maxLen = Math.max(maxLen, 1 + dfs(nr, nc));
            }
        }
        
        memo[r][c] = maxLen;
        return maxLen;
    }
}