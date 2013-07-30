package main

import (
  "fmt"
  "net/http"
  "route"
  "util"
)

func main() {
  route.Init()
  fmt.Printf("Port: %s listening...\n", util.Datacfg().Section.Port)
  http.ListenAndServe(":" + util.Datacfg().Section.Port, nil)
}
