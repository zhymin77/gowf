package main

import (
  "net/http"
  "route"
  "framework"
  "log"
)

func main() {
  framework.Register(route.Register)
  log.Printf("Port: %s listening...\n", framework.Datacfg().Glob.Port)
  http.ListenAndServe(":" + framework.Datacfg().Glob.Port, nil)
}
