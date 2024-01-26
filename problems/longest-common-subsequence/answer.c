#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int helper(char* text1, int i1, char* text2, int i2, int** cache) {
    if (cache[i1][i2] != -1) {
        return cache[i1][i2];
    }
    int res;
    if (text1[i1] == '\0' || text2[i2] == '\0') {
        res = 0;
    } else if (text1[i1] == text2[i2]) {
        res = 1 + helper(text1, i1 + 1, text2, i2 + 1, cache);
    } else {
        int res1 = helper(text1, i1 + 1, text2, i2, cache);
        int res2 = helper(text1, i1, text2, i2 + 1, cache);
        res = res1 >= res2 ? res1 : res2;
    }
    cache[i1][i2] = res;
    return res;
}

int** init2dArray(int rows, int cols, int initVal) {
    int** arr = (int**)malloc(rows * sizeof(int*));
    for (int i = 0; i < rows; i++) {
        arr[i] = (int*)malloc(cols * sizeof(int));
        memset(arr[i], -1, cols * sizeof(int));
    }
    return arr;
}

void free2dArray(int** arr, int rows) {
    for (int i = 0; i < rows; i++) {
        free(arr[i]);
    }
    free(arr);
}

int longestCommonSubsequence(char* text1, char* text2) {
    int rows = strlen(text1) + 1;
    int cols = strlen(text2) + 1;
    int** cache = init2dArray(rows, cols, -1);
    int res = helper(text1, 0, text2, 0, cache);
    free2dArray(cache, rows);
    return res;
}

int main() {
    char* text1;
    char* text2;
    int res;

    text1 = "abcde";
    text2 = "ace";
    res = longestCommonSubsequence(text1, text2);
    // 3
    printf("%d\n", res);

    text1 = "abc";
    text2 = "abc";
    res = longestCommonSubsequence(text1, text2);
    // 3
    printf("%d\n", res);

    text1 = "abc";
    text2 = "def";
    res = longestCommonSubsequence(text1, text2);
    // 0
    printf("%d\n", res);

    text1 = "bl";
    text2 = "yby";
    res = longestCommonSubsequence(text1, text2);
    // 1
    printf("%d\n", res);

    return 0;
}