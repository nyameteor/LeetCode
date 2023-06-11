# 1146. Snapshot Array

- Difficulty: Medium
- Topics: Array, Hash Table, Binary Search, Design
- Link: https://leetcode.com/problems/snapshot-array/

## Description

Implement a SnapshotArray that supports the following interface:

- `SnapshotArray(int length)` initializes an array-like data structure with the given length. **Initially, each element equals 0**.
- `void set(index, val)` sets the element at the given `index` to be equal to `val`.
- `int snap()` takes a snapshot of the array and returns the `snap_id`: the total number of times we called `snap()` minus `1`.
- `int get(index, snap_id)` returns the value at the given `index`, at the time we took the snapshot with the given `snap_id`

**Example 1:**

```
**Input:** ["SnapshotArray","set","snap","set","get"]
[[3],[0,5],[],[0,6],[0,0]]
**Output:** [null,null,0,null,5]
**Explanation:**
SnapshotArray snapshotArr = new SnapshotArray(3); // set the length to be 3
snapshotArr.set(0,5);  // Set array[0] = 5
snapshotArr.snap();  // Take a snapshot, return snap_id = 0
snapshotArr.set(0,6);
snapshotArr.get(0,0);  // Get the value of array[0] with snap_id = 0, return 5
```

**Constraints:**

- `1 <= length <= 5 * 10^4`
- `0 <= index < length`
- `0 <= val <= 10^9`
- `0 <= snap_id <` (the total number of times we call `snap()`)
- At most `5 * 10^4` calls will be made to `set`, `snap`, and `get`.

## Solution

References: https://leetcode.com/problems/snapshot-array/editorial/

Store `snapId, value` in a `snapshot` struct.

Store snapshots for every nums[i] in `snapshots[][]`.

- Set(index, val): Add a snaptshot to `snapshots[index]`
- Get(index, snap_id):
  - Use binary search to find the index of rightmost snapshot for the target `snap_id`, denote result as `snap_index`.
  - Return `snapshots[index][snap_index].value`.
