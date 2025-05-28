import java.Math;
import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;


class Solution {
    public int[] maxTargetNodes(int[][] edges1, int[][] edges2, int k) {
        final int n = edges1.length +1 , m = edges2.length + 1;
        final List<Integer>[] tree1 = createTree(edges1), tree2 = createTree(edges2);
        int[] res = new int[n];
        int maxTree2 = 0;

        for(int i=0;i<m;i++) maxTree2 = Math.max(maxTree2, dfs(tree2, i, -1, k-1));

        Arrays.fill(res, maxTree2);
        for(int i=0;i<n;i++) res[i] += dfs(tree1, i, -1, k);

        return res;
    }

    List<Integer>[] createTree(final int[][] edges) {
        final int n = edges.length + 1;
        List<Integer>[] tree = new List[n];
        Arrays.setAll(tree, e -> new ArrayList<>());
        for(int[] edge : edges) {
            int a = edge[0], b = edge[1];
            tree[a].add(b);
            tree[b].add(a);
        }
        return tree;
    }

    int dfs(final List<Integer>[] tree, int root, int parent, int depth) {
        if(depth<0) return 0;
        int res = 1;
        for(int neighbour : tree[root]) {
            if(neighbour != parent) res += dfs(tree, neighbour, root, depth -1);
        }
        return res;
    }
}