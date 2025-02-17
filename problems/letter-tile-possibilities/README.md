# 1079. Letter Tile Possibilities

- Difficulty: Medium
- Topics: Hash Table, String, Backtracking, Counting
- Link: https://leetcode.com/problems/letter-tile-possibilities/

## Description

You have `n`  `tiles`, where each tile has one letter `tiles[i]` printed on it.

Return *the number of possible non-empty sequences of letters* you can make using the letters printed on those `tiles`.

**Example 1:**

```
Input: tiles = "AAB"
Output: 8
Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".

```

**Example 2:**

```
Input: tiles = "AAABBC"
Output: 188

```

**Example 3:**

```
Input: tiles = "V"
Output: 1

```

**Constraints:**

- `1 <= tiles.length <= 7`
- `tiles` consists of uppercase English letters.

## Solution

### **Approach: Backtracking with Deduplication**

We use **backtracking** to explore all possible sequences by choosing each letter one by one while marking used letters to prevent repetition.

1. Use a set to store unique sequences.  

2. Define a recursive `dfs` function that explores all possible sequences by selecting letters one by one.  

3. For each letter in `tiles`:  
   - If it's not used yet, mark it as used.  
   - Add it to the current sequence (`path`).  
   - Recursively call `dfs` to continue building sequences.  
   - Backtrack by removing the letter and marking it as unused.  

4. If `path` is not empty, add it to the set before moving to the next step.  

5. Finally, return the number of unique sequences stored in the set.
