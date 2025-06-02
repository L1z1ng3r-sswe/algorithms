#include <vector>

int maxProfit(std::vector<int>& prices) {
  int res = 0;

  int minPrice = prices[0];
  int maxPrice = prices[0];

  for (size_t i = 1; i < prices.size(); ++i) {
    if (maxPrice < prices[i]) {
      maxPrice = prices[i];
      res = max(maxPrice, maxPrice - minPrice);
    } else if (prices[i] < minPrice) {
      minPrice = prices[i];
      maxPrice = prices[i];
    }
  }

  return res;
}

int max(int a, int b) {
  if (a > b) {
    return a;
  }
  return b;
}