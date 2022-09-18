# 210. Course Schedule II

- Difficulty: Medium
- Topics: Depth-First Search, Breadth-First Search, Graph, Topological Sort
- Link: https://leetcode.com/problems/course-schedule-ii/

## Description

There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that you **must** take course `bi` first if you want to take course `ai`.

- For example, the pair `[0, 1]`, indicates that to take course `0` you have to first take course `1`.

Return _the ordering of courses you should take to finish all courses_. If there are many valid answers, return **any** of them. If it is impossible to finish all courses, return **an empty array**.

**Example 1:**

```
Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].
```

**Example 2:**

```
Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].
```

**Example 3:**

```
Input: numCourses = 1, prerequisites = []
Output: [0]
```

**Constraints:**

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= numCourses * (numCourses - 1)`
- `prerequisites[i].length == 2`
- `0 <= ai, bi < numCourses`
- `ai != bi`
- All the pairs `[ai, bi]` are **distinct**.

## Solution

### Topological Sort

拓扑排序。给定一组有向图，将所有的顶点排序，使得所有的有向边均从排在前面的元素指向排在后面的元素（或者说明无法做到这一点）。

一种在排序时检测有向环的方式是，使用一个数组 indgree 记录排序过程中所有顶点的入度。

- 遍历图 G 中所有的顶点，设当前顶点为 v
  - 若 v 的入度为 0，则可以对该顶点进行广度/深度优先搜索：
    - 将 v 加入到结果序列 s 中
    - 将 v 的入度赋值为 -1 （将 v 标记为已访问）
    - 遍历 v 指向的所有顶点，设当前顶点为 w
      - 将 w 的入度 - 1 （w 少了一个前置的顶点 v，故入度需要减 1）
      - 若 w 的入度为 0，则将 w 加入到待遍历的队列/栈中
  - 若 v 的入度不为 0，跳过
- 检查结果序列 s
  - 若 s 包含所有顶点，则返回结果序列
  - 若 s 不包含所有顶点，则说明图中存在环，返回空
