#include "uthash.h"
#include <stdio.h>

typedef struct {
    int num; /* key */
    int idx;
    UT_hash_handle hh; /* makes this structure hashable */
} entry;

int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
    *returnSize = 2;
    int* res = (int*)malloc(*returnSize * sizeof(int));
    entry* entries = NULL;

    for (int i = 0; i < numsSize; i++) {
        entry* e;
        int otherNum = target - nums[i];
        HASH_FIND_INT(entries, &otherNum, e);
        if (e == NULL) {
            e = (entry*)malloc(sizeof(entry));
            e->num = nums[i];
            e->idx = i;
            HASH_ADD_INT(entries, num, e);
        } else {
            res[0] = e->idx;
            res[1] = i;
            return res;
        }
    }

    return res;
}

void printResult(int* res, int resSize) {
    for (int i = 0; i < resSize; i++) {
        printf("%d ", res[i]);
    }
    printf("\n");
}

int main() {
    int* res = NULL;

    int nums[5];
    int numsSize;
    int target;
    int returnSize;

    /* Output: [0,1] */
    nums[0] = 2;
    nums[1] = 7;
    nums[2] = 11;
    nums[3] = 15;
    numsSize = 4;
    target = 9;
    res = twoSum(nums, numsSize, target, &returnSize);
    printResult(res, returnSize);
    free(res);

    /* Output: [1,2] */
    nums[0] = 3;
    nums[1] = 2;
    nums[2] = 4;
    numsSize = 3;
    target = 6;
    res = twoSum(nums, numsSize, target, &returnSize);
    printResult(res, returnSize);
    free(res);

    /* Output: [0,1] */
    nums[0] = 3;
    nums[1] = 3;
    numsSize = 2;
    target = 6;
    res = twoSum(nums, numsSize, target, &returnSize);
    printResult(res, returnSize);
    free(res);

    return 0;
}