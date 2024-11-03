# Problem Description

### Example:
- Input: `n = 2`
- Output: `2`

- Input: `n = 3`
- Output: `3`

## Initial Approach
fibonacci numbers

## Actual Approach
fibonacci numbers

## What I Didn't Think About The First Time


## Solution (C++)
```cpp

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
};
```

## Complexity
n - input data
time: O(n)
space: O(1)