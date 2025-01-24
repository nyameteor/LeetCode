# 207. Course Schedule

- Difficulty: Medium
- Topics: Depth-First Search, Breadth-First Search, Graph, Topological Sort
- Link: https://leetcode.com/problems/course-schedule/

## Description

There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that you **must** take course `bi` first if you want to take course `ai`.

- For example, the pair `[0, 1]`, indicates that to take course `0` you have to first take course `1`.

Return `true` if you can finish all courses. Otherwise, return `false`.

**Example 1:**

```
**Input:** numCourses = 2, prerequisites = [[1,0]]
**Output:** true
**Explanation:** There are a total of 2 courses to take. 
To take course 1 you should have finished course 0. So it is possible.

```

**Example 2:**

```
**Input:** numCourses = 2, prerequisites = [[1,0],[0,1]]
**Output:** false
**Explanation:** There are a total of 2 courses to take. 
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.

```

**Constraints:**

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= 5000`
- `prerequisites[i].length == 2`
- `0 <= ai, bi < numCourses`
- All the pairs prerequisites[i] are **unique**.

## Solution

The problem can be modeled as a **directed graph**, where courses are nodes and prerequisites are edges. The goal is to check for a **cycle**, as a cycle prevents completing all courses.

### Approach: DFS

1. **Graph Representation**:
   - Use an adjacency list to represent course dependencies.

2. **Cycle Detection**:
   - Perform a **DFS** using two markers:
     - `visited`: Fully processed nodes.
     - `inStack`: Nodes in the current recursion stack.
   - A cycle is detected if a node is revisited while still in the current stack.

3. **Result**:
   - Return `false` if a cycle is detected, otherwise `true`.

### References

- https://en.wikipedia.org/wiki/Topological_sorting#Algorithms
