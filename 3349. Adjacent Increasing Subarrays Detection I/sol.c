#include <stdbool.h>

int maxIncreasingLength(int* nums, int numsSize, int start) {
    for (int i = 1; i < numsSize - start; i++) {
        if (nums[start + i - 1] >= nums[start + i])
            return i;
    }
    return numsSize - start;
}

bool hasIncreasingSubarrays(int* nums, int numsSize, int k) {
    if (numsSize < 2 * k)
        return false;
    for (int start = 0; start < numsSize - k;) {
        int firstSize = maxIncreasingLength(nums, numsSize, start);

        if (firstSize >= 2 * k)
            return true;

        if (firstSize < k) {
            start += firstSize;
        } else {
            int secondSize =
                maxIncreasingLength(nums, numsSize, start + firstSize);

            if (secondSize < k) {
                start += firstSize + secondSize;
            } else
                return true;
        }
    }
    return false;
}