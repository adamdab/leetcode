import java.util.HashMap;
import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Deque;
import java.util.ArrayDeque;
import java.Math;

class Solution {

    final int ALPHABET_COUNT = 26;

    public int largestPathValue(String colors, int[][] edges) {
        final int n = colors.length();

        int[][] tmp = new int[n][ALPHABET_COUNT];
        int[] incoming = new int[n];
        List<Integer>[] outgoing = new List[n];
        
        Arrays.setAll(outgoing, e -> new ArrayList<>());

        for(int i=0; i<edges.length; i++) {
            outgoing[edges[i][0]].add(edges[i][1]);
            incoming[edges[i][1]]++;
        }

        Deque<Integer> q = new ArrayDeque<>();
        int cnt = 0;
        int answer = -1;
        for(int i=0; i<n; i++) if(incoming[i]==0) {
            q.add(i);
        }

        while(!q.isEmpty()) {
            int root = q.remove();
            int colorIndex = colors.charAt(root) - 'a';
            tmp[root][colorIndex]+=1;
            answer = Math.max(answer, tmp[root][colorIndex]);
            cnt++;
            for(int child : outgoing[root]) {
                if(--incoming[child] == 0) q.add(child);
                for(int k=0; k<ALPHABET_COUNT; k++) tmp[child][k] = Math.max(tmp[child][k], tmp[root][k]);
            }
        }

        return cnt == n ? answer : -1;
        
    }
}