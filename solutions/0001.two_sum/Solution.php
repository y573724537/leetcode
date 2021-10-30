<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target) {
        $records = array();

        for ($i = 0; $i < count($nums); $i++) {
            $mate = $target - $nums[$i];
            if (isset($records[$mate])) {
                return array($records[$mate], $i);
            }

            $records[$nums[$i]] = $i;
        }

        return null;
    }
}
