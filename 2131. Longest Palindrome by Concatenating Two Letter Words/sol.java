import java.util.HashMap;
import java.util.Map;
import java.Math;


class Solution {
    public int longestPalindrome(String[] words) {
        Map<String, Integer> m = new HashMap<>();
        for (String word : words) {
            m.put(word, m.getOrDefault(word, 0) + 1);
        }

        int palindromeLength = 0, singleCenter = 0;

        for (Map.Entry<String, Integer> entry : m.entrySet()) {
            String word = entry.getKey();
            String reversed = new StringBuilder(word).reverse().toString();
            int count = entry.getValue();

    
            if (word.charAt(0) == word.charAt(1)) {
                singleCenter += count % 2;
                palindromeLength += (count / 2) * 2 * 2;
            } else {
                palindromeLength += Math.min(count, m.getOrDefault(reversed, 0)) * 2;
            }
        }
        palindromeLength += singleCenter > 0 ? 2 : 0;
        return palindromeLength;
    }
}
