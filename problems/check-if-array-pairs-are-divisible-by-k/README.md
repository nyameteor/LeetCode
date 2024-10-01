# 1497. Check If Array Pairs Are Divisible by k

- Difficulty: Medium
- Topics: Array, Hash Table, Counting
- Link: https://leetcode.com/problems/check-if-array-pairs-are-divisible-by-k/

## Description

Given an array of integers `arr` of even length `n` and an integer `k`.

We want to divide the array into exactly `n / 2` pairs such that the sum of each pair is divisible by `k`.

Return `true` *If you can find a way to do that or* `false` *otherwise*.

**Example 1:**

```
**Input:** arr = [1,2,3,4,5,10,6,7,8,9], k = 5
**Output:** true
**Explanation:** Pairs are (1,9),(2,8),(3,7),(4,6) and (5,10).
```

**Example 2:**

```
**Input:** arr = [1,2,3,4,5,6], k = 7
**Output:** true
**Explanation:** Pairs are (1,6),(2,5) and(3,4).
```

**Example 3:**

```
**Input:** arr = [1,2,3,4,5,6], k = 10
**Output:** false
**Explanation:** You can try all possible pairs to see that there is no way to divide arr into 3 pairs each with sum divisible by 10.
```

**Constraints:**

- `arr.length == n`
- `1 <= n <= 10^5`
- `n` is even.
- `-10^9 <= arr[i] <= 10^9`
- `1 <= k <= 10^5`

## Solution

We can use an array `freq` of size `k` to store the frequencies of the remainders for each element in `arr`.

To determine if the array pairs are divisible by `k`, we need to check two conditions:

- `freq[0]` must be even.
- For each `i` in `[1, k-1]`, `freq[i]` should equal `freq[k-i]`.

Since the elements in `arr` may be negative integers, we need to apply a positive modulo.
As explained in [this Stack Overflow answer](https://stackoverflow.com/a/58118871), we can implement the positive modulo function like this to keep the code readable:

```go
func positive_modulo(n, k int) int {
	r := n % k
	if r < 0 {
		r += k
	}
	return r
}
```
