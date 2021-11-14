# 1286. Iterator for Combination

- Difficulty: Medium
- Topics: String, Backtracking, Design, Iterator
- Link: https://leetcode.com/problems/iterator-for-combination/

## Description

Design the `CombinationIterator` class:

- `CombinationIterator(string characters, int combinationLength)` Initializes the object with a string `characters` of **sorted distinct** lowercase English letters and a number `combinationLength` as arguments.
- `next()` Returns the next combination of length `combinationLength` in **lexicographical order**.
- `hasNext()` Returns `true` if and only if there exists a next combination.

**Example 1:**

```
Input
["CombinationIterator", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
[["abc", 2], [], [], [], [], [], []]
Output
[null, "ab", true, "ac", true, "bc", false]

Explanation
CombinationIterator itr = new CombinationIterator("abc", 2);
itr.next();    // return "ab"
itr.hasNext(); // return True
itr.next();    // return "ac"
itr.hasNext(); // return True
itr.next();    // return "bc"
itr.hasNext(); // return False
```

**Constraints:**

- `1 <= combinationLength <= characters.length <= 15`
- All the characters of `characters` are **unique**.
- At most `10^4` calls will be made to `next` and `hasNext`.
- It's guaranteed that all calls of the function `next` are valid.

## Solution

next 和 hasNext 方法的实现则很简单，本题的重点是实现生成 Combinations 的方法。

### Depth-First Search

- DFS 搜索所有可能的 Combinations (本题按照词典顺序)
- 若满足搜索条件 combination.size() 等于 `combinationLenght`，则加入用于保存 Combinations 的序列。

```shell
# characters = "abcd", length = 3
        a
    /   |   \
   b    c    d
  / \   |
 c   d  d
 |
 d
        b
      /   \
     c     d
     |
     d

        c
        |
        d

        d

# combinations = {abc, abd, acd, bcd}
```
