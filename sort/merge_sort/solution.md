## 思路

### master公式
这里需要关注以下master公式，该公式用于计算递归函数的时间复杂度

T(N) = a(T/b) + O(N<sup>d</sup>)

* if log<sub>b</sub>a < d, T(N) = O(N<sup>d</sup>)
* if log<sub>b</sub>a > d, T(N) = O(N<sup>log<sub>b</sub>a</sup>)
* if log<sub>b</sub>a == d, T(N) = O(N<sup>d</sup> * log)