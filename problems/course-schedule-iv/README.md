# 1462. Course Schedule IV

- Difficulty: Medium
- Topics: Depth-First Search, Breadth-First Search, Graph, Topological Sort
- Link: https://leetcode.com/problems/course-schedule-iv/

## Description

There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that you **must** take course `ai` first if you want to take course `bi`.

- For example, the pair `[0, 1]` indicates that you have to take course `0` before you can take course `1`.

Prerequisites can also be **indirect**. If course `a` is a prerequisite of course `b`, and course `b` is a prerequisite of course `c`, then course `a` is a prerequisite of course `c`.

You are also given an array `queries` where `queries[j] = [uj, vj]`. For the `jth` query, you should answer whether course `uj` is a prerequisite of course `vj` or not.

Return *a boolean array* `answer`*, where* `answer[j]` *is the answer to the* `jth` *query.*

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/05/01/courses4-1-graph.jpg)

```
**Input:** numCourses = 2, prerequisites = [[1,0]], queries = [[0,1],[1,0]]
**Output:** [false,true]
**Explanation:** The pair [1, 0] indicates that you have to take course 1 before you can take course 0.
Course 0 is not a prerequisite of course 1, but the opposite is true.

```

**Example 2:**

```
**Input:** numCourses = 2, prerequisites = [], queries = [[1,0],[0,1]]
**Output:** [false,false]
**Explanation:** There are no prerequisites, and each course is independent.

```

**Example 3:**

![](https://assets.leetcode.com/uploads/2021/05/01/courses4-3-graph.jpg)

```
**Input:** numCourses = 3, prerequisites = [[1,2],[1,0],[2,0]], queries = [[1,0],[1,2]]
**Output:** [true,true]

```

**Constraints:**

- `2 <= numCourses <= 100`
- `0 <= prerequisites.length <= (numCourses * (numCourses - 1) / 2)`
- `prerequisites[i].length == 2`
- `0 <= ai, bi <= numCourses - 1`
- `ai != bi`
- All the pairs `[ai, bi]` are **unique**.
- The prerequisites graph has no cycles.
- `1 <= queries.length <= 104`
- `0 <= ui, vi <= numCourses - 1`
- `ui != vi`

## Solution

This problem is solved by representing courses and their prerequisites as a directed graph and using BFS to determine reachability:

1. **Graph Construction**:
   - Build a graph where each edge from course `a` to course `b` means `a` must be taken before `b`.
   - Track the number of prerequisites for each course with the `inDegrees` array.

2. **Reachability Matrix (`isReachable`)**:
   - `isReachable[u][v]` indicates whether course `u` is a prerequisite of course `v`. Initially, it's `false` for all pairs.
   - During BFS, update the matrix: if `u` can reach `v` (directly or indirectly), set `isReachable[u][v] = true`.

3. **BFS Propagation**:
   - Start with courses that have no prerequisites and propagate their reachability to neighbors, updating the matrix as we go.

4. **Query Answering**:
   - After BFS, simply check `isReachable[u][v]` for each query to determine if course `u` is a prerequisite of course `v`.
