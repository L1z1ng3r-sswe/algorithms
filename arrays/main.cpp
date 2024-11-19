// https://codeforces.com/contest/1805/problem/B

#include <iostream>
#include <string>
#include <algorithm>

int main() {
  int t;
  std::cin >> t;
  while (t--) {
    int n;
    std::cin >> n; 
    std::string s;
    std::cin >> s; 

    char minChar = 'z' + 1;
    int idx = -1;

    for (int i = 0; i < n; ++i) {
      char ch = s[i];
      if (minChar > ch) {
        minChar = ch;
        idx = i;
      } else if (minChar == s[i]) {
        idx = i;
      }
    }

    std::cout << s[idx] + s.substr(0, idx) + s.substr(idx+1, n) << std::endl;
  }
  return 0;
}