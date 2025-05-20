#include <stdio.h>
#include <stdbool.h>

bool isZeroArray(int* nums, int numsSize, int** queries, int queriesSize, int* queriesColSize) {
    int* tmp = (int*)malloc((numsSize + 1) * sizeof(int));
    for (int i = 0; i <= numsSize; i++)
        tmp[i] = 0;

    for (int i = 0; i < queriesSize; i++) {
        int li = *(*(queries + i)), ri = *(*(queries + i) + 1);
        tmp[li]++;
        tmp[ri + 1]--;
    }

    for (int i = 0, s = 0; i < numsSize; i++) {
        s += tmp[i];
        if (nums[i] > s) {
            return false;
        }
    }

    return true;
}