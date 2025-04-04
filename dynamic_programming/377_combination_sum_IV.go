func combinationSum4(nums []int, target int) int {
  sort.Ints(nums)
  
  dp := make(map[int]int)

  var dfs func(sum int) int
  dfs = func(sum int) int {
    if sum == target {
      return 1
    }a

    if val, ok := dp[sum]; ok {
      return val
    }

    res := 0

    for _, num := range nums {
      if sum + num <= target {
        res += dfs(sum+num)
      } else {
        break
      }
    }

    dp[sum] = res
    return res
  }

  return dfs(0)
}

// t = target
// n = len(nums)

// time: O(t*n)
// space: O(t)