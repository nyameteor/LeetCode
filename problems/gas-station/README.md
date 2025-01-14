# 134. Gas Station

- Difficulty: Medium
- Topics: Array, Greedy
- Link: https://leetcode.com/problems/gas-station/

## Description

There are `n` gas stations along a circular route, where the amount of gas at the `ith` station is `gas[i]`.

You have a car with an unlimited gas tank and it costs `cost[i]` of gas to travel from the `ith` station to its next `(i + 1)th` station. You begin the journey with an empty tank at one of the gas stations.

Given two integer arrays `gas` and `cost`, return *the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return* `-1`. If there exists a solution, it is **guaranteed** to be **unique**.

**Example 1:**

```
**Input:** gas = [1,2,3,4,5], cost = [3,4,5,1,2]
**Output:** 3
**Explanation:**
Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
Therefore, return 3 as the starting index.

```

**Example 2:**

```
**Input:** gas = [2,3,4], cost = [3,4,3]
**Output:** -1
**Explanation:**
You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 0. Your tank = 4 - 3 + 2 = 3
Travel to station 1. Your tank = 3 - 3 + 3 = 3
You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
Therefore, you can't travel around the circuit once no matter where you start.

```

**Constraints:**

- `n == gas.length == cost.length`
- `1 <= n <= 10^5`
- `0 <= gas[i], cost[i] <= 10^4`

## Solution

This solution uses a **prefix sum approach** to find the starting station for completing the circuit:

1. **Track Gas Balance**:
   - Calculate the cumulative gas balance (`totalGas`) by adding `gas[i]` and subtracting `cost[i]` as you traverse the stations.

2. **Find Critical Point**:
   - Track the minimum gas balance (`minGas`) and its index (`minIdx`) during the traversal. This identifies the point where the gas balance is at its lowest.

3. **Check Feasibility**:
   - If `totalGas` (final cumulative gas balance) is negative, return `-1` as it’s impossible to complete the circuit.

4. **Determine Starting Station**:
   - If `totalGas` is non-negative, the starting station is `(minIdx + 1) % n`, ensuring a valid circuit.

**Why This Works:**

- The **cumulative gas balance** (`totalGas`) represents the net gas available after traveling through all stations.
- The **minimum gas balance** (`minGas`) indicates the point where the gas tank would be at its lowest if starting from station `0`.
- Starting at the station immediately after `minIdx` ensures that the gas balance remains non-negative throughout the circuit, making it a valid starting point.

**Complexity:**

- **Time Complexity:** \(O(n)\) - Single traversal of the arrays.
- **Space Complexity:** \(O(1)\) - Constant extra space.
