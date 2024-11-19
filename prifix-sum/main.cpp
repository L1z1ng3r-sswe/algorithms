#include <iostream>
#include <vector>

int main() {
  int t;
  std::cin >> t;

  int prevSum = 0;
  while (t--) {
    int num;
    std::cin >> num;

    prevSum += num;
  }

  std::cout << prevSum;
}


