# 692. Top K Frequent Words

- Difficulty: Medium
- Topics: Hash Table, String, Trie, Sorting, Heap (Priority Queue), Bucket Sort, Counting
- Link: https://leetcode.com/problems/top-k-frequent-words/description/

## Description

Given an array of strings `words` and an integer `k`, return _the_ `k` _most frequent strings_.

Return the answer **sorted** by **the frequency** from highest to lowest. Sort the words with the same frequency by their **lexicographical order**.

**Example 1:**

```
Input: words = ["i","love","leetcode","i","love","coding"], k = 2
Output: ["i","love"]
Explanation: "i" and "love" are the two most frequent words.
Note that "i" comes before "love" due to a lower alphabetical order.
```

**Example 2:**

```
Input: words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
Output: ["the","is","sunny","day"]
Explanation: "the", "is", "sunny" and "day" are the four most frequent words, with the number of occurrence being 4, 3, 2 and 1 respectively.
```

**Constraints:**

- `1 <= words.length <= 500`
- `1 <= words[i].length <= 10`
- `words[i]` consists of lowercase English letters.
- `k` is in the range `[1, The number of unique words[i]]`

**Follow-up:** Could you solve it in `O(n log(k))` time and `O(n)` extra space?

## Solution

### Heap

**Max-Heap**

Use max-heap approach is straightfoward.

1. Build a max-heap sort by the `count` and `lexicographical order` of the word.
2. Push `n` items to heap.
3. Pop `k` times to get result.

We do `n` times push and `k` times pop with `len(heap)` is approximately equal to `n`.

Time: `O(n log(n))`, Space: `O(n)`

**Min-Heap**

Use min-heap approach is more efficient.

1. Build a min-heap sort by the `count` and `lexicographical order` of the word.
2. Push `n` items to heap, everytime the `len(heap) > k`, pop the top(minimum) item.
3. Pop `k` times to get result, notice to store items in reverse order.

We do `n` times push and `n` times pop with `len(heap)` is approximately equal to `k`.

Time: `O(n log(k))`, Space: `O(n)`

**References**

- [golang solution with hashmap and heap](https://leetcode.com/problems/top-k-frequent-words/solutions/321266/golang-solution-with-hashmap-and-heap/)
