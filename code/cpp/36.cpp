#include <iostream>
#include <unordered_set>
#include <vector>
using namespace std;

class Solution {
  public:
    bool isValidSudoku(vector<vector<char>> &board) {
        for (int i = 0; i < 9; i = i + 3) {
            for (int j = 0; j < 9; j = j + 3) {
                if (!checkSubBoard(i, j, board)) {
                    return false;
                }
            }
        }

        for (int i = 0; i < 9; i++) {
            if (!checkRow(i, board)) {
                return false;
            }
        }

        for (int j = 0; j < 9; j++) {
            if (!checkColumn(j, board)) {
                return false;
            }
        }

        return true;
    }

    // check 3 x 3 grid
    bool checkSubBoard(int r, int c, vector<vector<char>> &board) {
        unordered_set<char> digits;
        for (int i = 0; i < 3; i++) {
            for (int j = 0; j < 3; j++) {
                char x = board[r + i][c + j];
                if (isdigit(x)) {
                    if (!digits.count(x)) {
                        digits.insert(x);
                    } else {
                        return false;
                    }
                }
            }
        }
        return true;
    }

    bool checkRow(int r, vector<vector<char>> &board) {
        unordered_set<char> digits;
        for (int j = 0; j < 9; j++) {
            char x = board[r][j];
            if (isdigit(x)) {
                if (!digits.count(x)) {
                    digits.insert(x);
                } else {
                    return false;
                }
            }
        }
        return true;
    }

    bool checkColumn(int c, vector<vector<char>> &board) {
        unordered_set<char> digits;
        for (int i = 0; i < 9; i++) {
            char x = board[i][c];
            if (isdigit(x)) {
                if (!digits.count(x)) {
                    digits.insert(x);
                } else {
                    return false;
                }
            }
        }
        return true;
    }
};

int main() {
    Solution solution;
    vector<vector<char>> board;

    board = {{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
             {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
             {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
             {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
             {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
             {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
             {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
             {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
             {'.', '.', '.', '.', '8', '.', '.', '7', '9'}};
    cout << solution.isValidSudoku(board) << endl;

    board = {{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
             {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
             {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
             {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
             {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
             {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
             {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
             {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
             {'.', '.', '.', '.', '8', '.', '.', '7', '9'}};
    cout << solution.isValidSudoku(board) << endl;

    board = {{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
             {'.', '.', '.', '.', '.', '.', '.', '.', '.'},
             {'5', '.', '.', '.', '.', '.', '.', '9', '.'},
             {'.', '.', '.', '5', '6', '.', '.', '.', '.'},
             {'4', '.', '3', '.', '.', '.', '.', '.', '1'},
             {'.', '.', '.', '7', '.', '.', '.', '.', '.'},
             {'.', '.', '.', '5', '.', '.', '.', '.', '.'},
             {'.', '.', '.', '.', '.', '.', '.', '.', '.'},
             {'.', '.', '.', '.', '.', '.', '.', '.', '.'}};
    cout << solution.isValidSudoku(board) << endl;

    return 0;
}