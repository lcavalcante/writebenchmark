#include <errno.h>
#include <fcntl.h>  // for open
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#define BYTES 100000

int main(int argc, char *argv[]) {
    if (argc != 2) {
        write(2, "arg error", 9);
        exit(1);
    }

    int fp = open(argv[1], O_RDWR | O_CREAT);
    if (fp < 0) {
        write(2, "error opening file", 19);
        exit(1);
    }

    char text[BYTES] = "write";
    errno = 0;
    if (BYTES != write(fp, text, BYTES)) {
        char* error = strerror(errno);
        write(2, error, strlen(error));
        write(2, "\nerror write file", 17);
        exit(1);
    }

    close(fp);
    return 0;
}
