package photos

import (
  "appengine"
  "appengine/datastore"
)

func Add(photo Photo, c appengine.Context) {
  key := datastore.NewKey(c, "Photo", string(photo.File), 0, nil)
  datastore.Put(c, key, &photo)
}
