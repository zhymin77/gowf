package framework

import (
  "code.google.com/p/gcfg"
  "fmt"
)

// Config struct of cfg/*.cfg.
type Config struct {
  Glob struct {
    Port string
    ServerURL string
    MongoURL string
  }
  Path struct {
    Tmpl string
    Static string
  }
  Pagination struct {
    PageSize int
    ShowPage int
  }
}

// Datacfg load config once.
func Datacfg() *Config {
  cfgOnce.Do(LoadConfig)
  return &cfg
}

// LoadConfig implemention of loading configuration.
func LoadConfig() {
  path := "cfg/" + GetMode() + ".cfg"
  err := gcfg.ReadFileInto(&cfg, path)
  if err != nil {
    fmt.Println(err.Error())
  }
}
