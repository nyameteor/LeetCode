# 295. Find Median from Data Stream

- Difficulty: Hard
- Topics: Two Pointers, Design, Sorting, Heap (Priority Queue), Data Stream
- Link: https://leetcode.com/problems/find-median-from-data-stream/

## Description

The **median**
is the middle value in an ordered integer list. If the size of the list
is even, there is no middle value, and the median is the mean of the
two middle values.

- For example, for `arr = [2,3,4]`, the median is `3`.
- For example, for `arr = [2,3]`, the median is `(2 + 3) / 2 = 2.5`.

Implement the MedianFinder class:

- `MedianFinder()` initializes the `MedianFinder` object.
- `void addNum(int num)` adds the integer `num` from the data stream to the data structure.
- `double findMedian()` returns the median of all elements so far. Answers within `10-5` of the actual answer will be accepted.

**Example 1:**

```
Input
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
Output
[null, null, null, 1.5, null, 2.0]

Explanation
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0
```

**Constraints:**

- `-10^5 <= num <= 10^5`
- There will be at least one element in the data structure before calling `findMedian`.
- At most `5 * 10^4` calls will be made to `addNum` and `findMedian`.

**Follow up:**

- If all integer numbers from the stream are in the range `[0, 100]`, how would you optimize your solution?
- If `99%` of all integer numbers from the stream are in the range `[0, 100]`, how would you optimize your solution?

## Solution

### Approach: Two Heaps (Max Heap & Min Heap)

- Maintain two heaps:
  - `maxHeap` (max heap) stores the **smaller half** of numbers.
  - `minHeap` (min heap) stores the **larger half** of numbers.
- The goal is to keep both heaps balanced in size, or let `maxHeap` have one extra element when odd-sized.
- When adding a number:
  1. Push to `maxHeap`, then move the max element to `minHeap` to maintain order.
  2. If `minHeap` has more elements, move its min element back to `maxHeap`.
- The median is:
  - If both heaps are equal in size: **mean of both tops**.
  - If `maxHeap` has one more element: **top of `maxHeap`**.

**Complexity:**

- Time: `O(log n)` per `addNum()` due to heap operations.
- Space: `O(n)` to store numbers in heaps.

### References

- [Find Median from Running Data Stream](https://www.geeksforgeeks.org/median-of-stream-of-integers-running-integers/)
