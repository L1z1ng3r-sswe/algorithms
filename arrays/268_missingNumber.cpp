
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

class Solution {
 public:
  int missingNumber(std::vector<int> nums) {
    int n = nums.size();
    int estimate = n * (n + 1) / 2;

    int total = 0;

    for (const int& num : nums) {
      total += num;
    }

    return estimate - total;
  }
};