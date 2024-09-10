#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <time.h>

int main(int argc, char *argv[]) {
    int i;
    time_t start_time = time(NULL);
    time_t current_time;

    for (i = 0; i < 6; i++) {
        current_time = time(NULL);
        printf("Printing pid and elapsed time every 5 seconds: PID=%d, Elapsed Time=%ld seconds\n", getpid(), current_time - start_time);
        sleep(5);
    }

    printf("Arguments passed to the program:\n");
    for (i = 0; i < argc; i++) {
        printf("Argument %d: %s\n", i, argv[i]);
    }

    return 0;
}