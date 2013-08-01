package route

import (
  "net/http"
)

//
// alias int8 -> Index, reference @route/IRoute interface(implements R Method).
// must register @route.go/register
//
type Index int8

func (i *Index)R(m *map[string]func(http.ResponseWriter, *http.Request)) {

  // url'/test'
  (*m)["/test"] = func(w http.ResponseWriter, r *http.Request) {
    // render page
    renderTemplate(w, "body.tmpl", "a")
  }

}
