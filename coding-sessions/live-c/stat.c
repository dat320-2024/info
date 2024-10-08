#include <stdio.h>

void func() {
    //static int count = 0;
    int count=0;
    count++;
    printf("%d\n", count);
}
int main() {
    for (int i = 1; i <= 10; i++) {
        func();
    }
    return 0;
}