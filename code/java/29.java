class Solution {
  public int divide(int dividend, int divisor) {

    boolean negative = (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0);

    // avoid to handle edge case in Integer
    long a = dividend;
    long b = divisor;
    if (a < 0)
      a = -a;
    if (b < 0)
      b = -b;

    long quotient = div(a, b);

    if (!negative) {
      // edge case, avoid overflow
      if (quotient > (long) Integer.MAX_VALUE) {
        return Integer.MAX_VALUE;
      }
    }

    return negative ? (int) -quotient : (int) quotient;
  }

  // Recursion
  private long div(long dividend, long divisor) {
    // base case
    if (dividend - divisor < 0)
      return 0;
    if (dividend - divisor == 0)
      return 1;

    long quotient = 1;
    long divisor1 = divisor;

    // try to double quotient and divisor
    while (dividend - (divisor1 + divisor1) >= 0) {
      quotient += quotient;
      divisor1 += divisor1;
    }
    return quotient + div((dividend - divisor1), divisor);
  }

  public static void main(String[] args) {
    Solution solution = new Solution();
    assert solution.divide(10, 3) == 3;
    assert solution.divide(7, -3) == -2;
    assert solution.divide(0, 1) == 0;
    assert solution.divide(1, 1) == 1;
    assert solution.divide(Integer.MIN_VALUE, -1073741824) == 2;
    assert solution.divide(Integer.MAX_VALUE, Integer.MIN_VALUE) == 0;
  }
}