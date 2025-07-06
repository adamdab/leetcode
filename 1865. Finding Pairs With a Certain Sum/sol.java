import java.util.HashMap;
import java.util.Map;

class FindSumPairs {

    int[] nums2, nums1;
    Map<Integer, Integer> mapNums2;

    public FindSumPairs(int[] nums1, int[] nums2) {
        this.nums2 = nums2;
        this.nums1 = nums1;
        mapNums2 = new HashMap<>(nums2.length);
        for (int n : nums2) {
            mapNums2.put(n, mapNums2.getOrDefault(n, 0) + 1);
        }
    }

    //  Adds val to nums2[index], i.e., apply nums2[index] += val
    public void add(int index, int val) {
        int pre = nums2[index];
        int left = mapNums2.get(pre);
        if (left == 1)
            mapNums2.remove(pre);
        else
            mapNums2.put(pre, left - 1);
        mapNums2.put(pre + val, mapNums2.getOrDefault(pre + val, 0) + 1);
        nums2[index] += val;
    }

    // Returns the number of pairs (i, j) such that nums1[i] + nums2[j] == tot
    public int count(int tot) {
        int res = 0;
        for (int n : nums1) {
            res += mapNums2.getOrDefault(tot - n, 0);
        }
        return res;
    }
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * FindSumPairs obj = new FindSumPairs(nums1, nums2);
 * obj.add(index,val);
 * int param_2 = obj.count(tot);
 */