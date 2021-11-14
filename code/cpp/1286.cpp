#include <iostream>
#include <queue>
#include <vector>
using namespace std;

/**
 * Generate combinations with DFS
 */
class CombinationIterator {
  private:
    queue<string> combinations;

    // generate combinations
    void dfs(string path, string chars, int len) {
        if (path.size() == len) {
            combinations.push(path);
            return;
        }
        if (chars == "") {
            return;
        }
        // chars [i]: current node
        // chars from i+1 to the end: branches to traverse
        for (int i = 0; i < chars.size(); i++) {
            dfs(path + chars[i], chars.substr(i + 1), len);
        }
    }

  public:
    CombinationIterator(string characters, int combinationLength) {
        dfs("", characters, combinationLength);
    }

    string next() {
        string s = combinations.front();
        combinations.pop();
        return s;
    }

    bool hasNext() { return !combinations.empty(); }
};

int main() {
    CombinationIterator *ci = new CombinationIterator("abc", 2);
    cout << ci->next() << endl;    // "ab"
    cout << ci->hasNext() << endl; // true
    cout << ci->next() << endl;    // "ac"
    cout << ci->hasNext() << endl; // true
    cout << ci->next() << endl;    // "bc"
    cout << ci->hasNext() << endl; // false

    return 0;
}