package route

import (
  "net/http"
  "framework"
  "strconv"
)

// Register handlers.
func Register(m *map[string]func(http.ResponseWriter, *http.Request)) {
  framework.IRoute(new(Main)).R(m)
  http.Handle("/favicon.ico",
    http.StripPrefix("/favicon.ico", http.RedirectHandler("/assets/favicon.ico", http.StatusFound)))
}

func renderTemplate(r *http.Request, w http.ResponseWriter, tmpl string, data interface{}) {
  framework.RenderTemplate(w, tmpl, data)
}

func renderTemplateWithFuncMap(r *http.Request, w http.ResponseWriter, tmpl string, data interface{}, m *map[string]interface{}) {
  framework.RenderTemplateWithFuncMap(w, tmpl, data, m)
}

func renderWithoutLayout(r *http.Request, w http.ResponseWriter, tmpl string, d interface{}) {
  framework.RenderWithoutLayout(w, tmpl, d)
}

func renderWithoutLayoutWithFuncMap(r *http.Request, w http.ResponseWriter, tmpl string, d interface{}, m *map[string]interface{}) {
  framework.RenderWithoutLayoutWithFuncMap(w, tmpl, d, m)
}

func paramStr(key string, r *http.Request) string {
  return paramStrWithDef(key, r, "")
}

func paramStrWithDef(key string, r *http.Request, defv string) string {
  v := r.URL.Query()[key]
  if v == nil || len(v) == 0 {
    return defv
  }
  return string(v[0])
}

func paramInt(key string, r *http.Request) int {
  return paramIntWithDef(key, r, 0)
}

func paramIntWithDef(key string, r *http.Request, defv int) (rv int) {
  v := r.URL.Query()[key]
  if v == nil || len(v) == 0 {
    rv = defv
    return
  }
  rv, _ = strconv.Atoi(v[0])
  return
}
