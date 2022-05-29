# 1584. Min Cost to Connect All Points

- Difficulty: Medium
- Topics: Array, Union Find, Minimum Spanning Tree
- Link: https://leetcode.com/problems/min-cost-to-connect-all-points/

## Description

You are given an array `points` representing integer coordinates of some points on a 2D-plane, where `points[i] = [xi, yi]`.

The cost of connecting two points `[xi, yi]` and `[xj, yj]` is the **manhattan distance** between them: `|xi - xj| + |yi - yj|`, where `|val|` denotes the absolute value of `val`.

Return _the minimum cost to make all points connected._ All points are connected if there is **exactly one** simple path between any two points.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/08/26/d.png)

```
Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
Output: 20
Explanation:

We can connect the points as shown above to get the minimum cost of 20.
Notice that there is a unique path between every pair of points.
```

**Example 2:**

```
Input: points = [[3,12],[-2,5],[-4,1]]
Output: 18
```

**Constraints:**

- `1 <= points.length <= 1000`
- `-106 <= xi, yi <= 106`
- All pairs `(xi, yi)` are distinct.

## Solution

### Prim's Alogrithm

At any point in the algorithm:

- Find a new vertex to add to the tree by choosing the edge $(u, v)$ such that the cost of $(u,v)$ is the smallest among all edges, where $u$ is in the tree and $v$ is not.
- After a vertex $v$ is selected, for each $unknown \space w$ adjacent to $v$, $d_w = min(d_w, c_{w,v})$

Refer:

- https://en.wikipedia.org/wiki/Prim%27s_algorithm
- https://leetcode.com/problems/min-cost-to-connect-all-points/solution/
