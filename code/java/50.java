class Solution {
  // Recursion
  public double myPow(double x, int n) {
    if (n == 0)
      return 1;
    if (x == 1)
      return x;

    if (n < 0) {
      // edge case, avoid -n overflow
      if (n == Integer.MIN_VALUE) {
        n = Integer.MAX_VALUE;
        return 1.0D / (x * myPow(x, n));
      }
      return 1.0D / myPow(x, -n);
    }

    if (n % 2 == 0) {
      return myPow(x * x, n / 2);
    } else {
      return myPow(x * x, (n - 1) / 2) * x;
    }
  }

  public static void main(String[] args) {
    Solution solution = new Solution();
    assert solution.myPow(2.00000D, 10) == 1024D;
    assert solution.myPow(2.10000D, 3) == 9.261000000000001;
    assert solution.myPow(2.00000D, -2) == 0.25D;
    assert solution.myPow(1.00000D, -2147483648) == 1D;
    assert solution.myPow(2.00000D, -2147483648) == 0D;
  }
}