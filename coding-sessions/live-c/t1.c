#include <stdio.h>
#include <assert.h>
#include <pthread.h>

 int counter = 0;

void *DoCount(void *arg) {
    printf("%s: begin\n", (char *) arg);
    int i;
    for (i = 0; i < 10e5; i++) {
        counter++;
        // counter = counter + 1;
    }
    printf("%s: end\n", (char *)arg);
    return NULL;
}

int main(int argc, char *argv[]) {
    pthread_t p1, p2;
    printf("main: begin (counter = %d)\n", counter);
    int rc;
    rc = pthread_create(&p1, NULL, DoCount, "A");
    assert(rc == 0);
    rc = pthread_create(&p2, NULL, DoCount, "B");
    assert(rc == 0);
    rc = pthread_join(p1, NULL);
    assert(rc == 0);
    rc = pthread_join(p2, NULL);
    assert(rc == 0);
    printf("main: end (counter = %d)\n", counter);
    return 0;
}
