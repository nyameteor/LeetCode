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
- At most `2 * 10^5` calls will be made to `get` and `put`.

## Solution

The LRU (Least Recently Used) cache can be efficiently implemented using a combination of a **hash map** and a **doubly linked list**.

- **Hash Map**: Stores key-value pairs and provides O(1) access to nodes in the doubly linked list.
- **Doubly Linked List**: Tracks usage order, with the most recently used node at the front and the least recently used at the back. It allows O(1) insertion and removal of nodes.

### Operations

- **`get(key)`**: Returns the value if the key exists, moving the node to the front (most recently used). Returns `-1` if not found.
- **`put(key, value)`**: Updates the value or adds a new key-value pair. Evicts the least recently used node if the cache exceeds its capacity.

### Complexity

- **Time**: O(1) for both `get` and `put`.
- **Space**: O(capacity).

### References

Doubly linked list: https://en.wikipedia.org/wiki/Doubly_linked_list
