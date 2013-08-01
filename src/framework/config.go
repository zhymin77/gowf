package framework

import (
  "code.google.com/p/gcfg"
  "fmt"
)

/** config struct, reference to cfg/*.cfg files */
type Config struct {
  Glob struct {
    Port string
    ServerUrl string
  }
  Template struct {
    Path string
  }
}

/**
 * loading cfg, only once
 */
func Datacfg() *Config {
  cfgOnce.Do(LoadConfig)
  return &cfg
}

func LoadConfig() {
  path := "cfg/" + GetMode() + ".cfg"
  err := gcfg.ReadFileInto(&cfg, path)
  if err != nil {
    fmt.Println(err.Error())
  }
}
