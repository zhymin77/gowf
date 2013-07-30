package util

import (
  "code.google.com/p/gcfg"
  "fmt"
  "sync"
)

/** config struct, reference to cfg/*.cfg files */
type Config struct {
  Section struct {
    Port string
  }
  Template struct {
    Path string
  }
}

/**
 * loading cfg, only once
 */
var cfg Config
var once sync.Once
func Datacfg() *Config {
  once.Do(LoadConfig)
  return &cfg
}

func LoadConfig() {
  path := "cfg/" + GetMode() + ".cfg"
  err := gcfg.ReadFileInto(&cfg, path)
  if err != nil {
    fmt.Println(err.Error())
  }
}
