#include <iostream>
#include <queue>
#include <vector>
using namespace std;

// Backtracking
class Solution {
public:
    vector<string> generateParenthesis(int n) {
        vector<string> res;
        dfs(res, "", n, 0, 0);
        return res;
    }

    void dfs(vector<string> &paths, string path, int n, int open, int close) {
        if (open + close == 2 * n) {
            paths.push_back(path);
            return;
        }
        if (open < n) {
            dfs(paths, path + '(', n, open + 1, close);
        }
        if (open > close) {
            dfs(paths, path + ')', n, open, close + 1);
        }
    }
};

// Backtracking with strict definition
class Solution2 {
public:
    vector<string> generateParenthesis(int n) {
        vector<string> res;
        string path;
        dfs(res, path, n, n);
        return res;
    }

    /**
     * @param paths return result
     * @param path  path from root to any node, only one copy in whole search
     * @param open  number of open parentheses available
     * @param close number of close parentheses available
     */
    void dfs(vector<string> &paths, string &path, int open, int close) {
        if (open == 0 && close == 0) {
            paths.push_back(path);
            return;
        }

        /**
         * Tree Pruning
         * in a valid path, number of open parentheses can't be more than
         * close parentheses
         */
        if (open > close) {
            return;
        }

        if (open > 0) {
            path.push_back('(');
            dfs(paths, path, open - 1, close);
            path.pop_back();
        }
        if (close > 0) {
            path.push_back(')');
            dfs(paths, path, open, close - 1);
            path.pop_back();
        }
    }
};

void printResult(vector<string> res) {
    for (auto x : res) {
        cout << x << ' ';
    }
    cout << endl;
}

int main() {
    Solution solution;
    Solution2 solution2;

    printResult(solution.generateParenthesis(1));
    printResult(solution2.generateParenthesis(1));
    printResult(solution.generateParenthesis(3));
    printResult(solution2.generateParenthesis(3));
}