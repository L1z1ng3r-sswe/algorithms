
# Problem Description

### Example:
- Input: ``
- Output: ``

- Input: ``
- Output: ``

## Initial Approach
Iterating through the n and calculate the number of 1 for each number.

## Actual Approach
1 -> 001 000+1
2 -> 010 001+0
3 -> 011 001+1
4 -> 100 010+0 -> 1+0
5 -> 101 010+1 -> 1+1
6 -> 110 
7 -> 111
8 ->1000

they are simply carrying the same value when divided by 2 and adding num & 001 (to check the least significant bit (as we have shifted it))

## What I Didn't Think About The First Time
num>>1 - means getting rid of the least significant bit or division by 2
num<<1 - means getting rid of the most significant bit or multiplying by 2

instead of writing intToVec u could just use bitwise operators...

instead of this, u could use dynamic programming instead of calculating the shit. and calculate upword:
while (n!=0) {
    total1s += n & 1; // compares least significant bit
    n >>= 1; // equalent to / 2
}

## Solution (C++)
```cpp
class Solution {
public:
    std::vector<int> countBits(int n) {
        std::vector<int> res(n + 1, 0);

        for (int i = 1; i <= n; ++i) {
            res[i] = res[i >> 1] + (i & 1);
        }

        return res;
    }
};
```

## Complexity
time: O(n)
space: O(1)