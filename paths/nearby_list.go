package paths

import (
  "html/template"
  "strconv"
  "net/http"

  "appengine"

  "photos"
  "lib"
)

func HandleNearbyList(w http.ResponseWriter, r *http.Request) {
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

  t, err := template.ParseFiles("templates/list.html")
  if err != nil {
    lib.ServeError(c, w, err)
    return
  }
  t.Execute(w, allPhotos)
}
