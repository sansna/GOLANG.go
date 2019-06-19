#include <stdio.h>
#include "a.h"

int Lib() {
    // C print buffers per line, so the 2nd line is not visible.
    printf("%s\n", "hello world");
    //printf("%s", "hello world");
    return 0;
}

int main() {
    Lib();
    return 0;
}
