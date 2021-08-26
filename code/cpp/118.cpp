#include <cstdio>
#include <iostream>
#include <map>
#include <vector>

using namespace std;

/*

Pascal's Triangle

              0
            0   0
          0   1   0
        0   1   1   0
      0   1   2   1   0
    0   1   3   3   1   0
  0   1   4   6   4   1   0
0   0   0   0   0   0   0   0

*/

class Solution {
  public:
    vector<vector<int>> generate(int numRows) {
        vector<vector<int>> ans;

        for (int i = 0; i < numRows; i++) {
            // create current row
            vector<int> row;

            // edge case
            if (i == 0) {
                row = {1};
                ans.push_back(row);
            } else {
                // get last row
                vector<int> prev_row = ans.at(i - 1);

                row.push_back(1);
                for (int j = 1; j < prev_row.size(); j++) {
                    row.push_back(prev_row[j - 1] + prev_row[j]);
                }
                row.push_back(1);
                ans.push_back(row);
            }
        }

        return ans;
    }
};

int main() {
    Solution solution;

    /*
    1
    1 1
    1 2 1
    1 3 3 1
    1 4 6 4 1
    1 5 10 10 5 1
    1 6 15 20 15 6 1
    1 7 21 35 35 21 7 1
    */
    vector<vector<int>> result = solution.generate(8);
    for (int i = 0; i < result.size(); i++) {
        for (int j = 0; j < result[i].size(); j++) {
            cout << result[i][j] << " ";
        }
        cout << endl;
    }
}