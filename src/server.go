package main

import (
  "net/http"
  "route"
  "util"
  "log"
)

func main() {
  route.Init()
  log.Printf("Port: %s listening...\n", util.Datacfg().Section.Port)
  http.ListenAndServe(":" + util.Datacfg().Section.Port, nil)
}
