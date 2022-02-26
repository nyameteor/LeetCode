#include <iostream>
#include <string>
#include <vector>
using namespace std;

class Solution {
  public:
    int compareVersion(string version1, string version2) {
        vector<int> v1 = splitAndTrim(version1);
        vector<int> v2 = splitAndTrim(version2);

        for (int i = 0; i < max(size(v1), size(v2)); i++) {
            int x1 = i < size(v1) ? v1[i] : 0;
            int x2 = i < size(v2) ? v2[i] : 0;

            if (x1 < x2) {
                return -1;
            } else if (x1 > x2) {
                return 1;
            }
        }

        return 0;
    }

    /**
     * Split string with `.` and trim leading `0`, then convert to integer
     */
    vector<int> splitAndTrim(string str) {
        vector<int> v;
        for (int i = 0; i < size(str); i++) {
            string part = "";
            bool lead = true;
            while (str[i] != '.' && i < size(str)) {
                if (lead) {
                    if (str[i] == '0') {
                        i++;
                        continue;
                    } else {
                        lead = false;
                    }
                }
                part.push_back(str[i]);
                i++;
            }
            if (part != "") {
                v.push_back(stoi(part));
            } else {
                v.push_back(0);
            }
        }

        return v;
    }
};

int main() {

    Solution solution;
    string version1, version2;

    version1 = "1.01";
    version2 = "1.001";
    cout << solution.compareVersion(version1, version2) << endl;

    version1 = "1.0";
    version2 = "1.0.0";
    cout << solution.compareVersion(version1, version2) << endl;

    version1 = "0.1";
    version2 = "1.1";
    cout << solution.compareVersion(version1, version2) << endl;

    return 0;
}