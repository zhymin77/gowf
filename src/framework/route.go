package framework

import (
  "bytes"
  "html/template"
  "net/http"
  "path/filepath"
  "strings"
)

// IRoute url register must implement.
type IRoute interface {
  R(m *map[string]func(http.ResponseWriter, *http.Request))
}

// Register register url and implemention.
func Register(rfunc func(*map[string]func(http.ResponseWriter, *http.Request))) {
  regMap := make(map[string]func(http.ResponseWriter, *http.Request))
  rfunc(&regMap)
  for url, method := range regMap {
    http.HandleFunc(url, method)
  }
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(staticPath))))
}

// RenderTemplate render html page.
func RenderTemplate(w http.ResponseWriter, tmpl string, d interface{}) {
  RenderTemplateWithFuncMap(w, tmpl, d, nil)
}

// RenderTemplateWithFuncMap render html page.
func RenderTemplateWithFuncMap(w http.ResponseWriter, tmpl string, d interface{}, m *map[string]interface{}) {
  data := map[string]interface{}{"GlobServerURL": GlobServerURL, "D": d, "helper": GetHelperInstance()}
  var layout *template.Template
  var funcM template.FuncMap
  if m != nil {
    funcM = template.FuncMap(*m)
  } else {
    funcM = nil
  }
  layoutM := loadLayout(&data, &funcM)
  (*layoutM)["CONTENT"] = loadTmpl(tmpl, &data, &funcM)
  (*layoutM)["GlobServerURL"] = GlobServerURL
  if funcM != nil {
    layout = template.Must(
      template.New("layout.tmpl").Funcs(funcM).ParseFiles(tmplPath + "layout/layout.tmpl"))
  } else {
    layout = template.Must(template.New("layout.tmpl").ParseFiles(tmplPath + "layout/layout.tmpl"))
  }
  err := layout.Execute(w, &layoutM)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

// RenderWithoutLayout render without layout.
func RenderWithoutLayout(w http.ResponseWriter, tmpl string, d interface{}) {
  RenderWithoutLayoutWithFuncMap(w, tmpl, d, nil)
}

// RenderWithoutLayoutWithFuncMap render without layout.
func RenderWithoutLayoutWithFuncMap(w http.ResponseWriter, tmpl string, d interface{}, m *map[string]interface{}) {
  data := map[string]interface{}{"GlobServerURL": GlobServerURL, "D": d, "helper": GetHelperInstance()}
  var funcM template.FuncMap
  var temp *template.Template
  if m != nil {
    funcM = template.FuncMap(*m)
  } else {
    funcM = nil
  }
  if funcM != nil {
    temp = template.Must(template.ParseFiles(tmplPath + tmpl)).Funcs(funcM)
  } else {
    temp = template.Must(template.ParseFiles(tmplPath + tmpl))
  }
  err := temp.Execute(w, &data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

// load tmpl
func loadTmpl(tmpl string, data interface{}, funcM *template.FuncMap) template.HTML {
  var b bytes.Buffer
  if funcM != nil {
    template.Must(template.New(tmpl).Funcs(*funcM).ParseFiles(tmplPath + tmpl)).Execute(&b, &data)
  } else {
    template.Must(template.ParseFiles(tmplPath + tmpl)).Execute(&b, &data)
  }
  return template.HTML(b.String())
}

// load layout, escape layout/layout.tmpl
func loadLayout(data interface{}, funcM *template.FuncMap) *map[string]interface{} {
  layoutM := make(map[string]interface{})
  filenames, _ := filepath.Glob(tmplPath + "layout/*.tmpl")
  for _, filename := range filenames {
    baseName := filepath.Base(filename)
    if "layout.tmpl" != baseName {
      var b bytes.Buffer
      bytes := []byte(baseName)
      name := string(bytes[: len(bytes) - 5])
      if funcM != nil {
        template.Must(template.ParseFiles(filename)).Funcs(*funcM).Execute(&b, &data)
      } else {
        template.Must(template.ParseFiles(filename)).Execute(&b, &data)
      }
      layoutM[strings.ToUpper(name)] = template.HTML(b.String())
    }
  }
  return &layoutM
}
