import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

class Solution {
  public int[] twoSum(int[] nums, int target) {
    Map<Integer, Integer> map = new HashMap<>();
    int[] res = new int[2];

    for (int i = 0; i < nums.length; i++) {
      if (map.containsKey(target - nums[i])) {
        res[0] = map.get(target - nums[i]);
        res[1] = i;
        return res;
      } else {
        map.put(nums[i], i);
      }
    }
    return res;
  }

  public static void main(String[] args) {
    Solution solution = new Solution();
    assert Arrays.equals(solution.twoSum(new int[] { 2, 7, 11, 15 }, 9), new int[] { 0, 1 });
    assert Arrays.equals(solution.twoSum(new int[] { 3, 2, 4 }, 6), new int[] { 1, 2 });
    assert Arrays.equals(solution.twoSum(new int[] { 3, 3 }, 6), new int[] { 0, 1 });
  }
}