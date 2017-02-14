#include<stdio.h>

int smaller(int a, int b) {
  if (a <= b) {
    return a;
  }
  return b;
}

int main() {
  int a, b, min;
  a = 10;
  b = 2;
  min = smaller(a, b);
  printf("%d\n", min);
  return 0;
}
