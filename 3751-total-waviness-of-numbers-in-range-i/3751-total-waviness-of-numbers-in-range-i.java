class Solution {
    public int totalWaviness(int num1, int num2) {
        int totalSum = 0;
        for (int i = num1; i <= num2; i++) {
            totalSum += countWaviness(i);
        }
        return totalSum;
    }

    private int countWaviness(int n) {
        if (n < 100) {
            return 0; // 3자리 미만은 파형이 성립하지 않음
        }

        List<Integer> digits = new ArrayList<>();
        int temp = n;
        while (temp > 0) {
            digits.add(temp % 10);
            temp /= 10;
        }

        int count = 0;
        int m = digits.size();
        
        // 역순(오른쪽 자릿수부터) 채워지므로 양옆의 인덱스 비교가 유효합니다.
        for (int i = 1; i < m - 1; i++) {
            int prev = digits.get(i - 1);
            int curr = digits.get(i);
            int next = digits.get(i + 1);

            // Peak (고점) 조건
            if (curr > prev && curr > next) {
                count++;
            }
            // Valley (저점) 조건
            else if (curr < prev && curr < next) {
                count++;
            }
        }
        return count;
    }
}
