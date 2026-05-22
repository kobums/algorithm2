class Solution {
    public int[] findThePrefixCommonArray(int[] A, int[] B) {
        int n = A.length;
        int[] C = new int[n];
        int[] count = new int[n + 1];
        int commonCount = 0;
        
        for (int i = 0; i < n; i++) {
            count[A[i]]++;
            if (count[A[i]] == 2) {
                commonCount++;
            }
            
            count[B[i]]++;
            if (count[B[i]] == 2) {
                commonCount++;
            }
            
            C[i] = commonCount;
        }
        
        return C;
    }
}