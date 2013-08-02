package framework

import (
  "sync"
  "html/template"
)

const MODE_DEV = "dev"

var cfg Config
var cfgOnce, tmplsOnce, layoutOnce sync.Once
var templates map[string]*template.Template
var tmplPath = Datacfg().Path.Tmpl
var staticPath = Datacfg().Path.Static
var GlobServerUrl string = Datacfg().Glob.ServerUrl
var layoutProto *template.Template = nil
