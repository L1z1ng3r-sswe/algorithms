# Problem Description
Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

### Example:
- Input: `nums = [4,3,2,7,8,2,3,1]`
- Output: `[5,6]`

- Input: `nums = [1,1]`
- Output: `[1]`

## Initial Approach
Iterate through each cell and try to find the index of its value, but there is the possibility of a loop, and adding it directly would be incorrect as swapping would be in vain.

What if we try marking each cell with a value of -1 as already read using -abs(nums[-1])? After that, we can iterate through the vector and check all the positive indices.

## Actual Approach
Iterate through each cell and try to find the index of its value, but there is the possibility of a loop, and adding it directly would be incorrect as swapping would be in vain.

What if we try marking each cell with a value of -1 as already read using -abs(nums[-1])? After that, we can iterate through the vector and check all the positive indices.

## What I Didn't Think About The First Time
if swapping doesn't work try marking

## Solution (C++)
```cpp
class Solution {
  public:
    std::vector<int> findDisappearedNumbers(std::vector<int>& nums) {
      for (size_t i = 0; i < nums.size(); ++i) {
        size_t idx = std::abs(nums[i]);

        if (nums[idx-1] > 0) {
          nums[idx-1] = -nums[idx-1];
        }
      }

      std::vector<int> res;
      for (size_t i = 0; i < nums.size(); i++) {
        if (nums[i] > 0) {
          res.push_back(i+1);
        }
      }

      return res;
    }
};
```

## Complexity
n - length of nums
time: O(n)
space: O(1)