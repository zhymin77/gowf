package mongo

import (
  "framework"
  "labix.org/v2/mgo"
)

//Mongo db exec
func DBExec(db, collName string, mfunc func(coll *mgo.Collection)) {
  session, err := mgo.Dial(framework.Datacfg().Glob.MongoUrl)
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB(db).C(collName)
  mfunc(c)
}
