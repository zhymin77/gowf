package framework

import (
  "bytes"
  "html/template"
  "net/http"
  "path/filepath"
  "strings"
)

type Route struct {
  GlobServerUrl string
  D interface{}
}

type IRoute interface {
  R(m *map[string]func(http.ResponseWriter, *http.Request))
}

func Register(rfunc func(*map[string]func(http.ResponseWriter, *http.Request))) {
  regMap := make(map[string]func(http.ResponseWriter, *http.Request))
  rfunc(&regMap)
  for url, method := range regMap {
    http.HandleFunc(url, method)
  }
}

// Render template
func RenderTemplate(w http.ResponseWriter, tmpl string, d interface{}) {
  data := Route{GlobServerUrl: GlobServerUrl, D: d}
  var err error
  layoutM := loadLayout(&data)
  (*layoutM)["CONTENT"] = loadTmpl(tmpl, &data)
  layout := template.Must(template.New("layout.tmpl").ParseFiles(templatePath + "layout/layout.tmpl"))
  err = layout.Execute(w, &layoutM)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

// load tmpl
func loadTmpl(tmpl string, data interface{}) template.HTML {
  var b bytes.Buffer
  template.Must(template.ParseFiles(templatePath + tmpl)).Execute(&b, &data)
  return template.HTML(b.String())
}

// load layout, escape layout/layout.tmpl
func loadLayout(data interface{}) *map[string]template.HTML {
  layoutM := make(map[string]template.HTML)
  filenames, _ := filepath.Glob(templatePath + "layout/*.tmpl")
  for _, filename := range filenames {
    baseName := filepath.Base(filename)
    if "layout.tmpl" != baseName {
      var b bytes.Buffer
      bytes := []byte(baseName)
      name := string(bytes[: len(bytes) - 5])
      template.Must(template.ParseFiles(filename)).Execute(&b, &data)
      layoutM[strings.ToUpper(name)] = template.HTML(b.String())
    }
  }
  return &layoutM
}
