#include <functional>
#include <iostream>
#include <queue>
#include <unordered_map>
#include <vector>
using namespace std;

/**
 * Heap (Priority Queue)
 *
 * Idea:
 *  1. Count the frequency of each letter using unordered_map `freq`.
 *  2. To sort the frequencies, we use a priority_queue `q` with pairs (we need
 *     to create a custom compare funciton).
 *  3. Constructing the result is simple, just pop each pair from queue and add
 *     that number of chars to `res`.
 *
 * Refer:
 * https://leetcode.com/problems/sort-characters-by-frequency/discuss/1534608/C++-Two-Simple-and-Clean-Solutions-Detailed-Explanation
 * https://en.cppreference.com/w/cpp/container/priority_queue
 */
class Solution {
public:
    string frequencySort(string s) {
        unordered_map<char, int> freq;
        for (auto c : s) {
            freq[c]++;
        }

        auto cmp = [](pair<char, int> left, pair<char, int> right) {
            return left.second < right.second;
        };
        priority_queue<pair<char, int>, vector<pair<char, int>>, decltype(cmp)>
            q(cmp);

        for (auto pair : freq) {
            q.push(pair);
        }

        string res;
        while (!q.empty()) {
            pair<char, int> curr = q.top();
            q.pop();
            // append string: char * count
            res += string(curr.second, curr.first);
        }

        return res;
    }
};

int main() {
    Solution solution;

    cout << solution.frequencySort("tree") << endl;
    cout << solution.frequencySort("cccaaa") << endl;
    cout << solution.frequencySort("Aabb") << endl;

    return 0;
}