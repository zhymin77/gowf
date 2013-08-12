package framework

import (
  "sync"
  "html/template"
)

// const ModeDev value "dev"
const ModeDev = "dev"

var cfg Config
var cfgOnce, tmplsOnce, layoutOnce sync.Once
var templates map[string]*template.Template
var tmplPath = Datacfg().Path.Tmpl
var staticPath = Datacfg().Path.Static
// GlobServerURL extract value from *.cfg/glob/serverURL
var GlobServerURL = Datacfg().Glob.ServerURL
var layoutProto *template.Template
