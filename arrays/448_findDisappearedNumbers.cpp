
class Solution {
 public:
  std::vector<int> findDisappearedNumbers(std::vector<int>& nums) {
    for (size_t i = 0; i < nums.size(); ++i) {
      size_t idx = std::abs(nums[i]);

      if (nums[idx - 1] > 0) {
        nums[idx - 1] = -nums[idx - 1];
      }
    }

    std::vector<int> res;
    for (size_t i = 0; i < nums.size(); i++) {
      if (nums[i] > 0) {
        res.push_back(i + 1);
      }
    }

    return res;
  }
};