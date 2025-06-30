import java.util.Arrays;
import java.Math.*;

class Solution {
    public int findLHS(int[] nums) {
        Arrays.sort(nums);
        int prev = nums[0], max = 0, left = 0, right = 0;
        boolean wasChanged = false;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == prev)
                left++;
            else if (nums[i] - prev == 1) {
                wasChanged = true;
                right++;
            } else {
                if (wasChanged) {
                    max = Math.max(max, left + right);
                }
                if (nums[i] - nums[i - 1] == 1) {
                    left = right;
                    right = 1;
                    prev = nums[i - 1];
                    wasChanged = true;
                } else {
                    left = 1;
                    right = 0;
                    prev = nums[i];
                    wasChanged = false;
                }
            }
        }
        if (wasChanged) {
            max = Math.max(max, left + right);
        }
        return max;
    }
}