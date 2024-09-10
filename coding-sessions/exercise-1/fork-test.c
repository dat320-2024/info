#include <stdio.h>
#include <unistd.h>

int main() {
   int pid = fork();

    if (pid < 0) {
        fprintf(stderr, "Fork failed\n");
        return 1;
    } else if (pid == 0) {
        printf("Child process\n %d", getpid());
    } else {
        printf("Parent process\n");
    }

    printf("Fork return value: %d\n", pid);

    return 0;
}