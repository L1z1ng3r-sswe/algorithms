
class Solution {
 public:
  Solution(std::vector<int>& nums) {
    _nums = nums;
    _pref_sum.resize(nums.size() + 1);
    _pref_sum[0] = 0;

    for (int i = 0; i < nums.size(); ++i) {
      _pref_sum[i + 1] = _pref_sum[i] + nums[i];
    }
  };

  int sumRange(int left, int right) {
    return _pref_sum[right + 1] - (_pref_sum[left + 1] - _nums[left]);
  };

 private:
  std::vector<int> _nums;
  std::vector<int> _pref_sum;
};