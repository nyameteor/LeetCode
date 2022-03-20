#include <cassert>
#include <stack>
#include <unordered_map>
using namespace std;

class FreqStack {
  private:
    unordered_map<int, int> freq;
    unordered_map<int, stack<int>> group;
    int max_freq;

  public:
    FreqStack() { max_freq = 0; }

    void push(int val) {
        freq[val]++;
        int f = freq[val];
        if (f > max_freq) {
            max_freq = f;
        }

        group[f].push(val);
    }

    int pop() {
        int x = group[max_freq].top();
        group[max_freq].pop();

        freq[x]--;
        if (size(group[max_freq]) == 0) {
            max_freq--;
        }

        return x;
    }
};

int main() {

    FreqStack freqStack;

    freqStack.push(5); // The stack is [5]
    freqStack.push(7); // The stack is [5,7]
    freqStack.push(5); // The stack is [5,7,5]
    freqStack.push(7); // The stack is [5,7,5,7]
    freqStack.push(4); // The stack is [5,7,5,7,4]
    freqStack.push(5); // The stack is [5,7,5,7,4,5]

    // return 5, as 5 is the most frequent. The stack becomes [5,7,5,7,4].
    assert(freqStack.pop() == 5);

    // return 7, as 5 and 7 is the most frequent, but 7 is closest to the top.
    // The stack becomes [5,7,5,4].
    assert(freqStack.pop() == 7);

    // return 5, as 5 is the most frequent. The stack becomes [5,7,4].
    assert(freqStack.pop() == 5);

    // return 4, as 4, 5 and 7 is the most frequent, but 4 is closest to the
    // top. The stack becomes [5,7].
    assert(freqStack.pop() == 4);

    return 0;
}