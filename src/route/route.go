package route

import (
  "html/template"
  "net/http"
  "sync"
  "util"
  "path/filepath"
)

//
// register route.
//
func register(m *map[string]func(http.ResponseWriter, *http.Request)) {
  IRoute(new(Index)).R(m)
}

type IRoute interface {
  R(m *map[string]func(http.ResponseWriter, *http.Request))
}

func Init() {
  regMap := make(map[string]func(http.ResponseWriter, *http.Request))
  register(&regMap)
  for url, method := range regMap {
    http.HandleFunc(url, method)
  }
}


var once sync.Once
var path = util.Datacfg().Template.Path
var templates map[string]*template.Template
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
  var err error
  if util.GetMode() == util.MODE_DEV {
    err = loadTmpl(tmpl).Execute(w, data)
  } else {
    once.Do(cachingTmpls)
    err = templates[tmpl].Execute(w, data)
  }
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func cachingTmpls() {
  filenames, _ := filepath.Glob(path + "/*.tmpl")
  templates = make(map[string]*template.Template, len(filenames))
  for _, filename := range filenames {
    name := filepath.Base(filename)
    templates[name] = loadTmpl(name)
  }
}

func loadTmpl(tmpl string) *template.Template {
  return template.Must(layoutTmpl().ParseFiles(path + tmpl))
}
var layoutProto *template.Template = nil
func layoutTmpl() *template.Template {
  if util.GetMode() == util.MODE_DEV {
    return template.Must(template.New("layout").ParseGlob(path + "layout/*.tmpl"))
  } else {
    once.Do(func(){ layoutProto = template.Must(template.New("layout").ParseGlob(path + "layout/*.tmpl"))})
    tmpl, err := layoutProto.Clone()
    if err != nil {
      panic(err.Error())
    }
    return tmpl
  }
}
