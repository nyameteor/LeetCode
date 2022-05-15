#include <cmath>
#include <iostream>
#include <stack>
#include <vector>
using namespace std;

class Solution {
public:
    int largestRectangleArea(vector<int> &heights) {
        int maxArea = 0;

        // stack for height
        stack<int> stH;
        // stack for the left cloest index that greater or equal
        // than correspond height
        stack<int> stX;
        int x;
        int h;
        int curArea;
        for (int i = 0; i < size(heights); i++) {
            x = i;
            h = heights[i];
            // Pop off any elements greater than incoming to keep stack montonic
            while (!stH.empty() && h <= stH.top()) {
                x = stX.top();
                curArea = stH.top() * (i - stX.top());
                maxArea = max(maxArea, curArea);
                stX.pop();
                stH.pop();
            }
            stX.push(x);
            stH.push(h);
        }

        // calculate area of rectangles left on stack
        while (!stH.empty()) {
            curArea = stH.top() * (size(heights) - stX.top());
            maxArea = max(maxArea, curArea);
            stX.pop();
            stH.pop();
        }

        return maxArea;
    }
};

int main() {
    Solution solution;
    vector<int> heights;

    heights = {2, 1, 5, 6, 2, 3};
    cout << solution.largestRectangleArea(heights) << endl;

    heights = {2, 4};
    cout << solution.largestRectangleArea(heights) << endl;

    return 0;
}