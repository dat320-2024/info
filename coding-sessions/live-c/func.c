
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <time.h>

int* add(int a, int b) {
    int sum = a + b;
    printf("call add ( %d + %d )= %d\n", a ,b,sum);
    
    int * ptr= malloc(sizeof(int));
    printf("Address of sum in add function: %p\n", ptr);
    *ptr = sum;
    return ptr;
    
}
int main() {
    int *sum_ptr= add(10, 20);
    printf("-------------------\n");
    printf("Address of sum_ptr in main function: %p\n", sum_ptr);
    printf("Value of sum_ptr in main function: %d\n", *sum_ptr);     
    sleep(5);
    printf("-------------------\n");
    printf("Address of sum_ptr after sleep : %p\n", sum_ptr);
    printf("Value of sum_pt after sleep: %d\n", *sum_ptr);
    free(sum_ptr);
     sleep(5);
    printf("-------------------\n");
    printf("Address of sum_ptr after free : %p\n", sum_ptr);
    printf("Value of sum_pt after free: %d\n", *sum_ptr);

    return 0;
}
