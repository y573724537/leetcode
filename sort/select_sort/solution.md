## solution

对于数组arr，从0到len(arr) - 1每个位置，分别选出此数组中最小的元素，放置到该位置

例如对于位置0，从数组0到len(arr) - 1，找出最小的元素，假定位置i，则位置0上的数据与位置i上的数据互换位置，则0位置数据为最小

对于所有位置[0, len(arr) - 1]，都是当前位置范围[i, len(arr) - 1]中最小的，则整个数组是有序的