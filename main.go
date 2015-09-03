package main

import (
  "io/ioutil"
  // "time"
  // "strings"
  "bytes"
  // "fmt"
  "os/exec"
  // for os.Args
  "os"
  "log"
)

func  main() {

  // taking args from command line
  arg := os.Args
  cmd := exec.Command("curl", arg[1])
  var out bytes.Buffer
  cmd.Stdout = &out
  err := cmd.Run();if err != nil { log.Fatal(err) }
  // lines := strings.Split(out.String(), "\n")
  err = ioutil.WriteFile(arg[2], out.Bytes(), 0644)
  // for _, line := range lines {
  //   fmt.Println(line)
    // if strings.Contains(line, "css"){
    //   fmt.Println(line)
    //   time.Sleep(1 * time.Second )
    // }
  // }
}
