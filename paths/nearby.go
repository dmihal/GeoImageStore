package paths

import (
    "fmt"
    "strconv"
    "net/http"

    "appengine"

    "photos"
)

func HandleNearby(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)

  lat, _ := strconv.ParseFloat(r.FormValue("lat"), 64)
  lng, _ := strconv.ParseFloat(r.FormValue("lng"), 64)

  point := appengine.GeoPoint{
    Lat: lat,
    Lng: lng,
  }
  allPhotos, _ := photos.GetNearest(point, c)

  fmt.Fprintf(w, "Found %d photos\n", len(allPhotos))
  for _, photo := range allPhotos {
    fmt.Fprint(w, photo.Name, "\n")
  }
}
