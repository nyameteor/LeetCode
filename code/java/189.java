import java.util.Arrays;

class Solution {
  // Array Flip-over
  public void rotate(int[] nums, int k) {
    k = k % nums.length;
    reverse(nums, 0, nums.length);
    reverse(nums, 0, k);
    reverse(nums, k, nums.length);
  }

  private void reverse(int[] nums, int start, int end) {
    for (int i = start; i < start + (end - start) / 2; i++) {
      int temp = nums[end - 1 - (i - start)];
      nums[end - 1 - (i - start)] = nums[i];
      nums[i] = temp;
    }
  }

  public static void main(String[] args) {
    Solution solution = new Solution();
    int[] nums;

    nums = new int[] { 1, 2, 3, 4, 5, 6, 7 };
    solution.rotate(nums, 3);
    assert Arrays.equals(nums, new int[] { 5, 6, 7, 1, 2, 3, 4 });

    nums = new int[] { -1, -100, 3, 99 };
    solution.rotate(nums, 2);
    assert Arrays.equals(nums, new int[] { 3, 99, -1, -100 });
  }
}