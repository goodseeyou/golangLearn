package main

import (
    "time"
    "runtime"
    "fmt"
)


const stopIterNumber = 6

func recursion(i int, bs []byte) int {
    runtime.GC()
    PrintMemUsage()
    if i >= stopIterNumber {
        return 0
    }

    var a,b []byte
    for index := range bs {
        if index % 3 == 0 {
            a = append(a, bs[index])
        } else {
            b = append(b, bs[index])
        }
    }

    for len(a) < len(bs) {
        a = append(a, 0)
    }


    time.Sleep(5 * time.Second)

    i++
    bs = nil
    b = nil
    return recursion(i, a)
}

func PrintMemUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
        fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

func main() {
    var tmpByte = make([]byte, 100<<20)
    // make pysical memory allocation
    for index := range tmpByte {
        tmpByte[index] = 0
    }

    recursion(0, tmpByte)
    tmpByte = nil
    runtime.GC()
    time.Sleep(5 * time.Second)
    PrintMemUsage()
}
