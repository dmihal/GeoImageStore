package paths

import (
  "strconv"
  "time"
  "net/http"

  "appengine"
  "appengine/blobstore"

  "lib"
  "photos"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  blobs, values, err := blobstore.ParseUpload(r)
  if err != nil {
    lib.ServeError(c, w, err)
    return
  }
  file := blobs["file"]
  if len(file) == 0 {
    c.Errorf("no file uploaded")
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }

  lat, _ := strconv.ParseFloat(values.Get("lat"), 64)
  lng, _ := strconv.ParseFloat(values.Get("lng"), 64)
  photo := photos.Photo{
    Name: values.Get("title"),
    File: file[0].BlobKey,
    Location: appengine.GeoPoint{
      Lat: lat,
      Lng: lng,
    },
    Date: time.Now(),
  }
  err = photos.Add(photo, c)
  if err != nil {
    lib.ServeError(c, w, err)
    return
  }

  http.Redirect(w, r, "/serve/?blobKey="+string(file[0].BlobKey), http.StatusFound)
}
