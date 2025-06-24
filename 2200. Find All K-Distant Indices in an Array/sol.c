/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
inline int max(int a, int b) { return a > b ? a : b; }
inline int min(int a, int b) { return a < b ? a : b; };

int* findKDistantIndices(int* nums, int numsSize, int key, int k,
                         int* returnSize) {
    int* result = (int*)malloc((numsSize) * sizeof(int));
    int maxLeftIndex = -1;
    int counter = 0;
    for (int i = 0; i < numsSize; i++) {
        if (nums[i] == key) {
            maxLeftIndex = max(maxLeftIndex, i - k - 1);
            for (int j = maxLeftIndex + 1; j < i; j++) {
                result[counter++] = j;
            }
            int start = max(i, maxLeftIndex + 1);
            maxLeftIndex = min(i + k, numsSize - 1);
            for (int j = start; j < maxLeftIndex + 1; j++) {
                result[counter++] = j;
            }
        }
    }
    *returnSize = counter;
    return result;
}