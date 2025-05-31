import java.util.ArrayDeque;
import java.util.Deque;
import java.Math;

class Solution {
    private int N;

    public int snakesAndLadders(int[][] board) {
        N = board.length; 
        final int start = 1;
        final int target = N * N;
        Deque<Integer> queue = new ArrayDeque<>();
        queue.offer(start);
        boolean[] visited = new boolean[target + 1];
        visited[start] = true; 
        int moves = 0; 
        
        while (!queue.isEmpty()) {
            for (int i = queue.size(); i > 0; --i) {
                int current = queue.poll();
                if (current == target) return moves;
                for (int k = current + 1; k <= Math.min(current + 6, target); ++k) {
                    int[] position = convertToPosition(k);
                    int next = k;
                    if (board[position[0]][position[1]] != -1) {
                        next = board[position[0]][position[1]];
                    }
                    if (!visited[next]) {
                        visited[next] = true;
                        queue.offer(next);
                    }
                }
            }
            moves++;
        }
        return -1;
    }

    private int[] convertToPosition(int squareNum) {
        int row = (squareNum - 1) / N;
        int col = (squareNum - 1) % N;
        if (row % 2 == 1) {
            col = N - 1 - col; 
        }
        return new int[] {N - 1 - row, col};
    }
}