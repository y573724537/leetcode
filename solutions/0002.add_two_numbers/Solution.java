/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode head = new ListNode(0);
        ListNode prev = head;
        int carry = 0, curVal = 0;

        while (l1 != null && l2 != null) {
            int val = l1.val + l2.val + carry;
            carry = val / 10;
            curVal = val % 10;

            ListNode cur = new ListNode(curVal);

            prev.next = cur;
            prev = cur;
            l1 = l1.next;
            l2 = l2.next;
        }

        while (l1 != null) {
            int val = l1.val + carry;
            carry = val / 10;
            curVal = val % 10;

            ListNode cur = new ListNode(curVal);

            prev.next = cur;
            prev = cur;
            l1 = l1.next;
        }

        while (l2 != null) {
            int val = l2.val + carry;
            carry = val / 10;
            curVal = val % 10;

            ListNode cur = new ListNode(curVal);

            prev.next = cur;
            prev = cur;
            l2 = l2.next;
        }

        if (carry != 0) {
            prev.next = new ListNode(carry);
        }

        return head.next;
    }
}
