class Solution {
    public String convert(String s, int numRows) {
        if (numRows < 2) {
            return s;
        }

        int length = Math.min(s.length(), numRows);

        StringBuilder[] containers = new StringBuilder[length];

        for (int i = 0; i < length; i++) {
            containers[i] = new StringBuilder();
        }

        boolean isDown = false;
        for (int i = 0, j = 0; i < s.length(); i++) {
            if (j == 0 || j == length - 1) {
                isDown = !isDown;
            }

            containers[j].append(s.substring(i, i + 1));

            if (isDown) {
                j++;
            } else {
                j--;
            }
        }

        StringBuilder convert = new StringBuilder();
        for (int i = 0; i < length; i++) {
            convert.append(containers[i]);
        }

        return convert.toString();
    }
}
