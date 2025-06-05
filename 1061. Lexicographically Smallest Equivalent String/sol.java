import java.Math.*;
import java.util.*;


class Solution {

    char[] chars = "abcdefghijklmnopqrstuvwxyz".toCharArray();

    void swap(char from, char to) {
        int fromIndex = from - 'a';
        int toIndex = to - 'a';
        char currentMin = chars[fromIndex];
        char mappedTo = chars[toIndex];
        if(mappedTo < currentMin) {
            for(int i=0; i< chars.length; i++) {
                if(chars[i] == currentMin) {
                    chars[i] = mappedTo;
                }
            }
        }
    }

    public String smallestEquivalentString(String s1, String s2, String baseStr) {
        for(int i = 0; i< s1.length(); i++) {
            swap(s1.charAt(i), s2.charAt(i));
            swap(s2.charAt(i), s1.charAt(i));
        }

        StringBuilder sb = new StringBuilder();
        for(char c : baseStr.toCharArray()) {
            int encoding = c - 'a';
            sb.append(chars[encoding]);
        }

        return sb.toString();
    }
}