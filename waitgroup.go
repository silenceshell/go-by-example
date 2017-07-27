package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
   wg := sync.WaitGroup{}

   for i:=0; i<10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            fmt.Printf("I am goroutine %d\n", i)
            time.Sleep(time.Second)
        }(i)
   }

   fmt.Println("waiting..")
   wg.Wait()
   fmt.Println("Everything done..")

}

