# 380. Insert Delete GetRandom O(1)

- Difficulty: Medium
- Topics: Array, Hash Table, Math, Design, Randomized
- Link: https://leetcode.com/problems/insert-delete-getrandom-o1/

## Description

Implement the `RandomizedSet` class:

- `RandomizedSet()` Initializes the `RandomizedSet` object.
- `bool insert(int val)` Inserts an item `val` into the set if not present. Returns `true` if the item was not present, `false` otherwise.
- `bool remove(int val)` Removes an item `val` from the set if present. Returns `true` if the item was present, `false` otherwise.
- `int getRandom()` Returns a random element from the
  current set of elements (it's guaranteed that at least one element
  exists when this method is called). Each element must have the **same probability** of being returned.

You must implement the functions of the class such that each function works in **average** `O(1)` time complexity.

**Example 1:**

```
**Input**
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
**Output**
[null, true, false, true, 2, true, false, 2]

**Explanation**
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
randomizedSet.insert(2); // 2 was already in the set, so return false.
randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.
```

**Constraints:**

- `-2^31 <= val <= 2^31 - 1`
- At most `2 *10^5` calls will be made to `insert`, `remove`, and `getRandom`.
- There will be **at least one** element in the data structure when `getRandom` is called.

## Solution

To get a random element, we can store all values in an array `vals`.

To get whether an element is present, we can store indexes of values in `vals` in an map `indices`.

The key is how to remove an element in a list in `O(1)`. There is only one way, which is remove the last element, so we can:

- Get the index of the value that needs to be removed `iVal`, and the index of last element `iLast`.
- Swap the value of `vals[iVal]` and `vals[iLast]`.
- Remove the last element of `vals`.
- Don't forget to update `indices`. update `indices[vals[iLast]]` to `iVal`, and delete `indices[iVal]`.

References:

https://leetcode.com/problems/insert-delete-getrandom-o1/solutions/2858286/python3-hashmap-list-got-this-in-google-phone-interview/
