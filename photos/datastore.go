package photos

import (
  "appengine"
  "appengine/datastore"
)

func Add(photo Photo, c appengine.Context) {
  key := datastore.NewKey(c, "Photo", string(photo.File), 0, nil)
  datastore.Put(c, key, &photo)
}
func GetAll(c appengine.Context) []Photo {
  query := datastore.NewQuery("Photo")

  var photos []Photo

  for t := query.Run(c); ; {
    var photo Photo
    _, err := t.Next(&photo)
    if err == datastore.Done {
      break
    }
    if err != nil {
      return nil
    }
    photos = append(photos, photo)
  }
  return photos
}
