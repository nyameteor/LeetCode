# 146. LRU Cache

- Difficulty: Medium
- Topics: Hash Table, Linked List, Design, Doubly-Linked List
- Link: https://leetcode.com/problems/lru-cache/

## Description

Design a data structure that follows the constraints of a **[Least Recently Used (LRU) cache](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU)**.

Implement the `LRUCache` class:

- `LRUCache(int capacity)` Initialize the LRU cache with **positive** size `capacity`.
- `int get(int key)` Return the value of the `key` if the key exists, otherwise return `-1`.
- `void put(int key, int value)` Update the value of the `key` if the `key` exists. Otherwise, add the `key-value` pair to the cache. If the number of keys exceeds the `capacity` from this operation, **evict** the least recently used key.

The functions `get` and `put` must each run in `O(1)` average time complexity.

**Example 1:**

```
Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4
```

**Constraints:**

- `1 <= capacity <= 3000`
- `0 <= key <= 10^4`
- `0 <= value <= 10^5`
- At most 2` * 10^5` calls will be made to `get` and `put`.

## Solution

经典的模拟题，通过最久未使用算法（LRU）实现缓存替换机制。

### Hash Table + Queue

比较粗暴易懂的方法，优化使用的数据结构后也只快过 5% 的提交。

- 使用一个 Hash Table 保存 key-value 作为缓存
- 使用一个 Queue 保存 key，作为存储 `Recently Used` 的序列，Queue 的队头为 `Least Recently Used`
- 按以下条件调整 Queue 中元素的顺序：
  - 当 LRUCache 的 get 和 put 方法访问到 key 时，更新 LRU（将 key 先删除后插入，使 Queue 中的 key 移动至队尾）
  - 当 LRUCache 的 put 方法未访问到 key 且 capacity 已满时，置换 LRU (Queue 队头出队，后插入 key 到 Queue）
