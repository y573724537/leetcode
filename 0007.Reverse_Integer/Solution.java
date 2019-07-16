class Solution {
    public int reverse(int x) {
        if (x == Integer.MIN_VALUE) {
            return 0;
        }

        int sign = 1;
        int reverse = 0;

        if (x < 0) {
            sign = -1;
            x = -x;
        }

        while (x != 0) {
            if (reverse * 10.0 + x % 10 > Integer.MAX_VALUE) {
                return 0;
            }

            reverse = reverse * 10 + x % 10;
            x /= 10;
        }

        return reverse * sign;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.reverse(-2147483648));
    }
}
