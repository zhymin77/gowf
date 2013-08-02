package route

import (
  "net/http"
  "framework/mongo"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "fmt"
)

//
// alias int8 -> Index, reference @route/IRoute interface(implements R Method).
// must register @route.go/register
//

type PC struct {
  Id  string "_id"
  EndTime int "endTime"
  Name string
  StartTime int "startTime"
  Operate string
}

type Index int8

func (i *Index)R(m *map[string]func(http.ResponseWriter, *http.Request)) {

  // url'/test'
  (*m)["/test"] = func(w http.ResponseWriter, r *http.Request) {
    var result []PC
    mongo.DBExec("autoupdate", "ck", func(coll *mgo.Collection) {
      coll.Find(nil).All(&result)
      fmt.Println(result[0].StartTime)
      fmt.Println(result[0].EndTime)
      fmt.Println(result[0].Operate)
      fmt.Println(result[0].Name)
      fmt.Println(result[0].Id)
    })
    // render page
    renderTemplate(w, "body.tmpl", &result)
  }

  (*m)["/string"] = func(w http.ResponseWriter, r *http.Request) {
    var bm bson.M
    mongo.DBExec("autoupdate", "ck", func(coll *mgo.Collection) {
      coll.Find(bson.M{"_id":"51e75b53559a6e80ac2d5e72"}).One(&bm)
    })
    fmt.Fprint(w, bm)
  }

}
