import java.util.Map;
import java.util.HashMap;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        char[] array = s.toCharArray();

        Map<Character, Integer> dict = new HashMap<>();
        int max = 0, preLen = 0, curLen = 0;

        for (int i = 0; i < array.length; i++) {
            if (null == dict.get(array[i])) {
                curLen = preLen + 1;
            } else {
                curLen = Math.min(preLen + 1, i - dict.get(array[i]));
            }

            max = Math.max(max, curLen);
            dict.put(array[i], i);
            preLen = curLen;
        }

        return max;
    }
}
