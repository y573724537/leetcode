<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target) {
        $dict = array();

        for ($i = 0; $i < count($nums); $i++) {
            $another = $target - $nums[$i];
            if (isset($dict[$another])) {
                return array($dict[$another], $i);
            }

            $dict[$nums[$i]] = $i;
        }

        return null;
    }
}
