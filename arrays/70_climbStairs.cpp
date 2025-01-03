
class Solution {
 public:
  int climbStairs(int n) {
    if (n <= 3) {
      return n;
    }

    int prevPrev = 2;
    int prev = 3;

    int res = 0;
    for (n -= 3; n > 0; --n) {
      res = prev + prevPrev;
      prevPrev = prev;
      prev = res;
    }

    return res;
  }
}