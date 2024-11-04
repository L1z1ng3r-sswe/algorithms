# Problem Description
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

### Example:
- Input: `prices = [7,1,5,3,6,4]`
- Output: `5`

- Input: `[7,6,4,3,1]`
- Output: `0`

## Initial Approach
```go
if len of prices < 2 return 0
initialize min and max to 0 and 1 respectively // we cannot set them initially because we need to choose the time of selling

for i := 2; i < prices.size(); i++ {
  res = std::max(max - min, res); // even if nothing changes, we update res - overhead

  if (max < prices[i]) {
    max = prices[i];
  } else if (min > prices[i]) { // doesn't work because we need an initial value for min
    min = prices[i];
    // the next values could be less than max, and min could be so low that it is greater than res.
  }
}
```

## Actual Approach
```go
  initialize min and max to prices[0].
  
  if max < prices[i] {
    max = prices[i]
    res = max(res, max - min) // because min varies, and the next max could be less than the previous (as we change max to min)
  } else if min > prices[i] { // it works because previously we were initializing them to different data.
    min = prices[i]
    max = min // because the next values could be less than max, and min could be so low that it is greater than res.
  }
```

## What I Didn't Think About The First Time
Consider initial variable values correctly. How will they be used? Avoid unnecessary checks for updating res, and update only on changes to avoid overhead. Edge cases involving 0-1 values could be handled inside the main logic.

## Solution (C++)
```cpp
class Solution {
  public:
    int maxProfit(std::vector<int> prices) {
      int min = prices[0];
      int max = prices[0];
      int res = 0;

      for (size_t i = 1; i < prices.size(); ++i) {
        if (max < prices[i]) {
          max = prices[i];
          res = std::max(res, max-min);
        } else if (min > prices[i]) {
          min = prices[i];
          max = min;
        }
      }
      return res;
    }
};
```

## Complexity
n - lenght of prices
time: O(n)
space: O(1)