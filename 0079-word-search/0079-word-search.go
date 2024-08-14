func exist(board [][]byte, word string) bool {
    rows := len(board)
    cols := len(board[0])

    // Helper function for backtracking
    var backtrack func(row, col, index int) bool
    backtrack = func(row, col, index int) bool {
        // Base case: if the entire word is found
        if index == len(word) {
            return true
        }

        // Check the boundaries and if the current cell matches the current character in word
        if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != word[index] {
            return false
        }

        // Mark the cell as visited by temporarily changing its value
        temp := board[row][col]
        board[row][col] = '#'

        // Explore all four directions (up, down, left, right)
        found := backtrack(row-1, col, index+1) || // up
            backtrack(row+1, col, index+1) || // down
            backtrack(row, col-1, index+1) || // left
            backtrack(row, col+1, index+1) // right

        // Restore the cell's original value
        board[row][col] = temp

        return found
    }

    // Iterate over each cell in the board as the starting point
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if backtrack(i, j, 0) {
                return true
            }
        }
    }

    return false
}
