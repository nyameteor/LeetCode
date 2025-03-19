# 48. Rotate Image

- Difficulty: Medium
- Topics: Array, Math, Matrix
- Link: https://leetcode.com/problems/rotate-image/

## Description

You are given an `n x n` 2D `matrix` representing an image, rotate the image by **90** degrees (clockwise).

You have to rotate the image [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm), which means you have to modify the input 2D matrix directly. **DO NOT** allocate another 2D matrix and do the rotation.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/08/28/mat1.jpg)

```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2020/08/28/mat2.jpg)

```
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
```

**Constraints:**

- `n == matrix.length == matrix[i].length`
- `1 <= n <= 20`
- `-1000 <= matrix[i][j] <= 1000`

## Solution

### Approach: Transpose + Reverse

To rotate the matrix **90 degrees clockwise** in-place:

1. **Transpose** the matrix by swapping rows and columns.
2. **Reverse** the order of elements in each row.

This approach modifies the matrix directly without using extra space, making it memory efficient.

### References

Many image processing libraries (e.g., **PIL**) use a similar technique for rotating images:

- [PIL Image.rotate Implementation](https://github.com/python-pillow/Pillow/blob/e1bf0f647f769fd42ed7e73c0d73672896204e2a/src/PIL/Image.py#L2376)
