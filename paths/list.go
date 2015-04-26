package paths

import (
    "fmt"
    "net/http"

    "appengine"

    "photos"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)

  allPhotos := photos.GetAll(c)

  for _, photo := range allPhotos {
    fmt.Fprint(w, photo.Name, "\n")
  }
}
