class Solution {
    public int maxProfit(int[] prices) {
        int n = prices.length;
        if (n <= 1)
            return 0;

        int[][] table = new int[n][2];

        table[0][0] = 0;
        table[0][1] = -prices[0];

        for (int i = 1; i < n; i++) {
            table[i][0] = Math.max(table[i - 1][0], table[i - 1][1] + prices[i]);
            table[i][1] = Math.max(table[i - 1][1], table[i - 1][0] - prices[i]);
        }

        return table[n - 1][0];
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        assert solution.maxProfit(new int[] { 7, 1, 5, 3, 6, 4 }) == 7;
        assert solution.maxProfit(new int[] { 1, 2, 3, 4, 5 }) == 4;
        assert solution.maxProfit(new int[] { 7, 6, 4, 3, 1 }) == 0;
    }
}