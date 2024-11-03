# Problem Description
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

### Example:
- Input: `nums = [2,2,1]`
- Output: `1`

- Input: `nums = [4,1,2,1,2]`
- Output: `4`

- Input: `nums = [1]`
- Output: `1`

## Initial Approach
I was stucked, because needed linear time complexity and constant space complexity,  nums indices themselves because nums[i] withing -int32 to +int32.

## Actual Approach
Iterate through numbers and xor with a result each value.

## What I Didn't Think About The First Time
If u need to find all the unique elements consider using xor

## Solution (C++)
```cpp
class Solution {
  public:
    int singleNumber(std::vector<int>& nums) {
      int res = 0;
      for (const int& num : nums) {
        res ^= num;
      }
      return res;
    }
};
```

## Complexity
n - length of nums
time: O(n)
space: O(n)