#include <iostream>
#include <queue>
#include <vector>
using namespace std;

/**
 * Generate combinations with DFS
 *
 * TODO: Optimize
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
        for (int i = 0; i < chars.size(); i++) {
            string branches = chars;
            branches.erase(0, i + 1);
            dfs(path + chars[i], branches, len);
        }
    }

  public:
    CombinationIterator(string characters, int combinationLength) {
        for (int i = 0; i < characters.size(); i++) {
            string path = "";
            string branches = characters;
            branches.erase(0, i + 1);
            dfs(path + characters[i], branches, combinationLength);
        }
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