package photos

import (
  "time"
  "appengine"
)

type Photo struct {
  Name     string
  Date     time.Time
  Account  string
  Location appengine.GeoPoint
  File     appengine.BlobKey
}

type searchPhoto struct {
  Date     time.Time
  Location appengine.GeoPoint
}

func makeSearchPhoto(photo Photo) searchPhoto {
  sPhoto := searchPhoto{
    Date: photo.Date,
    Location: photo.Location,
  }
  return sPhoto
}
