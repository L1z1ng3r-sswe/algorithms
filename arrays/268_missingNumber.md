# Problem Description
Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

### Example:
- Input: `nums = [9,6,4,2,3,5,7,0,1]`
- Output: `8`

- Input: `nums = [3,0,1]`
- Output: `2`

- Input: `nums = [0,1]`
- Output: `2`

## Initial Approach
```cpp
// step 1: iterate through nums. if num != n ? swap : continue
// step 2: iterate through nums and compare idx with value
// return n
```

## Actual Approach
```cpp
// step 1: iterate through each num and place it in a while loop
// step 2: iterate through nums and compare idx with value
// step 3: return n

// step 1: calculate estimated sum using n(n+1)/2
// step 2: iterate through nums and sum up total
// step 3: return estimated - total
```

## What I Didn't Think About The First Time
First approach:
Sum Formula: The formula n(n+1)/2 works because consecutive numbers pair up symmetrically to derive 
n+1.

Second approach:
Swapping with while Loop: This ensures that each element is moved to its correct position. The while loop keeps iterating to place each element correctly, and it stops once the element is positioned or can't be swapped further (e.g., when a previously placed value is smaller).

## Solution (C++)
```cpp
// swapping approach
class Solution {
  public:
    int missingNumber(std::vector<int>& nums) {
      int n = nums.size();

      for (int i = 0; i < n; ++i) {
        while (nums[i] < n && nums[i] != nums[nums[i]]) {
          std::swap(nums[i], nums[nums[i]]);
        }
      }

      for (int i = 0; i < n; ++i) {
        if (i != nums[i]) {
           return i;
        }
      }
    
      return n;
    }
};

// Sum of the First n Natural Numbers n(n+1)/2
class Solution {
  public:
    int missingNumber(std::vector<int> nums) {
      int n = nums.size();
      int estimate = n * (n+1) / 2;

      int total = 0;

      for (const int& num : nums) {
        total+=num;
      }

      return estimate - total;
    }
};
```

## Complexity
n - length of nums
time: O(n)
space: O(1)


