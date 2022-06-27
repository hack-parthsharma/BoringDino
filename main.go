package main


import (
  "fmt"
  "time"
  "os"
  "os/signal"
  "syscall"
  bg "github.com/onealmond/boringdino/dinosaur"
)

func main() {
  quit := make(chan int)
  sink := make(chan os.Signal, 1)
  signal.Notify(sink, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
  go func(){
    for {
      sig := <-sink
      switch (sig) {
        case syscall.SIGINT:
          fallthrough
        case syscall.SIGTERM:
          fallthrough
        case syscall.SIGQUIT:
          fmt.Println("Bye~")
          quit <- 0
      }
    }
  }()

  dino := bg.NewDinosaur()
  // Action simulation
  go func() {
    for {
      dino.Measure()
      time.Sleep(2 * time.Second)
    }
  }()

  code := <-quit
  dino.Close()
  os.Exit(code)
}
