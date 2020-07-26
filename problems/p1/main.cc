
#include <iostream>

using namespace std;

int getTotal(int limit) {
  int total = 0;
  for (int i = 0; i < limit; i++) {
    if (i % 3 == 0 || i % 5 == 0) {
      total += i;
    }
  }

  return total;
}

int main(void) {
  int t = getTotal(10);
  cout << t << endl;
}
