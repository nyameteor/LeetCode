# 1335. Minimum Difficulty of a Job Schedule

- Difficulty: Hard
- Topics: Array, Dynamic Programming
- Link: https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/

## Description

You want to schedule a list of jobs in `d` days. Jobs are dependent (i.e To work on the `ith` job, you have to finish all the jobs `j` where `0 <= j < i`).

You have to finish **at least** one task every day. The difficulty of a job schedule is the sum of difficulties of each day of the `d` days. The difficulty of a day is the maximum difficulty of a job done on that day.

You are given an integer array `jobDifficulty` and an integer `d`. The difficulty of the `ith` job is `jobDifficulty[i]`.

Return _the minimum difficulty of a job schedule_. If you cannot find a schedule for the jobs return `-1`.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/01/16/untitled.png)

```
Input: jobDifficulty = [6,5,4,3,2,1], d = 2
Output: 7
Explanation: First day you can finish the first 5 jobs, total difficulty = 6.
Second day you can finish the last job, total difficulty = 1.
The difficulty of the schedule = 6 + 1 = 7
```

**Example 2:**

```
Input: jobDifficulty = [9,9,9], d = 4
Output: -1
Explanation: If you finish a job per day you will still have a free day. you cannot find a schedule for the given jobs.
```

**Example 3:**

```
Input: jobDifficulty = [1,1,1], d = 3
Output: 3
Explanation: The schedule is one job per day. total difficulty will be 3.
```

**Constraints:**

- `1 <= jobDifficulty.length <= 300`
- `0 <= jobDifficulty[i] <= 1000`
- `1 <= d <= 10`

## Solution

Similar to [10. Regular Expression Matching][10] and [44. Wildcard Matching][44], We can use 2D DP.

Here is a virtualization for `jobDifficulty = [6, 5, 4, 3, 2, 1]`, `d = 3`:

```text
|   0   |   1   |   2   |   3   |   4   |   5   |
|   6   |   5   |   4   |   3   |   2   |   1   |
|  dp(i,j)  |   dp(i+1, j+1)  |  max(job[i:])   |
|   6   |   5   |               4               |
|   6   |       5       |           3           |
|   6   |           5           |       2       |
|   6   |               5               |   1   |
|       6       |           4           |   1   |
|           6           |       3       |   1   |
|               6               |   2   |   1   |
```

**References**

- [[Go] DP | Unit-test][go-dp]
- [[LeetCode The Hard Way] Explained Line By Line][explain]

[10]: https://leetcode.com/problems/regular-expression-matching/
[44]: https://leetcode.com/problems/wildcard-matching/
[go-dp]: https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/solutions/2709149/go-dp-unit-test/
[explain]: https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/solutions/2708161/leetcode-the-hard-way-explained-line-by-line/
