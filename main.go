package main

import (
  "bytes"
  "fmt"
  "os/exec"
  // for os.Args
  "os"
  "log"
)

func  main() {

  arg := os.Args
  cmd := exec.Command(arg[1])
  var out bytes.Buffer
  cmd.Stdout = &out
  err := cmd.Run();if err != nil { log.Fatal(err) }
  fmt.Println(out.String())

}
