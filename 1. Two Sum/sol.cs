public class Solution {
    public int[] TwoSum(int[] nums, int target) {
        Dictionary<int, int> dict = new Dictionary<int, int>();
        for(int i=0;i<nums.Length; i++) {
            int remainder = target - nums[i];
            if(dict.ContainsKey(remainder)) {
                return new int[] {dict[remainder], i};
            } else if(!dict.ContainsKey(nums[i])) {
                dict.Add(nums[i], i);
            }
        }
        throw new ArgumentException("Invalid argument");
    }
}