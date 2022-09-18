#include <cstdio>
#include <iostream>
#include <map>
#include <queue>
#include <vector>

using namespace std;

/**
 * Heap
 */
class Solution {
public:
    vector<int> topKFrequent(vector<int> &nums, int k) {
        map<int, int> count;
        for (auto num : nums) {
            count[num]++;
        }

        // initialize a heap with most frequent elements at the top
        auto cmp = [](const auto &p1, const auto &p2) {
            return p1.second > p2.second;
        };
        priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(cmp)>
            pq(cmp);

        // keep k elements in the heap
        for (auto p : count) {
            pq.push(p);
        }
        while (size(pq) > k) {
            pq.pop();
        }

        // build an result array
        vector<int> top(k);
        for (int i = k - 1; i >= 0; i--) {
            top[i] = pq.top().first;
            pq.pop();
        }

        return top;
    }
};

/**
 * Bucket Sort
 */
class Solution2 {
public:
    vector<int> topKFrequent(vector<int> &nums, int k) {

        map<int, int> cnts;
        int max_cnt = 0;

        // store count of every element to Map
        // and get the max count
        // (max count is the maximum key value for buckets)
        for (int i = 0; i < nums.size(); i++) {
            cnts[nums[i]]++;
            max_cnt = max(max_cnt, cnts[nums[i]]);
        }

        // use buckets to store element by count
        vector<vector<int>> buckets(max_cnt + 1);
        for (auto it = cnts.begin(); it != cnts.end(); it++) {
            buckets[it->second].push_back(it->first);
        }

        // get k most frequent elements
        vector<int> res;
        for (int i = buckets.size() - 1; i > 0; i--) {
            // insert elements in buckets[i] to result
            for (int j = 0; j < buckets[i].size(); j++) {
                res.push_back(buckets[i][j]);
            }
            if (res.size() == k) {
                break;
            }
        }
        return res;
    }
};

void printVector(vector<int> &num) {
    for (int i = 0; i < num.size(); i++) {
        cout << num[i] << " ";
    }
    cout << endl;
}

int main() {
    Solution solution;
    vector<int> nums;
    vector<int> res;

    nums = {1, 1, 1, 2, 2, 3};
    res = solution.topKFrequent(nums, 2);
    printVector(res);

    nums = {1, 1, 1, 2, 2, 3, 4, 4, 5};
    res = solution.topKFrequent(nums, 3);
    printVector(res);

    nums = {1};
    res = solution.topKFrequent(nums, 1);
    printVector(res);
}