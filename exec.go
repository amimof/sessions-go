package main

import (
  "log"
  "os/exec"
  "io"
  "os"
  "flag"
  "time"
  "fmt"
)

func main() {
  var username string
  var host string

  flag.StringVar(&username, "u", "root", "Username")
  flag.StringVar(&host, "h", "localhost", "Host")

  flag.Parse()

  //cmd := exec.Command("ssh", host, "-l", username)
  cmd := exec.Command("ping", "-n", "3", "localhost")

  stdout, err := cmd.StdoutPipe()
  if err != nil {
    log.Fatal(err)
  }

  stderr, err := cmd.StderrPipe()
  if err != nil {
    log.Fatal(err)
  }

  cmd.Stdin = os.Stdin

  err = cmd.Start()
  if err != nil {
    log.Fatal(err)
  }

  go io.Copy(os.Stdout, stdout)
  go io.Copy(os.Stderr, stderr)

  log.Printf("Command START")
  fmt.Println("-----------------------------------------")
  startt := time.Now()
  cmd.Wait()
  fmt.Println("\n-----------------------------------------")
  log.Printf("Command END")
  log.Printf("Command lasted for %v", time.Since(startt))

}
