import java.util.PriorityQueue;
import java.util.Queue;

class Solution {
    // Min Heap from JDK
    public int findKthLargest(int[] nums, int k) {
        Queue<Integer> pq = new PriorityQueue<>();
        for (int num : nums) {
            if (pq.size() == k) {
                if (pq.peek() < num) {
                    pq.poll();
                    pq.offer(num);
                } else {
                    continue;
                }
            } else {
                pq.offer(num);
            }
        }
        return pq.poll();
    }
}