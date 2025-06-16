int maximumDifference(int* nums, int numsSize) {
    int res = -1;
    int min = nums[0];
    for(int i=1; i< numsSize; i++) {
        if(min > nums[i]) min = nums[i];
        if(min < nums[i]) res = res < nums[i] - min ? nums[i] - min : res;
    }
    return res;
}