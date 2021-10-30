import java.util.Map;
import java.util.HashMap;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> records = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            int mate = target - nums[i];
            if (null != records.get(mate)) {
                return new int[]{records.get(mate), i};
            }

            records.put(nums[i], i);
        }

        return null;
    }
}
