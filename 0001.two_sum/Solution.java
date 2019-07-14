import java.util.Map;
import java.util.HashMap;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> dict = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            int another = target - nums[i];
            if (null != dict.get(another)) {
                return new int[]{dict.get(another), i};
            }

            dict.put(nums[i], i);
        }

        return null;
    }
}
