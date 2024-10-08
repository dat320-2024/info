#include <stdio.h>
#include <assert.h>
#include <pthread.h>

void *someMethod(void *arg) {
     for (int i = 0; i < 10; i++) {
    printf("%s at %d \n", (char *) arg, i);
     }
    return NULL;
}

int main(int argc, char *argv[]) {
    pthread_t p1, p2;
    int rc;
    printf("main: begin\n");
    rc = pthread_create(&p1, NULL, someMethod, "A");
    assert(rc == 0);
    rc = pthread_create(&p2, NULL, someMethod, "B");
    assert(rc == 0);
    // join waits for the threads to finish
    rc = pthread_join(p2, NULL);
    assert(rc == 0);
    rc = pthread_join(p1, NULL);
    assert(rc == 0);
    printf("main: end\n");
    return 0;
}