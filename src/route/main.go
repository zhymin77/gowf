package route

import (
  "net/http"
  "util"
)

// Main handle detail page request.
// Need route.go register: framework.IRoute(new(Main)).R(m)
type Main int8

func (d *Main)R(m *map[string]func(http.ResponseWriter, *http.Request)) {

  (*m)["/"] = func(w http.ResponseWriter, r *http.Request) {
    if r.URL.String() != "/" {
      http.NotFound(w, r)
    } else {
      data := map[string]interface{}{
        "tip" : "Render page with layout",
        "say" : "Hello Gowf!"}
      renderTemplate(r, w, "l.tmpl", &data)
    }
  }

  (*m)["/lf"] = func(w http.ResponseWriter, r *http.Request) {
    data := map[string]interface{}{
      "tip" : "Render page with layout and function map",
      "say" : "Hello Gowf!", "time": int32(1385023552)}
    renderTemplateWithFuncMap(r, w, "lf.tmpl", &data, &map[string]interface{}{"ConvertUnix2String": util.ConvertUnix2String})
  }

  (*m)["/wolf"] = func(w http.ResponseWriter, r *http.Request) {
    data := map[string]interface{}{
      "tip" : "Render page without layout",
      "say" : "Hello Gowf!"}
    renderWithoutLayout(r, w, "l.tmpl", &data)
  }
}
