<?php

/**
 * Definition for a singly-linked list.
 */
class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}

class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2) {
        $head = new ListNode(0);
        $prev = $head;
        $carry = 0;
        $curVal = 0;

        while ($l1 != null && $l2 != null) {
            $val = $l1->val + $l2->val + $carry;
            $curVal = $val % 10;
            $carry = floor($val / 10);

            $cur = new ListNode($curVal);
            $prev->next = $cur;
            $prev = $cur;
            $l1 = $l1->next;
            $l2 = $l2->next;
        }

        while ($l1 != null) {
            $val = $l1->val + $carry;
            $curVal = $val % 10;
            $carry = floor($val / 10);

            $cur = new ListNode($curVal);
            $prev->next = $cur;
            $prev = $cur;
            $l1 = $l1->next;
        }

        while ($l2 != null) {
            $val = $l2->val + $carry;
            $curVal = $val % 10;
            $carry = floor($val / 10);

            $cur = new ListNode($curVal);
            $prev->next = $cur;
            $prev = $cur;
            $l2 = $l2->next;
        }

        if ($carry != 0) {
            $prev->next = new ListNode($carry);
        }

        return $head->next;
    }
}
