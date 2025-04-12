# 470. Implement Rand10() Using Rand7()

- Difficulty: Medium
- Topics: Math, Rejection Sampling, Randomized, Probability and Statistics
- Link: https://leetcode.com/problems/implement-rand10-using-rand7/

## Description

Given the **API** `rand7()` that generates a uniform random integer in the range `[1, 7]`, write a function `rand10()` that generates a uniform random integer in the range `[1, 10]`. You can only call the API `rand7()`, and you shouldn't call any other API. Please **do not** use a language's built-in random API.

Each test case will have one **internal** argument `n`, the number of times that your implemented function `rand10()` will be called while testing. Note that this is **not an argument** passed to `rand10()`.

**Follow up:**

- What is the [expected value](https://en.wikipedia.org/wiki/Expected_value) for the number of calls to `rand7()` function?
- Could you minimize the number of calls to `rand7()`?

**Example 1:**

```
Input: n = 1
Output: [2]
```

**Example 2:**

```
Input: n = 2
Output: [2,8]
```

**Example 3:**

```
Input: n = 3
Output: [3,8,10]
```

**Constraints:**

- `1 <= n <= 10^5`

## Solution

### Key Idea

Use **rejection sampling** to generate uniform random numbers in the range [1, 10] by first generating values from a larger range and rejecting values that fall outside the desired range.

### Steps

1. **Generate Larger Range**: Two calls to `rand7()` produce 49 unique outcomes (7 * 7 = 49). These outcomes are uniformly distributed.
2. **Rejection Sampling**: Only outcomes in the range [0, 39] are used, mapping to [1, 10]. Results >= 40 are rejected and retried.
3. **Uniform Distribution**: By rejecting values >= 40, each number in the range [1, 10] has an equal chance of being chosen.

### References

- [Rejection sampling](https://en.wikipedia.org/wiki/Rejection_sampling)
