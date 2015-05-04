package paths

import (
    "fmt"
    "strconv"
    "net/http"

    "appengine"

    "photos"
    "lib"
)

func HandleNearby(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)

  lat, _ := strconv.ParseFloat(r.FormValue("lat"), 64)
  lng, _ := strconv.ParseFloat(r.FormValue("lng"), 64)

  point := appengine.GeoPoint{
    Lat: lat,
    Lng: lng,
  }
  allPhotos, err := photos.GetNearest(point, c)

  if err != nil {
    lib.ServeError(c, w, err)
    return
  }

  json, _ := photos.EncodeJSON(allPhotos)

  w.Header().Set("Content-Type", "application/json")
  fmt.Fprint(w, json)
}
