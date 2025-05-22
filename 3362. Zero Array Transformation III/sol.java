import java.util.PriorityQueue;
import java.util.HashMap;
import java.util.List;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Map;



class Solution {
    public int maxRemoval(int[] nums, int[][] queries) {
        final int n = nums.length, q = queries.length;
        PriorityQueue<Integer> pqueue = new PriorityQueue<>(Collections.reverseOrder());
        Map<Integer, List<Integer>> ends = new HashMap<>(n);
        int[] tmp = new int[n+1];
        for(int i=0; i<n; i++) ends.put(i, new ArrayList<>());
        for(int i=0;i <q; i++) {
            ends.get(queries[i][0]).add(queries[i][1]);
        }

        for(int i=0, s=0; i<n; i++) {
            final int x = nums[i];
            s+=tmp[i];
            for(int j : ends.get(i)) {
                pqueue.add(j);
            }
            for(; x>s && pqueue.size()>0 && pqueue.peek()>=i; s++) {
                tmp[pqueue.poll()+1]--;
            }
            if(x>s) return -1;
        }

        return pqueue.size();
    }
}