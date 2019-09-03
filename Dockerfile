FROM gcc:latest 
#as base

COPY src write-test/src
WORKDIR write-test
RUN mkdir bin
RUN gcc src/write.c -o bin/write

CMD ["./bin/write", "file"] 