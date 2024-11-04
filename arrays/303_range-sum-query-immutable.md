# Problem Description
Given an integer array nums, handle multiple queries of the following type:

Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
Implement the NumArray class:

NumArray(int[] nums) Initializes the object with the integer array nums.
int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).

### Example:
- Input: `["NumArray", "sumRange", "sumRange", "sumRange"][[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]`

- Output: `[null, 1, -1, -3]`

## Initial Approach
Combine the prefix sum and return the subtraction of the left-1 prefix sum (left prefix sum minus left nums) from the right prefix sum.

## Actual Approach
Combine the prefix sum and return the subtraction of the left-1 prefix sum (left prefix sum minus left nums) from the right prefix sum.

## What I Didn't Think About The First Time


## Solution (C++)
```cpp
class Solution {
public:
    Solution(std::vector<int>& nums) {
      _nums = nums;
      _pref_sum.resize(nums.size()+1);
      _pref_sum[0] = 0;

      for (int i = 0; i < nums.size(); ++i ) {
        _pref_sum[i+1] = _pref_sum[i] + nums[i];
      }  
    };
    
    int sumRange(int left, int right) {
      return _pref_sum[right+1] - (_pref_sum[left+1] - _nums[left]);
    };
private:
  std::vector<int> _nums;
  std::vector<int> _pref_sum;
};
```

## Complexity
insert: O(n)
pop: O(1)