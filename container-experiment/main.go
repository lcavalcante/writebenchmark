package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    // "syscall"
    "time"

    "github.com/containerd/containerd"
    "github.com/containerd/containerd/cio"
    "github.com/containerd/containerd/namespaces"
    "github.com/containerd/containerd/oci"
    specs "github.com/opencontainers/runtime-spec/specs-go"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
}

func main() {
    n, err := strconv.Atoi(os.Args[1])
    check(err)

    f, err := os.Create("output-" + os.Args[1] + ".csv")
    check(err)
    defer f.Close()

    respond := make(chan time.Duration, n)

    for i := 878; i <= n; i++ {
        //value := strconv.Itoa(i)
        var str strings.Builder
        str.WriteString("/dev/rbd/scalable_sgx/")
        rbd := fmt.Sprintf("test-%dk", i)
        str.WriteString(rbd)
        go writeTest(respond, str.String(), rbd)
    }

    for i := 878; i <= n; i++ {
        diff := <-respond
        b_wrote, err := f.WriteString(fmt.Sprintf("%d, %d, %v\n", n, i, diff))
        check(err)
        log.Printf("wrote %d bytes \n", b_wrote)
        f.Sync()
    }
    //if err := writeTest(); err != nil {
    //  log.Fatal(err)
    //}
}

func printTest(a string) error {
    log.Printf("%s\n", a)

    return nil
}

func writeTest(respond chan<- time.Duration, mp string, rbd string) {
    starttime := time.Now()
    client, err := containerd.New("/run/containerd/containerd.sock")
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    ctx := namespaces.WithNamespace(context.Background(), "example4")
    image, err := client.Pull(ctx, "docker.io/lucasmc/writetest:latest", containerd.WithPullUnpack)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Successfully pulled %s image\n", image.Name())

    mnt := []specs.Mount{{Destination: "/data", Type: "bind", Source: "/aoba", Options:[]string{"rw","rbind"}}}
    container, err := client.NewContainer(
            ctx,
            rbd,
            containerd.WithNewSnapshot(rbd, image),
            containerd.WithNewSpec(oci.WithImageConfig(image), oci.WithMounts(mnt)),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer container.Delete(ctx, containerd.WithSnapshotCleanup)
    log.Printf("Successfully created container with ID %s and snapshot with ID write-test-snapshot", container.ID())

    task, err := container.NewTask(ctx, cio.NewCreator(cio.WithStdio))
    if err != nil {
        log.Fatal(err)
    }
    defer task.Delete(ctx)

    // make sure we wait before calling start
    exitStatusC, err := task.Wait(ctx)
    if err != nil {
        fmt.Println(err)
    }

    // call start on the task to execute the redis server
    if err := task.Start(ctx); err != nil {
        log.Fatal(err)
    }

    // kill the process and get the exit status
    // if err := task.Kill(ctx, syscall.SIGTERM); err != nil {
    //     log.Fatal(err)
    // }

    // wait for the process to fully exit and print out the exit status
    status := <-exitStatusC
    code, _, err := status.Result()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("write-test exited with status: %d\n", code)
    endtime := time.Now()
    diff := endtime.Sub(starttime)
    //log.Println(diff)
    respond <- diff
}
