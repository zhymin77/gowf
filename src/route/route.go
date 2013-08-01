package route

import (
  "net/http"
  "framework"
)

//
// register route.
//
func Register(m *map[string]func(http.ResponseWriter, *http.Request)) {
  framework.IRoute(new(Index)).R(m)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
  framework.RenderTemplate(w, tmpl, data)
}
