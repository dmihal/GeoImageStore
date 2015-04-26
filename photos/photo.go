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
