# 2381. Shifting Letters II

- Difficulty: Medium
- Topics: Array, String, Prefix Sum
- Link: https://leetcode.com/problems/shifting-letters-ii/

## Description

You are given a string `s` of lowercase English letters and a 2D integer array `shifts` where `shifts[i] = [starti, endi, directioni]`. For every `i`, **shift** the characters in `s` from the index `starti` to the index `endi` (**inclusive**) forward if `directioni = 1`, or shift the characters backward if `directioni = 0`.

Shifting a character **forward** means replacing it with the **next** letter in the alphabet (wrapping around so that `'z'` becomes `'a'`). Similarly, shifting a character **backward** means replacing it with the **previous** letter in the alphabet (wrapping around so that `'a'` becomes `'z'`).

Return *the final string after all such shifts to* `s` *are applied*.

**Example 1:**

```
**Input:** s = "abc", shifts = [[0,1,0],[1,2,1],[0,2,1]]
**Output:** "ace"
**Explanation:** Firstly, shift the characters from index 0 to index 1 backward. Now s = "zac".
Secondly, shift the characters from index 1 to index 2 forward. Now s = "zbd".
Finally, shift the characters from index 0 to index 2 forward. Now s = "ace".
```

**Example 2:**

```
**Input:** s = "dztz", shifts = [[0,0,0],[1,1,1]]
**Output:** "catz"
**Explanation:** Firstly, shift the characters from index 0 to index 0 backward. Now s = "cztz".
Finally, shift the characters from index 1 to index 1 forward. Now s = "catz".

```

**Constraints:**

- `1 <= s.length, shifts.length <= 5 * 10^4`
- `shifts[i].length == 3`
- `0 <= starti <= endi < s.length`
- `0 <= directioni <= 1`
- `s` consists of lowercase English letters.

## Solution

### Approach: Difference Array

The **Difference Array** is a powerful technique for efficiently handling range updates by recording changes at the boundaries. To get the final result, compute the prefix sum of the difference array.

#### Steps

1. **Initialize the Difference Array**:  
   `diff[i] = arr[i] - arr[i-1]` for `i > 0`, and `diff[0] = arr[0]`.

2. **Apply Range Updates**:  
   To add `x` to the range `[l, r]`:  
   - Increment `diff[l]` by `x`.  
   - Decrement `diff[r+1]` by `x` (if within bounds).

3. **Reconstruct the Array**:  
   Compute the prefix sum of `diff` to get the updated array.

#### Example

Given `arr = [2, 4, 5, 7]`:

1. **Difference Array**:  
   `diff = [2, 2, 1, 2, 0]`

2. **Apply Update** (`Add 3 to range [1, 2]`):  
   `diff = [2, 5, 1, -1, 0]`

3. **Reconstruct Array**:  
   `arr = [2, 7, 8, 7]`

#### Advantages

- **Efficient**: Range updates in `O(1)`, array reconstruction in `O(n)`.
- **Simpler**: Avoids iterating over entire ranges for each update.

This technique is ideal for problems with frequent range updates, such as **Shifting Letters**.
