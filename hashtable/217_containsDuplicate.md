# Problem Description
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

### Example:
- Input: `nums = [1,2,3,1]`
- Output: `true` 

- Input: `nums = [1,2,3,4]`
- Output: `false`

- Input: `nums = [1,1,1,3,3,4,3,2,4,2]`
- Output: `true`

## Initial Approach
```cpp
// step 1: define map<int:int>
// step 2: iterate through nums
// step 3: on each iteration check apperance ? return true : instert to map
// step 4: return false if each value is distinct
```

## Actual Approach 
```cpp
// step 1: define set
// step 2: iterate through nums
// step 3: on each iteration check apperance ? return true : inster to set
// step 4: return false if each value is distinct
```

## What I Didn't Think About The First Time
Use a set instead of a map if you only need to check for the apperance of an element without storing anything as a value.

## Solution (C++)
```cpp
class Solution {
  public:
    bool containsDuplicate(const std::vector<int>& nums) {
      std::unordered_set<int> numsSet;
      numsSet.reserve(nums.size()/2);

      for (const int& num : nums) {
        if (numsSet.find(num) != numsSet.end()) {
          return true;
        }
        numsSet.insert(num);
      }

      return false;
    }
};
```

## Complexity
n - length of nums
time: O(n)
space: O(n)