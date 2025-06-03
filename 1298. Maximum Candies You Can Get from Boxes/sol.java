import java.util.HashSet;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Set;


class Solution {
    
    // stauts -> 1/0 open/close status of boxes
    // candies -> num of candies in boxes
    // keys -> list of keys you can open with i-th box
    // containedBoxes -> list of boxes you find in box

    public int maxCandies(int[] status, int[] candies, int[][] keys, int[][] containedBoxes, int[] initialBoxes) {
        int count = 0;
        Set<Integer> avaliableBoxes = new HashSet<Integer>();
        Queue<Integer> boxesToProcess = new LinkedList<>();
        for (int box : initialBoxes) {
            if (status[box] == 1)
                boxesToProcess.add(box);
            else
                avaliableBoxes.add(box);
        }
        while (!boxesToProcess.isEmpty()) {
            int box = boxesToProcess.remove();
            // unlock boxes
            for (int key : keys[box]) {
                status[key] = 1;
            }

            // add all contained boxes
            for (int b : containedBoxes[box]) {
                // if contained box is unlocked unlock it and add to processable
                if (status[b] == 1)
                    boxesToProcess.add(b);
                else
                    avaliableBoxes.add(b);
            }

            // add now unlocked avaliable boxes
            for (int b : avaliableBoxes)
                if (status[b] == 1)
                    boxesToProcess.add(b);
            avaliableBoxes.remove(box);
            avaliableBoxes.removeIf(b -> boxesToProcess.contains(b));

            // add candy
            count += candies[box];
        }

        return count;
    }
}