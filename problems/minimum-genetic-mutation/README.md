# 433. Minimum Genetic Mutation

- Difficulty: Medium
- Topics: Hash Table, String, Breadth-First Search
- Link: https://leetcode.com/problems/minimum-genetic-mutation/

## Description

A gene string can be represented by an 8-character long string, with choices from `'A'`, `'C'`, `'G'`, and `'T'`.

Suppose we need to investigate a mutation from a gene string `startGene` to a gene string `endGene` where one mutation is defined as one single character changed in the gene string.

- For example, `"AACCGGTT" --> "AACCGGTA"` is one mutation.

There is also a gene bank `bank` that records all the valid gene mutations. A gene must be in `bank` to make it a valid gene string.

Given the two gene strings `startGene` and `endGene` and the gene bank `bank`, return *the minimum number of mutations needed to mutate from* `startGene` *to* `endGene`. If there is no such a mutation, return `-1`.

Note that the starting point is assumed to be valid, so it might not be included in the bank.

**Example 1:**

```
Input: startGene = "AACCGGTT", endGene = "AACCGGTA", bank = ["AACCGGTA"]
Output: 1
```

**Example 2:**

```
Input: startGene = "AACCGGTT", endGene = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
Output: 2
```

**Constraints:**

- `0 <= bank.length <= 10`
- `startGene.length == endGene.length == bank[i].length == 8`
- `startGene`, `endGene`, and `bank[i]` consist of only the characters `['A', 'C', 'G', 'T']`.

## Solution

### Approach: Breadth-First Search

1. **Graph Representation**:
   - Treat each gene as a node. Connect nodes if they differ by exactly one character (valid mutation).

2. **BFS Traversal**:
   - Start from `startGene`, explore all valid mutations using BFS.
   - Track visited nodes and return the BFS level when `endGene` is reached.

3. **Termination**:
   - If `endGene` is found, return the number of mutations. Otherwise, return `-1`.

**Complexity**:

- **Time**: `O(n^2)`, for graph construction and BFS traversal.
- **Space**: `O(n^2)`, for the graph and BFS queue.
