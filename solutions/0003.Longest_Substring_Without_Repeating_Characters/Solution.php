<?php

class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function lengthOfLongestSubstring($s) {
        $dict = array();
        $max = $preLen = $curLen = 0;

        for ($i = 0; $i < strlen($s); $i++) {
            if (!isset($dict[$s[$i]])) {
                $curLen = $preLen + 1;
            } else {
                $conLen = $preLen + 1;
                $posLen = $i - $dict[$s[$i]];
                $curLen = $conLen < $posLen ? $conLen : $posLen;
            }

            $max = $max < $curLen ? $curLen : $max;
            $dict[$s[$i]] = $i;
            $preLen = $curLen;
        }

        return $max;
    }
}
